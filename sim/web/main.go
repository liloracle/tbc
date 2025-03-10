package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"strings"
	"syscall"
	"time"

	dist "github.com/wowsims/tbc/binary_dist"
	"github.com/wowsims/tbc/sim"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"

	googleProto "google.golang.org/protobuf/proto"
)

func init() {
	sim.RegisterAll()
}

func main() {
	var useFS = flag.Bool("usefs", false, "Use local file system for client files. Set to true during development.")
	var wasm = flag.Bool("wasm", false, "Use wasm for sim instead of web server apis. Can only be used with usefs=true")
	var simName = flag.String("sim", "", "Name of simulator to launch (ex: balance_druid, elemental_shaman, etc)")
	var host = flag.String("host", ":3333", "URL to host the interface on.")
	var launch = flag.Bool("launch", true, "auto launch browser")

	flag.Parse()

	runServer(*useFS, *host, *launch, *simName, *wasm, bufio.NewReader(os.Stdin))
}

func runServer(useFS bool, host string, launchBrowser bool, simName string, wasm bool, inputReader *bufio.Reader) {
	var fs http.Handler
	if useFS {
		log.Printf("Using local file system for development.")
		fs = http.FileServer(http.Dir("./dist"))
	} else {
		log.Printf("Embedded file server running.")
		fs = http.FileServer(http.FS(dist.FS))
	}

	http.HandleFunc("/statWeights", handleAPI)
	http.HandleFunc("/computeStats", handleAPI)
	http.HandleFunc("/individualSim", handleAPI)
	http.HandleFunc("/raidSim", handleAPI)
	http.HandleFunc("/gearList", handleAPI)

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, "/tbc/") {
			resp.Write([]byte(`
				<html><body><a href="/tbc/elemental_shaman">Elemental Shaman Sim</a"><br>
				<html><body><a href="/tbc/enhancement_shaman">Enhancement Shaman Sim</a"><br>
				<a href="/tbc/balance_druid">Balance Druid Sim</a"><br>
				<a href="/tbc/shadow_priest">Shadow Priest Sim</a"></body></html>
		    `))
			return
		}

		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		if !useFS || (useFS && !wasm) {
			if strings.HasSuffix(req.URL.Path, "sim_worker.js") {
				req.URL.Path = strings.Replace(req.URL.Path, "sim_worker.js", "net_worker.js", 1)
			}
		}
		fs.ServeHTTP(resp, req)
	})

	if launchBrowser {
		url := fmt.Sprintf("http://localhost%s/tbc/%s", host, simName)
		log.Printf("Launching interface on %s", url)
		go func() {
			var cmd *exec.Cmd
			if runtime.GOOS == "windows" {
				cmd = exec.Command("start", "msedge", url)
			} else if runtime.GOOS == "darwin" {
				cmd = exec.Command("open", url)
			} else if runtime.GOOS == "linux" {
				cmd = exec.Command("xdg-open", url)
			}
			err := cmd.Start()
			if err != nil {
				log.Printf("Error launching browser: %#v", err.Error())
			}
		}()
	}

	go func() {
		// Launch server!
		log.Printf("Closing: %s", http.ListenAndServe(host, nil))
	}()

	// used to read a CTRL+C
	c := make(chan os.Signal, 10)
	signal.Notify(c, syscall.SIGINT)

	go func() {
		<-c
		log.Printf("Shutting down")
		os.Exit(0)
	}()
	fmt.Printf("Enter Command... '?' for list\n")
	for {
		fmt.Printf("> ")
		text, err := inputReader.ReadString('\n')
		if err != nil {
			// block forever
			<-c
			os.Exit(-1)
		}
		if len(text) == 0 {
			continue
		}
		command := strings.TrimSpace(text)
		switch command {
		case "profile":
			filename := fmt.Sprintf("profile_%d.cpu", time.Now().Unix())
			fmt.Printf("Running profiling for 15 seconds, output to %s\n", filename)
			f, err := os.Create(filename)
			if err != nil {
				log.Fatal("could not create CPU profile: ", err)
			}
			if err := pprof.StartCPUProfile(f); err != nil {
				log.Fatal("could not start CPU profile: ", err)
			}
			go func() {
				time.Sleep(time.Second * 15)
				pprof.StopCPUProfile()
				f.Close()
				fmt.Printf("Profiling complete.\n> ")
			}()
		case "quit":
			os.Exit(1)
		case "?":
			fmt.Printf("Commands:\n\tprofile - start a CPU profile for debugging performance\n\tquit - exits\n\n")
		case "":
			// nothing.
		default:
			fmt.Printf("Unknown command: '%s'", command)
		}
	}
}

type apiHandler struct {
	msg    func() googleProto.Message
	handle func(googleProto.Message) googleProto.Message
}

// Handlers to decode and handle each proto function
var handlers = map[string]apiHandler{
	"/raidSim": {msg: func() googleProto.Message { return &proto.RaidSimRequest{} }, handle: func(msg googleProto.Message) googleProto.Message {
		return core.RunRaidSim(msg.(*proto.RaidSimRequest))
	}},
	"/individualSim": {msg: func() googleProto.Message { return &proto.IndividualSimRequest{} }, handle: func(msg googleProto.Message) googleProto.Message {
		return core.RunIndividualSim(msg.(*proto.IndividualSimRequest))
	}},
	"/statWeights": {msg: func() googleProto.Message { return &proto.StatWeightsRequest{} }, handle: func(msg googleProto.Message) googleProto.Message {
		return core.StatWeights(msg.(*proto.StatWeightsRequest))
	}},
	"/computeStats": {msg: func() googleProto.Message { return &proto.ComputeStatsRequest{} }, handle: func(msg googleProto.Message) googleProto.Message {
		return core.ComputeStats(msg.(*proto.ComputeStatsRequest))
	}},
	"/gearList": {msg: func() googleProto.Message { return &proto.GearListRequest{} }, handle: func(msg googleProto.Message) googleProto.Message {
		return core.GetGearList(msg.(*proto.GearListRequest))
	}},
}

// handleAPI is generic handler for any api function using protos.
func handleAPI(w http.ResponseWriter, r *http.Request) {
	endpoint := r.URL.Path

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		return
	}
	handler, ok := handlers[endpoint]
	if !ok {
		log.Printf("Invalid Endpoint: %s", endpoint)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	msg := handler.msg()
	if err := googleProto.Unmarshal(body, msg); err != nil {
		log.Printf("Failed to parse request: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := handler.handle(msg)

	outbytes, err := googleProto.Marshal(result)
	if err != nil {
		log.Printf("[ERROR] Failed to marshal result: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/x-protobuf")
	w.Write(outbytes)
}

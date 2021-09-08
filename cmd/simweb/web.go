package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/wowsims/tbc/api"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/ui"
)

func main() {
	var useFS = flag.Bool("usefs", false, "Use local file system and wasm. Set to true for dev")
	var host = flag.String("host", ":3333", "URL to host the interface on.")

	flag.Parse()

	var fs http.Handler
	if *useFS {
		log.Printf("Using local file system for development.")
		fs = http.FileServer(http.Dir("."))
	} else {
		log.Printf("Embedded file server running.")
		fs = http.FileServer(http.FS(ui.FS))
	}

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	})

	url := fmt.Sprintf("http://localhost%s/ui", *host)
	log.Printf("Launching interface on %s", url)

	go func() {
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("explorer", url)
		} else if runtime.GOOS == "darwin" {
			cmd = exec.Command("open", url)
		} else if runtime.GOOS == "linux" {
			cmd = exec.Command("xdg-open", url)
		}
		err := cmd.Start()
		if err != nil {
			log.Printf("Error launching browser: %#v", err.Error())
		}
		log.Printf("Closing: %s", http.ListenAndServe(*host, nil))
	}()

	fmt.Printf("Enter Command... '?' for list\n")
	for {
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
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

// result := GetGearList(*request.GearList)
// return ApiResult{GearList: &result}
// } else if request.ComputeStats != nil {
// result := ComputeStats(*request.ComputeStats)
// return ApiResult{ComputeStats: &result}
// } else if request.StatWeights != nil {
// result := StatWeights(*request.StatWeights)
// return ApiResult{StatWeights: &result}
// } else if request.Sim != nil {

func handleIndividualSim(w http.ResponseWriter, r *http.Request) {
	isr := api.IndividualSimRequest{}

	req := core.SimRequest{}

	result := core.RunSimulation(req)
}
func handleRaidSim(w http.ResponseWriter, r *http.Request) {

}
func handleGearList(w http.ResponseWriter, r *http.Request) {

}
func handleComputeStats(w http.ResponseWriter, r *http.Request) {

}
func handleStatWeights(w http.ResponseWriter, r *http.Request) {
}

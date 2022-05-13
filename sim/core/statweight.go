package core

import (
	"math"
	"math/rand"
	"sync"
	"sync/atomic"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	googleProto "google.golang.org/protobuf/proto"
)

type StatWeightValues struct {
	Weights       stats.Stats
	WeightsStdev  stats.Stats
	EpValues      stats.Stats
	EpValuesStdev stats.Stats
}

func (swv StatWeightValues) ToProto() *proto.StatWeightValues {
	return &proto.StatWeightValues{
		Weights:       swv.Weights[:],
		WeightsStdev:  swv.WeightsStdev[:],
		EpValues:      swv.EpValues[:],
		EpValuesStdev: swv.EpValuesStdev[:],
	}
}

type StatWeightsResult struct {
	Dps StatWeightValues
}

func (swr StatWeightsResult) ToProto() *proto.StatWeightsResult {
	return &proto.StatWeightsResult{
		Dps: swr.Dps.ToProto(),
	}
}

func CalcStatWeight(swr proto.StatWeightsRequest, statsToWeigh []stats.Stat, referenceStat stats.Stat, progress chan *proto.ProgressMetrics) StatWeightsResult {
	if swr.Player.BonusStats == nil {
		swr.Player.BonusStats = make([]float64, stats.Len)
	}

	raidProto := SinglePlayerRaidProto(swr.Player, swr.PartyBuffs, swr.RaidBuffs)
	baseStatsResult := ComputeStats(&proto.ComputeStatsRequest{
		Raid: raidProto,
	})
	baseStats := baseStatsResult.RaidStats.Parties[0].Players[0].FinalStats

	baseSimRequest := &proto.RaidSimRequest{
		Raid:       raidProto,
		Encounter:  swr.Encounter,
		SimOptions: swr.SimOptions,
	}
	baselineResult := RunRaidSim(baseSimRequest)
	baselineDpsMetrics := baselineResult.RaidMetrics.Parties[0].Players[0].Dps

	var waitGroup sync.WaitGroup

	// Do half the iterations with a positive, and half with a negative value for better accuracy.
	resultLow := StatWeightsResult{}
	resultHigh := StatWeightsResult{}
	dpsHistsLow := [stats.Len]map[int32]int32{}
	dpsHistsHigh := [stats.Len]map[int32]int32{}

	var iterationsTotal int32
	var iterationsDone int32
	var simsTotal int32
	var simsCompleted int32

	doStat := func(stat stats.Stat, value float64, isLow bool) {
		defer waitGroup.Done()

		simRequest := googleProto.Clone(baseSimRequest).(*proto.RaidSimRequest)
		simRequest.Raid.Parties[0].Players[0].BonusStats[stat] += value
		simRequest.SimOptions.Iterations /= 2 // Cut in half since we're doing above and below separately.

		reporter := make(chan *proto.ProgressMetrics, 10)
		go RunSim(*simRequest, reporter) // RunRaidSim(simRequest)

		var localIterations int32
		var simResult *proto.RaidSimResult
	statsim:
		for {
			select {
			case metrics, ok := <-reporter:
				if !ok {
					break statsim
				}
				atomic.AddInt32(&iterationsDone, (metrics.CompletedIterations - localIterations))
				localIterations = metrics.CompletedIterations
				if metrics.FinalRaidResult != nil {
					atomic.AddInt32(&simsCompleted, 1)
					simResult = metrics.FinalRaidResult
				}
				if progress != nil {
					progress <- &proto.ProgressMetrics{
						TotalIterations:     atomic.LoadInt32(&iterationsTotal),
						CompletedIterations: atomic.LoadInt32(&iterationsDone),
						CompletedSims:       atomic.LoadInt32(&simsCompleted),
						TotalSims:           atomic.LoadInt32(&simsTotal),
					}
				}
				if metrics.FinalRaidResult != nil {
					break statsim
				}
			}
		}
		dpsMetrics := simResult.RaidMetrics.Parties[0].Players[0].Dps
		dpsDiff := (dpsMetrics.Avg - baselineDpsMetrics.Avg) / value

		if isLow {
			resultLow.Dps.Weights[stat] = dpsDiff
			dpsHistsLow[stat] = dpsMetrics.Hist
		} else {
			resultHigh.Dps.Weights[stat] = dpsDiff
			dpsHistsHigh[stat] = dpsMetrics.Hist
		}
	}

	const defaultStatMod = 50.0
	statModsLow := stats.Stats{}
	statModsHigh := stats.Stats{}

	// Make sure reference stat is included.
	statModsLow[referenceStat] = defaultStatMod
	statModsHigh[referenceStat] = defaultStatMod

	for _, stat := range statsToWeigh {
		statMod := defaultStatMod
		statModsHigh[stat] = statMod
		statModsLow[stat] = -statMod
	}

	for stat, _ := range statModsLow {
		if statModsLow[stat] == 0 {
			continue
		}
		waitGroup.Add(2)
		atomic.AddInt32(&iterationsTotal, swr.SimOptions.Iterations)
		atomic.AddInt32(&simsTotal, 2)

		go doStat(stats.Stat(stat), statModsLow[stat], true)
		go doStat(stats.Stat(stat), statModsHigh[stat], false)
	}

	waitGroup.Wait()

	for _, stat := range statsToWeigh {
		if stat == stats.SpellHit || stat == stats.MeleeHit {
			// Check for hard caps.
			if resultHigh.Dps.Weights[stat] < 0.05 {
				statModsHigh[stat] = 0
			} else if baseStats[stat] > 80 {
				// For spell/melee hit, don't use the high comparison if we might be anywhere
				// remotely close to cap.
				resultHigh.Dps.Weights[stat] = resultLow.Dps.Weights[stat]
				statModsHigh[stat] = statModsLow[stat]
			}
		}
	}

	result := StatWeightsResult{}
	for statIdx, _ := range statModsLow {
		stat := stats.Stat(statIdx)
		if statModsLow[stat] == 0 || statModsHigh[stat] == 0 {
			continue
		}

		result.Dps.Weights[stat] = (resultLow.Dps.Weights[stat] + resultHigh.Dps.Weights[stat]) / 2
	}

	for statIdx, _ := range statModsLow {
		stat := stats.Stat(statIdx)
		if statModsLow[stat] == 0 || statModsHigh[stat] == 0 {
			continue
		}

		result.Dps.EpValues[stat] = result.Dps.Weights[stat] / result.Dps.Weights[referenceStat]

		weightStdevLow := computeStDevFromHists(swr.SimOptions.Iterations/2, statModsLow[stat], dpsHistsLow[stat], baselineDpsMetrics.Hist, nil, statModsLow[referenceStat])
		weightStdevHigh := computeStDevFromHists(swr.SimOptions.Iterations/2, statModsHigh[stat], dpsHistsHigh[stat], baselineDpsMetrics.Hist, nil, statModsHigh[referenceStat])
		result.Dps.WeightsStdev[stat] = (weightStdevLow + weightStdevHigh) / 2

		epStdevLow := computeStDevFromHists(swr.SimOptions.Iterations/2, statModsLow[stat], dpsHistsLow[stat], baselineDpsMetrics.Hist, dpsHistsLow[referenceStat], statModsLow[referenceStat])
		epStdevHigh := computeStDevFromHists(swr.SimOptions.Iterations/2, statModsHigh[stat], dpsHistsHigh[stat], baselineDpsMetrics.Hist, dpsHistsHigh[referenceStat], statModsHigh[referenceStat])
		result.Dps.EpValuesStdev[stat] = (epStdevLow + epStdevHigh) / 2
	}

	return result
}

func computeStDevFromHists(iters int32, modValue float64, moddedStatDpsHist map[int32]int32, baselineDpsHist map[int32]int32, referenceDpsHist map[int32]int32, referenceModValue float64) float64 {
	sum := 0.0
	sumSquared := 0.0
	n := iters * 10
	for i := int32(0); i < n; {
		denominator := 1.0
		if referenceDpsHist != nil {
			denominator = float64(sampleFromDpsHist(referenceDpsHist, iters)-sampleFromDpsHist(baselineDpsHist, iters)) / referenceModValue
		}

		if denominator != 0 {
			ep := (float64(sampleFromDpsHist(moddedStatDpsHist, iters)-sampleFromDpsHist(baselineDpsHist, iters)) / modValue) / denominator
			sum += ep
			sumSquared += ep * ep
			i++
		}
	}
	epAvg := sum / float64(n)
	epStDev := math.Sqrt((sumSquared / float64(n)) - (epAvg * epAvg))
	return epStDev
}

// Picks a random value from a histogram, taking into account the bucket sizes.
func sampleFromDpsHist(hist map[int32]int32, histNumSamples int32) int32 {
	r := rand.Float64()
	sampleIdx := int32(math.Floor(float64(histNumSamples) * r))

	var curSampleIdx int32
	for roundedDps, count := range hist {
		curSampleIdx += count
		if curSampleIdx >= sampleIdx {
			return roundedDps
		}
	}

	panic("Invalid dps histogram")
}

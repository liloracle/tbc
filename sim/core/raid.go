package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Party struct {
	Raid  *Raid
	Index int

	Players []Agent
	Pets    []PetAgent // Cached list of all the pets in the party.

	PlayersAndPets []Agent // Cached list of players + pets, concatenated.

	dpsMetrics DistributionMetrics
}

func NewParty(raid *Raid, index int, partyConfig proto.Party) *Party {
	party := &Party{
		Raid:       raid,
		Index:      index,
		dpsMetrics: NewDistributionMetrics(),
	}

	for playerIndex, playerConfig := range partyConfig.Players {
		if playerConfig != nil && playerConfig.Class != proto.Class_ClassUnknown {
			party.Players = append(party.Players, NewAgent(party, playerIndex, *playerConfig))
		}
	}

	return party
}

func (party *Party) Size() int {
	return len(party.Players)
}

func (party *Party) IsFull() bool {
	return party.Size() >= 5
}

func (party *Party) GetPartyBuffs(basePartyBuffs *proto.PartyBuffs) proto.PartyBuffs {
	// Compute the full party buffs for this party.
	partyBuffs := proto.PartyBuffs{}
	if basePartyBuffs != nil {
		partyBuffs = *basePartyBuffs
	}
	for _, player := range party.Players {
		player.AddPartyBuffs(&partyBuffs)
		player.GetCharacter().AddPartyBuffs(&partyBuffs)
	}
	return partyBuffs
}

func (party *Party) AddStats(newStats stats.Stats) {
	for _, agent := range party.Players {
		agent.GetCharacter().AddStats(newStats)
	}
}

func (party *Party) AddStat(stat stats.Stat, amount float64) {
	for _, agent := range party.Players {
		agent.GetCharacter().AddStat(stat, amount)
	}
}

func (party *Party) reset(sim *Simulation) {
	for _, agent := range party.Players {
		agent.GetCharacter().reset(sim, agent)

		agent.GetCharacter().SetGCDTimer(sim, 0)
		for _, petAgent := range agent.GetCharacter().Pets {
			if petAgent.GetPet().initialEnabled {
				petAgent.GetPet().Enable(sim, petAgent)
			}
		}
	}

	party.dpsMetrics.reset()
}

func (party *Party) doneIteration(sim *Simulation) {
	for _, agent := range party.Players {
		agent.GetCharacter().doneIteration(sim)
		party.dpsMetrics.Total += agent.GetCharacter().Metrics.dps.Total
	}

	party.dpsMetrics.doneIteration(sim.Duration.Seconds())
}

func (party *Party) GetMetrics(numIterations int32) *proto.PartyMetrics {
	metrics := &proto.PartyMetrics{
		Dps: party.dpsMetrics.ToProto(numIterations),
	}

	playerIdx := 0
	i := 0
	for playerIdx < len(party.Players) {
		player := party.Players[playerIdx]
		if player.GetCharacter().PartyIndex == i {
			metrics.Players = append(metrics.Players, player.GetCharacter().GetMetricsProto(numIterations))
			playerIdx++
		} else {
			metrics.Players = append(metrics.Players, &proto.UnitMetrics{})
		}
		i++
	}

	return metrics
}
func (party *Party) GetStats() *proto.PartyStats {
	partyStats := &proto.PartyStats{}

	playerIdx := 0
	i := 0
	for playerIdx < len(party.Players) {
		player := party.Players[playerIdx]
		if player.GetCharacter().PartyIndex == i {
			partyStats.Players = append(partyStats.Players, player.GetCharacter().GetStatsProto())
			playerIdx++
		} else {
			partyStats.Players = append(partyStats.Players, &proto.PlayerStats{})
		}
		i++
	}

	return partyStats
}

type Raid struct {
	Parties []*Party

	dpsMetrics DistributionMetrics

	AllUnits []*Unit // Cached list of all Units (players and pets) in the raid.
}

// Makes a new raid.
func NewRaid(raidConfig proto.Raid) *Raid {
	raid := &Raid{
		dpsMetrics: NewDistributionMetrics(),
	}

	if raidConfig.StaggerStormstrikes {
		enhanceShaman := RaidPlayersWithSpec(raidConfig, proto.Spec_SpecEnhancementShaman)
		if len(enhanceShaman) > 1 {
			stagger := time.Duration(float64(time.Second*10) / float64(len(enhanceShaman)))
			for i, shaman := range enhanceShaman {
				delay := stagger * time.Duration(i)
				shaman.Spec.(*proto.Player_EnhancementShaman).EnhancementShaman.Rotation.FirstStormstrikeDelay = delay.Seconds()
			}
		}
	}

	for partyIndex, partyConfig := range raidConfig.Parties {
		if partyConfig != nil {
			raid.Parties = append(raid.Parties, NewParty(raid, partyIndex, *partyConfig))
		}
	}

	return raid
}

func (raid *Raid) Size() int {
	totalPlayers := 0
	for _, party := range raid.Parties {
		totalPlayers += party.Size()
	}
	return totalPlayers
}

func (raid *Raid) IsFull() bool {
	return raid.Size() >= 25
}

func (raid *Raid) GetRaidBuffs(baseRaidBuffs *proto.RaidBuffs) proto.RaidBuffs {
	// Compute the full raid buffs from the raid.
	raidBuffs := proto.RaidBuffs{}
	if baseRaidBuffs != nil {
		raidBuffs = *baseRaidBuffs
	}
	for _, party := range raid.Parties {
		for _, player := range party.Players {
			player.AddRaidBuffs(&raidBuffs)
			player.GetCharacter().AddRaidBuffs(&raidBuffs)
		}
	}
	return raidBuffs
}

// Precompute the playersAndPets array for each party.
func (raid *Raid) updatePlayersAndPets() {
	raidPlayers := []*Unit{}
	raidPets := []*Unit{}

	for _, party := range raid.Parties {
		party.Pets = []PetAgent{}
		for _, player := range party.Players {
			for _, petAgent := range player.GetCharacter().Pets {
				party.Pets = append(party.Pets, petAgent)
				raidPets = append(raidPets, &petAgent.GetPet().Unit)
			}
		}
		party.PlayersAndPets = make([]Agent, len(party.Players)+len(party.Pets))
		for i, player := range party.Players {
			party.PlayersAndPets[i] = player
			raidPlayers = append(raidPlayers, &player.GetCharacter().Unit)
		}
		for i, pet := range party.Pets {
			party.PlayersAndPets[len(party.Players)+i] = pet
		}
	}

	raid.AllUnits = append(raidPlayers, raidPets...)
}

func (raid *Raid) applyCharacterEffects(raidConfig proto.Raid) {
	raidBuffs := raid.GetRaidBuffs(raidConfig.Buffs)

	for partyIdx, party := range raid.Parties {
		partyConfig := *raidConfig.Parties[partyIdx]
		partyBuffs := party.GetPartyBuffs(partyConfig.Buffs)

		// Apply all buffs to the players in this party.
		for playerIdx, player := range party.Players {
			playerConfig := *partyConfig.Players[playerIdx]
			individualBuffs := proto.IndividualBuffs{}
			if playerConfig.Buffs != nil {
				individualBuffs = *playerConfig.Buffs
			}

			player.GetCharacter().applyAllEffects(player, raidBuffs, partyBuffs, individualBuffs)
		}
	}
}

func (raid Raid) AddStats(s stats.Stats) {
	for _, party := range raid.Parties {
		party.AddStats(s)
	}
}

func (raid Raid) GetPlayerFromRaidTarget(raidTarget proto.RaidTarget) Agent {
	raidIndex := raidTarget.TargetIndex

	partyIndex := int(raidIndex / 5)
	playerIndex := int(raidIndex % 5)

	if partyIndex < 0 || partyIndex >= len(raid.Parties) {
		return nil
	}

	party := raid.Parties[partyIndex]

	if playerIndex < 0 || playerIndex >= len(party.Players) {
		return nil
	}

	return party.Players[playerIndex]
}

func (raid *Raid) reset(sim *Simulation) {
	for _, party := range raid.Parties {
		party.reset(sim)
	}
	raid.dpsMetrics.reset()
}

func (raid *Raid) doneIteration(sim *Simulation) {
	for _, party := range raid.Parties {
		party.doneIteration(sim)
		raid.dpsMetrics.Total += party.dpsMetrics.Total
	}

	raid.dpsMetrics.doneIteration(sim.Duration.Seconds())
}

func (raid *Raid) GetMetrics(numIterations int32) *proto.RaidMetrics {
	metrics := &proto.RaidMetrics{
		Dps: raid.dpsMetrics.ToProto(numIterations),
	}
	for _, party := range raid.Parties {
		metrics.Parties = append(metrics.Parties, party.GetMetrics(numIterations))
	}
	return metrics
}

func (raid *Raid) GetStats() *proto.RaidStats {
	raidStats := &proto.RaidStats{}
	for _, party := range raid.Parties {
		raidStats.Parties = append(raidStats.Parties, party.GetStats())
	}
	return raidStats
}

func SinglePlayerRaidProto(player *proto.Player, partyBuffs *proto.PartyBuffs, raidBuffs *proto.RaidBuffs) *proto.Raid {
	return &proto.Raid{
		Parties: []*proto.Party{
			&proto.Party{
				Players: []*proto.Player{
					player,
				},
				Buffs: partyBuffs,
			},
		},
		Buffs: raidBuffs,
	}
}

func RaidPlayersWithSpec(raid proto.Raid, spec proto.Spec) []*proto.Player {
	var specPlayers []*proto.Player
	for _, party := range raid.Parties {
		for _, player := range party.Players {
			if player != nil && player.GetSpec() != nil && PlayerProtoToSpec(*player) == spec {
				specPlayers = append(specPlayers, player)
			}
		}
	}
	return specPlayers
}

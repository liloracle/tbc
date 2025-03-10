package shaman

import (
	"log"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemEffect(19344, ApplyNaturalAlignmentCrystal)
	core.AddItemEffect(30663, ApplyFathomBroochOfTheTidewalker)
	core.AddItemEffect(33506, ApplySkycallTotem)

	core.AddItemSet(ItemSetTidefury)
	core.AddItemSet(ItemSetCycloneRegalia)
	core.AddItemSet(ItemSetCataclysmRegalia)
	core.AddItemSet(ItemSetSkyshatterRegalia)

	// Even though these item effects are handled elsewhere, add them so they are
	// detected for automatic testing.
	core.AddItemEffect(TotemOfThePulsingEarth, func(core.Agent) {})
}

var Tidefury2PcAuraID = core.NewAuraID()
var ItemSetTidefury = core.ItemSet{
	Name:  "Tidefury Raiment",
	Items: map[int32]struct{}{28231: {}, 27510: {}, 28349: {}, 27909: {}, 27802: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Handled in chain_lightning.go
		},
		4: func(agent core.Agent) {
			shamanAgent, ok := agent.(ShamanAgent)
			if !ok {
				log.Fatalf("Non-shaman attempted to activate shaman cyclone set bonus.")
			}
			shaman := shamanAgent.GetShaman()

			if shaman.SelfBuffs.WaterShield {
				shaman.AddStat(stats.MP5, 3)
			}
		},
	},
}

var Cyclone4PcAuraID = core.NewAuraID()
var Cyclone4PcManaRegainAuraID = core.NewAuraID()
var ItemSetCycloneRegalia = core.ItemSet{
	Name:  "Cyclone Regalia",
	Items: map[int32]struct{}{29033: {}, 29035: {}, 29034: {}, 29036: {}, 29037: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Handled in shaman.go
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID:   Cyclone4PcAuraID,
					Name: "Cyclone 4pc Bonus",
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						if !spellEffect.Crit || sim.RandomFloat("cycl4p") > 0.11 {
							return // if not a crit or didn't proc, don't activate
						}
						character.AddAura(sim, core.Aura{
							ID:   Cyclone4PcManaRegainAuraID,
							Name: "Cyclone Mana Cost Reduction",
							OnCast: func(sim *core.Simulation, cast *core.Cast) {
								// TODO: how to make sure this goes in before clearcasting?
								cast.ManaCost -= 270
							},
							OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
								character.RemoveAura(sim, Cyclone4PcManaRegainAuraID)
							},
						})
					},
				}
			})
		},
	},
}

var Cataclysm4PcAuraID = core.NewAuraID()
var ItemSetCataclysmRegalia = core.ItemSet{
	Name:  "Cataclysm Regalia",
	Items: map[int32]struct{}{30169: {}, 30170: {}, 30171: {}, 30172: {}, 30173: {}},
	Bonuses: map[int32]core.ApplyEffect{
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID:   Cataclysm4PcAuraID,
					Name: "Cataclysm 4pc Bonus",
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						if !spellEffect.Crit || sim.RandomFloat("cata4p") > 0.25 {
							return
						}
						character.AddMana(sim, 120, "Cataclysm 4p Bonus", false)
					},
				}
			})
		},
	},
}

var ItemSetSkyshatterRegalia = core.ItemSet{
	Name:  "Skyshatter Regalia",
	Items: map[int32]struct{}{34437: {}, 31017: {}, 34542: {}, 31008: {}, 31014: {}, 31020: {}, 31023: {}, 34566: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MP5, 15)
			agent.GetCharacter().AddStat(stats.SpellCrit, 35)
			agent.GetCharacter().AddStat(stats.SpellPower, 45)
		},
		4: func(agent core.Agent) {
			// Increases damage done by Lightning Bolt by 5%.
			// Implemented in lightning_bolt.go.
		},
	},
}

var NaturalAlignmentCrystalCooldownID = core.NewCooldownID()

func ApplyNaturalAlignmentCrystal(agent core.Agent) {
	const sp = 250
	const dur = time.Second * 20
	const cd = time.Minute * 5

	agent.GetCharacter().AddMajorCooldown(core.MajorCooldown{
		ActionID:         core.ActionID{ItemID: 19344},
		CooldownID:       NaturalAlignmentCrystalCooldownID,
		Cooldown:         cd,
		SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		SharedCooldown:   dur,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				character.SetCD(NaturalAlignmentCrystalCooldownID, sim.CurrentTime+cd)
				character.AddStat(stats.SpellPower, sp)
				character.Metrics.AddInstantCast(core.ActionID{SpellID: 23734})

				character.AddAura(sim, core.Aura{
					ID:      core.OffensiveTrinketActiveAuraID,
					Name:    "Natural Alignment Crystal",
					SpellID: 23734,
					Expires: sim.CurrentTime + dur,
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						cast.ManaCost *= 1.2
					},
					OnExpire: func(sim *core.Simulation) {
						character.AddStat(stats.SpellPower, -sp)
					},
				})
			}
		},
	})
}

// ActivateFathomBrooch adds an aura that has a chance on cast of nature spell
//  to restore 335 mana. 40s ICD
var FathomBroochOfTheTidewalkerAuraID = core.NewAuraID()

func ApplyFathomBroochOfTheTidewalker(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		icd := core.NewICD()
		const icdDur = time.Second * 40

		return core.Aura{
			ID:   FathomBroochOfTheTidewalkerAuraID,
			Name: "Fathom-Brooch of the Tidewalker",
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				if icd.IsOnCD(sim) {
					return
				}
				if cast.SpellSchool != stats.NatureSpellPower {
					return
				}
				if sim.RandomFloat("Fathom-Brooch of the Tidewalker") > 0.15 {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				character.AddMana(sim, 335, "Fathom-Brooch of the Tidewalker", false)
			},
		}
	})
}

var SkycallTotemAuraID = core.NewAuraID()
var EnergizedAuraID = core.NewAuraID()

func ApplySkycallTotem(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 101
		const dur = time.Second * 10

		return core.Aura{
			ID:      SkycallTotemAuraID,
			Name:    "Skycall Totem",
			Expires: core.NeverExpires,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				if cast.ActionID.SpellID != SpellIDLB12 || sim.RandomFloat("Skycall Totem") > 0.15 {
					return
				}
				character.AddAuraWithTemporaryStats(sim, EnergizedAuraID, 0, "Energized", stats.SpellHaste, hasteBonus, dur)
			},
		}
	})
}

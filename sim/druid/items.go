package druid

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemSet(ItemSetMalorne)
	core.AddItemSet(ItemSetNordrassil)
	core.AddItemSet(ItemSetThunderheart)
}

var Malorne2PcAuraID = core.NewAuraID()

var ItemSetMalorne = core.ItemSet{
	Name:  "Malorne Rainment",
	Items: map[int32]struct{}{29093: {}, 29094: {}, 29091: {}, 29092: {}, 29095: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID:   Malorne2PcAuraID,
					Name: "Malorne 2pc Bonus",
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						if sim.RandomFloat("malorne 2p") < 0.05 {
							spellCast.Character.AddStat(stats.Mana, 120)
						}
					},
				}
			})
		},
		4: func(agent core.Agent) {
			// Currently this is handled in druid.go (reducing CD of innervate)
		},
	},
}

var Nordrassil4pAuraID = core.NewAuraID()

var ItemSetNordrassil = core.ItemSet{
	Name:  "Nordrassil Regalia",
	Items: map[int32]struct{}{30231: {}, 30232: {}, 30233: {}, 30234: {}, 30235: {}},
	Bonuses: map[int32]core.ApplyEffect{
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID:   Nordrassil4pAuraID,
					Name: "Nordrassil 4p Bonus",
					OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						agent, ok := agent.(Agent)
						if !ok {
							panic("why is a non-druid using nordassil regalia")
						}
						druid := agent.GetDruid()
						if spellCast.ActionID.SpellID == SpellIDSF8 || spellCast.ActionID.SpellID == SpellIDSF6 {
							// Check if moonfire/insectswarm is ticking on the target.
							// TODO: in a raid simulator we need to be able to see which dots are ticking from other druids.
							if (druid.MoonfireSpell.DotInput.IsTicking(sim) && druid.MoonfireSpell.Target.Index == spellEffect.Target.Index) ||
								(druid.InsectSwarmSpell.DotInput.IsTicking(sim) && druid.InsectSwarmSpell.Target.Index == spellEffect.Target.Index) {
								spellEffect.DamageMultiplier *= 1.1
							}
						}
					},
				}
			})
		},
	},
}

var ItemSetThunderheart = core.ItemSet{
	Name:  "Thunderheart Regalia",
	Items: map[int32]struct{}{31043: {}, 31035: {}, 31040: {}, 31046: {}, 31049: {}, 34572: {}, 34446: {}, 34555: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// handled in moonfire.go in template construction
		},
		4: func(agent core.Agent) {
			// handled in starfire.go in template construction
		},
	},
}

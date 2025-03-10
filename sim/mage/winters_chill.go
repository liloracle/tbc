package mage

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDWintersChill int32 = 28595

// Winters Chill has a separate hit check from frostbolt, so it needs its own spell.
func (mage *Mage) newWintersChillTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:        "Winters Chill",
				SpellSchool: stats.FrostSpellPower,
				Character:   &mage.Character,
				ActionID: core.ActionID{
					SpellID: SpellIDWintersChill,
				},
				IgnoreCooldowns: true,
				IgnoreManaCost:  true,
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{},
		},
	}

	spell.SpellHitEffect.SpellEffect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance

	spell.SpellHitEffect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		// Don't overwrite the permanent version.
		if spellEffect.Target.RemainingAuraDuration(sim, core.WintersChillDebuffID) == core.NeverExpires {
			return
		}

		newNumStacks := core.MinInt32(5, spellEffect.Target.NumStacks(core.WintersChillDebuffID)+1)
		spellEffect.Target.ReplaceAura(sim, core.WintersChillAura(sim, newNumStacks))
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) procWintersChill(sim *core.Simulation, target *core.Target) {
	// Initialize cast from precomputed template.
	wintersChill := &mage.wintersChillSpell
	mage.wintersChillCastTemplate.Apply(wintersChill)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	wintersChill.Target = target
	wintersChill.Init(sim)
	wintersChill.Cast(sim)
}

var WintersChillAuraID = core.NewAuraID()

func (mage *Mage) applyWintersChill() {
	if mage.Talents.WintersChill == 0 {
		return
	}

	procChance := float64(mage.Talents.WintersChill) / 5.0

	mage.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: WintersChillAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellCast.SpellSchool == stats.FrostSpellPower && spellCast.ActionID.SpellID != SpellIDWintersChill {
					if procChance != 1.0 && sim.RandomFloat("Winters Chill") > procChance {
						return
					}

					mage.procWintersChill(sim, spellEffect.Target)
				}
			},
		}
	})
}

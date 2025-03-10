package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDScorch int32 = 27074

func (mage *Mage) newScorchTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:           "Scorch",
				CritMultiplier: 1.5 + 0.125*float64(mage.Talents.SpellPower),
				SpellSchool:    stats.FireSpellPower,
				Character:      &mage.Character,
				BaseManaCost:   180,
				ManaCost:       180,
				CastTime:       time.Millisecond * 1500,
				ActionID: core.ActionID{
					SpellID: SpellIDScorch,
				},
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: mage.spellDamageMultiplier,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage:    305,
				MaxBaseDamage:    361,
				SpellCoefficient: 1.5 / 3.5,
			},
		},
	}

	spell.SpellHitEffect.SpellEffect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	spell.SpellHitEffect.SpellEffect.BonusSpellCritRating += float64(mage.Talents.Incineration) * 2 * core.SpellCritRatingPerCritChance
	spell.SpellHitEffect.SpellEffect.BonusSpellCritRating += float64(mage.Talents.CriticalMass) * 2 * core.SpellCritRatingPerCritChance
	spell.SpellHitEffect.SpellEffect.BonusSpellCritRating += float64(mage.Talents.Pyromaniac) * 1 * core.SpellCritRatingPerCritChance
	spell.ManaCost -= spell.BaseManaCost * float64(mage.Talents.Pyromaniac) * 0.01
	spell.SpellHitEffect.SpellEffect.StaticDamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	if mage.Talents.ImprovedScorch > 0 {
		procChance := float64(mage.Talents.ImprovedScorch) / 3.0
		spell.SpellHitEffect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			// Don't overwrite the permanent version.
			if spellEffect.Target.RemainingAuraDuration(sim, core.ImprovedScorchDebuffID) == core.NeverExpires {
				return
			}

			if procChance != 1.0 || sim.RandomFloat("Improved Scorch") > procChance {
				return
			}

			newNumStacks := core.MinInt32(5, spellEffect.Target.NumStacks(core.ImprovedScorchDebuffID)+1)
			spellEffect.Target.ReplaceAura(sim, core.ImprovedScorchAura(sim, newNumStacks))
		}
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) NewScorch(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	scorch := &mage.scorchSpell
	mage.scorchCastTemplate.Apply(scorch)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	scorch.Target = target
	scorch.Init(sim)

	return scorch
}

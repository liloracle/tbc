package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDVampiricTouch int32 = 34917

var VampiricTouchDebuffID = core.NewDebuffID()

func (priest *Priest) newVampiricTouchTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		Name:           "Vampiric Touch",
		CritMultiplier: 1.5,
		SpellSchool:    stats.ShadowSpellPower,
		Character:      &priest.Character,
		BaseManaCost:   425,
		ManaCost:       425,
		CastTime:       time.Millisecond * 1500,
		ActionID: core.ActionID{
			SpellID: SpellIDVampiricTouch,
		},
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:        5,
			TickLength:           time.Second * 3,
			TickBaseDamage:       650 / 5,
			TickSpellCoefficient: 0.2,
			DebuffID:             VampiricTouchDebuffID,
			SpellID:              SpellIDVampiricTouch,
		},
	}

	priest.applyTalentsToShadowSpell(&baseCast, &effect)

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		SpellHitEffect: effect,
	})
}

func (priest *Priest) NewVampiricTouch(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := priest.VTSpellCasting

	priest.vtCastTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Target = target
	mf.Init(sim)

	return mf
}

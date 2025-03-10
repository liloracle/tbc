package enhancement

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/shaman"
)

func RegisterEnhancementShaman() {
	core.RegisterAgentFactory(
		proto.Player_EnhancementShaman{},
		func(character core.Character, options proto.Player) core.Agent {
			return NewEnhancementShaman(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_EnhancementShaman)
			if !ok {
				panic("Invalid spec value for Enhancement Shaman!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewEnhancementShaman(character core.Character, options proto.Player) *EnhancementShaman {
	enhOptions := options.GetEnhancementShaman()

	selfBuffs := shaman.SelfBuffs{}

	if enhOptions.Rotation.Totems != nil {
		selfBuffs.ManaSpring = enhOptions.Rotation.Totems.Water == proto.WaterTotem_ManaSpringTotem
		selfBuffs.EarthTotem = enhOptions.Rotation.Totems.Earth
		selfBuffs.AirTotem = enhOptions.Rotation.Totems.Air
		selfBuffs.NextTotemDropType[shaman.AirTotem] = int32(enhOptions.Rotation.Totems.Air)
		selfBuffs.FireTotem = enhOptions.Rotation.Totems.Fire
		selfBuffs.NextTotemDropType[shaman.FireTotem] = int32(enhOptions.Rotation.Totems.Fire)

		selfBuffs.TwistWindfury = enhOptions.Rotation.Totems.TwistWindfury
		if selfBuffs.TwistWindfury {
			selfBuffs.NextTotemDropType[shaman.AirTotem] = int32(proto.AirTotem_WindfuryTotem)
			selfBuffs.NextTotemDrops[shaman.AirTotem] = 0 // drop windfury immediately
		}

		selfBuffs.TwistFireNova = enhOptions.Rotation.Totems.TwistFireNova
		if selfBuffs.TwistFireNova {
			selfBuffs.NextTotemDropType[shaman.FireTotem] = int32(proto.FireTotem_FireNovaTotem) // start by dropping nova, then alternating.
		}
	}
	enh := &EnhancementShaman{
		Shaman: shaman.NewShaman(character, *enhOptions.Talents, selfBuffs),
	}
	// Enable Auto Attacks for this spec
	enh.EnableAutoAttacks()

	// TODO: de-sync dual weapons swing timers?

	// Modify auto attacks multiplier from weapon mastery.
	enh.AutoAttacks.DamageMultiplier *= 1 + 0.02*float64(enhOptions.Talents.WeaponMastery)
	shaman.ApplyWindfuryImbue(enh.Shaman, true, true)

	return enh
}

type EnhancementShaman struct {
	*shaman.Shaman
}

func (enh *EnhancementShaman) GetShaman() *shaman.Shaman {
	return enh.Shaman
}

func (enh *EnhancementShaman) Reset(sim *core.Simulation) {
	enh.Shaman.Reset(sim)
}

func (enh *EnhancementShaman) Act(sim *core.Simulation) time.Duration {
	// Redrop totems when needed.
	dropTime := enh.TryDropTotems(sim)
	if dropTime > 0 {
		return enh.AutoAttacks.TimeUntil(sim, nil, nil, dropTime)
	}

	success := true
	cost := 0.0
	const manaReserve = 1000 // if mana goes under 1000 we will need more soon. Pop shamanistic rage.
	if enh.CurrentMana() < manaReserve && enh.TryActivateShamanisticRage(sim) {
		// Just wait for GCD
		return enh.AutoAttacks.TimeUntil(sim, nil, nil, 0)
	} else if enh.GetRemainingCD(shaman.StormstrikeCD, sim.CurrentTime) == 0 {
		ss := enh.NewStormstrike(sim, sim.GetPrimaryTarget())
		cost = ss.Cost.Value
		if success = ss.Attack(sim); success {
			return enh.AutoAttacks.TimeUntil(sim, nil, ss, 0)
		}
	} else if enh.GetRemainingCD(shaman.ShockCooldownID, sim.CurrentTime) == 0 {
		shock := enh.NewEarthShock(sim, sim.GetPrimaryTarget())
		cost = shock.ManaCost
		if success = shock.Cast(sim); success {
			return enh.AutoAttacks.TimeUntil(sim, shock, nil, 0)
		}
	}
	if !success {
		regenTime := enh.TimeUntilManaRegen(cost)
		enh.Character.Metrics.MarkOOM(sim, &enh.Character, regenTime)
		return sim.CurrentTime + regenTime
	}

	// Do nothing, just swing axes until next CD available
	nextCD := enh.GetRemainingCD(shaman.StormstrikeCD, sim.CurrentTime)
	shockCD := enh.GetRemainingCD(shaman.ShockCooldownID, sim.CurrentTime)
	if shockCD < nextCD {
		nextCD = shockCD
	}
	return enh.AutoAttacks.TimeUntil(sim, nil, nil, sim.CurrentTime+nextCD)
}

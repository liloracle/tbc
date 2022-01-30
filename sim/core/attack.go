package core

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// OnBeforeMHSwing is called right before an auto attack fires
//  if false is returned the weapon swing is cancelled.
//  This allows for abilities that convert a white attack into yellow attack.
type OnBeforeMHSwing func(sim *Simulation) bool

// OnBeforeMelee is invoked once for each ability, even if there are multiple hits.
//  This should be used for any effects that adjust the stats / multipliers of the attack.
type OnBeforeMelee func(sim *Simulation, ability *ActiveMeleeAbility)

// OnBeforeMelee is invoked before the hit/dmg rolls are made.
//  This is invoked on both auto attacks and melee abilities.
//  This should be used for any effects that adjust the stats / multipliers of the attack.
type OnBeforeMeleeHit func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect)

// OnMeleeAttack is invoked on auto attacks and abilities.
//  This should be used for any on-hit procs.
type OnMeleeAttack func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect)

type ResourceCost struct {
	Type  stats.Stat // stats.Mana, stats.Energy, stats.Rage
	Value float64
}

type MeleeAbility struct {
	// ID for the action.
	ActionID

	// The character performing this action.
	Character *Character

	// If set, this action will start a cooldown using its cooldown ID.
	// Note that the GCD CD will be activated even if this is not set.
	Cooldown time.Duration

	// If set, this will be used as the GCD instead of the default value (1.5s).
	GCDCooldown time.Duration

	// If set, CD for this action and GCD CD will be ignored, and this action
	// will not set new values for those CDs either.
	IgnoreCooldowns bool

	// If set, this spell will have its resource cost ignored.
	IgnoreCost bool

	Cost ResourceCost

	CastTime time.Duration // most melee skills are instant... are there any with a cast time?

	// E.g. for nature spells, set to stats.NatureSpellPower.
	SpellSchool stats.Stat

	// How much to multiply damage by, if this cast crits.
	CritMultiplier float64

	// If true, will force the cast to crit (if it doesnt miss).
	GuaranteedCrit bool

	// Internal field only, used to prevent pool objects from being used by
	// multiple attacks simultaneously.
	objectInUse bool
}

type MeleeHitType byte

const (
	MeleeHitTypeMiss MeleeHitType = iota
	MeleeHitTypeDodge
	MeleeHitTypeParry
	MeleeHitTypeGlance
	MeleeHitTypeBlock
	MeleeHitTypeCrit
	MeleeHitTypeHit
)

type AbilityEffect struct {
	// Target of the spell.
	Target *Target

	// Bonus stats to be added to the attack.
	BonusHitRating        float64
	BonusAttackPower      float64
	BonusCritRating       float64
	BonusExpertiseRating  float64
	BonusArmorPenetration float64
	BonusWeaponDamage     float64

	IsWhiteHit bool

	// Causes the first roll for this hit to be copied from ActiveMeleeAbility.Effects[0].HitType.
	// This is only used by Shaman Stormstrike.
	ReuseMainHitRoll bool

	// Additional multiplier that is always applied.
	DamageMultiplier float64

	// applies fixed % increases to damage at cast time.
	//  Only use multipliers that don't change for the lifetime of the sim.
	//  This should probably only be mutated in a template and not changed in auras.
	StaticDamageMultiplier float64

	// Multiplier for all threat generated by this effect.
	ThreatMultiplier float64

	// Adds a fixed amount of threat to this spell, before multipliers.
	FlatThreatBonus float64

	// Ignores reduction from target armor.
	IgnoreArmor bool

	// The type of hit this was, i.e. miss/dodge/block/crit/hit.
	HitType MeleeHitType

	// The damage done by this effect.
	Damage float64
}

// Represents a generic weapon. Pets / unarmed / various other cases dont use
// actual weapon items so this is an abstraction of a Weapon.
type Weapon struct {
	BaseDamageMin float64
	BaseDamageMax float64
	SwingSpeed    float64
	SwingDuration time.Duration // Duration between 2 swings.
}

func newWeaponFromUnarmed() Weapon {
	// These numbers are probably wrong but nobody cares.
	return Weapon{
		BaseDamageMin: 0,
		BaseDamageMax: 0,
		SwingSpeed:    1,
		SwingDuration: time.Second,
	}
}

func newWeaponFromItem(item items.Item) Weapon {
	return Weapon{
		BaseDamageMin: item.WeaponDamageMin,
		BaseDamageMax: item.WeaponDamageMax,
		SwingSpeed:    item.SwingSpeed,
		SwingDuration: time.Duration(item.SwingSpeed * float64(time.Second)),
	}
}

// Returns weapon stats using the main hand equipped weapon.
func (character *Character) WeaponFromMainHand() Weapon {
	if weapon := character.GetMHWeapon(); weapon != nil {
		return newWeaponFromItem(*weapon)
	} else {
		return newWeaponFromUnarmed()
	}
}

// Returns weapon stats using the off hand equipped weapon.
func (character *Character) WeaponFromOffHand() Weapon {
	if weapon := character.GetOHWeapon(); weapon != nil {
		return newWeaponFromItem(*weapon)
	} else {
		return Weapon{}
	}
}

// Returns weapon stats using the off hand equipped weapon.
func (character *Character) WeaponFromRanged() Weapon {
	if weapon := character.GetRangedWeapon(); weapon != nil {
		return newWeaponFromItem(*weapon)
	} else {
		return Weapon{}
	}
}

func (weapon Weapon) BaseDamage(sim *Simulation) float64 {
	return weapon.BaseDamageMin + (weapon.BaseDamageMax-weapon.BaseDamageMin)*sim.RandomFloat("melee")
}

func (weapon Weapon) calculateSwingDamage(sim *Simulation, attackPower float64) float64 {
	return weapon.BaseDamage(sim) + (weapon.SwingSpeed*attackPower)/MeleeAttackRatingPerDamage
}

// If MainHand or Offhand is non-zero the associated ability will create a weapon swing.
type WeaponDamageInput struct {
	// Whether this input corresponds to the OH weapon.
	// It's important that this be 'IsOH' instead of 'IsMH' so that MH is the default.
	// This should be mutually exclusive with isRanged.
	IsOH bool

	// Whether this input corresponds to the ranged weapon.
	// This should be mutually exclusive with isOH.
	IsRanged bool

	DamageMultiplier float64 // Damage multiplier on weapon damage.
	FlatDamageBonus  float64 // Flat bonus added to swing.

	// If set, skips the normal calc for weapon damage and uses this function instead.
	CalculateDamage func(attackPower float64, bonusWeaponDamage float64) float64
}

type AbilityHitEffect struct {
	AbilityEffect
	DirectInput DirectDamageInput
	WeaponInput WeaponDamageInput
}

type ActiveMeleeAbility struct {
	MeleeAbility

	OnMeleeAttack OnMeleeAttack

	HitType     MeleeHitType // Hit roll result
	Hits        int32
	Misses      int32
	Crits       int32
	Dodges      int32
	Glances     int32
	Parries     int32
	Blocks      int32
	TotalDamage float64 // Damage done by this cast.
	TotalThreat float64 // Threat generated by this cast.

	// For abilities that only have 1 hit. Either this or Effects should be filled, not both.
	Effect AbilityHitEffect

	// For abilities that have more than 1 hit. Either this of Effect should be filled, not both.
	Effects []AbilityHitEffect
}

func (effect *AbilityEffect) Landed() bool {
	return effect.HitType != MeleeHitTypeMiss && effect.HitType != MeleeHitTypeDodge && effect.HitType != MeleeHitTypeParry
}

func (effect *AbilityEffect) String() string {
	if effect.HitType == MeleeHitTypeMiss {
		return "Miss"
	} else if effect.HitType == MeleeHitTypeDodge {
		return "Dodge"
	} else if effect.HitType == MeleeHitTypeParry {
		return "Parry"
	}

	var sb strings.Builder

	if effect.HitType == MeleeHitTypeHit {
		sb.WriteString("Hit")
	} else if effect.HitType == MeleeHitTypeCrit {
		sb.WriteString("Crit")
	} else if effect.HitType == MeleeHitTypeGlance {
		sb.WriteString("Glance")
	} else { // Block
		sb.WriteString("Block")
	}

	fmt.Fprintf(&sb, " for %0.3f damage", effect.Damage)

	return sb.String()
}

// Computes an attack result using the white-hit table formula (single roll).
func (ahe *AbilityHitEffect) WhiteHitTableResult(sim *Simulation, ability *ActiveMeleeAbility) MeleeHitType {
	// 1. Single roll -> Miss				Dodge	Parry	Glance	Block	Crit / Hit
	// 3 				8.0%(9.0% hit cap)	6.5%	14.0%	24% 	5%		-4.8%

	// TODO: many calculations in here can be cached. For now its just written out fully.
	//  Once everything is working we can start caching values.
	character := ability.Character

	roll := sim.RandomFloat("auto attack")
	level := float64(ahe.Target.Level)
	skill := 350.0 // assume max skill level for now.

	// Difference between attacker's waepon skill and target's defense skill.
	skillDifference := (level * 5) - skill

	// Miss
	missChance := 0.05 + skillDifference*0.002
	if ahe.IsWhiteHit && character.AutoAttacks.IsDualWielding {
		missChance += 0.19
	}
	hitSuppression := (skillDifference - 10) * 0.002
	hitBonus := ((character.stats[stats.MeleeHit] + ahe.BonusHitRating) / (MeleeHitRatingPerHitChance * 100)) - hitSuppression
	if hitBonus > 0 {
		missChance = math.Max(0, missChance-hitBonus)
	}

	chance := missChance
	if roll < chance {
		return MeleeHitTypeMiss
	}

	if !ahe.IsRanged() { // Ranged hits can't be dodged/glance, and are always 2-roll
		// Dodge
		dodge := 0.05 + skillDifference*0.001
		expertisePercentage := math.Min(math.Floor((character.stats[stats.Expertise]+ahe.BonusExpertiseRating)/(ExpertisePerQuarterPercentReduction))/400, dodge)
		chance += dodge - expertisePercentage
		if roll < chance {
			return MeleeHitTypeDodge
		}

		// Parry (if in front)
		// If the target is a mob and defense minus weapon skill is 11 or more:
		// ParryChance = 5% + (TargetLevel*5 - AttackerSkill) * 0.6%

		// If the target is a mob and defense minus weapon skill is 10 or less:
		// ParryChance = 5% + (TargetLevel*5 - AttackerSkill) * 0.1%

		// Block (if in front)
		// If the target is a mob:
		// BlockChance = MIN(5%, 5% + (TargetLevel*5 - AttackerSkill) * 0.1%)
		// If we actually implement blocks, ranged hits can be blocked.

		// Glance
		chance += math.Max(0.06+skillDifference*0.012, 0)
		if roll < chance {
			return MeleeHitTypeGlance
		}

		// Crit
		critChance := ((character.stats[stats.MeleeCrit] + ahe.BonusCritRating) / (MeleeCritRatingPerCritChance * 100)) - skillDifference*0.002 - 0.018
		chance += critChance
		if roll < chance {
			return MeleeHitTypeCrit
		}
	}

	return MeleeHitTypeHit
}

func (ahe *AbilityHitEffect) calculateDamage(sim *Simulation, ability *ActiveMeleeAbility) {
	character := ability.Character

	if ahe.AbilityEffect.ReuseMainHitRoll {
		ahe.HitType = ability.Effects[0].HitType
	} else {
		ahe.HitType = ahe.WhiteHitTableResult(sim, ability)
	}

	if !ahe.Landed() {
		ahe.Damage = 0
		return
	}

	var attackPower float64
	var bonusWeaponDamage float64
	if ahe.IsRanged() {
		attackPower = character.stats[stats.RangedAttackPower] + ahe.BonusAttackPower
		bonusWeaponDamage = character.PseudoStats.BonusRangedDamage + ahe.BonusWeaponDamage
	} else {
		attackPower = character.stats[stats.AttackPower] + ahe.BonusAttackPower
		bonusWeaponDamage = character.PseudoStats.BonusMeleeDamage + ahe.BonusWeaponDamage
	}

	dmg := 0.0
	if ahe.IsWeaponHit() {
		if ahe.WeaponInput.CalculateDamage != nil {
			dmg += ahe.WeaponInput.CalculateDamage(attackPower, bonusWeaponDamage)
		} else {
			// Bonus weapon damage applies after OH penalty: https://www.youtube.com/watch?v=bwCIU87hqTs
			if ahe.IsRanged() {
				dmg += character.AutoAttacks.Ranged.calculateSwingDamage(sim, attackPower) + bonusWeaponDamage
			} else if ahe.IsMH() {
				dmg += character.AutoAttacks.MH.calculateSwingDamage(sim, attackPower) + bonusWeaponDamage
			} else {
				dmg += character.AutoAttacks.OH.calculateSwingDamage(sim, attackPower)*0.5 + bonusWeaponDamage
			}
			dmg *= ahe.WeaponInput.DamageMultiplier
			dmg += ahe.WeaponInput.FlatDamageBonus
		}
	}

	// Add damage from DirectInput
	if ahe.DirectInput.MinBaseDamage != 0 {
		dmg += ahe.DirectInput.MinBaseDamage + (ahe.DirectInput.MaxBaseDamage-ahe.DirectInput.MinBaseDamage)*sim.RandomFloat("Melee Direct Input")
	}
	dmg += attackPower * ahe.DirectInput.SpellCoefficient
	dmg += ahe.DirectInput.FlatDamageBonus

	// If this is a yellow attack, need a 2nd roll to decide crit. Otherwise just use existing hit result.
	if !ahe.AbilityEffect.IsWhiteHit || ahe.IsRanged() {
		skill := 350.0
		level := float64(ahe.Target.Level)
		critChance := ((character.stats[stats.MeleeCrit] + ahe.BonusCritRating) / (MeleeCritRatingPerCritChance * 100)) + ((skill - (level * 5)) * 0.002) - 0.018

		roll := sim.RandomFloat("weapon swing")
		if roll < critChance {
			ahe.HitType = MeleeHitTypeCrit
		} else {
			ahe.HitType = MeleeHitTypeHit
		}
	}

	if ahe.HitType == MeleeHitTypeCrit {
		dmg *= ability.CritMultiplier
	} else if ahe.HitType == MeleeHitTypeGlance {
		dmg *= 0.75
	}

	// Apply armor reduction.
	if !ahe.IgnoreArmor {
		dmg *= 1 - ahe.Target.ArmorDamageReduction(character.stats[stats.ArmorPenetration]+ahe.BonusArmorPenetration)
		//if sim.Log != nil {
		//	character.Log(sim, "Target armor: %0.2f\n", ahe.Target.currentArmor)
		//}
	}

	// Apply all other effect multipliers.
	dmg *= ahe.DamageMultiplier * ahe.StaticDamageMultiplier

	ahe.Damage = dmg
}

// Returns whether this hit effect is associated with one of the character's
// weapons. This check is necessary to decide whether certains effects are eligible.
func (ahe *AbilityHitEffect) IsWeaponHit() bool {
	return ahe.WeaponInput.DamageMultiplier != 0 || ahe.WeaponInput.CalculateDamage != nil
}

// Returns whether this hit effect is associated with the main-hand weapon.
func (ahe *AbilityHitEffect) IsMH() bool {
	return !ahe.WeaponInput.IsOH && !ahe.WeaponInput.IsRanged
}

// Returns whether this hit effect is associated with the off-hand weapon.
func (ahe *AbilityHitEffect) IsOH() bool {
	return ahe.WeaponInput.IsOH
}

// Returns whether this hit effect is associated with either melee weapon.
func (ahe *AbilityHitEffect) IsMelee() bool {
	return !ahe.WeaponInput.IsRanged
}

// Returns whether this hit effect is associated with the ranged weapon.
func (ahe *AbilityHitEffect) IsRanged() bool {
	return ahe.WeaponInput.IsRanged
}

// Returns whether this hit effect matches the hand in which a weapon is equipped.
func (ahe *AbilityHitEffect) IsEquippedHand(mh bool, oh bool) bool {
	return (ahe.IsMH() && mh) || (ahe.IsOH() && oh)
}

// It appears that TBC does not do hasted GCD for abilities.
//  Leaving this option here in case we want it in the future.
const EnableAbilityHaste = false

func (ability *ActiveMeleeAbility) CalculatedGCD(char *Character) time.Duration {
	baseGCD := GCDDefault
	if !EnableAbilityHaste {
		// TODO: Other classes have different GCD defaults for abilites.
		//  Probably want to just have all abilities specify the GCD explicitly?
		return baseGCD
	}
	if ability.GCDCooldown != 0 {
		baseGCD = ability.GCDCooldown
	}
	return MaxDuration(GCDMin, time.Duration(float64(baseGCD)/char.SwingSpeed()))
}

// Attack will perform the attack
//  Returns false if unable to attack (due to resource lacking)
func (ability *ActiveMeleeAbility) Attack(sim *Simulation) bool {
	if !ability.IgnoreCooldowns && ability.Character.GetRemainingCD(GCDCooldownID, sim.CurrentTime) > 0 {
		log.Fatalf("Ability used while on GCD\n-------\nAbility %s: %#v\n", ability.ActionID, ability)
	}
	if ability.MeleeAbility.Cost.Type != 0 {
		if ability.MeleeAbility.Cost.Type == stats.Mana {
			if ability.Character.CurrentMana() < ability.MeleeAbility.Cost.Value {
				return false
			}
			ability.Character.SpendMana(sim, ability.MeleeAbility.Cost.Value, ability.MeleeAbility.ActionID)
		} else if ability.MeleeAbility.Cost.Type == stats.Rage {
			if ability.Character.CurrentRage() < ability.MeleeAbility.Cost.Value {
				return false
			}
			ability.Character.SpendRage(sim, ability.MeleeAbility.Cost.Value, ability.MeleeAbility.ActionID)
		} else {
			if ability.Character.CurrentEnergy() < ability.MeleeAbility.Cost.Value {
				return false
			}
			ability.Character.SpendEnergy(sim, ability.MeleeAbility.Cost.Value, ability.MeleeAbility.ActionID)
		}
	}

	ability.Character.OnBeforeMelee(sim, ability)

	if len(ability.Effects) == 0 {
		ability.Effect.performAttack(sim, ability)
	} else {
		for i, _ := range ability.Effects {
			ahe := &ability.Effects[i]
			ahe.performAttack(sim, ability)
		}
	}

	if !ability.IgnoreCooldowns {
		gcdCD := MaxDuration(ability.CalculatedGCD(ability.Character), ability.CastTime)
		ability.Character.SetGCDTimer(sim, sim.CurrentTime+gcdCD)

		if ability.ActionID.CooldownID != 0 {
			ability.Character.SetCD(ability.ActionID.CooldownID, sim.CurrentTime+ability.Cooldown)
		}
	}
	ability.Character.Metrics.AddMeleeAbility(ability)
	return true
}

func (ahe *AbilityHitEffect) performAttack(sim *Simulation, ability *ActiveMeleeAbility) {
	ability.Character.OnBeforeMeleeHit(sim, ability, ahe)
	ahe.Target.OnBeforeMeleeHit(sim, ability, ahe)

	ahe.calculateDamage(sim, ability)

	if ahe.HitType == MeleeHitTypeMiss {
		ability.Misses++
	} else if ahe.HitType == MeleeHitTypeDodge {
		ability.Dodges++
	} else if ahe.HitType == MeleeHitTypeGlance {
		ability.Glances++
	} else if ahe.HitType == MeleeHitTypeCrit {
		ability.Crits++
	} else if ahe.HitType == MeleeHitTypeHit {
		ability.Hits++
	} else if ahe.HitType == MeleeHitTypeParry {
		ability.Parries++
	} else if ahe.HitType == MeleeHitTypeBlock {
		ability.Blocks++
	}
	ability.TotalDamage += ahe.Damage
	ability.TotalThreat += (ahe.Damage + ahe.FlatThreatBonus) * ahe.ThreatMultiplier * ability.Character.PseudoStats.ThreatMultiplier

	if sim.Log != nil {
		ability.Character.Log(sim, "%s %s", ability.ActionID, ahe)
	}

	ability.Character.OnMeleeAttack(sim, ability, ahe)
	ahe.Target.OnMeleeAttack(sim, ability, ahe)
	if ability.OnMeleeAttack != nil {
		ability.OnMeleeAttack(sim, ability, ahe)
	}
}

type AutoAttacks struct {
	// initialized
	agent     Agent
	character *Character
	MH        Weapon
	OH        Weapon
	Ranged    Weapon

	IsDualWielding bool

	// If true, core engine will handle calling SwingMelee(). Set to false to manually manage
	// swings, for example for hunter melee weaving.
	AutoSwingMelee bool

	// If true, core engine will handle calling SwingRanged(). Unless you're a hunter, don't
	// use this.
	AutoSwingRanged bool

	// Set this to true to use the OH delay macro, mostly used by enhance shamans.
	// This will intentionally delay OH swings to that they always fall within the
	// 0.5s window following a MH swing.
	DelayOHSwings bool

	MainhandSwingAt time.Duration
	OffhandSwingAt  time.Duration
	RangedSwingAt   time.Duration

	ActiveMeleeAbility // Parameters for auto attacks.

	RangedAuto ActiveMeleeAbility // Parameters for ranged auto attacks.

	OnBeforeMHSwing OnBeforeMHSwing

	// The time at which the last MH swing occurred.
	previousMHSwingAt time.Duration

	// PendingAction which handles auto attacks.
	autoSwingAction *PendingAction
}

// Options for initializing auto attacks.
type AutoAttackOptions struct {
	MainHand        Weapon
	OffHand         Weapon
	Ranged          Weapon
	AutoSwingMelee  bool // If true, core engine will handle calling SwingMelee() for you.
	AutoSwingRanged bool // If true, core engine will handle calling SwingRanged() for you.
	DelayOHSwings   bool
}

func (character *Character) EnableAutoAttacks(agent Agent, options AutoAttackOptions) {
	aa := AutoAttacks{
		agent:           agent,
		character:       character,
		MH:              options.MainHand,
		OH:              options.OffHand,
		Ranged:          options.Ranged,
		AutoSwingMelee:  options.AutoSwingMelee,
		AutoSwingRanged: options.AutoSwingRanged,
		DelayOHSwings:   options.DelayOHSwings,
		ActiveMeleeAbility: ActiveMeleeAbility{
			MeleeAbility: MeleeAbility{
				ActionID:        ActionID{OtherID: proto.OtherAction_OtherActionAttack},
				Character:       character,
				SpellSchool:     stats.AttackPower,
				CritMultiplier:  2,
				IgnoreCooldowns: true,
				IgnoreCost:      true,
			},
			Effect: AbilityHitEffect{
				AbilityEffect: AbilityEffect{
					IsWhiteHit:             true,
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
					ThreatMultiplier:       1,
				},
				WeaponInput: WeaponDamageInput{
					DamageMultiplier: 1,
				},
			},
		},
		RangedAuto: ActiveMeleeAbility{
			MeleeAbility: MeleeAbility{
				ActionID:        ActionID{OtherID: proto.OtherAction_OtherActionShoot},
				Character:       character,
				SpellSchool:     stats.AttackPower,
				CritMultiplier:  2,
				IgnoreCooldowns: true,
				IgnoreCost:      true,
			},
			Effect: AbilityHitEffect{
				AbilityEffect: AbilityEffect{
					IsWhiteHit:             true,
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
					ThreatMultiplier:       1,
				},
				WeaponInput: WeaponDamageInput{
					IsRanged:         true,
					DamageMultiplier: 1,
				},
			},
		},
	}

	if options.AutoSwingMelee && options.AutoSwingRanged {
		panic("Cant auto swing both melee and ranged!")
	}

	aa.IsDualWielding = aa.MH.SwingSpeed != 0 && aa.OH.SwingSpeed != 0

	character.AutoAttacks = aa
}

func (aa *AutoAttacks) IsEnabled() bool {
	return aa.MH.SwingSpeed != 0
}

// Empty handler so Agents don't have to provide one if they have no logic to add.
func (character *Character) OnAutoAttack(sim *Simulation) {}

func (aa *AutoAttacks) reset(sim *Simulation) {
	if !aa.IsEnabled() {
		return
	}

	aa.MainhandSwingAt = 0
	aa.OffhandSwingAt = 0
	aa.RangedSwingAt = 0

	// Set a fake value for previousMHSwing so that offhand swing delay works
	// properly at the start.
	aa.previousMHSwingAt = time.Second * -1

	// Apply random delay of 0 - 50% swing time, to one of the weapons.
	delay := time.Duration(sim.RandomFloat("SwingResetDelay") * float64(aa.MH.SwingDuration/2))
	isMHDelay := sim.RandomFloat("SwingResetWeapon") < 0.5

	if isMHDelay {
		aa.MainhandSwingAt = delay
	} else if aa.IsDualWielding {
		aa.OffhandSwingAt = delay
	}

	aa.autoSwingAction = nil
	aa.resetAutoSwing(sim)
}

func (aa *AutoAttacks) resetAutoSwing(sim *Simulation) {
	if !aa.AutoSwingMelee && !aa.AutoSwingRanged {
		return
	}

	if aa.autoSwingAction != nil {
		aa.autoSwingAction.Cancel(sim)
	}

	pa := &PendingAction{
		Priority:     ActionPriorityAuto,
		NextActionAt: 0, // First auto is always at 0
	}
	pa.OnAction = func(sim *Simulation) {
		if aa.AutoSwingMelee {
			aa.SwingMelee(sim, sim.GetPrimaryTarget())
			pa.NextActionAt = aa.NextAttackAt()
		} else {
			aa.SwingRanged(sim, sim.GetPrimaryTarget())
			pa.NextActionAt = aa.RangedSwingAt
		}

		// Cancelled means we made a new one because of a swing speed change.
		if !pa.cancelled {
			sim.AddPendingAction(pa)
		}
	}
	aa.autoSwingAction = pa
	sim.AddPendingAction(pa)
}

func (aa *AutoAttacks) MainhandSwingSpeed() time.Duration {
	return time.Duration(float64(aa.MH.SwingDuration) / aa.character.SwingSpeed())
}

func (aa *AutoAttacks) OffhandSwingSpeed() time.Duration {
	return time.Duration(float64(aa.OH.SwingDuration) / aa.character.SwingSpeed())
}

func (aa *AutoAttacks) RangedSwingSpeed() time.Duration {
	return time.Duration(float64(aa.Ranged.SwingDuration) / aa.character.RangedSwingSpeed())
}

// SwingMelee will check any swing timers if they are up, and if so, swing!
func (aa *AutoAttacks) SwingMelee(sim *Simulation, target *Target) {
	aa.TrySwingMH(sim, target)
	aa.TrySwingOH(sim, target)
}

func (aa *AutoAttacks) SwingRanged(sim *Simulation, target *Target) {
	aa.TrySwingRanged(sim, target)
}

// Performs an autoattack using the main hand weapon, if the MH CD is ready.
func (aa *AutoAttacks) TrySwingMH(sim *Simulation, target *Target) {
	if aa.MainhandSwingAt > sim.CurrentTime {
		return
	}

	if aa.OnBeforeMHSwing != nil {
		// Allow MH swing to be overridden for abilities like Heroic Strike.
		doSwing := aa.OnBeforeMHSwing(sim)
		if !doSwing {
			return
		}
	}

	ama := aa.ActiveMeleeAbility
	ama.ActionID.Tag = 1
	ama.Effect.Target = target
	ama.Effect.WeaponInput.IsOH = false
	ama.Attack(sim)
	aa.MainhandSwingAt = sim.CurrentTime + aa.MainhandSwingSpeed()
	aa.previousMHSwingAt = sim.CurrentTime
	aa.agent.OnAutoAttack(sim)
}

// Performs an autoattack using the main hand weapon, if the OH CD is ready.
func (aa *AutoAttacks) TrySwingOH(sim *Simulation, target *Target) {
	if !aa.IsDualWielding || aa.OffhandSwingAt > sim.CurrentTime {
		return
	}

	if aa.DelayOHSwings && (sim.CurrentTime-aa.previousMHSwingAt) > time.Millisecond*500 {
		// Delay the OH swing for later, so it follows the MH swing.
		aa.OffhandSwingAt = aa.MainhandSwingAt + time.Millisecond*100
		if sim.Log != nil {
			aa.character.Log(sim, "Delaying OH swing by %s", aa.OffhandSwingAt-sim.CurrentTime)
		}
		return
	}

	ama := aa.ActiveMeleeAbility
	ama.ActionID.Tag = 2
	ama.Effect.Target = target
	ama.Effect.WeaponInput.IsOH = true
	ama.Attack(sim)
	aa.OffhandSwingAt = sim.CurrentTime + aa.OffhandSwingSpeed()
	aa.agent.OnAutoAttack(sim)
}

// Performs an autoattack using the ranged weapon, if the ranged CD is ready.
func (aa *AutoAttacks) TrySwingRanged(sim *Simulation, target *Target) {
	if aa.RangedSwingAt > sim.CurrentTime {
		return
	}

	ama := aa.RangedAuto
	ama.Effect.Target = target
	ama.Attack(sim)
	aa.RangedSwingAt = sim.CurrentTime + aa.RangedSwingSpeed()
}

func (aa *AutoAttacks) ModifySwingTime(sim *Simulation, amount float64) {
	if !aa.IsEnabled() {
		return
	}

	mhSwingTime := aa.MainhandSwingAt - sim.CurrentTime
	if mhSwingTime > 1 { // If its 1 we end up rounding down to 0 and causing a panic.
		aa.MainhandSwingAt = sim.CurrentTime + time.Duration(float64(mhSwingTime)/amount)
	}

	if aa.OH.SwingSpeed != 0 {
		ohSwingTime := aa.OffhandSwingAt - sim.CurrentTime
		if ohSwingTime > 1 {
			newTime := time.Duration(float64(ohSwingTime) / amount)
			if newTime > 0 {
				aa.OffhandSwingAt = sim.CurrentTime + newTime
			}
		}
	}

	if aa.Ranged.SwingSpeed != 0 {
		rangedSwingTime := aa.RangedSwingAt - sim.CurrentTime
		if rangedSwingTime > 1 {
			newTime := time.Duration(float64(rangedSwingTime) / amount)
			if newTime > 0 {
				aa.RangedSwingAt = sim.CurrentTime + newTime
			}
		}
	}

	aa.resetAutoSwing(sim)
}

// Returns the time at which the next attack will occur.
func (aa *AutoAttacks) NextAttackAt() time.Duration {
	nextAttack := aa.MainhandSwingAt
	if aa.OH.SwingSpeed != 0 {
		nextAttack = MinDuration(nextAttack, aa.OffhandSwingAt)
	}
	return nextAttack
}

// Returns the time at which the next event will occur, considering both autos and the gcd.
func (aa *AutoAttacks) NextEventAt(sim *Simulation) time.Duration {
	if aa.NextAttackAt() == sim.CurrentTime {
		panic(fmt.Sprintf("Returned 0 from next attack at %s, mh: %s, oh: %s", sim.CurrentTime, aa.MainhandSwingAt, aa.OffhandSwingAt))
	}
	return MinDuration(
		sim.CurrentTime+aa.Character.GetRemainingCD(GCDCooldownID, sim.CurrentTime),
		aa.NextAttackAt())
}

type PPMManager struct {
	mhProcChance     float64
	ohProcChance     float64
	rangedProcChance float64
}

// For manually overriding proc chance.
func (ppmm *PPMManager) SetProcChance(isMH bool, newChance float64) {
	if isMH {
		ppmm.mhProcChance = newChance
	} else {
		ppmm.ohProcChance = newChance
	}
}
func (ppmm *PPMManager) SetRangedChance(newChance float64) {
	ppmm.rangedProcChance = newChance
}

// Returns whether the effect procced.
func (ppmm *PPMManager) Proc(sim *Simulation, isMH bool, isRanged bool, label string) bool {
	if isMH {
		return ppmm.ProcMH(sim, label)
	} else if !isRanged {
		return ppmm.ProcOH(sim, label)
	} else {
		return ppmm.ProcRanged(sim, label)
	}
}

// Returns whether the effect procced, assuming MH.
func (ppmm *PPMManager) ProcMH(sim *Simulation, label string) bool {
	return ppmm.mhProcChance > 0 && sim.RandomFloat(label) < ppmm.mhProcChance
}

// Returns whether the effect procced, assuming OH.
func (ppmm *PPMManager) ProcOH(sim *Simulation, label string) bool {
	return ppmm.ohProcChance > 0 && sim.RandomFloat(label) < ppmm.ohProcChance
}

// Returns whether the effect procced, assuming Ranged.
func (ppmm *PPMManager) ProcRanged(sim *Simulation, label string) bool {
	return ppmm.rangedProcChance > 0 && sim.RandomFloat(label) < ppmm.rangedProcChance
}

// PPMToChance converts a character proc-per-minute into mh/oh proc chances
func (aa *AutoAttacks) NewPPMManager(ppm float64) PPMManager {
	if aa.MH.SwingSpeed == 0 {
		// Means this character didn't enable autoattacks.
		return PPMManager{
			mhProcChance:     0,
			ohProcChance:     0,
			rangedProcChance: 0,
		}
	}

	return PPMManager{
		mhProcChance:     (aa.MH.SwingSpeed * ppm) / 60.0,
		ohProcChance:     (aa.OH.SwingSpeed * ppm) / 60.0,
		rangedProcChance: (aa.Ranged.SwingSpeed * ppm) / 60.0,
	}
}

type MeleeAbilityTemplate struct {
	template ActiveMeleeAbility
	effects  []AbilityHitEffect
}

func (template *MeleeAbilityTemplate) Apply(newAction *ActiveMeleeAbility) {
	if newAction.objectInUse {
		panic(fmt.Sprintf("Melee ability (%s) already in use", newAction.ActionID))
	}
	*newAction = template.template
	newAction.Effects = template.effects
	copy(newAction.Effects, template.template.Effects)
}

// Takes in a cast template and returns a template, so you don't need to keep track of which things to allocate yourself.
func NewMeleeAbilityTemplate(abilityTemplate ActiveMeleeAbility) MeleeAbilityTemplate {
	if len(abilityTemplate.Effects) > 0 && abilityTemplate.Effect.StaticDamageMultiplier != 0 {
		panic("Cannot use both Effect and Effects, pick one!")
	}

	return MeleeAbilityTemplate{
		template: abilityTemplate,
		effects:  make([]AbilityHitEffect, len(abilityTemplate.Effects)),
	}
}

package core

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Applies buffs that affect individual players.
func applyBuffEffects(agent Agent, raidBuffs proto.RaidBuffs, partyBuffs proto.PartyBuffs, individualBuffs proto.IndividualBuffs) {
	character := agent.GetCharacter()

	if raidBuffs.ArcaneBrilliance {
		character.AddStats(stats.Stats{
			stats.Intellect: 40,
		})
	}

	gotwAmount := GetTristateValueFloat(raidBuffs.GiftOfTheWild, 14.0, 14.0*1.35)
	character.AddStats(stats.Stats{
		stats.Stamina:   gotwAmount,
		stats.Agility:   gotwAmount,
		stats.Strength:  gotwAmount,
		stats.Intellect: gotwAmount,
		stats.Spirit:    gotwAmount,
	})

	character.AddStats(stats.Stats{
		stats.SpellCrit: GetTristateValueFloat(partyBuffs.MoonkinAura, 5*SpellCritRatingPerCritChance, 5*SpellCritRatingPerCritChance+20),
	})
	character.AddStats(stats.Stats{
		stats.MeleeCrit: GetTristateValueFloat(partyBuffs.LeaderOfThePack, 5*MeleeCritRatingPerCritChance, 5*MeleeCritRatingPerCritChance+20),
	})

	if partyBuffs.TrueshotAura {
		character.AddStats(stats.Stats{
			stats.AttackPower:       125,
			stats.RangedAttackPower: 125,
		})
	}

	if partyBuffs.FerociousInspiration > 0 {
		multiplier := math.Pow(1.03, float64(partyBuffs.FerociousInspiration))
		character.PseudoStats.DamageDealtMultiplier *= multiplier
	}

	if partyBuffs.DraeneiRacialMelee {
		character.AddStats(stats.Stats{
			stats.MeleeHit: 1 * MeleeHitRatingPerHitChance,
		})
	}

	if partyBuffs.DraeneiRacialCaster {
		character.AddStats(stats.Stats{
			stats.SpellHit: 1 * SpellHitRatingPerHitChance,
		})
	}
	character.AddStats(stats.Stats{
		stats.Spirit: GetTristateValueFloat(raidBuffs.DivineSpirit, 50.0, 50.0),
	})
	if raidBuffs.DivineSpirit == proto.TristateEffect_TristateEffectImproved {
		character.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Spirit,
			ModifiedStat: stats.SpellPower,
			Modifier: func(spirit float64, spellPower float64) float64 {
				return spellPower + spirit*0.1
			},
		})
	}

	if individualBuffs.ShadowPriestDps > 0 {
		character.AddStats(stats.Stats{
			stats.MP5: float64(individualBuffs.ShadowPriestDps) * 0.25,
		})
	}

	character.AddStats(stats.Stats{
		stats.MP5: GetTristateValueFloat(individualBuffs.BlessingOfWisdom, 42.0, 50.0),
	})

	character.AddStats(stats.Stats{
		stats.AttackPower:       GetTristateValueFloat(individualBuffs.BlessingOfMight, 220, 264),
		stats.RangedAttackPower: GetTristateValueFloat(individualBuffs.BlessingOfMight, 220, 264),
	})

	if individualBuffs.BlessingOfKings {
		bokStats := [5]stats.Stat{
			stats.Agility,
			stats.Strength,
			stats.Stamina,
			stats.Intellect,
			stats.Spirit,
		}

		for _, stat := range bokStats {
			character.AddStatDependency(stats.StatDependency{
				SourceStat:   stat,
				ModifiedStat: stat,
				Modifier: func(curValue float64, _ float64) float64 {
					return curValue * 1.1
				},
			})
		}
	}

	if individualBuffs.BlessingOfSalvation {
		character.PseudoStats.ThreatMultiplier *= 0.7
	}

	if partyBuffs.SanctityAura == proto.TristateEffect_TristateEffectImproved {
		character.AddPermanentAura(func(*Simulation) *Aura { return SanctityAura(character, 2) })
	} else if partyBuffs.SanctityAura == proto.TristateEffect_TristateEffectRegular {
		character.AddPermanentAura(func(*Simulation) *Aura { return SanctityAura(character, 0) })
	}

	if partyBuffs.BattleShout != proto.TristateEffect_TristateEffectMissing {
		talentMultiplier := GetTristateValueFloat(partyBuffs.BattleShout, 1, 1.25)

		character.AddStats(stats.Stats{
			stats.AttackPower: 306 * talentMultiplier,
		})
		if partyBuffs.BsSolarianSapphire {
			partyBuffs.SnapshotBsSolarianSapphire = false
			character.AddStats(stats.Stats{
				stats.AttackPower: 70 * talentMultiplier,
			})
		}

		snapshotAp := 0.0
		if partyBuffs.SnapshotBsSolarianSapphire {
			snapshotAp += 70 * talentMultiplier
		}
		if partyBuffs.SnapshotBsT2 {
			snapshotAp += 30 * talentMultiplier
		}
		if snapshotAp > 0 {
			character.AddPermanentAuraWithOptions(PermanentAura{
				AuraFactory:     func(*Simulation) *Aura { return SnapshotBattleShoutAura(character, snapshotAp) },
				RespectDuration: true,
			})
		}
	}
	character.AddStats(stats.Stats{
		stats.Health: GetTristateValueFloat(partyBuffs.CommandingShout, 1080, 1080*1.25),
	})

	if partyBuffs.TotemOfWrath > 0 {
		character.AddStats(stats.Stats{
			stats.SpellCrit: 3 * SpellCritRatingPerCritChance * float64(partyBuffs.TotemOfWrath),
			stats.SpellHit:  3 * SpellHitRatingPerHitChance * float64(partyBuffs.TotemOfWrath),
		})
	}
	character.AddStats(stats.Stats{
		stats.SpellPower: GetTristateValueFloat(partyBuffs.WrathOfAirTotem, 101.0, 121.0),
	})
	if partyBuffs.WrathOfAirTotem == proto.TristateEffect_TristateEffectRegular && partyBuffs.SnapshotImprovedWrathOfAirTotem {
		character.AddPermanentAuraWithOptions(PermanentAura{
			AuraFactory:     func(*Simulation) *Aura { return SnapshotImprovedWrathOfAirTotemAura(character) },
			RespectDuration: true,
		})
	}
	character.AddStats(stats.Stats{
		stats.Agility: GetTristateValueFloat(partyBuffs.GraceOfAirTotem, 77.0, 88.55),
	})
	switch partyBuffs.StrengthOfEarthTotem {
	case proto.StrengthOfEarthType_Basic:
		character.AddStat(stats.Strength, 86)
	case proto.StrengthOfEarthType_CycloneBonus:
		character.AddStat(stats.Strength, 98)
	case proto.StrengthOfEarthType_EnhancingTotems:
		character.AddStat(stats.Strength, 98.9)
	case proto.StrengthOfEarthType_EnhancingAndCyclone:
		character.AddStat(stats.Strength, 110.9)
	}
	if (partyBuffs.StrengthOfEarthTotem == proto.StrengthOfEarthType_Basic || partyBuffs.StrengthOfEarthTotem == proto.StrengthOfEarthType_EnhancingTotems) && partyBuffs.SnapshotImprovedStrengthOfEarthTotem {
		character.AddPermanentAuraWithOptions(PermanentAura{
			AuraFactory:     func(*Simulation) *Aura { return SnapshotImprovedStrengthOfEarthTotemAura(character) },
			RespectDuration: true,
		})
	}
	character.AddStats(stats.Stats{
		stats.MP5: GetTristateValueFloat(partyBuffs.ManaSpringTotem, 50, 62.5),
	})
	if partyBuffs.WindfuryTotemRank > 0 && IsEligibleForWindfuryTotem(character) {
		character.HasWFTotem = true
		character.AddPermanentAura(func(*Simulation) *Aura {
			return WindfuryTotemAura(character, partyBuffs.WindfuryTotemRank, partyBuffs.WindfuryTotemIwt)
		})
	}
	if partyBuffs.TranquilAirTotem {
		character.PseudoStats.ThreatMultiplier *= 0.8
	}

	if individualBuffs.UnleashedRage {
		character.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.AttackPower,
			ModifiedStat: stats.AttackPower,
			Modifier: func(ap float64, _ float64) float64 {
				return ap * 1.1
			},
		})
	}

	registerBloodlustCD(agent, partyBuffs.Bloodlust)
	registerPowerInfusionCD(agent, individualBuffs.PowerInfusions)
	registerManaTideTotemCD(agent, partyBuffs.ManaTideTotems)
	registerInnervateCD(agent, individualBuffs.Innervates)

	character.AddStats(stats.Stats{
		stats.SpellCrit: 28 * float64(partyBuffs.AtieshMage),
	})
	character.AddStats(stats.Stats{
		stats.SpellPower:   33 * float64(partyBuffs.AtieshWarlock),
		stats.HealingPower: 33 * float64(partyBuffs.AtieshWarlock),
	})

	if partyBuffs.BraidedEterniumChain {
		character.AddStats(stats.Stats{stats.MeleeCrit: 28})
	}
	if partyBuffs.EyeOfTheNight {
		character.AddStats(stats.Stats{stats.SpellPower: 34})
	}
	if partyBuffs.JadePendantOfBlasting {
		character.AddStats(stats.Stats{stats.SpellPower: 15})
	}
	if partyBuffs.ChainOfTheTwilightOwl {
		character.AddStats(stats.Stats{stats.SpellCrit: 2 * SpellCritRatingPerCritChance})
	}
}

// Applies buffs to pets.
func applyPetBuffEffects(petAgent PetAgent, raidBuffs proto.RaidBuffs, partyBuffs proto.PartyBuffs, individualBuffs proto.IndividualBuffs) {
	// Summoned pets, like Mage Water Elemental, aren't around to receive raid buffs.
	if !petAgent.GetPet().initialEnabled {
		return
	}

	// We need to modify the buffs a bit because some things are applied to pets by
	// the owner during combat (Bloodlust) or don't make sense for a pet.
	partyBuffs.Bloodlust = 0
	partyBuffs.Drums = proto.Drums_DrumsUnknown
	individualBuffs.Innervates = 0
	individualBuffs.PowerInfusions = 0

	// For some reason pets don't benefit from buffs that are ratings, e.g. crit rating or haste rating.
	partyBuffs.LeaderOfThePack = MinTristate(partyBuffs.LeaderOfThePack, proto.TristateEffect_TristateEffectRegular)
	partyBuffs.BraidedEterniumChain = false

	applyBuffEffects(petAgent, raidBuffs, partyBuffs, individualBuffs)
}

func SnapshotImprovedStrengthOfEarthTotemAura(character *Character) *Aura {
	return character.NewTemporaryStatsAura("Strength of Earth Totem Snapshot", ActionID{SpellID: 37223}, stats.Stats{stats.Strength: 12}, time.Second*110)
}

func SnapshotImprovedWrathOfAirTotemAura(character *Character) *Aura {
	return character.NewTemporaryStatsAura("Wrath of Air Totem Snapshot", ActionID{SpellID: 37212}, stats.Stats{stats.SpellPower: 20}, time.Second*110)
}

func SnapshotBattleShoutAura(character *Character, snapshotAp float64) *Aura {
	return character.NewTemporaryStatsAura("Battle Shout Snapshot", ActionID{SpellID: 2048, Tag: 1}, stats.Stats{stats.AttackPower: snapshotAp}, time.Second*110)
}

func SanctityAura(character *Character, level float64) *Aura {
	return character.GetOrRegisterAura(&Aura{
		Label:    "Sanctity Aura",
		ActionID: ActionID{SpellID: 31870},
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.HolyDamageDealtMultiplier *= 1.1
			aura.Unit.PseudoStats.DamageDealtMultiplier *= 1 + 0.01*level
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.HolyDamageDealtMultiplier /= 1.1
			aura.Unit.PseudoStats.DamageDealtMultiplier /= 1 + 0.01*level
		},
	})
}

var (
	WindfuryTotemSpellRanks = []int32{
		8512,
		10613,
		10614,
		25585,
		25587,
	}

	windfuryBuffSpellRanks = []int32{
		8516,
		10608,
		10610,
		25583,
		25584,
	}

	windfuryAPBonuses = []float64{
		122,
		229,
		315,
		375,
		445,
	}
)

func IsEligibleForWindfuryTotem(character *Character) bool {
	return character.AutoAttacks.IsEnabled() &&
		character.HasMHWeapon() &&
		!character.HasMHWeaponImbue
}

func WindfuryTotemAura(character *Character, rank int32, iwtTalentPoints int32) *Aura {
	buffActionID := ActionID{SpellID: windfuryBuffSpellRanks[rank-1]}
	apBonus := windfuryAPBonuses[rank-1]
	apBonus *= 1 + 0.15*float64(iwtTalentPoints)

	var charges int32

	wfBuffAura := character.NewTemporaryStatsAuraWrapped("Windfury Buff", buffActionID, stats.Stats{stats.AttackPower: apBonus}, time.Millisecond*1500, func(config *Aura) {
		config.OnSpellHit = func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			if !spellEffect.OutcomeRollCategory.Matches(OutcomeRollCategoryWhite) {
				return
			}
			charges--
			if charges == 0 {
				aura.Deactivate(sim)
			}
		}
	})

	wfTemplate := character.AutoAttacks.MHAuto.Template
	wfTemplate.ActionID = ActionID{SpellID: windfuryBuffSpellRanks[rank-1]} // temporary buff ("Windfury Attack") spell id

	wfSpell := character.GetOrRegisterSpell(SpellConfig{
		Template:   wfTemplate,
		ModifyCast: ModifyCastAssignTarget,
	})

	const procChance = 0.2

	icd := NewICD()
	const icdDur = time.Duration(1)

	return character.GetOrRegisterAura(&Aura{
		Label:    "Windfury Totem",
		ActionID: ActionID{SpellID: WindfuryTotemSpellRanks[rank-1]}, // totem spell id ("Windfury Totem")
		OnSpellHit: func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(ProcMaskMeleeMHAuto) {
				return
			}

			if wfBuffAura.IsActive() {
				return
			}
			if icd.IsOnCD(sim) {
				// Checking for WF buff aura isn't quite enough now that we refactored auras.
				// TODO: Clean this up to remove the need for an instant ICD.
				return
			}

			if sim.RandomFloat("Windfury Totem") > procChance {
				return
			}

			// TODO: the current proc system adds auras after cast and damage, in game they're added after cast
			startCharges := int32(2)
			if !spellEffect.ProcMask.Matches(ProcMaskMeleeMHSpecial) {
				startCharges--
			}
			charges = startCharges
			wfBuffAura.Activate(sim)
			icd = InternalCD(sim.CurrentTime + icdDur)

			wfSpell.Cast(sim, spellEffect.Target)
		},
	})
}

// Used for approximating cooldowns applied by other players to you, such as
// bloodlust, innervate, power infusion, etc. This is specifically for buffs
// which can be consecutively applied multiple times to a single player.
type externalConsecutiveCDApproximation struct {
	ActionID         ActionID
	AuraTag          string
	CooldownID       CooldownID
	CooldownPriority float64
	Type             int32
	AuraDuration     time.Duration
	AuraCD           time.Duration

	// Callback for any special initialization.
	Init func(sim *Simulation, character *Character)

	// Callback for extra activation conditions.
	ShouldActivate CooldownActivationCondition

	// Applies the buff.
	AddAura CooldownActivation
}

// numSources is the number of other players assigned to apply the buff to this player.
// E.g. the number of other shaman in the group using bloodlust.
func registerExternalConsecutiveCDApproximation(agent Agent, config externalConsecutiveCDApproximation, numSources int32) {
	if numSources == 0 {
		return
	}

	externalCDs := make([]InternalCD, numSources)
	nextExternalIndex := 0

	agent.GetCharacter().AddMajorCooldown(MajorCooldown{
		ActionID:   config.ActionID,
		CooldownID: config.CooldownID,
		Cooldown:   config.AuraDuration, // Assumes that multiple buffs are different sources.
		Priority:   config.CooldownPriority,
		Type:       config.Type,

		CanActivate: func(sim *Simulation, character *Character) bool {
			if externalCDs[nextExternalIndex].IsOnCD(sim) {
				return false
			}

			if character.HasActiveAuraWithTag(config.AuraTag) {
				return false
			}

			return true
		},
		ShouldActivate: config.ShouldActivate,

		ActivationFactory: func(sim *Simulation) CooldownActivation {
			for i := 0; i < int(numSources); i++ {
				externalCDs[i] = NewICD()
			}
			nextExternalIndex = 0

			if config.Init != nil {
				config.Init(sim, agent.GetCharacter())
			}

			return func(sim *Simulation, character *Character) {
				config.AddAura(sim, character)

				externalCDs[nextExternalIndex] = InternalCD(sim.CurrentTime + config.AuraCD)
				nextExternalIndex = (nextExternalIndex + 1) % len(externalCDs)

				if externalCDs[nextExternalIndex].IsOnCD(sim) {
					character.SetCD(config.CooldownID, sim.CurrentTime+externalCDs[nextExternalIndex].GetRemainingCD(sim))
				} else {
					character.SetCD(config.CooldownID, sim.CurrentTime+config.AuraDuration)
				}
			}
		},
	})
}

const BloodlustAuraTag = "Bloodlust"

var sharedBloodlustCooldownID = NewCooldownID() // Different from shaman bloodlust CD.
const BloodlustDuration = time.Second * 40
const BloodlustCD = time.Minute * 10

func registerBloodlustCD(agent Agent, numBloodlusts int32) {
	var bloodlustAura *Aura

	registerExternalConsecutiveCDApproximation(
		agent,
		externalConsecutiveCDApproximation{
			ActionID:         ActionID{SpellID: 2825, Tag: -1},
			AuraTag:          BloodlustAuraTag,
			CooldownID:       sharedBloodlustCooldownID,
			CooldownPriority: CooldownPriorityBloodlust,
			AuraDuration:     BloodlustDuration,
			AuraCD:           BloodlustCD,
			Type:             CooldownTypeDPS,

			Init: func(sim *Simulation, character *Character) {
				bloodlustAura = BloodlustAura(character, -1)
			},
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// Haste portion doesn't stack with Power Infusion, so prefer to wait.
				return !character.HasActiveAuraWithTag(PowerInfusionAuraTag)
			},
			AddAura: func(sim *Simulation, character *Character) { bloodlustAura.Activate(sim) },
		},
		numBloodlusts)
}

func BloodlustAura(character *Character, actionTag int32) *Aura {
	const bonus = 1.3
	const inverseBonus = 1 / bonus
	actionID := ActionID{SpellID: 2825, Tag: actionTag}

	return character.GetOrRegisterAura(&Aura{
		Label:    "Bloodlust-" + actionID.String(),
		Tag:      BloodlustAuraTag,
		ActionID: actionID,
		Duration: BloodlustDuration,
		OnGain: func(aura *Aura, sim *Simulation) {
			if len(character.GetAurasWithTag(PowerInfusionAuraTag)) > 0 {
				character.PseudoStats.CastSpeedMultiplier /= 1.2
			}
			character.PseudoStats.CastSpeedMultiplier *= bonus
			character.MultiplyAttackSpeed(sim, bonus)

			if len(character.Pets) > 0 {
				for _, petAgent := range character.Pets {
					pet := petAgent.GetPet()
					if pet.IsEnabled() {
						BloodlustAura(&pet.Character, actionTag).Activate(sim)
					}
				}
			}
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			if len(character.GetAurasWithTag(PowerInfusionAuraTag)) > 0 {
				character.PseudoStats.CastSpeedMultiplier *= 1.2
			}
			character.PseudoStats.CastSpeedMultiplier *= inverseBonus
			character.MultiplyAttackSpeed(sim, inverseBonus)
		},
	})
}

var PowerInfusionAuraTag = "PowerInfusion"
var sharedPowerInfusionCooldownID = NewCooldownID() // Different from priest PI CD.
const PowerInfusionDuration = time.Second * 15
const PowerInfusionCD = time.Minute * 3

func registerPowerInfusionCD(agent Agent, numPowerInfusions int32) {
	var piAura *Aura

	registerExternalConsecutiveCDApproximation(
		agent,
		externalConsecutiveCDApproximation{
			ActionID:         ActionID{SpellID: 10060, Tag: -1},
			AuraTag:          PowerInfusionAuraTag,
			CooldownID:       sharedPowerInfusionCooldownID,
			CooldownPriority: CooldownPriorityDefault,
			AuraDuration:     PowerInfusionDuration,
			AuraCD:           PowerInfusionCD,
			Type:             CooldownTypeDPS,

			Init: func(sim *Simulation, character *Character) {
				piAura = PowerInfusionAura(character, -1)
			},
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// Haste portion doesn't stack with Bloodlust, so prefer to wait.
				return !character.HasActiveAuraWithTag(BloodlustAuraTag)
			},
			AddAura: func(sim *Simulation, character *Character) { piAura.Activate(sim) },
		},
		numPowerInfusions)
}

func PowerInfusionAura(character *Character, actionTag int32) *Aura {
	const bonus = 1.2
	const inverseBonus = 1 / bonus
	actionID := ActionID{SpellID: 10060, Tag: actionTag}

	return character.GetOrRegisterAura(&Aura{
		Label:    "PowerInfusion-" + actionID.String(),
		Tag:      PowerInfusionAuraTag,
		ActionID: actionID,
		Duration: PowerInfusionDuration,
		OnGain: func(aura *Aura, sim *Simulation) {
			if character.HasManaBar() {
				// TODO: Double-check this is how the calculation works.
				character.PseudoStats.CostMultiplier *= 0.8
			}
			if len(character.GetAurasWithTag(BloodlustAuraTag)) == 0 {
				character.PseudoStats.CastSpeedMultiplier *= bonus
			}
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			if character.HasManaBar() {
				character.PseudoStats.CostMultiplier /= 0.8
			}
			if len(character.GetAurasWithTag(BloodlustAuraTag)) == 0 {
				character.PseudoStats.CastSpeedMultiplier *= inverseBonus
			}
		},
	})
}

var InnervateAuraTag = "Innervate"
var sharedInnervateCooldownID = NewCooldownID()

const InnervateDuration = time.Second * 20
const InnervateCD = time.Minute * 6

func InnervateManaThreshold(character *Character) float64 {
	if character.Class == proto.Class_ClassMage {
		// Mages burn mana really fast so they need a higher threshold.
		return character.MaxMana() * 0.7
	} else {
		return 1000
	}
}

func registerInnervateCD(agent Agent, numInnervates int32) {
	innervateThreshold := 0.0
	expectedManaPerInnervate := 0.0
	remainingInnervateUsages := 0
	var innervateAura *Aura

	registerExternalConsecutiveCDApproximation(
		agent,
		externalConsecutiveCDApproximation{
			ActionID:         ActionID{SpellID: 29166, Tag: -1},
			AuraTag:          InnervateAuraTag,
			CooldownID:       sharedInnervateCooldownID,
			CooldownPriority: CooldownPriorityDefault,
			AuraDuration:     InnervateDuration,
			AuraCD:           InnervateCD,
			Type:             CooldownTypeMana,
			Init: func(sim *Simulation, character *Character) {
				innervateThreshold = InnervateManaThreshold(character)
				expectedManaPerInnervate = character.SpiritManaRegenPerSecond() * 5 * 20
				remainingInnervateUsages = int(1 + (MaxDuration(0, sim.Duration))/InnervateCD)
				character.ExpectedBonusMana += expectedManaPerInnervate * float64(remainingInnervateUsages)
				innervateAura = InnervateAura(character, expectedManaPerInnervate, -1)
			},
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// Only cast innervate when very low on mana, to make sure all other mana CDs are prioritized.
				if character.CurrentMana() > innervateThreshold {
					return false
				}
				return true
			},
			AddAura: func(sim *Simulation, character *Character) {
				innervateAura.Activate(sim)

				newRemainingUsages := int(sim.GetRemainingDuration() / InnervateCD)
				// AddInnervateAura already accounts for 1 usage, which is why we subtract 1 less.
				character.ExpectedBonusMana -= expectedManaPerInnervate * MaxFloat(0, float64(remainingInnervateUsages-newRemainingUsages-1))
				remainingInnervateUsages = newRemainingUsages
			},
		},
		numInnervates)
}

func InnervateAura(character *Character, expectedBonusManaReduction float64, actionTag int32) *Aura {
	actionID := ActionID{SpellID: 29166, Tag: actionTag}
	return character.GetOrRegisterAura(&Aura{
		Label:    "Innervate-" + actionID.String(),
		Tag:      InnervateAuraTag,
		ActionID: actionID,
		Duration: InnervateDuration,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.ForceFullSpiritRegen = true
			character.PseudoStats.SpiritRegenMultiplier *= 5.0
			character.UpdateManaRegenRates()

			expectedBonusManaPerTick := expectedBonusManaReduction / 10
			StartPeriodicAction(sim, PeriodicActionOptions{
				Period:   InnervateDuration / 10,
				NumTicks: 10,
				OnAction: func(sim *Simulation) {
					character.ExpectedBonusMana -= expectedBonusManaPerTick
					character.Metrics.BonusManaGained += expectedBonusManaPerTick
				},
			})
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.ForceFullSpiritRegen = false
			character.PseudoStats.SpiritRegenMultiplier /= 5.0
			character.UpdateManaRegenRates()
		},
	})
}

var ManaTideTotemAuraTag = "ManaTideTotem"
var sharedManaTideTotemCooldownID = NewCooldownID()

const ManaTideTotemDuration = time.Second * 12
const ManaTideTotemCD = time.Minute * 5

func ManaTideTotemAmount(character *Character) float64 {
	// Subtract 120 mana to simulate the loss of mana spring while MTT is active.
	// This isn't correct for multi-resto shaman groups, but that isnt a common case.
	return character.MaxMana()*0.24 - 120
}

func registerManaTideTotemCD(agent Agent, numManaTideTotems int32) {
	expectedManaPerManaTideTotem := 0.0
	remainingManaTideTotemUsages := 0
	initialDelay := time.Duration(0)
	var mttAura *Aura

	registerExternalConsecutiveCDApproximation(
		agent,
		externalConsecutiveCDApproximation{
			ActionID:         ActionID{SpellID: 16190, Tag: -1},
			AuraTag:          ManaTideTotemAuraTag,
			CooldownID:       sharedManaTideTotemCooldownID,
			CooldownPriority: CooldownPriorityDefault,
			AuraDuration:     ManaTideTotemDuration,
			AuraCD:           ManaTideTotemCD,
			Type:             CooldownTypeMana,
			Init: func(sim *Simulation, character *Character) {
				// Use first MTT at 60s, or halfway through the fight, whichever comes first.
				initialDelay = MinDuration(sim.Duration/2, time.Second*60)

				expectedManaPerManaTideTotem = ManaTideTotemAmount(character)
				remainingManaTideTotemUsages = int(1 + MaxDuration(0, sim.Duration-initialDelay)/ManaTideTotemCD)
				character.ExpectedBonusMana += expectedManaPerManaTideTotem * float64(remainingManaTideTotemUsages)
				mttAura = ManaTideTotemAura(character, -1)
			},
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// A normal resto shaman would wait to use MTT.
				if sim.CurrentTime < initialDelay {
					return false
				}
				return true
			},
			AddAura: func(sim *Simulation, character *Character) {
				mttAura.Activate(sim)

				newRemainingUsages := int(sim.GetRemainingDuration() / ManaTideTotemCD)
				// AddManaTideTotemAura already accounts for 1 usage, which is why we subtract 1 less.
				character.ExpectedBonusMana -= expectedManaPerManaTideTotem * MaxFloat(0, float64(remainingManaTideTotemUsages-newRemainingUsages-1))
				remainingManaTideTotemUsages = newRemainingUsages
			},
		},
		numManaTideTotems)
}

func ManaTideTotemAura(character *Character, actionTag int32) *Aura {
	actionID := ActionID{SpellID: 16190, Tag: actionTag}

	return character.GetOrRegisterAura(&Aura{
		Label:    "ManaTideTotem-" + actionID.String(),
		Tag:      ManaTideTotemAuraTag,
		ActionID: actionID,
		Duration: ManaTideTotemDuration,
		OnGain: func(aura *Aura, sim *Simulation) {
			if character.HasManaBar() {
				manaPerTick := ManaTideTotemAmount(character) / 4
				StartPeriodicAction(sim, PeriodicActionOptions{
					Period:   ManaTideTotemDuration / 4,
					NumTicks: 4,
					OnAction: func(sim *Simulation) {
						character.AddMana(sim, manaPerTick, actionID, true)
						character.ExpectedBonusMana -= manaPerTick
					},
				})
			}
		},
	})
}

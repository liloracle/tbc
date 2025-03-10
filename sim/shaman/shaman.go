package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func NewShaman(character core.Character, talents proto.ShamanTalents, selfBuffs SelfBuffs) *Shaman {
	shaman := &Shaman{
		Character: character,
		Talents:   talents,
		SelfBuffs: selfBuffs,
	}

	if shaman.Talents.NaturesGuidance > 0 {
		shaman.AddStat(stats.SpellHit, float64(shaman.Talents.NaturesGuidance)*1*core.SpellHitRatingPerHitChance)
		shaman.AddStat(stats.MeleeHit, float64(shaman.Talents.NaturesGuidance)*1*core.MeleeHitRatingPerHitChance)
	}

	if shaman.Talents.ThunderingStrikes > 0 {
		shaman.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*1*float64(shaman.Talents.ThunderingStrikes))
	}

	if shaman.Talents.DualWieldSpecialization > 0 {
		// TODO: Check that player is actually dual wielding
		shaman.AddStat(stats.MeleeHit, core.MeleeHitRatingPerHitChance*2*float64(shaman.Talents.DualWieldSpecialization))
	}

	if shaman.Talents.UnleashedRage > 0 {
		shaman.applyUnleashedRage(shaman.Talents.UnleashedRage)
	}
	if shaman.Talents.ShamanisticFocus {
		shaman.applyShamanisticFocus()
	}
	if shaman.Talents.Flurry > 0 {
		shaman.applyFlurry(shaman.Talents.Flurry)
	}

	// Add Shaman stat dependencies
	shaman.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect/78.1)*core.SpellCritRatingPerCritChance
		},
	})

	shaman.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Strength,
		ModifiedStat: stats.AttackPower,
		Modifier: func(strength float64, attackPower float64) float64 {
			return attackPower + strength*2
		},
	})

	shaman.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.MeleeCrit,
		Modifier: func(agility float64, meleeCrit float64) float64 {
			return meleeCrit + (agility/25)*core.MeleeCritRatingPerCritChance
		},
	})

	if shaman.Talents.UnrelentingStorm > 0 {
		coeff := 0.02 * float64(shaman.Talents.UnrelentingStorm)
		shaman.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.MP5,
			Modifier: func(intellect float64, mp5 float64) float64 {
				return mp5 + intellect*coeff
			},
		})
	}

	if shaman.Talents.AncestralKnowledge > 0 {
		coeff := 0.01 * float64(shaman.Talents.AncestralKnowledge)
		shaman.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Mana,
			ModifiedStat: stats.Mana,
			Modifier: func(mana float64, _ float64) float64 {
				return mana + mana*coeff
			},
		})
	}

	if shaman.Talents.MentalQuickness > 0 {
		coeff := 0.1 * float64(shaman.Talents.MentalQuickness)
		shaman.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.AttackPower,
			ModifiedStat: stats.SpellPower,
			Modifier: func(attackPower float64, spellPower float64) float64 {
				return spellPower + attackPower*coeff
			},
		})
	}

	if shaman.Talents.NaturesBlessing > 0 {
		coeff := 0.1 * float64(shaman.Talents.NaturesBlessing)
		shaman.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.SpellPower,
			Modifier: func(intellect float64, spellPower float64) float64 {
				return spellPower + intellect*coeff
			},
		})
	}

	if selfBuffs.WaterShield {
		shaman.AddStat(stats.MP5, 50)
	}

	shaman.registerBloodlustCD()
	shaman.registerElementalMasteryCD()
	shaman.registerNaturesSwiftnessCD()

	return shaman
}

// Which buffs this shaman is using.
type SelfBuffs struct {
	Bloodlust   bool
	WaterShield bool

	ManaSpring    bool
	EarthTotem    proto.EarthTotem
	AirTotem      proto.AirTotem
	TwistWindfury bool // if true will cast WF every 10s and then GoA

	FireTotem proto.FireTotem
	// If true Fire Nova will be dropped on CD.
	// After it will recast whatever other fire totem is available
	TwistFireNova bool

	// Mutated state on each run.
	NextTotemDrops    [4]time.Duration // track when to drop totems
	NextTotemDropType [4]int32         // track what totem to drop next to support twisting
}

// Indexes into NextTotemDrops for self buffs
const (
	AirTotem int = iota
	EarthTotem
	FireTotem
	WaterTotem
)

// Shaman represents a shaman character.
type Shaman struct {
	core.Character

	Talents   proto.ShamanTalents
	SelfBuffs SelfBuffs

	ElementalFocusStacks byte
	FlurryStacks         byte
	Focused              bool // If Shamanistic Focus is active

	// "object pool" for shaman spells that are currently being cast.
	lightningBoltSpell   core.SimpleSpell
	lightningBoltSpellLO core.SimpleSpell

	chainLightningSpell    core.MultiTargetDirectDamageSpell
	chainLightningSpellLOs []core.MultiTargetDirectDamageSpell

	// Precomputed templated cast generator for quickly resetting cast fields.
	lightningBoltCastTemplate   core.SimpleSpellTemplate
	lightningBoltLOCastTemplate core.SimpleSpellTemplate

	chainLightningCastTemplate    core.MultiTargetDirectDamageSpellTemplate
	chainLightningLOCastTemplates []core.MultiTargetDirectDamageSpellTemplate

	stormstrikeTemplate core.MeleeAbilittyTemplate
	stormstrikeSpell    core.ActiveMeleeAbility

	// Shocks
	shockSpell         core.SimpleSpell
	frostShockTemplate core.SimpleSpellTemplate
	earthShockTemplate core.SimpleSpellTemplate

	// Fire Totems
	FireTotemSpell       core.SimpleSpell
	searingTotemTemplate core.SimpleSpellTemplate
	magmaTotemTemplate   core.SimpleSpellTemplate
	novaTotemTemplate    core.SimpleSpellTemplate

	unleashedRages []core.Aura
}

// Implemented by each Shaman spec.
type ShamanAgent interface {
	core.Agent

	// The Shaman controlled by this Agent.
	GetShaman() *Shaman
}

func (shaman *Shaman) GetCharacter() *core.Character {
	return &shaman.Character
}

func (shaman *Shaman) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
}
func (shaman *Shaman) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
	if shaman.Talents.TotemOfWrath && shaman.SelfBuffs.FireTotem == proto.FireTotem_TotemOfWrath {
		partyBuffs.TotemOfWrath += 1
	}

	if shaman.SelfBuffs.ManaSpring {
		partyBuffs.ManaSpringTotem = core.MaxTristate(partyBuffs.ManaSpringTotem, proto.TristateEffect_TristateEffectRegular)
		if shaman.Talents.RestorativeTotems == 5 {
			partyBuffs.ManaSpringTotem = proto.TristateEffect_TristateEffectImproved
		}
	}

	switch shaman.SelfBuffs.AirTotem {
	case proto.AirTotem_WrathOfAirTotem:
		woaValue := proto.TristateEffect_TristateEffectRegular
		if ItemSetCycloneRegalia.CharacterHasSetBonus(shaman.GetCharacter(), 2) {
			woaValue = proto.TristateEffect_TristateEffectImproved
		}
		partyBuffs.WrathOfAirTotem = core.MaxTristate(partyBuffs.WrathOfAirTotem, woaValue)
	case proto.AirTotem_GraceOfAirTotem:
		value := proto.TristateEffect_TristateEffectRegular
		if shaman.Talents.EnhancingTotems == 2 {
			value = proto.TristateEffect_TristateEffectImproved
		}
		partyBuffs.GraceOfAirTotem = core.MaxTristate(partyBuffs.GraceOfAirTotem, value)

	}

	if shaman.SelfBuffs.EarthTotem == proto.EarthTotem_StrengthOfEarthTotem {
		value := proto.TristateEffect_TristateEffectRegular
		if shaman.Talents.EnhancingTotems == 2 {
			value = proto.TristateEffect_TristateEffectImproved
		}
		partyBuffs.StrengthOfEarthTotem = core.MaxTristate(partyBuffs.StrengthOfEarthTotem, value)
	}
}

func (shaman *Shaman) Init(sim *core.Simulation) {
	// Precompute all the spell templates.
	shaman.stormstrikeTemplate = shaman.newStormstrikeTemplate(sim)
	shaman.lightningBoltCastTemplate = shaman.newLightningBoltTemplate(sim, false)
	shaman.lightningBoltLOCastTemplate = shaman.newLightningBoltTemplate(sim, true)

	shaman.chainLightningCastTemplate = shaman.newChainLightningTemplate(sim, false)

	numHits := core.MinInt32(3, sim.GetNumTargets())
	shaman.chainLightningSpellLOs = make([]core.MultiTargetDirectDamageSpell, numHits)
	shaman.chainLightningLOCastTemplates = []core.MultiTargetDirectDamageSpellTemplate{}
	for i := int32(0); i < numHits; i++ {
		shaman.chainLightningLOCastTemplates = append(shaman.chainLightningLOCastTemplates, shaman.newChainLightningTemplate(sim, true))
	}
	shaman.frostShockTemplate = shaman.newFrostShockTemplate(sim)
	shaman.earthShockTemplate = shaman.newEarthShockTemplate(sim)

	shaman.searingTotemTemplate = shaman.newSearingTotemTemplate(sim)
	shaman.magmaTotemTemplate = shaman.newMagmaTotemTemplate(sim)
	shaman.novaTotemTemplate = shaman.newNovaTotemTemplate(sim)
}

func (shaman *Shaman) Reset(sim *core.Simulation) {
	// Check to see if we are casting a totem to set its expire time.
	for i := range shaman.SelfBuffs.NextTotemDrops {
		shaman.SelfBuffs.NextTotemDrops[i] = core.NeverExpires
		switch i {
		case AirTotem:
			if shaman.SelfBuffs.AirTotem != proto.AirTotem_NoAirTotem {
				shaman.SelfBuffs.NextTotemDrops[i] = time.Second * 120 // 2 min until drop totems
				shaman.SelfBuffs.NextTotemDropType[i] = int32(shaman.SelfBuffs.AirTotem)
			}
			if shaman.SelfBuffs.TwistWindfury {
				shaman.SelfBuffs.NextTotemDropType[i] = int32(proto.AirTotem_WindfuryTotem)
				shaman.SelfBuffs.NextTotemDrops[i] = time.Second * 10 // gotta recast windfury after 10s
			}
		case EarthTotem:
			if shaman.SelfBuffs.EarthTotem != proto.EarthTotem_NoEarthTotem {
				shaman.SelfBuffs.NextTotemDrops[i] = time.Second * 120 // 2 min until drop totems
				shaman.SelfBuffs.NextTotemDropType[i] = int32(shaman.SelfBuffs.EarthTotem)
			}
		case FireTotem:
			if shaman.SelfBuffs.FireTotem != proto.FireTotem_NoFireTotem {
				shaman.SelfBuffs.NextTotemDropType[i] = int32(shaman.SelfBuffs.FireTotem)
				if shaman.SelfBuffs.TwistFireNova {
					shaman.SelfBuffs.NextTotemDropType[FireTotem] = int32(proto.FireTotem_FireNovaTotem) // start by dropping nova, then alternating.
				}
				shaman.SelfBuffs.NextTotemDrops[i] = time.Second * 120 // 2 min until drop totems
				if shaman.SelfBuffs.FireTotem != proto.FireTotem_TotemOfWrath {
					shaman.SelfBuffs.NextTotemDrops[i] = 0 // attack totems we drop immediately
				}
			}
		case WaterTotem:
			if shaman.SelfBuffs.ManaSpring {
				shaman.SelfBuffs.NextTotemDrops[i] = time.Second * 120 // 2 min until drop totems
			}
		}
	}

	// Reset stacks and unleashed rage auras
	shaman.unleashedRages = shaman.unleashedRages[0:]
	shaman.ElementalFocusStacks = 0
}

func (shaman *Shaman) Advance(sim *core.Simulation, elapsedTime time.Duration) {
	// Enh shaman could have a 5s window without casting, use longer regen function
	shaman.Character.RegenMana(sim, elapsedTime)
}

var BloodlustCooldownID = core.NewCooldownID()

func (shaman *Shaman) registerBloodlustCD() {
	if !shaman.SelfBuffs.Bloodlust {
		return
	}

	shaman.AddMajorCooldown(core.MajorCooldown{
		ActionID:   core.ActionID{SpellID: 2825, Tag: int32(shaman.RaidIndex)},
		CooldownID: BloodlustCooldownID,
		Cooldown:   core.BloodlustCD,
		Priority:   core.CooldownPriorityBloodlust,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Need to check if any party member has lust, not just self, because of
			// major CD ordering issues with the shared bloodlust.
			for _, partyMember := range character.Party.Players {
				if partyMember.GetCharacter().HasAura(core.BloodlustAuraID) {
					return false
				}
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				for _, partyMember := range character.Party.Players {
					core.AddBloodlustAura(sim, partyMember.GetCharacter())
				}
				character.SetCD(BloodlustCooldownID, sim.CurrentTime+core.BloodlustCD)
				character.Metrics.AddInstantCast(core.ActionID{SpellID: 2825})
			}
		},
	})
}

var ElementalMasteryAuraID = core.NewAuraID()
var ElementalMasteryCooldownID = core.NewCooldownID()

func (shaman *Shaman) registerElementalMasteryCD() {
	if !shaman.Talents.ElementalMastery {
		return
	}

	shaman.AddMajorCooldown(core.MajorCooldown{
		ActionID:   core.ActionID{SpellID: 16166},
		CooldownID: ElementalMasteryCooldownID,
		Cooldown:   time.Minute * 3,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				character.Metrics.AddInstantCast(core.ActionID{SpellID: 16166})

				character.AddAura(sim, core.Aura{
					ID:      ElementalMasteryAuraID,
					SpellID: 16166,
					Name:    "Elemental Mastery",
					Expires: core.NeverExpires,
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						cast.ManaCost = 0
						cast.BonusCritRating = 100.0 * core.SpellCritRatingPerCritChance
					},
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						// Remove the buff and put skill on CD
						character.SetCD(ElementalMasteryCooldownID, sim.CurrentTime+time.Minute*3)
						character.RemoveAura(sim, ElementalMasteryAuraID)
						character.UpdateMajorCooldowns()
					},
				})
			}
		},
	})
}

var NaturesSwiftnessAuraID = core.NewAuraID()
var NaturesSwiftnessCooldownID = core.NewCooldownID()

func (shaman *Shaman) registerNaturesSwiftnessCD() {
	if !shaman.Talents.NaturesSwiftness {
		return
	}

	shaman.AddMajorCooldown(core.MajorCooldown{
		ActionID:   core.ActionID{SpellID: 16188},
		CooldownID: NaturesSwiftnessCooldownID,
		Cooldown:   time.Minute * 3,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Don't use NS unless we're casting a full-length lightning bolt, which is
			// the only spell shamans have with a cast longer than GCD.
			if character.HasTemporarySpellCastSpeedIncrease() {
				return false
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				character.AddAura(sim, core.Aura{
					ID:      NaturesSwiftnessAuraID,
					Name:    "Nature's Swiftness",
					Expires: core.NeverExpires,
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						if cast.ActionID.SpellID != SpellIDLB12 {
							return
						}

						cast.CastTime = 0
					},
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						if cast.ActionID.SpellID != SpellIDLB12 {
							return
						}

						// Remove the buff and put skill on CD
						character.SetCD(NaturesSwiftnessCooldownID, sim.CurrentTime+time.Minute*3)
						character.RemoveAura(sim, NaturesSwiftnessAuraID)
						character.UpdateMajorCooldowns()
						character.Metrics.AddInstantCast(core.ActionID{SpellID: 16188})
					},
				})
			}
		},
	})
}

var UnleasedRageTalentAuraID = core.NewAuraID()
var UnleasedRageProcAuraID = core.NewAuraID()

func (shaman *Shaman) applyUnleashedRage(level int32) {
	if len(shaman.unleashedRages) == 0 {
		shaman.unleashedRages = make([]core.Aura, 5) //pre-fill up to 5 auras
		for i := range shaman.unleashedRages {
			shaman.unleashedRages[i] = core.Aura{
				ID:   UnleasedRageProcAuraID,
				Name: "Unleased Rage",
			}
		}
	}
	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		dur := time.Second * 10
		bonus := 0.02 * float64(level)
		return core.Aura{
			ID:   UnleasedRageTalentAuraID,
			Name: "Unleashed Rage Talent",
			OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
				if result != core.MeleeHitTypeCrit {
					return
				}
				for i, player := range shaman.GetCharacter().Party.Players {
					char := player.GetCharacter()

					// Set new expire time.
					shaman.unleashedRages[i].Expires = sim.CurrentTime + dur

					if char.HasAura(UnleasedRageProcAuraID) {
						// Renew existing
						char.ReplaceAura(sim, shaman.unleashedRages[i])
						continue
					}

					// Update aura with new OnExpire
					ap := char.GetStat(stats.AttackPower) * bonus
					char.AddStat(stats.AttackPower, ap)
					shaman.unleashedRages[i].OnExpire = func(sim *core.Simulation) {
						char.AddStat(stats.AttackPower, -ap)
					}
					char.AddAura(sim, shaman.unleashedRages[i])
				}
			},
		}
	})
}

var ShamanisticFocusTalentAuraID = core.NewAuraID()

func (shaman *Shaman) applyShamanisticFocus() {
	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID:   ShamanisticFocusTalentAuraID,
			Name: "Shamanistic Focus Talent",
			OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
				if result != core.MeleeHitTypeCrit {
					return
				}
				shaman.Focused = true
			},
		}
	})
}

func (shaman *Shaman) applyFocusedEffect(cast *core.SimpleSpell) {
	cast.ManaCost -= cast.BaseManaCost * 0.6
	cast.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		shaman.Focused = false
	}
}

var FlurryTalentAuraID = core.NewAuraID()
var FlurryProcAuraID = core.NewAuraID()

func (shaman *Shaman) applyFlurry(level int32) {
	shaman.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		shaman.FlurryStacks = 0 // this function is invoked on reset, so this resets the stacks.
		icdDur := time.Millisecond * 500
		var icd core.InternalCD

		bonus := 1.3
		inverseBonus := 1 / 1.3
		return core.Aura{
			ID:   FlurryTalentAuraID,
			Name: "Flurry Talent",
			OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
				if result != core.MeleeHitTypeCrit {
					if ability == nil {
						// Remove a stack from auto attacks
						if shaman.FlurryStacks > 0 && !icd.IsOnCD(sim) {
							icd = core.InternalCD(sim.CurrentTime + icdDur)
							shaman.FlurryStacks--
							// Remove aura will reset attack speed
							if shaman.FlurryStacks == 0 {
								shaman.RemoveAura(sim, FlurryProcAuraID)
							}
						}
					}
					return
				}
				if shaman.FlurryStacks == 0 {
					shaman.MultiplyMeleeSpeed(sim, bonus)
					shaman.AddAura(sim, core.Aura{
						ID:      FlurryProcAuraID,
						SpellID: 16280,
						Name:    "Flurry",
						Expires: core.NeverExpires,
						OnExpire: func(sim *core.Simulation) {
							shaman.MultiplyMeleeSpeed(sim, inverseBonus)
						},
					})
				}
				icd = 0
				shaman.FlurryStacks = 3
			},
		}
	})
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDraenei, Class: proto.Class_ClassShaman}] = stats.Stats{
		stats.Strength:    103,
		stats.Agility:     61,
		stats.Stamina:     113,
		stats.Intellect:   109,
		stats.Spirit:      122,
		stats.Mana:        2958,
		stats.SpellCrit:   47.89,
		stats.AttackPower: 120,
		stats.MeleeCrit:   50.16,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceOrc, Class: proto.Class_ClassShaman}] = stats.Stats{
		stats.Strength:    105,
		stats.Agility:     61,
		stats.Stamina:     116,
		stats.Intellect:   105,
		stats.Spirit:      123,
		stats.Mana:        2958,
		stats.SpellCrit:   47.89,
		stats.AttackPower: 120,
		stats.MeleeCrit:   50.16,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTauren, Class: proto.Class_ClassShaman}] = stats.Stats{
		stats.Strength:    107,
		stats.Agility:     59,
		stats.Stamina:     116,
		stats.Intellect:   103,
		stats.Spirit:      122,
		stats.Mana:        2958,
		stats.SpellCrit:   47.89,
		stats.AttackPower: 120,
		stats.MeleeCrit:   49.72,
	}

	trollStats := stats.Stats{
		stats.Strength:    103,
		stats.Agility:     66,
		stats.Stamina:     115,
		stats.Intellect:   104,
		stats.Spirit:      121,
		stats.Mana:        2958,
		stats.SpellCrit:   47.89,
		stats.AttackPower: 120,
		stats.MeleeCrit:   51.23,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll10, Class: proto.Class_ClassShaman}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll30, Class: proto.Class_ClassShaman}] = trollStats
}

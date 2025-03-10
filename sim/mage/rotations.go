package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func (mage *Mage) Act(sim *core.Simulation) time.Duration {
	// If a major cooldown uses the GCD, it might already be on CD when Act() is called.
	if mage.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
		return sim.CurrentTime + mage.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	}

	var spell *core.SimpleSpell
	if mage.RotationType == proto.Mage_Rotation_Arcane {
		spell = mage.doArcaneRotation(sim)
	} else if mage.RotationType == proto.Mage_Rotation_Fire {
		spell = mage.doFireRotation(sim)
	} else {
		spell = mage.doFrostRotation(sim)
	}

	actionSuccessful := spell.Cast(sim)

	if !actionSuccessful {
		regenTime := mage.TimeUntilManaRegen(spell.GetManaCost())
		// Waiting too long can give us enough mana to pick less mana-effecient spells.
		waitTime := core.MinDuration(regenTime, time.Second*1)
		waitAction := core.NewWaitAction(sim, mage.GetCharacter(), waitTime, core.WaitReasonOOM)
		waitAction.Cast(sim)
		return sim.CurrentTime + waitAction.GetDuration()
	}

	return sim.CurrentTime + core.MaxDuration(
		mage.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime),
		spell.GetDuration())
}

func (mage *Mage) doArcaneRotation(sim *core.Simulation) *core.SimpleSpell {
	// Only arcane rotation cares about mana tracking so update it here.
	// Don't need to update tracker because we only use certain functions.
	//mage.manaTracker.Update(sim, mage.GetCharacter())

	target := sim.GetPrimaryTarget()

	// Create an AB object because we use its mana cost / cast time in many of our calculations.
	arcaneBlast, numStacks := mage.NewArcaneBlast(sim, target)
	willDropStacks := mage.willDropArcaneBlastStacks(sim, arcaneBlast, numStacks)

	mage.isBlastSpamming = mage.canBlast(sim, arcaneBlast, numStacks, willDropStacks)
	if mage.isBlastSpamming {
		return arcaneBlast
	}

	currentManaPercent := mage.CurrentManaPercent()

	if mage.isDoingRegenRotation {
		// Check if we should stop regen rotation.
		if currentManaPercent > mage.ArcaneRotation.StopRegenRotationPercent && willDropStacks {
			mage.isDoingRegenRotation = false
		}
	} else {
		// Check if we should start regen rotation.
		startThreshold := mage.ArcaneRotation.StartRegenRotationPercent
		if mage.HasAura(core.BloodlustAuraID) {
			startThreshold = core.MinFloat(0.1, startThreshold)
		}

		if currentManaPercent < startThreshold {
			mage.isDoingRegenRotation = true
			mage.tryingToDropStacks = true
			mage.numCastsDone = 0
		}
	}

	if !mage.isDoingRegenRotation {
		return arcaneBlast
	}

	if mage.tryingToDropStacks {
		if willDropStacks {
			mage.tryingToDropStacks = false
			mage.numCastsDone = 1 // 1 to count the blast we're about to return
			return arcaneBlast
		} else {
			// Do a filler spell while waiting for stacks to drop.
			arcaneBlast.Cancel(sim)
			mage.numCastsDone++
			switch mage.ArcaneRotation.Filler {
			case proto.Mage_Rotation_ArcaneRotation_Frostbolt:
				return mage.NewFrostbolt(sim, target)
			case proto.Mage_Rotation_ArcaneRotation_ArcaneMissiles:
				return mage.NewArcaneMissiles(sim, target)
			case proto.Mage_Rotation_ArcaneRotation_Scorch:
				return mage.NewScorch(sim, target)
			case proto.Mage_Rotation_ArcaneRotation_Fireball:
				return mage.NewFireball(sim, target)
			case proto.Mage_Rotation_ArcaneRotation_ArcaneMissilesFrostbolt:
				if mage.numCastsDone%2 == 1 {
					return mage.NewArcaneMissiles(sim, target)
				} else {
					return mage.NewFrostbolt(sim, target)
				}
			case proto.Mage_Rotation_ArcaneRotation_ArcaneMissilesScorch:
				if mage.numCastsDone%2 == 1 {
					return mage.NewArcaneMissiles(sim, target)
				} else {
					return mage.NewScorch(sim, target)
				}
			case proto.Mage_Rotation_ArcaneRotation_ScorchTwoFireball:
				if mage.numCastsDone%3 == 1 {
					return mage.NewScorch(sim, target)
				} else {
					return mage.NewFireball(sim, target)
				}
			default:
				return mage.NewFrostbolt(sim, target)
			}
		}
	} else {
		mage.numCastsDone++
		if mage.numCastsDone >= mage.ArcaneRotation.ArcaneBlastsBetweenFillers {
			mage.tryingToDropStacks = true
			mage.numCastsDone = 0
		}
		return arcaneBlast
	}
}

func (mage *Mage) doFireRotation(sim *core.Simulation) *core.SimpleSpell {
	target := sim.GetPrimaryTarget()

	if mage.FireRotation.MaintainImprovedScorch && (target.NumStacks(core.ImprovedScorchDebuffID) < 5 || target.RemainingAuraDuration(sim, core.ImprovedScorchDebuffID) < time.Millisecond*5500) {
		return mage.NewScorch(sim, target)
	}

	if mage.FireRotation.WeaveFireBlast && !mage.IsOnCD(FireBlastCooldownID, sim.CurrentTime) {
		return mage.NewFireBlast(sim, target)
	}

	if mage.FireRotation.PrimarySpell == proto.Mage_Rotation_FireRotation_Fireball {
		return mage.NewFireball(sim, target)
	} else {
		return mage.NewScorch(sim, target)
	}
}

func (mage *Mage) doFrostRotation(sim *core.Simulation) *core.SimpleSpell {
	target := sim.GetPrimaryTarget()
	spell := mage.NewFrostbolt(sim, target)
	return spell
}

import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Class } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { Drums } from '/tbc/core/proto/common.js';
import { Encounter } from '/tbc/core/proto/common.js';
import { ItemSlot } from '/tbc/core/proto/common.js';
import { MobType } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { TristateEffect } from '/tbc/core/proto/common.js'
import { Player } from '/tbc/core/player.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { Sim } from '/tbc/core/sim.js';
import { IndividualSimUI } from '/tbc/core/individual_sim_ui.js';

import { ElementalShaman, ElementalShaman_Rotation as ElementalShamanRotation, ElementalShaman_Options as ElementalShamanOptions } from '/tbc/core/proto/shaman.js';

import * as IconInputs from '/tbc/core/components/icon_inputs.js';
import * as OtherInputs from '/tbc/core/components/other_inputs.js';
import * as Mechanics from '/tbc/core/constants/mechanics.js';
import * as Tooltips from '/tbc/core/constants/tooltips.js';

import * as ShamanInputs from './inputs.js';
import * as Presets from './presets.js';

export class ElementalShamanSimUI extends IndividualSimUI<Spec.SpecElementalShaman> {
  constructor(parentElem: HTMLElement, player: Player<Spec.SpecElementalShaman>) {
		super(parentElem, player, {
			cssClass: 'elemental-shaman-sim-ui',
			// List any known bugs / issues here and they'll be shown on the site.
			knownIssues: [
			],

			// All stats for which EP should be calculated.
			epStats: [
				Stat.StatIntellect,
				Stat.StatSpellPower,
				Stat.StatNatureSpellPower,
				Stat.StatSpellHit,
				Stat.StatSpellCrit,
				Stat.StatSpellHaste,
				Stat.StatMP5,
			],
			// Reference stat against which to calculate EP. I think all classes use either spell power or attack power.
			epReferenceStat: Stat.StatSpellPower,
			// Which stats to display in the Character Stats section, at the bottom of the left-hand sidebar.
			displayStats: [
				Stat.StatStamina,
				Stat.StatIntellect,
				Stat.StatSpellPower,
				Stat.StatNatureSpellPower,
				Stat.StatSpellHit,
				Stat.StatSpellCrit,
				Stat.StatSpellHaste,
				Stat.StatMP5,
			],
			modifyDisplayStats: (player: Player<Spec.SpecElementalShaman>, stats: Stats) => {
				stats = stats.withStat(Stat.StatSpellHit,
						stats.getStat(Stat.StatSpellHit)
						+ player.getTalents().elementalPrecision * 2 * Mechanics.SPELL_HIT_RATING_PER_HIT_CHANCE);

				stats = stats.withStat(Stat.StatSpellCrit,
						stats.getStat(Stat.StatSpellCrit)
						+ player.getTalents().lightningMastery * 1 * Mechanics.SPELL_CRIT_RATING_PER_CRIT_CHANCE
						+ player.getTalents().tidalMastery * 1 * Mechanics.SPELL_CRIT_RATING_PER_CRIT_CHANCE);

				return stats;
			},

			defaults: {
				// Default equipped gear.
				gear: Presets.P1_PRESET.gear,
				// Default EP weights for sorting gear in the gear picker.
				epWeights: Stats.fromMap({
					[Stat.StatIntellect]: 0.33,
					[Stat.StatSpellPower]: 1,
					[Stat.StatNatureSpellPower]: 1,
					[Stat.StatSpellCrit]: 0.78,
					[Stat.StatSpellHaste]: 1.25,
					[Stat.StatMP5]: 0.08,
				}),
				// Default consumes settings.
				consumes: Presets.DefaultConsumes,
				// Default rotation settings.
				rotation: Presets.DefaultRotation,
				// Default talents.
				talents: Presets.StandardTalents.data,
				// Default spec-specific settings.
				specOptions: Presets.DefaultOptions,
				// Default raid/party buffs settings.
				raidBuffs: RaidBuffs.create({
					arcaneBrilliance: true,
					divineSpirit: TristateEffect.TristateEffectImproved,
					giftOfTheWild: TristateEffect.TristateEffectImproved,
				}),
				partyBuffs: PartyBuffs.create({
				}),
				individualBuffs: IndividualBuffs.create({
					blessingOfKings: true,
					blessingOfWisdom: 2,
				}),
				debuffs: Debuffs.create({
					judgementOfWisdom: true,
					misery: true,
				}),
			},

			// IconInputs to include in the 'Self Buffs' section on the settings tab.
			selfBuffInputs: [
				ShamanInputs.IconWaterShield,
				ShamanInputs.IconBloodlust,
				ShamanInputs.IconWrathOfAirTotem,
				ShamanInputs.IconTotemOfWrath,
				ShamanInputs.IconManaSpringTotem,
				IconInputs.DrumsOfBattleConsume,
				IconInputs.DrumsOfRestorationConsume,
			],
			// IconInputs to include in the 'Other Buffs' section on the settings tab.
			raidBuffInputs: [
				IconInputs.ArcaneBrilliance,
				IconInputs.DivineSpirit,
				IconInputs.GiftOfTheWild,
			],
			partyBuffInputs: [
				IconInputs.MoonkinAura,
				IconInputs.DrumsOfBattleBuff,
				IconInputs.DrumsOfRestorationBuff,
				IconInputs.Bloodlust,
				IconInputs.WrathOfAirTotem,
				IconInputs.TotemOfWrath,
				IconInputs.ManaSpringTotem,
				IconInputs.ManaTideTotem,
				IconInputs.EyeOfTheNight,
				IconInputs.ChainOfTheTwilightOwl,
				IconInputs.JadePendantOfBlasting,
				IconInputs.AtieshWarlock,
				IconInputs.AtieshMage,
			],
			playerBuffInputs: [
				IconInputs.BlessingOfKings,
				IconInputs.BlessingOfWisdom,
				IconInputs.Innervate,
				IconInputs.PowerInfusion,
			],
			// IconInputs to include in the 'Debuffs' section on the settings tab.
			debuffInputs: [
				IconInputs.JudgementOfWisdom,
				IconInputs.ImprovedSealOfTheCrusader,
				IconInputs.Misery,
			],
			// IconInputs to include in the 'Consumes' section on the settings tab.
			consumeInputs: [
				IconInputs.DefaultSuperManaPotion,
				IconInputs.DefaultDestructionPotion,
				IconInputs.DefaultDarkRune,
				IconInputs.FlaskOfBlindingLight,
				IconInputs.FlaskOfSupremePower,
				IconInputs.AdeptsElixir,
				IconInputs.ElixirOfMajorMageblood,
				IconInputs.ElixirOfDraenicWisdom,
				IconInputs.BrilliantWizardOil,
				IconInputs.SuperiorWizardOil,
				IconInputs.BlackenedBasilisk,
				IconInputs.SkullfishSoup,
			],
			// Inputs to include in the 'Rotation' section on the settings tab.
			rotationInputs: ShamanInputs.ElementalShamanRotationConfig,
			// Inputs to include in the 'Other' section on the settings tab.
			otherInputs: {
				inputs: [
					OtherInputs.ShadowPriestDPS,
					OtherInputs.StartingPotion,
					OtherInputs.NumStartingPotions,
				],
			},
			encounterPicker: {
				// Whether to include 'Target Armor' in the 'Encounter' section of the settings tab.
				showTargetArmor: false,
				// Whether to include 'Execute Duration (%)' in the 'Encounter' section of the settings tab.
				showExecuteProportion: false,
				// Whether to include 'Num Targets' in the 'Encounter' section of the settings tab.
				showNumTargets: true,
			},

			// If true, the talents on the talents tab will not be individually modifiable by the user.
			// Note that the use can still pick between preset talents, if there is more than 1.
			freezeTalents: false,

			presets: {
				// Preset talents that the user can quickly select.
				talents: [
					Presets.StandardTalents,
				],
				// Preset gear configurations that the user can quickly select.
				gear: [
					Presets.P1_PRESET,
					Presets.P2_PRESET,
					Presets.P3_PRESET,
				],
			},
		});
	}
}

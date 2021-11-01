import { Player } from '/tbc/core/player.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';

import { DruidTalentsPicker } from './druid.js';
import { MageTalentsPicker } from './mage.js';
import { PaladinTalentsPicker } from './paladin.js';
import { PriestTalentsPicker } from './priest.js';
import { ShamanTalentsPicker } from './shaman.js';
import { WarlockTalentsPicker } from './warlock.js';
import { TalentsPicker } from './talents_picker.js';

export function newTalentsPicker<SpecType extends Spec>(spec: Spec, parent: HTMLElement, player: Player<SpecType>): TalentsPicker<SpecType> {
  switch (spec) {
    case Spec.SpecBalanceDruid:
      return new DruidTalentsPicker(parent, sim as Sim<Spec.SpecBalanceDruid>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecElementalShaman:
      return new ShamanTalentsPicker(parent, player as Player<Spec.SpecElementalShaman>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecMage:
      return new MageTalentsPicker(parent, sim as Sim<Spec.SpecMage>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecRetributionPaladin:
      return new PaladinTalentsPicker(parent, sim as Sim<Spec.SpecRetributionPaladin>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecShadowPriest:
      return new PriestTalentsPicker(parent, sim as Sim<Spec.SpecShadowPriest>) as TalentsPicker<SpecType>;
      break;
    case Spec.SpecWarlock:
      return new WarlockTalentsPicker(parent, sim as Sim<Spec.SpecWarlock>) as TalentsPicker<SpecType>;
      break;
    default:
      const playerClass = specToClass[spec];
      throw new Error('Unimplemented class talents: ' + playerClass);
  }
}

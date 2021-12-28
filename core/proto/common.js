import { WireType } from '/tbc/protobuf-ts/index.js';
import { UnknownFieldHandler } from '/tbc/protobuf-ts/index.js';
import { reflectionMergePartial } from '/tbc/protobuf-ts/index.js';
import { MESSAGE_TYPE } from '/tbc/protobuf-ts/index.js';
import { MessageType } from '/tbc/protobuf-ts/index.js';
/**
 * @generated from protobuf enum proto.Spec
 */
export var Spec;
(function (Spec) {
    /**
     * @generated from protobuf enum value: SpecBalanceDruid = 0;
     */
    Spec[Spec["SpecBalanceDruid"] = 0] = "SpecBalanceDruid";
    /**
     * @generated from protobuf enum value: SpecElementalShaman = 1;
     */
    Spec[Spec["SpecElementalShaman"] = 1] = "SpecElementalShaman";
    /**
     * @generated from protobuf enum value: SpecEnhancementShaman = 9;
     */
    Spec[Spec["SpecEnhancementShaman"] = 9] = "SpecEnhancementShaman";
    /**
     * @generated from protobuf enum value: SpecHunter = 8;
     */
    Spec[Spec["SpecHunter"] = 8] = "SpecHunter";
    /**
     * @generated from protobuf enum value: SpecMage = 2;
     */
    Spec[Spec["SpecMage"] = 2] = "SpecMage";
    /**
     * @generated from protobuf enum value: SpecRetributionPaladin = 3;
     */
    Spec[Spec["SpecRetributionPaladin"] = 3] = "SpecRetributionPaladin";
    /**
     * @generated from protobuf enum value: SpecRogue = 7;
     */
    Spec[Spec["SpecRogue"] = 7] = "SpecRogue";
    /**
     * @generated from protobuf enum value: SpecShadowPriest = 4;
     */
    Spec[Spec["SpecShadowPriest"] = 4] = "SpecShadowPriest";
    /**
     * @generated from protobuf enum value: SpecWarlock = 5;
     */
    Spec[Spec["SpecWarlock"] = 5] = "SpecWarlock";
    /**
     * @generated from protobuf enum value: SpecWarrior = 6;
     */
    Spec[Spec["SpecWarrior"] = 6] = "SpecWarrior";
})(Spec || (Spec = {}));
/**
 * @generated from protobuf enum proto.Race
 */
export var Race;
(function (Race) {
    /**
     * @generated from protobuf enum value: RaceUnknown = 0;
     */
    Race[Race["RaceUnknown"] = 0] = "RaceUnknown";
    /**
     * @generated from protobuf enum value: RaceBloodElf = 1;
     */
    Race[Race["RaceBloodElf"] = 1] = "RaceBloodElf";
    /**
     * @generated from protobuf enum value: RaceDraenei = 2;
     */
    Race[Race["RaceDraenei"] = 2] = "RaceDraenei";
    /**
     * @generated from protobuf enum value: RaceDwarf = 3;
     */
    Race[Race["RaceDwarf"] = 3] = "RaceDwarf";
    /**
     * @generated from protobuf enum value: RaceGnome = 4;
     */
    Race[Race["RaceGnome"] = 4] = "RaceGnome";
    /**
     * @generated from protobuf enum value: RaceHuman = 5;
     */
    Race[Race["RaceHuman"] = 5] = "RaceHuman";
    /**
     * @generated from protobuf enum value: RaceNightElf = 6;
     */
    Race[Race["RaceNightElf"] = 6] = "RaceNightElf";
    /**
     * @generated from protobuf enum value: RaceOrc = 7;
     */
    Race[Race["RaceOrc"] = 7] = "RaceOrc";
    /**
     * @generated from protobuf enum value: RaceTauren = 8;
     */
    Race[Race["RaceTauren"] = 8] = "RaceTauren";
    /**
     * @generated from protobuf enum value: RaceTroll10 = 9;
     */
    Race[Race["RaceTroll10"] = 9] = "RaceTroll10";
    /**
     * @generated from protobuf enum value: RaceTroll30 = 10;
     */
    Race[Race["RaceTroll30"] = 10] = "RaceTroll30";
    /**
     * @generated from protobuf enum value: RaceUndead = 11;
     */
    Race[Race["RaceUndead"] = 11] = "RaceUndead";
})(Race || (Race = {}));
/**
 * @generated from protobuf enum proto.Class
 */
export var Class;
(function (Class) {
    /**
     * @generated from protobuf enum value: ClassUnknown = 0;
     */
    Class[Class["ClassUnknown"] = 0] = "ClassUnknown";
    /**
     * @generated from protobuf enum value: ClassDruid = 1;
     */
    Class[Class["ClassDruid"] = 1] = "ClassDruid";
    /**
     * @generated from protobuf enum value: ClassHunter = 2;
     */
    Class[Class["ClassHunter"] = 2] = "ClassHunter";
    /**
     * @generated from protobuf enum value: ClassMage = 3;
     */
    Class[Class["ClassMage"] = 3] = "ClassMage";
    /**
     * @generated from protobuf enum value: ClassPaladin = 4;
     */
    Class[Class["ClassPaladin"] = 4] = "ClassPaladin";
    /**
     * @generated from protobuf enum value: ClassPriest = 5;
     */
    Class[Class["ClassPriest"] = 5] = "ClassPriest";
    /**
     * @generated from protobuf enum value: ClassRogue = 6;
     */
    Class[Class["ClassRogue"] = 6] = "ClassRogue";
    /**
     * @generated from protobuf enum value: ClassShaman = 7;
     */
    Class[Class["ClassShaman"] = 7] = "ClassShaman";
    /**
     * @generated from protobuf enum value: ClassWarlock = 8;
     */
    Class[Class["ClassWarlock"] = 8] = "ClassWarlock";
    /**
     * @generated from protobuf enum value: ClassWarrior = 9;
     */
    Class[Class["ClassWarrior"] = 9] = "ClassWarrior";
})(Class || (Class = {}));
/**
 * @generated from protobuf enum proto.Stat
 */
export var Stat;
(function (Stat) {
    /**
     * @generated from protobuf enum value: StatStrength = 0;
     */
    Stat[Stat["StatStrength"] = 0] = "StatStrength";
    /**
     * @generated from protobuf enum value: StatAgility = 1;
     */
    Stat[Stat["StatAgility"] = 1] = "StatAgility";
    /**
     * @generated from protobuf enum value: StatStamina = 2;
     */
    Stat[Stat["StatStamina"] = 2] = "StatStamina";
    /**
     * @generated from protobuf enum value: StatIntellect = 3;
     */
    Stat[Stat["StatIntellect"] = 3] = "StatIntellect";
    /**
     * @generated from protobuf enum value: StatSpirit = 4;
     */
    Stat[Stat["StatSpirit"] = 4] = "StatSpirit";
    /**
     * @generated from protobuf enum value: StatSpellPower = 5;
     */
    Stat[Stat["StatSpellPower"] = 5] = "StatSpellPower";
    /**
     * @generated from protobuf enum value: StatHealingPower = 6;
     */
    Stat[Stat["StatHealingPower"] = 6] = "StatHealingPower";
    /**
     * @generated from protobuf enum value: StatArcaneSpellPower = 7;
     */
    Stat[Stat["StatArcaneSpellPower"] = 7] = "StatArcaneSpellPower";
    /**
     * @generated from protobuf enum value: StatFireSpellPower = 8;
     */
    Stat[Stat["StatFireSpellPower"] = 8] = "StatFireSpellPower";
    /**
     * @generated from protobuf enum value: StatFrostSpellPower = 9;
     */
    Stat[Stat["StatFrostSpellPower"] = 9] = "StatFrostSpellPower";
    /**
     * @generated from protobuf enum value: StatHolySpellPower = 10;
     */
    Stat[Stat["StatHolySpellPower"] = 10] = "StatHolySpellPower";
    /**
     * @generated from protobuf enum value: StatNatureSpellPower = 11;
     */
    Stat[Stat["StatNatureSpellPower"] = 11] = "StatNatureSpellPower";
    /**
     * @generated from protobuf enum value: StatShadowSpellPower = 12;
     */
    Stat[Stat["StatShadowSpellPower"] = 12] = "StatShadowSpellPower";
    /**
     * @generated from protobuf enum value: StatMP5 = 13;
     */
    Stat[Stat["StatMP5"] = 13] = "StatMP5";
    /**
     * @generated from protobuf enum value: StatSpellHit = 14;
     */
    Stat[Stat["StatSpellHit"] = 14] = "StatSpellHit";
    /**
     * @generated from protobuf enum value: StatSpellCrit = 15;
     */
    Stat[Stat["StatSpellCrit"] = 15] = "StatSpellCrit";
    /**
     * @generated from protobuf enum value: StatSpellHaste = 16;
     */
    Stat[Stat["StatSpellHaste"] = 16] = "StatSpellHaste";
    /**
     * @generated from protobuf enum value: StatSpellPenetration = 17;
     */
    Stat[Stat["StatSpellPenetration"] = 17] = "StatSpellPenetration";
    /**
     * @generated from protobuf enum value: StatAttackPower = 18;
     */
    Stat[Stat["StatAttackPower"] = 18] = "StatAttackPower";
    /**
     * @generated from protobuf enum value: StatMeleeHit = 19;
     */
    Stat[Stat["StatMeleeHit"] = 19] = "StatMeleeHit";
    /**
     * @generated from protobuf enum value: StatMeleeCrit = 20;
     */
    Stat[Stat["StatMeleeCrit"] = 20] = "StatMeleeCrit";
    /**
     * @generated from protobuf enum value: StatMeleeHaste = 21;
     */
    Stat[Stat["StatMeleeHaste"] = 21] = "StatMeleeHaste";
    /**
     * @generated from protobuf enum value: StatArmorPenetration = 22;
     */
    Stat[Stat["StatArmorPenetration"] = 22] = "StatArmorPenetration";
    /**
     * @generated from protobuf enum value: StatExpertise = 23;
     */
    Stat[Stat["StatExpertise"] = 23] = "StatExpertise";
    /**
     * @generated from protobuf enum value: StatMana = 24;
     */
    Stat[Stat["StatMana"] = 24] = "StatMana";
    /**
     * @generated from protobuf enum value: StatEnergy = 25;
     */
    Stat[Stat["StatEnergy"] = 25] = "StatEnergy";
    /**
     * @generated from protobuf enum value: StatRage = 26;
     */
    Stat[Stat["StatRage"] = 26] = "StatRage";
    /**
     * @generated from protobuf enum value: StatArmor = 27;
     */
    Stat[Stat["StatArmor"] = 27] = "StatArmor";
})(Stat || (Stat = {}));
/**
 * Does not correspond to anything in-game; just our own label to help filter
 * items in the UI.
 *
 * @generated from protobuf enum proto.ItemCategory
 */
export var ItemCategory;
(function (ItemCategory) {
    /**
     * @generated from protobuf enum value: ItemCategoryUnknown = 0;
     */
    ItemCategory[ItemCategory["ItemCategoryUnknown"] = 0] = "ItemCategoryUnknown";
    /**
     * @generated from protobuf enum value: ItemCategoryCaster = 1;
     */
    ItemCategory[ItemCategory["ItemCategoryCaster"] = 1] = "ItemCategoryCaster";
    /**
     * @generated from protobuf enum value: ItemCategoryMelee = 2;
     */
    ItemCategory[ItemCategory["ItemCategoryMelee"] = 2] = "ItemCategoryMelee";
    /**
     * @generated from protobuf enum value: ItemCategoryHybrid = 3;
     */
    ItemCategory[ItemCategory["ItemCategoryHybrid"] = 3] = "ItemCategoryHybrid";
    /**
     * @generated from protobuf enum value: ItemCategoryHealer = 4;
     */
    ItemCategory[ItemCategory["ItemCategoryHealer"] = 4] = "ItemCategoryHealer";
})(ItemCategory || (ItemCategory = {}));
/**
 * @generated from protobuf enum proto.ItemType
 */
export var ItemType;
(function (ItemType) {
    /**
     * @generated from protobuf enum value: ItemTypeUnknown = 0;
     */
    ItemType[ItemType["ItemTypeUnknown"] = 0] = "ItemTypeUnknown";
    /**
     * @generated from protobuf enum value: ItemTypeHead = 1;
     */
    ItemType[ItemType["ItemTypeHead"] = 1] = "ItemTypeHead";
    /**
     * @generated from protobuf enum value: ItemTypeNeck = 2;
     */
    ItemType[ItemType["ItemTypeNeck"] = 2] = "ItemTypeNeck";
    /**
     * @generated from protobuf enum value: ItemTypeShoulder = 3;
     */
    ItemType[ItemType["ItemTypeShoulder"] = 3] = "ItemTypeShoulder";
    /**
     * @generated from protobuf enum value: ItemTypeBack = 4;
     */
    ItemType[ItemType["ItemTypeBack"] = 4] = "ItemTypeBack";
    /**
     * @generated from protobuf enum value: ItemTypeChest = 5;
     */
    ItemType[ItemType["ItemTypeChest"] = 5] = "ItemTypeChest";
    /**
     * @generated from protobuf enum value: ItemTypeWrist = 6;
     */
    ItemType[ItemType["ItemTypeWrist"] = 6] = "ItemTypeWrist";
    /**
     * @generated from protobuf enum value: ItemTypeHands = 7;
     */
    ItemType[ItemType["ItemTypeHands"] = 7] = "ItemTypeHands";
    /**
     * @generated from protobuf enum value: ItemTypeWaist = 8;
     */
    ItemType[ItemType["ItemTypeWaist"] = 8] = "ItemTypeWaist";
    /**
     * @generated from protobuf enum value: ItemTypeLegs = 9;
     */
    ItemType[ItemType["ItemTypeLegs"] = 9] = "ItemTypeLegs";
    /**
     * @generated from protobuf enum value: ItemTypeFeet = 10;
     */
    ItemType[ItemType["ItemTypeFeet"] = 10] = "ItemTypeFeet";
    /**
     * @generated from protobuf enum value: ItemTypeFinger = 11;
     */
    ItemType[ItemType["ItemTypeFinger"] = 11] = "ItemTypeFinger";
    /**
     * @generated from protobuf enum value: ItemTypeTrinket = 12;
     */
    ItemType[ItemType["ItemTypeTrinket"] = 12] = "ItemTypeTrinket";
    /**
     * @generated from protobuf enum value: ItemTypeWeapon = 13;
     */
    ItemType[ItemType["ItemTypeWeapon"] = 13] = "ItemTypeWeapon";
    /**
     * @generated from protobuf enum value: ItemTypeRanged = 14;
     */
    ItemType[ItemType["ItemTypeRanged"] = 14] = "ItemTypeRanged";
})(ItemType || (ItemType = {}));
/**
 * @generated from protobuf enum proto.ArmorType
 */
export var ArmorType;
(function (ArmorType) {
    /**
     * @generated from protobuf enum value: ArmorTypeUnknown = 0;
     */
    ArmorType[ArmorType["ArmorTypeUnknown"] = 0] = "ArmorTypeUnknown";
    /**
     * @generated from protobuf enum value: ArmorTypeCloth = 1;
     */
    ArmorType[ArmorType["ArmorTypeCloth"] = 1] = "ArmorTypeCloth";
    /**
     * @generated from protobuf enum value: ArmorTypeLeather = 2;
     */
    ArmorType[ArmorType["ArmorTypeLeather"] = 2] = "ArmorTypeLeather";
    /**
     * @generated from protobuf enum value: ArmorTypeMail = 3;
     */
    ArmorType[ArmorType["ArmorTypeMail"] = 3] = "ArmorTypeMail";
    /**
     * @generated from protobuf enum value: ArmorTypePlate = 4;
     */
    ArmorType[ArmorType["ArmorTypePlate"] = 4] = "ArmorTypePlate";
})(ArmorType || (ArmorType = {}));
/**
 * @generated from protobuf enum proto.WeaponType
 */
export var WeaponType;
(function (WeaponType) {
    /**
     * @generated from protobuf enum value: WeaponTypeUnknown = 0;
     */
    WeaponType[WeaponType["WeaponTypeUnknown"] = 0] = "WeaponTypeUnknown";
    /**
     * @generated from protobuf enum value: WeaponTypeAxe = 1;
     */
    WeaponType[WeaponType["WeaponTypeAxe"] = 1] = "WeaponTypeAxe";
    /**
     * @generated from protobuf enum value: WeaponTypeDagger = 2;
     */
    WeaponType[WeaponType["WeaponTypeDagger"] = 2] = "WeaponTypeDagger";
    /**
     * @generated from protobuf enum value: WeaponTypeFist = 3;
     */
    WeaponType[WeaponType["WeaponTypeFist"] = 3] = "WeaponTypeFist";
    /**
     * @generated from protobuf enum value: WeaponTypeMace = 4;
     */
    WeaponType[WeaponType["WeaponTypeMace"] = 4] = "WeaponTypeMace";
    /**
     * @generated from protobuf enum value: WeaponTypeOffHand = 5;
     */
    WeaponType[WeaponType["WeaponTypeOffHand"] = 5] = "WeaponTypeOffHand";
    /**
     * @generated from protobuf enum value: WeaponTypePolearm = 6;
     */
    WeaponType[WeaponType["WeaponTypePolearm"] = 6] = "WeaponTypePolearm";
    /**
     * @generated from protobuf enum value: WeaponTypeShield = 7;
     */
    WeaponType[WeaponType["WeaponTypeShield"] = 7] = "WeaponTypeShield";
    /**
     * @generated from protobuf enum value: WeaponTypeStaff = 8;
     */
    WeaponType[WeaponType["WeaponTypeStaff"] = 8] = "WeaponTypeStaff";
    /**
     * @generated from protobuf enum value: WeaponTypeSword = 9;
     */
    WeaponType[WeaponType["WeaponTypeSword"] = 9] = "WeaponTypeSword";
})(WeaponType || (WeaponType = {}));
/**
 * @generated from protobuf enum proto.HandType
 */
export var HandType;
(function (HandType) {
    /**
     * @generated from protobuf enum value: HandTypeUnknown = 0;
     */
    HandType[HandType["HandTypeUnknown"] = 0] = "HandTypeUnknown";
    /**
     * @generated from protobuf enum value: HandTypeMainHand = 1;
     */
    HandType[HandType["HandTypeMainHand"] = 1] = "HandTypeMainHand";
    /**
     * @generated from protobuf enum value: HandTypeOneHand = 2;
     */
    HandType[HandType["HandTypeOneHand"] = 2] = "HandTypeOneHand";
    /**
     * @generated from protobuf enum value: HandTypeOffHand = 3;
     */
    HandType[HandType["HandTypeOffHand"] = 3] = "HandTypeOffHand";
    /**
     * @generated from protobuf enum value: HandTypeTwoHand = 4;
     */
    HandType[HandType["HandTypeTwoHand"] = 4] = "HandTypeTwoHand";
})(HandType || (HandType = {}));
/**
 * @generated from protobuf enum proto.RangedWeaponType
 */
export var RangedWeaponType;
(function (RangedWeaponType) {
    /**
     * @generated from protobuf enum value: RangedWeaponTypeUnknown = 0;
     */
    RangedWeaponType[RangedWeaponType["RangedWeaponTypeUnknown"] = 0] = "RangedWeaponTypeUnknown";
    /**
     * @generated from protobuf enum value: RangedWeaponTypeBow = 1;
     */
    RangedWeaponType[RangedWeaponType["RangedWeaponTypeBow"] = 1] = "RangedWeaponTypeBow";
    /**
     * @generated from protobuf enum value: RangedWeaponTypeCrossbow = 2;
     */
    RangedWeaponType[RangedWeaponType["RangedWeaponTypeCrossbow"] = 2] = "RangedWeaponTypeCrossbow";
    /**
     * @generated from protobuf enum value: RangedWeaponTypeGun = 3;
     */
    RangedWeaponType[RangedWeaponType["RangedWeaponTypeGun"] = 3] = "RangedWeaponTypeGun";
    /**
     * @generated from protobuf enum value: RangedWeaponTypeIdol = 4;
     */
    RangedWeaponType[RangedWeaponType["RangedWeaponTypeIdol"] = 4] = "RangedWeaponTypeIdol";
    /**
     * @generated from protobuf enum value: RangedWeaponTypeLibram = 5;
     */
    RangedWeaponType[RangedWeaponType["RangedWeaponTypeLibram"] = 5] = "RangedWeaponTypeLibram";
    /**
     * @generated from protobuf enum value: RangedWeaponTypeThrown = 6;
     */
    RangedWeaponType[RangedWeaponType["RangedWeaponTypeThrown"] = 6] = "RangedWeaponTypeThrown";
    /**
     * @generated from protobuf enum value: RangedWeaponTypeTotem = 7;
     */
    RangedWeaponType[RangedWeaponType["RangedWeaponTypeTotem"] = 7] = "RangedWeaponTypeTotem";
    /**
     * @generated from protobuf enum value: RangedWeaponTypeWand = 8;
     */
    RangedWeaponType[RangedWeaponType["RangedWeaponTypeWand"] = 8] = "RangedWeaponTypeWand";
})(RangedWeaponType || (RangedWeaponType = {}));
/**
 * All slots on the gear menu where a single item can be worn.
 *
 * @generated from protobuf enum proto.ItemSlot
 */
export var ItemSlot;
(function (ItemSlot) {
    /**
     * @generated from protobuf enum value: ItemSlotHead = 0;
     */
    ItemSlot[ItemSlot["ItemSlotHead"] = 0] = "ItemSlotHead";
    /**
     * @generated from protobuf enum value: ItemSlotNeck = 1;
     */
    ItemSlot[ItemSlot["ItemSlotNeck"] = 1] = "ItemSlotNeck";
    /**
     * @generated from protobuf enum value: ItemSlotShoulder = 2;
     */
    ItemSlot[ItemSlot["ItemSlotShoulder"] = 2] = "ItemSlotShoulder";
    /**
     * @generated from protobuf enum value: ItemSlotBack = 3;
     */
    ItemSlot[ItemSlot["ItemSlotBack"] = 3] = "ItemSlotBack";
    /**
     * @generated from protobuf enum value: ItemSlotChest = 4;
     */
    ItemSlot[ItemSlot["ItemSlotChest"] = 4] = "ItemSlotChest";
    /**
     * @generated from protobuf enum value: ItemSlotWrist = 5;
     */
    ItemSlot[ItemSlot["ItemSlotWrist"] = 5] = "ItemSlotWrist";
    /**
     * @generated from protobuf enum value: ItemSlotHands = 6;
     */
    ItemSlot[ItemSlot["ItemSlotHands"] = 6] = "ItemSlotHands";
    /**
     * @generated from protobuf enum value: ItemSlotWaist = 7;
     */
    ItemSlot[ItemSlot["ItemSlotWaist"] = 7] = "ItemSlotWaist";
    /**
     * @generated from protobuf enum value: ItemSlotLegs = 8;
     */
    ItemSlot[ItemSlot["ItemSlotLegs"] = 8] = "ItemSlotLegs";
    /**
     * @generated from protobuf enum value: ItemSlotFeet = 9;
     */
    ItemSlot[ItemSlot["ItemSlotFeet"] = 9] = "ItemSlotFeet";
    /**
     * @generated from protobuf enum value: ItemSlotFinger1 = 10;
     */
    ItemSlot[ItemSlot["ItemSlotFinger1"] = 10] = "ItemSlotFinger1";
    /**
     * @generated from protobuf enum value: ItemSlotFinger2 = 11;
     */
    ItemSlot[ItemSlot["ItemSlotFinger2"] = 11] = "ItemSlotFinger2";
    /**
     * @generated from protobuf enum value: ItemSlotTrinket1 = 12;
     */
    ItemSlot[ItemSlot["ItemSlotTrinket1"] = 12] = "ItemSlotTrinket1";
    /**
     * @generated from protobuf enum value: ItemSlotTrinket2 = 13;
     */
    ItemSlot[ItemSlot["ItemSlotTrinket2"] = 13] = "ItemSlotTrinket2";
    /**
     * can be 1h or 2h
     *
     * @generated from protobuf enum value: ItemSlotMainHand = 14;
     */
    ItemSlot[ItemSlot["ItemSlotMainHand"] = 14] = "ItemSlotMainHand";
    /**
     * @generated from protobuf enum value: ItemSlotOffHand = 15;
     */
    ItemSlot[ItemSlot["ItemSlotOffHand"] = 15] = "ItemSlotOffHand";
    /**
     * @generated from protobuf enum value: ItemSlotRanged = 16;
     */
    ItemSlot[ItemSlot["ItemSlotRanged"] = 16] = "ItemSlotRanged";
})(ItemSlot || (ItemSlot = {}));
/**
 * @generated from protobuf enum proto.ItemQuality
 */
export var ItemQuality;
(function (ItemQuality) {
    /**
     * @generated from protobuf enum value: ItemQualityJunk = 0;
     */
    ItemQuality[ItemQuality["ItemQualityJunk"] = 0] = "ItemQualityJunk";
    /**
     * @generated from protobuf enum value: ItemQualityCommon = 1;
     */
    ItemQuality[ItemQuality["ItemQualityCommon"] = 1] = "ItemQualityCommon";
    /**
     * @generated from protobuf enum value: ItemQualityUncommon = 2;
     */
    ItemQuality[ItemQuality["ItemQualityUncommon"] = 2] = "ItemQualityUncommon";
    /**
     * @generated from protobuf enum value: ItemQualityRare = 3;
     */
    ItemQuality[ItemQuality["ItemQualityRare"] = 3] = "ItemQualityRare";
    /**
     * @generated from protobuf enum value: ItemQualityEpic = 4;
     */
    ItemQuality[ItemQuality["ItemQualityEpic"] = 4] = "ItemQualityEpic";
    /**
     * @generated from protobuf enum value: ItemQualityLegendary = 5;
     */
    ItemQuality[ItemQuality["ItemQualityLegendary"] = 5] = "ItemQualityLegendary";
})(ItemQuality || (ItemQuality = {}));
/**
 * @generated from protobuf enum proto.GemColor
 */
export var GemColor;
(function (GemColor) {
    /**
     * @generated from protobuf enum value: GemColorUnknown = 0;
     */
    GemColor[GemColor["GemColorUnknown"] = 0] = "GemColorUnknown";
    /**
     * @generated from protobuf enum value: GemColorMeta = 1;
     */
    GemColor[GemColor["GemColorMeta"] = 1] = "GemColorMeta";
    /**
     * @generated from protobuf enum value: GemColorRed = 2;
     */
    GemColor[GemColor["GemColorRed"] = 2] = "GemColorRed";
    /**
     * @generated from protobuf enum value: GemColorBlue = 3;
     */
    GemColor[GemColor["GemColorBlue"] = 3] = "GemColorBlue";
    /**
     * @generated from protobuf enum value: GemColorYellow = 4;
     */
    GemColor[GemColor["GemColorYellow"] = 4] = "GemColorYellow";
    /**
     * @generated from protobuf enum value: GemColorGreen = 5;
     */
    GemColor[GemColor["GemColorGreen"] = 5] = "GemColorGreen";
    /**
     * @generated from protobuf enum value: GemColorOrange = 6;
     */
    GemColor[GemColor["GemColorOrange"] = 6] = "GemColorOrange";
    /**
     * @generated from protobuf enum value: GemColorPurple = 7;
     */
    GemColor[GemColor["GemColorPurple"] = 7] = "GemColorPurple";
    /**
     * @generated from protobuf enum value: GemColorPrismatic = 8;
     */
    GemColor[GemColor["GemColorPrismatic"] = 8] = "GemColorPrismatic";
})(GemColor || (GemColor = {}));
/**
 * @generated from protobuf enum proto.TristateEffect
 */
export var TristateEffect;
(function (TristateEffect) {
    /**
     * @generated from protobuf enum value: TristateEffectMissing = 0;
     */
    TristateEffect[TristateEffect["TristateEffectMissing"] = 0] = "TristateEffectMissing";
    /**
     * @generated from protobuf enum value: TristateEffectRegular = 1;
     */
    TristateEffect[TristateEffect["TristateEffectRegular"] = 1] = "TristateEffectRegular";
    /**
     * @generated from protobuf enum value: TristateEffectImproved = 2;
     */
    TristateEffect[TristateEffect["TristateEffectImproved"] = 2] = "TristateEffectImproved";
})(TristateEffect || (TristateEffect = {}));
/**
 * @generated from protobuf enum proto.Drums
 */
export var Drums;
(function (Drums) {
    /**
     * @generated from protobuf enum value: DrumsUnknown = 0;
     */
    Drums[Drums["DrumsUnknown"] = 0] = "DrumsUnknown";
    /**
     * @generated from protobuf enum value: DrumsOfBattle = 1;
     */
    Drums[Drums["DrumsOfBattle"] = 1] = "DrumsOfBattle";
    /**
     * @generated from protobuf enum value: DrumsOfRestoration = 2;
     */
    Drums[Drums["DrumsOfRestoration"] = 2] = "DrumsOfRestoration";
})(Drums || (Drums = {}));
/**
 * @generated from protobuf enum proto.Potions
 */
export var Potions;
(function (Potions) {
    /**
     * @generated from protobuf enum value: UnknownPotion = 0;
     */
    Potions[Potions["UnknownPotion"] = 0] = "UnknownPotion";
    /**
     * @generated from protobuf enum value: DestructionPotion = 1;
     */
    Potions[Potions["DestructionPotion"] = 1] = "DestructionPotion";
    /**
     * @generated from protobuf enum value: SuperManaPotion = 2;
     */
    Potions[Potions["SuperManaPotion"] = 2] = "SuperManaPotion";
})(Potions || (Potions = {}));
/**
 * @generated from protobuf enum proto.MobType
 */
export var MobType;
(function (MobType) {
    /**
     * @generated from protobuf enum value: MobTypeUnknown = 0;
     */
    MobType[MobType["MobTypeUnknown"] = 0] = "MobTypeUnknown";
    /**
     * @generated from protobuf enum value: MobTypeBeast = 1;
     */
    MobType[MobType["MobTypeBeast"] = 1] = "MobTypeBeast";
    /**
     * @generated from protobuf enum value: MobTypeDemon = 2;
     */
    MobType[MobType["MobTypeDemon"] = 2] = "MobTypeDemon";
    /**
     * @generated from protobuf enum value: MobTypeDragonkin = 3;
     */
    MobType[MobType["MobTypeDragonkin"] = 3] = "MobTypeDragonkin";
    /**
     * @generated from protobuf enum value: MobTypeElemental = 4;
     */
    MobType[MobType["MobTypeElemental"] = 4] = "MobTypeElemental";
    /**
     * @generated from protobuf enum value: MobTypeGiant = 5;
     */
    MobType[MobType["MobTypeGiant"] = 5] = "MobTypeGiant";
    /**
     * @generated from protobuf enum value: MobTypeHumanoid = 6;
     */
    MobType[MobType["MobTypeHumanoid"] = 6] = "MobTypeHumanoid";
    /**
     * @generated from protobuf enum value: MobTypeMechanical = 7;
     */
    MobType[MobType["MobTypeMechanical"] = 7] = "MobTypeMechanical";
    /**
     * @generated from protobuf enum value: MobTypeUndead = 8;
     */
    MobType[MobType["MobTypeUndead"] = 8] = "MobTypeUndead";
})(MobType || (MobType = {}));
/**
 * Extra enum for describing which items are eligible for an enchant, when
 * ItemType alone is not enough.
 *
 * @generated from protobuf enum proto.EnchantType
 */
export var EnchantType;
(function (EnchantType) {
    /**
     * @generated from protobuf enum value: EnchantTypeNormal = 0;
     */
    EnchantType[EnchantType["EnchantTypeNormal"] = 0] = "EnchantTypeNormal";
    /**
     * @generated from protobuf enum value: EnchantTypeTwoHanded = 1;
     */
    EnchantType[EnchantType["EnchantTypeTwoHanded"] = 1] = "EnchantTypeTwoHanded";
    /**
     * @generated from protobuf enum value: EnchantTypeShield = 2;
     */
    EnchantType[EnchantType["EnchantTypeShield"] = 2] = "EnchantTypeShield";
})(EnchantType || (EnchantType = {}));
// @generated message type with reflection information, may provide speed optimized methods
class RaidBuffs$Type extends MessageType {
    constructor() {
        super("proto.RaidBuffs", [
            { no: 1, name: "arcane_brilliance", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "divine_spirit", kind: "enum", T: () => ["proto.TristateEffect", TristateEffect] },
            { no: 5, name: "gift_of_the_wild", kind: "enum", T: () => ["proto.TristateEffect", TristateEffect] }
        ]);
    }
    create(value) {
        const message = { arcaneBrilliance: false, divineSpirit: 0, giftOfTheWild: 0 };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool arcane_brilliance */ 1:
                    message.arcaneBrilliance = reader.bool();
                    break;
                case /* proto.TristateEffect divine_spirit */ 4:
                    message.divineSpirit = reader.int32();
                    break;
                case /* proto.TristateEffect gift_of_the_wild */ 5:
                    message.giftOfTheWild = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* bool arcane_brilliance = 1; */
        if (message.arcaneBrilliance !== false)
            writer.tag(1, WireType.Varint).bool(message.arcaneBrilliance);
        /* proto.TristateEffect divine_spirit = 4; */
        if (message.divineSpirit !== 0)
            writer.tag(4, WireType.Varint).int32(message.divineSpirit);
        /* proto.TristateEffect gift_of_the_wild = 5; */
        if (message.giftOfTheWild !== 0)
            writer.tag(5, WireType.Varint).int32(message.giftOfTheWild);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.RaidBuffs
 */
export const RaidBuffs = new RaidBuffs$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PartyBuffs$Type extends MessageType {
    constructor() {
        super("proto.PartyBuffs", [
            { no: 1, name: "bloodlust", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "moonkin_aura", kind: "enum", T: () => ["proto.TristateEffect", TristateEffect] },
            { no: 3, name: "draenei_racial_melee", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "draenei_racial_caster", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "drums", kind: "enum", T: () => ["proto.Drums", Drums] },
            { no: 6, name: "atiesh_mage", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "atiesh_warlock", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 8, name: "braided_eternium_chain", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 9, name: "eye_of_the_night", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 10, name: "chain_of_the_twilight_owl", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 11, name: "jade_pendant_of_blasting", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 12, name: "mana_spring_totem", kind: "enum", T: () => ["proto.TristateEffect", TristateEffect] },
            { no: 17, name: "mana_tide_totems", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 13, name: "totem_of_wrath", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 14, name: "wrath_of_air_totem", kind: "enum", T: () => ["proto.TristateEffect", TristateEffect] },
            { no: 15, name: "grace_of_air_totem", kind: "enum", T: () => ["proto.TristateEffect", TristateEffect] },
            { no: 16, name: "strength_of_earth_totem", kind: "enum", T: () => ["proto.TristateEffect", TristateEffect] }
        ]);
    }
    create(value) {
        const message = { bloodlust: 0, moonkinAura: 0, draeneiRacialMelee: false, draeneiRacialCaster: false, drums: 0, atieshMage: 0, atieshWarlock: 0, braidedEterniumChain: false, eyeOfTheNight: false, chainOfTheTwilightOwl: false, jadePendantOfBlasting: false, manaSpringTotem: 0, manaTideTotems: 0, totemOfWrath: 0, wrathOfAirTotem: 0, graceOfAirTotem: 0, strengthOfEarthTotem: 0 };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 bloodlust */ 1:
                    message.bloodlust = reader.int32();
                    break;
                case /* proto.TristateEffect moonkin_aura */ 2:
                    message.moonkinAura = reader.int32();
                    break;
                case /* bool draenei_racial_melee */ 3:
                    message.draeneiRacialMelee = reader.bool();
                    break;
                case /* bool draenei_racial_caster */ 4:
                    message.draeneiRacialCaster = reader.bool();
                    break;
                case /* proto.Drums drums */ 5:
                    message.drums = reader.int32();
                    break;
                case /* int32 atiesh_mage */ 6:
                    message.atieshMage = reader.int32();
                    break;
                case /* int32 atiesh_warlock */ 7:
                    message.atieshWarlock = reader.int32();
                    break;
                case /* bool braided_eternium_chain */ 8:
                    message.braidedEterniumChain = reader.bool();
                    break;
                case /* bool eye_of_the_night */ 9:
                    message.eyeOfTheNight = reader.bool();
                    break;
                case /* bool chain_of_the_twilight_owl */ 10:
                    message.chainOfTheTwilightOwl = reader.bool();
                    break;
                case /* bool jade_pendant_of_blasting */ 11:
                    message.jadePendantOfBlasting = reader.bool();
                    break;
                case /* proto.TristateEffect mana_spring_totem */ 12:
                    message.manaSpringTotem = reader.int32();
                    break;
                case /* int32 mana_tide_totems */ 17:
                    message.manaTideTotems = reader.int32();
                    break;
                case /* int32 totem_of_wrath */ 13:
                    message.totemOfWrath = reader.int32();
                    break;
                case /* proto.TristateEffect wrath_of_air_totem */ 14:
                    message.wrathOfAirTotem = reader.int32();
                    break;
                case /* proto.TristateEffect grace_of_air_totem */ 15:
                    message.graceOfAirTotem = reader.int32();
                    break;
                case /* proto.TristateEffect strength_of_earth_totem */ 16:
                    message.strengthOfEarthTotem = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* int32 bloodlust = 1; */
        if (message.bloodlust !== 0)
            writer.tag(1, WireType.Varint).int32(message.bloodlust);
        /* proto.TristateEffect moonkin_aura = 2; */
        if (message.moonkinAura !== 0)
            writer.tag(2, WireType.Varint).int32(message.moonkinAura);
        /* bool draenei_racial_melee = 3; */
        if (message.draeneiRacialMelee !== false)
            writer.tag(3, WireType.Varint).bool(message.draeneiRacialMelee);
        /* bool draenei_racial_caster = 4; */
        if (message.draeneiRacialCaster !== false)
            writer.tag(4, WireType.Varint).bool(message.draeneiRacialCaster);
        /* proto.Drums drums = 5; */
        if (message.drums !== 0)
            writer.tag(5, WireType.Varint).int32(message.drums);
        /* int32 atiesh_mage = 6; */
        if (message.atieshMage !== 0)
            writer.tag(6, WireType.Varint).int32(message.atieshMage);
        /* int32 atiesh_warlock = 7; */
        if (message.atieshWarlock !== 0)
            writer.tag(7, WireType.Varint).int32(message.atieshWarlock);
        /* bool braided_eternium_chain = 8; */
        if (message.braidedEterniumChain !== false)
            writer.tag(8, WireType.Varint).bool(message.braidedEterniumChain);
        /* bool eye_of_the_night = 9; */
        if (message.eyeOfTheNight !== false)
            writer.tag(9, WireType.Varint).bool(message.eyeOfTheNight);
        /* bool chain_of_the_twilight_owl = 10; */
        if (message.chainOfTheTwilightOwl !== false)
            writer.tag(10, WireType.Varint).bool(message.chainOfTheTwilightOwl);
        /* bool jade_pendant_of_blasting = 11; */
        if (message.jadePendantOfBlasting !== false)
            writer.tag(11, WireType.Varint).bool(message.jadePendantOfBlasting);
        /* proto.TristateEffect mana_spring_totem = 12; */
        if (message.manaSpringTotem !== 0)
            writer.tag(12, WireType.Varint).int32(message.manaSpringTotem);
        /* int32 mana_tide_totems = 17; */
        if (message.manaTideTotems !== 0)
            writer.tag(17, WireType.Varint).int32(message.manaTideTotems);
        /* int32 totem_of_wrath = 13; */
        if (message.totemOfWrath !== 0)
            writer.tag(13, WireType.Varint).int32(message.totemOfWrath);
        /* proto.TristateEffect wrath_of_air_totem = 14; */
        if (message.wrathOfAirTotem !== 0)
            writer.tag(14, WireType.Varint).int32(message.wrathOfAirTotem);
        /* proto.TristateEffect grace_of_air_totem = 15; */
        if (message.graceOfAirTotem !== 0)
            writer.tag(15, WireType.Varint).int32(message.graceOfAirTotem);
        /* proto.TristateEffect strength_of_earth_totem = 16; */
        if (message.strengthOfEarthTotem !== 0)
            writer.tag(16, WireType.Varint).int32(message.strengthOfEarthTotem);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.PartyBuffs
 */
export const PartyBuffs = new PartyBuffs$Type();
// @generated message type with reflection information, may provide speed optimized methods
class IndividualBuffs$Type extends MessageType {
    constructor() {
        super("proto.IndividualBuffs", [
            { no: 1, name: "blessing_of_kings", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "blessing_of_wisdom", kind: "enum", T: () => ["proto.TristateEffect", TristateEffect] },
            { no: 3, name: "blessing_of_might", kind: "enum", T: () => ["proto.TristateEffect", TristateEffect] },
            { no: 4, name: "shadow_priest_dps", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "innervates", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "power_infusions", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value) {
        const message = { blessingOfKings: false, blessingOfWisdom: 0, blessingOfMight: 0, shadowPriestDps: 0, innervates: 0, powerInfusions: 0 };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool blessing_of_kings */ 1:
                    message.blessingOfKings = reader.bool();
                    break;
                case /* proto.TristateEffect blessing_of_wisdom */ 2:
                    message.blessingOfWisdom = reader.int32();
                    break;
                case /* proto.TristateEffect blessing_of_might */ 3:
                    message.blessingOfMight = reader.int32();
                    break;
                case /* int32 shadow_priest_dps */ 4:
                    message.shadowPriestDps = reader.int32();
                    break;
                case /* int32 innervates */ 5:
                    message.innervates = reader.int32();
                    break;
                case /* int32 power_infusions */ 6:
                    message.powerInfusions = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* bool blessing_of_kings = 1; */
        if (message.blessingOfKings !== false)
            writer.tag(1, WireType.Varint).bool(message.blessingOfKings);
        /* proto.TristateEffect blessing_of_wisdom = 2; */
        if (message.blessingOfWisdom !== 0)
            writer.tag(2, WireType.Varint).int32(message.blessingOfWisdom);
        /* proto.TristateEffect blessing_of_might = 3; */
        if (message.blessingOfMight !== 0)
            writer.tag(3, WireType.Varint).int32(message.blessingOfMight);
        /* int32 shadow_priest_dps = 4; */
        if (message.shadowPriestDps !== 0)
            writer.tag(4, WireType.Varint).int32(message.shadowPriestDps);
        /* int32 innervates = 5; */
        if (message.innervates !== 0)
            writer.tag(5, WireType.Varint).int32(message.innervates);
        /* int32 power_infusions = 6; */
        if (message.powerInfusions !== 0)
            writer.tag(6, WireType.Varint).int32(message.powerInfusions);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.IndividualBuffs
 */
export const IndividualBuffs = new IndividualBuffs$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Consumes$Type extends MessageType {
    constructor() {
        super("proto.Consumes", [
            { no: 1, name: "flask_of_blinding_light", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "flask_of_mighty_restoration", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "flask_of_pure_death", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "flask_of_supreme_power", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "adepts_elixir", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "elixir_of_major_fire_power", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 7, name: "elixir_of_major_frost_power", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 8, name: "elixir_of_major_shadow_power", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 9, name: "elixir_of_draenic_wisdom", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 10, name: "elixir_of_major_mageblood", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 11, name: "brilliant_wizard_oil", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 12, name: "superior_wizard_oil", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 13, name: "blackened_basilisk", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 14, name: "skullfish_soup", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 20, name: "kreegsStoutBeatdown", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 15, name: "default_potion", kind: "enum", T: () => ["proto.Potions", Potions] },
            { no: 16, name: "starting_potion", kind: "enum", T: () => ["proto.Potions", Potions] },
            { no: 17, name: "num_starting_potions", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 18, name: "dark_rune", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 19, name: "drums", kind: "enum", T: () => ["proto.Drums", Drums] }
        ]);
    }
    create(value) {
        const message = { flaskOfBlindingLight: false, flaskOfMightyRestoration: false, flaskOfPureDeath: false, flaskOfSupremePower: false, adeptsElixir: false, elixirOfMajorFirePower: false, elixirOfMajorFrostPower: false, elixirOfMajorShadowPower: false, elixirOfDraenicWisdom: false, elixirOfMajorMageblood: false, brilliantWizardOil: false, superiorWizardOil: false, blackenedBasilisk: false, skullfishSoup: false, kreegsStoutBeatdown: false, defaultPotion: 0, startingPotion: 0, numStartingPotions: 0, darkRune: false, drums: 0 };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool flask_of_blinding_light */ 1:
                    message.flaskOfBlindingLight = reader.bool();
                    break;
                case /* bool flask_of_mighty_restoration */ 2:
                    message.flaskOfMightyRestoration = reader.bool();
                    break;
                case /* bool flask_of_pure_death */ 3:
                    message.flaskOfPureDeath = reader.bool();
                    break;
                case /* bool flask_of_supreme_power */ 4:
                    message.flaskOfSupremePower = reader.bool();
                    break;
                case /* bool adepts_elixir */ 5:
                    message.adeptsElixir = reader.bool();
                    break;
                case /* bool elixir_of_major_fire_power */ 6:
                    message.elixirOfMajorFirePower = reader.bool();
                    break;
                case /* bool elixir_of_major_frost_power */ 7:
                    message.elixirOfMajorFrostPower = reader.bool();
                    break;
                case /* bool elixir_of_major_shadow_power */ 8:
                    message.elixirOfMajorShadowPower = reader.bool();
                    break;
                case /* bool elixir_of_draenic_wisdom */ 9:
                    message.elixirOfDraenicWisdom = reader.bool();
                    break;
                case /* bool elixir_of_major_mageblood */ 10:
                    message.elixirOfMajorMageblood = reader.bool();
                    break;
                case /* bool brilliant_wizard_oil */ 11:
                    message.brilliantWizardOil = reader.bool();
                    break;
                case /* bool superior_wizard_oil */ 12:
                    message.superiorWizardOil = reader.bool();
                    break;
                case /* bool blackened_basilisk */ 13:
                    message.blackenedBasilisk = reader.bool();
                    break;
                case /* bool skullfish_soup */ 14:
                    message.skullfishSoup = reader.bool();
                    break;
                case /* bool kreegsStoutBeatdown */ 20:
                    message.kreegsStoutBeatdown = reader.bool();
                    break;
                case /* proto.Potions default_potion */ 15:
                    message.defaultPotion = reader.int32();
                    break;
                case /* proto.Potions starting_potion */ 16:
                    message.startingPotion = reader.int32();
                    break;
                case /* int32 num_starting_potions */ 17:
                    message.numStartingPotions = reader.int32();
                    break;
                case /* bool dark_rune */ 18:
                    message.darkRune = reader.bool();
                    break;
                case /* proto.Drums drums */ 19:
                    message.drums = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* bool flask_of_blinding_light = 1; */
        if (message.flaskOfBlindingLight !== false)
            writer.tag(1, WireType.Varint).bool(message.flaskOfBlindingLight);
        /* bool flask_of_mighty_restoration = 2; */
        if (message.flaskOfMightyRestoration !== false)
            writer.tag(2, WireType.Varint).bool(message.flaskOfMightyRestoration);
        /* bool flask_of_pure_death = 3; */
        if (message.flaskOfPureDeath !== false)
            writer.tag(3, WireType.Varint).bool(message.flaskOfPureDeath);
        /* bool flask_of_supreme_power = 4; */
        if (message.flaskOfSupremePower !== false)
            writer.tag(4, WireType.Varint).bool(message.flaskOfSupremePower);
        /* bool adepts_elixir = 5; */
        if (message.adeptsElixir !== false)
            writer.tag(5, WireType.Varint).bool(message.adeptsElixir);
        /* bool elixir_of_major_fire_power = 6; */
        if (message.elixirOfMajorFirePower !== false)
            writer.tag(6, WireType.Varint).bool(message.elixirOfMajorFirePower);
        /* bool elixir_of_major_frost_power = 7; */
        if (message.elixirOfMajorFrostPower !== false)
            writer.tag(7, WireType.Varint).bool(message.elixirOfMajorFrostPower);
        /* bool elixir_of_major_shadow_power = 8; */
        if (message.elixirOfMajorShadowPower !== false)
            writer.tag(8, WireType.Varint).bool(message.elixirOfMajorShadowPower);
        /* bool elixir_of_draenic_wisdom = 9; */
        if (message.elixirOfDraenicWisdom !== false)
            writer.tag(9, WireType.Varint).bool(message.elixirOfDraenicWisdom);
        /* bool elixir_of_major_mageblood = 10; */
        if (message.elixirOfMajorMageblood !== false)
            writer.tag(10, WireType.Varint).bool(message.elixirOfMajorMageblood);
        /* bool brilliant_wizard_oil = 11; */
        if (message.brilliantWizardOil !== false)
            writer.tag(11, WireType.Varint).bool(message.brilliantWizardOil);
        /* bool superior_wizard_oil = 12; */
        if (message.superiorWizardOil !== false)
            writer.tag(12, WireType.Varint).bool(message.superiorWizardOil);
        /* bool blackened_basilisk = 13; */
        if (message.blackenedBasilisk !== false)
            writer.tag(13, WireType.Varint).bool(message.blackenedBasilisk);
        /* bool skullfish_soup = 14; */
        if (message.skullfishSoup !== false)
            writer.tag(14, WireType.Varint).bool(message.skullfishSoup);
        /* bool kreegsStoutBeatdown = 20; */
        if (message.kreegsStoutBeatdown !== false)
            writer.tag(20, WireType.Varint).bool(message.kreegsStoutBeatdown);
        /* proto.Potions default_potion = 15; */
        if (message.defaultPotion !== 0)
            writer.tag(15, WireType.Varint).int32(message.defaultPotion);
        /* proto.Potions starting_potion = 16; */
        if (message.startingPotion !== 0)
            writer.tag(16, WireType.Varint).int32(message.startingPotion);
        /* int32 num_starting_potions = 17; */
        if (message.numStartingPotions !== 0)
            writer.tag(17, WireType.Varint).int32(message.numStartingPotions);
        /* bool dark_rune = 18; */
        if (message.darkRune !== false)
            writer.tag(18, WireType.Varint).bool(message.darkRune);
        /* proto.Drums drums = 19; */
        if (message.drums !== 0)
            writer.tag(19, WireType.Varint).int32(message.drums);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Consumes
 */
export const Consumes = new Consumes$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Debuffs$Type extends MessageType {
    constructor() {
        super("proto.Debuffs", [
            { no: 1, name: "judgement_of_wisdom", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "improved_seal_of_the_crusader", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "misery", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "curse_of_elements", kind: "enum", T: () => ["proto.TristateEffect", TristateEffect] },
            { no: 5, name: "isb_uptime", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value) {
        const message = { judgementOfWisdom: false, improvedSealOfTheCrusader: false, misery: false, curseOfElements: 0, isbUptime: 0 };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool judgement_of_wisdom */ 1:
                    message.judgementOfWisdom = reader.bool();
                    break;
                case /* bool improved_seal_of_the_crusader */ 2:
                    message.improvedSealOfTheCrusader = reader.bool();
                    break;
                case /* bool misery */ 3:
                    message.misery = reader.bool();
                    break;
                case /* proto.TristateEffect curse_of_elements */ 4:
                    message.curseOfElements = reader.int32();
                    break;
                case /* double isb_uptime */ 5:
                    message.isbUptime = reader.double();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* bool judgement_of_wisdom = 1; */
        if (message.judgementOfWisdom !== false)
            writer.tag(1, WireType.Varint).bool(message.judgementOfWisdom);
        /* bool improved_seal_of_the_crusader = 2; */
        if (message.improvedSealOfTheCrusader !== false)
            writer.tag(2, WireType.Varint).bool(message.improvedSealOfTheCrusader);
        /* bool misery = 3; */
        if (message.misery !== false)
            writer.tag(3, WireType.Varint).bool(message.misery);
        /* proto.TristateEffect curse_of_elements = 4; */
        if (message.curseOfElements !== 0)
            writer.tag(4, WireType.Varint).int32(message.curseOfElements);
        /* double isb_uptime = 5; */
        if (message.isbUptime !== 0)
            writer.tag(5, WireType.Bit64).double(message.isbUptime);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Debuffs
 */
export const Debuffs = new Debuffs$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Target$Type extends MessageType {
    constructor() {
        super("proto.Target", [
            { no: 1, name: "armor", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 4, name: "level", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "mob_type", kind: "enum", T: () => ["proto.MobType", MobType] },
            { no: 2, name: "debuffs", kind: "message", T: () => Debuffs }
        ]);
    }
    create(value) {
        const message = { armor: 0, level: 0, mobType: 0 };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* double armor */ 1:
                    message.armor = reader.double();
                    break;
                case /* int32 level */ 4:
                    message.level = reader.int32();
                    break;
                case /* proto.MobType mob_type */ 3:
                    message.mobType = reader.int32();
                    break;
                case /* proto.Debuffs debuffs */ 2:
                    message.debuffs = Debuffs.internalBinaryRead(reader, reader.uint32(), options, message.debuffs);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* double armor = 1; */
        if (message.armor !== 0)
            writer.tag(1, WireType.Bit64).double(message.armor);
        /* int32 level = 4; */
        if (message.level !== 0)
            writer.tag(4, WireType.Varint).int32(message.level);
        /* proto.MobType mob_type = 3; */
        if (message.mobType !== 0)
            writer.tag(3, WireType.Varint).int32(message.mobType);
        /* proto.Debuffs debuffs = 2; */
        if (message.debuffs)
            Debuffs.internalBinaryWrite(message.debuffs, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Target
 */
export const Target = new Target$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Encounter$Type extends MessageType {
    constructor() {
        super("proto.Encounter", [
            { no: 1, name: "duration", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 2, name: "targets", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Target }
        ]);
    }
    create(value) {
        const message = { duration: 0, targets: [] };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* double duration */ 1:
                    message.duration = reader.double();
                    break;
                case /* repeated proto.Target targets */ 2:
                    message.targets.push(Target.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* double duration = 1; */
        if (message.duration !== 0)
            writer.tag(1, WireType.Bit64).double(message.duration);
        /* repeated proto.Target targets = 2; */
        for (let i = 0; i < message.targets.length; i++)
            Target.internalBinaryWrite(message.targets[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Encounter
 */
export const Encounter = new Encounter$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ItemSpec$Type extends MessageType {
    constructor() {
        super("proto.ItemSpec", [
            { no: 2, name: "id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "enchant", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "gems", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value) {
        const message = { id: 0, enchant: 0, gems: [] };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 id */ 2:
                    message.id = reader.int32();
                    break;
                case /* int32 enchant */ 3:
                    message.enchant = reader.int32();
                    break;
                case /* repeated int32 gems */ 4:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.gems.push(reader.int32());
                    else
                        message.gems.push(reader.int32());
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* int32 id = 2; */
        if (message.id !== 0)
            writer.tag(2, WireType.Varint).int32(message.id);
        /* int32 enchant = 3; */
        if (message.enchant !== 0)
            writer.tag(3, WireType.Varint).int32(message.enchant);
        /* repeated int32 gems = 4; */
        if (message.gems.length) {
            writer.tag(4, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.gems.length; i++)
                writer.int32(message.gems[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.ItemSpec
 */
export const ItemSpec = new ItemSpec$Type();
// @generated message type with reflection information, may provide speed optimized methods
class EquipmentSpec$Type extends MessageType {
    constructor() {
        super("proto.EquipmentSpec", [
            { no: 1, name: "items", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => ItemSpec }
        ]);
    }
    create(value) {
        const message = { items: [] };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated proto.ItemSpec items */ 1:
                    message.items.push(ItemSpec.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* repeated proto.ItemSpec items = 1; */
        for (let i = 0; i < message.items.length; i++)
            ItemSpec.internalBinaryWrite(message.items[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.EquipmentSpec
 */
export const EquipmentSpec = new EquipmentSpec$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Item$Type extends MessageType {
    constructor() {
        super("proto.Item", [
            { no: 1, name: "id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 16, name: "wowhead_id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 14, name: "categories", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["proto.ItemCategory", ItemCategory] },
            { no: 15, name: "class_allowlist", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["proto.Class", Class] },
            { no: 3, name: "type", kind: "enum", T: () => ["proto.ItemType", ItemType] },
            { no: 4, name: "armor_type", kind: "enum", T: () => ["proto.ArmorType", ArmorType] },
            { no: 5, name: "weapon_type", kind: "enum", T: () => ["proto.WeaponType", WeaponType] },
            { no: 6, name: "hand_type", kind: "enum", T: () => ["proto.HandType", HandType] },
            { no: 7, name: "ranged_weapon_type", kind: "enum", T: () => ["proto.RangedWeaponType", RangedWeaponType] },
            { no: 8, name: "stats", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 9, name: "gem_sockets", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["proto.GemColor", GemColor] },
            { no: 10, name: "socketBonus", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 11, name: "phase", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 12, name: "quality", kind: "enum", T: () => ["proto.ItemQuality", ItemQuality] },
            { no: 13, name: "unique", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { id: 0, wowheadId: 0, name: "", categories: [], classAllowlist: [], type: 0, armorType: 0, weaponType: 0, handType: 0, rangedWeaponType: 0, stats: [], gemSockets: [], socketBonus: [], phase: 0, quality: 0, unique: false };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 id */ 1:
                    message.id = reader.int32();
                    break;
                case /* int32 wowhead_id */ 16:
                    message.wowheadId = reader.int32();
                    break;
                case /* string name */ 2:
                    message.name = reader.string();
                    break;
                case /* repeated proto.ItemCategory categories */ 14:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.categories.push(reader.int32());
                    else
                        message.categories.push(reader.int32());
                    break;
                case /* repeated proto.Class class_allowlist */ 15:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.classAllowlist.push(reader.int32());
                    else
                        message.classAllowlist.push(reader.int32());
                    break;
                case /* proto.ItemType type */ 3:
                    message.type = reader.int32();
                    break;
                case /* proto.ArmorType armor_type */ 4:
                    message.armorType = reader.int32();
                    break;
                case /* proto.WeaponType weapon_type */ 5:
                    message.weaponType = reader.int32();
                    break;
                case /* proto.HandType hand_type */ 6:
                    message.handType = reader.int32();
                    break;
                case /* proto.RangedWeaponType ranged_weapon_type */ 7:
                    message.rangedWeaponType = reader.int32();
                    break;
                case /* repeated double stats */ 8:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.stats.push(reader.double());
                    else
                        message.stats.push(reader.double());
                    break;
                case /* repeated proto.GemColor gem_sockets */ 9:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.gemSockets.push(reader.int32());
                    else
                        message.gemSockets.push(reader.int32());
                    break;
                case /* repeated double socketBonus */ 10:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.socketBonus.push(reader.double());
                    else
                        message.socketBonus.push(reader.double());
                    break;
                case /* int32 phase */ 11:
                    message.phase = reader.int32();
                    break;
                case /* proto.ItemQuality quality */ 12:
                    message.quality = reader.int32();
                    break;
                case /* bool unique */ 13:
                    message.unique = reader.bool();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* int32 id = 1; */
        if (message.id !== 0)
            writer.tag(1, WireType.Varint).int32(message.id);
        /* int32 wowhead_id = 16; */
        if (message.wowheadId !== 0)
            writer.tag(16, WireType.Varint).int32(message.wowheadId);
        /* string name = 2; */
        if (message.name !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.name);
        /* repeated proto.ItemCategory categories = 14; */
        if (message.categories.length) {
            writer.tag(14, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.categories.length; i++)
                writer.int32(message.categories[i]);
            writer.join();
        }
        /* repeated proto.Class class_allowlist = 15; */
        if (message.classAllowlist.length) {
            writer.tag(15, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.classAllowlist.length; i++)
                writer.int32(message.classAllowlist[i]);
            writer.join();
        }
        /* proto.ItemType type = 3; */
        if (message.type !== 0)
            writer.tag(3, WireType.Varint).int32(message.type);
        /* proto.ArmorType armor_type = 4; */
        if (message.armorType !== 0)
            writer.tag(4, WireType.Varint).int32(message.armorType);
        /* proto.WeaponType weapon_type = 5; */
        if (message.weaponType !== 0)
            writer.tag(5, WireType.Varint).int32(message.weaponType);
        /* proto.HandType hand_type = 6; */
        if (message.handType !== 0)
            writer.tag(6, WireType.Varint).int32(message.handType);
        /* proto.RangedWeaponType ranged_weapon_type = 7; */
        if (message.rangedWeaponType !== 0)
            writer.tag(7, WireType.Varint).int32(message.rangedWeaponType);
        /* repeated double stats = 8; */
        if (message.stats.length) {
            writer.tag(8, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.stats.length; i++)
                writer.double(message.stats[i]);
            writer.join();
        }
        /* repeated proto.GemColor gem_sockets = 9; */
        if (message.gemSockets.length) {
            writer.tag(9, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.gemSockets.length; i++)
                writer.int32(message.gemSockets[i]);
            writer.join();
        }
        /* repeated double socketBonus = 10; */
        if (message.socketBonus.length) {
            writer.tag(10, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.socketBonus.length; i++)
                writer.double(message.socketBonus[i]);
            writer.join();
        }
        /* int32 phase = 11; */
        if (message.phase !== 0)
            writer.tag(11, WireType.Varint).int32(message.phase);
        /* proto.ItemQuality quality = 12; */
        if (message.quality !== 0)
            writer.tag(12, WireType.Varint).int32(message.quality);
        /* bool unique = 13; */
        if (message.unique !== false)
            writer.tag(13, WireType.Varint).bool(message.unique);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Item
 */
export const Item = new Item$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Enchant$Type extends MessageType {
    constructor() {
        super("proto.Enchant", [
            { no: 1, name: "id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "effect_id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 10, name: "is_spell_id", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "type", kind: "enum", T: () => ["proto.ItemType", ItemType] },
            { no: 9, name: "enchant_type", kind: "enum", T: () => ["proto.EnchantType", EnchantType] },
            { no: 7, name: "stats", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 8, name: "quality", kind: "enum", T: () => ["proto.ItemQuality", ItemQuality] }
        ]);
    }
    create(value) {
        const message = { id: 0, effectId: 0, name: "", isSpellId: false, type: 0, enchantType: 0, stats: [], quality: 0 };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 id */ 1:
                    message.id = reader.int32();
                    break;
                case /* int32 effect_id */ 2:
                    message.effectId = reader.int32();
                    break;
                case /* string name */ 3:
                    message.name = reader.string();
                    break;
                case /* bool is_spell_id */ 10:
                    message.isSpellId = reader.bool();
                    break;
                case /* proto.ItemType type */ 4:
                    message.type = reader.int32();
                    break;
                case /* proto.EnchantType enchant_type */ 9:
                    message.enchantType = reader.int32();
                    break;
                case /* repeated double stats */ 7:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.stats.push(reader.double());
                    else
                        message.stats.push(reader.double());
                    break;
                case /* proto.ItemQuality quality */ 8:
                    message.quality = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* int32 id = 1; */
        if (message.id !== 0)
            writer.tag(1, WireType.Varint).int32(message.id);
        /* int32 effect_id = 2; */
        if (message.effectId !== 0)
            writer.tag(2, WireType.Varint).int32(message.effectId);
        /* string name = 3; */
        if (message.name !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.name);
        /* bool is_spell_id = 10; */
        if (message.isSpellId !== false)
            writer.tag(10, WireType.Varint).bool(message.isSpellId);
        /* proto.ItemType type = 4; */
        if (message.type !== 0)
            writer.tag(4, WireType.Varint).int32(message.type);
        /* proto.EnchantType enchant_type = 9; */
        if (message.enchantType !== 0)
            writer.tag(9, WireType.Varint).int32(message.enchantType);
        /* repeated double stats = 7; */
        if (message.stats.length) {
            writer.tag(7, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.stats.length; i++)
                writer.double(message.stats[i]);
            writer.join();
        }
        /* proto.ItemQuality quality = 8; */
        if (message.quality !== 0)
            writer.tag(8, WireType.Varint).int32(message.quality);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Enchant
 */
export const Enchant = new Enchant$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Gem$Type extends MessageType {
    constructor() {
        super("proto.Gem", [
            { no: 1, name: "id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "stats", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 1 /*ScalarType.DOUBLE*/ },
            { no: 4, name: "color", kind: "enum", T: () => ["proto.GemColor", GemColor] },
            { no: 5, name: "phase", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "quality", kind: "enum", T: () => ["proto.ItemQuality", ItemQuality] },
            { no: 7, name: "unique", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { id: 0, name: "", stats: [], color: 0, phase: 0, quality: 0, unique: false };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 id */ 1:
                    message.id = reader.int32();
                    break;
                case /* string name */ 2:
                    message.name = reader.string();
                    break;
                case /* repeated double stats */ 3:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.stats.push(reader.double());
                    else
                        message.stats.push(reader.double());
                    break;
                case /* proto.GemColor color */ 4:
                    message.color = reader.int32();
                    break;
                case /* int32 phase */ 5:
                    message.phase = reader.int32();
                    break;
                case /* proto.ItemQuality quality */ 6:
                    message.quality = reader.int32();
                    break;
                case /* bool unique */ 7:
                    message.unique = reader.bool();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* int32 id = 1; */
        if (message.id !== 0)
            writer.tag(1, WireType.Varint).int32(message.id);
        /* string name = 2; */
        if (message.name !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.name);
        /* repeated double stats = 3; */
        if (message.stats.length) {
            writer.tag(3, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.stats.length; i++)
                writer.double(message.stats[i]);
            writer.join();
        }
        /* proto.GemColor color = 4; */
        if (message.color !== 0)
            writer.tag(4, WireType.Varint).int32(message.color);
        /* int32 phase = 5; */
        if (message.phase !== 0)
            writer.tag(5, WireType.Varint).int32(message.phase);
        /* proto.ItemQuality quality = 6; */
        if (message.quality !== 0)
            writer.tag(6, WireType.Varint).int32(message.quality);
        /* bool unique = 7; */
        if (message.unique !== false)
            writer.tag(7, WireType.Varint).bool(message.unique);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.Gem
 */
export const Gem = new Gem$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RaidTarget$Type extends MessageType {
    constructor() {
        super("proto.RaidTarget", [
            { no: 1, name: "target_index", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value) {
        const message = { targetIndex: 0 };
        Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 target_index */ 1:
                    message.targetIndex = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* int32 target_index = 1; */
        if (message.targetIndex !== 0)
            writer.tag(1, WireType.Varint).int32(message.targetIndex);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.RaidTarget
 */
export const RaidTarget = new RaidTarget$Type();

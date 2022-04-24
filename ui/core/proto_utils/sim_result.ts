import { ActionMetrics as ActionMetricsProto } from '/tbc/core/proto/api.js';
import { AuraMetrics as AuraMetricsProto } from '/tbc/core/proto/api.js';
import { DistributionMetrics as DistributionMetricsProto } from '/tbc/core/proto/api.js';
import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { EncounterMetrics as EncounterMetricsProto } from '/tbc/core/proto/api.js';
import { Party as PartyProto } from '/tbc/core/proto/api.js';
import { PartyMetrics as PartyMetricsProto } from '/tbc/core/proto/api.js';
import { Player as PlayerProto } from '/tbc/core/proto/api.js';
import { UnitMetrics as UnitMetricsProto } from '/tbc/core/proto/api.js';
import { Raid as RaidProto } from '/tbc/core/proto/api.js';
import { RaidMetrics as RaidMetricsProto } from '/tbc/core/proto/api.js';
import { ResourceMetrics as ResourceMetricsProto, ResourceType } from '/tbc/core/proto/api.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { SimRun } from '/tbc/core/proto/ui.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { classColors } from '/tbc/core/proto_utils/utils.js';
import { getTalentTreeIcon } from '/tbc/core/proto_utils/utils.js';
import { playerToSpec } from '/tbc/core/proto_utils/utils.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';
import { bucket } from '/tbc/core/utils.js';
import { sum } from '/tbc/core/utils.js';

import {
	AuraUptimeLog,
	CastLog,
	DamageDealtLog,
	DpsLog,
	Entity,
	MajorCooldownUsedLog,
	ResourceChangedLogGroup,
	SimLog,
	ThreatLogGroup,
} from './logs_parser.js';

export interface SimResultFilter {
	// Raid index of the player to display, or null for all players.
	player?: number | null;

	// Target index of the target to display, or null for all targets.
	target?: number | null;
}

class SimResultData {
	readonly request: RaidSimRequest;
	readonly result: RaidSimResult;

	constructor(request: RaidSimRequest, result: RaidSimResult) {
		this.request = request;
		this.result = result;
	}

	get iterations() {
		return this.request.simOptions?.iterations || 1;
	}

	get duration() {
		return this.request.encounter?.duration || 1;
	}

	get firstIterationDuration() {
		return this.result.firstIterationDuration || 1;
	}
}

// Holds all the data from a simulation call, and provides helper functions
// for parsing it.
export class SimResult {
	readonly request: RaidSimRequest;
	readonly result: RaidSimResult;

	readonly raidMetrics: RaidMetrics;
	readonly encounterMetrics: EncounterMetrics;
	readonly logs: Array<SimLog>;

	private constructor(request: RaidSimRequest, result: RaidSimResult, raidMetrics: RaidMetrics, encounterMetrics: EncounterMetrics, logs: Array<SimLog>) {
		this.request = request;
		this.result = result;
		this.raidMetrics = raidMetrics;
		this.encounterMetrics = encounterMetrics;
		this.logs = logs;
	}

	getPlayers(filter?: SimResultFilter): Array<UnitMetrics> {
		if (filter?.player || filter?.player === 0) {
			const player = this.getPlayerWithRaidIndex(filter.player);
			return player ? [player] : [];
		} else {
			return this.raidMetrics.parties.map(party => party.players).flat();
		}
	}

	// Returns the first player, regardless of which party / raid slot its in.
	getFirstPlayer(): UnitMetrics | null {
		return this.getPlayers()[0] || null;
	}

	getPlayerWithRaidIndex(raidIndex: number): UnitMetrics | null {
		return this.getPlayers().find(player => player.index == raidIndex) || null;
	}

	getTargets(filter?: SimResultFilter): Array<UnitMetrics> {
		if (filter?.target || filter?.target === 0) {
			const target = this.getTargetWithIndex(filter.target);
			return target ? [target] : [];
		} else {
			return this.encounterMetrics.targets.slice();
		}
	}

	getTargetWithIndex(index: number): UnitMetrics | null {
		return this.getTargets().find(target => target.index == index) || null;
	}

	getDamageMetrics(filter: SimResultFilter): DistributionMetricsProto {
		if (filter.player || filter.player === 0) {
			return this.getPlayerWithRaidIndex(filter.player)?.dps || DistributionMetricsProto.create();
		}

		return this.raidMetrics.dps;
	}

	getActionMetrics(filter: SimResultFilter): Array<ActionMetrics> {
		return ActionMetrics.joinById(this.getPlayers(filter).map(player => player.getPlayerAndPetActions()).flat());
	}

	getSpellMetrics(filter: SimResultFilter): Array<ActionMetrics> {
		return this.getActionMetrics(filter).filter(e => e.hitAttempts != 0 && !e.isMeleeAction)
	}

	getMeleeMetrics(filter: SimResultFilter): Array<ActionMetrics> {
		return this.getActionMetrics(filter).filter(e => e.hitAttempts != 0 && e.isMeleeAction);
	}

	getResourceMetrics(filter: SimResultFilter, resourceType: ResourceType): Array<ResourceMetrics> {
		return ResourceMetrics.joinById(this.getPlayers(filter).map(player => player.resources.filter(resource => resource.type == resourceType)).flat());
	}

	getBuffMetrics(filter: SimResultFilter): Array<AuraMetrics> {
		return AuraMetrics.joinById(this.getPlayers(filter).map(player => player.auras).flat());
	}

	getDebuffMetrics(filter: SimResultFilter): Array<AuraMetrics> {
		return AuraMetrics.joinById(this.getTargets(filter).map(target => target.auras).flat()).filter(aura => aura.uptimePercent != 0);
	}

	toProto(): SimRun {
		return SimRun.create({
			request: this.request,
			result: this.result,
		});
	}

	static async fromProto(proto: SimRun): Promise<SimResult> {
		return SimResult.makeNew(proto.request || RaidSimRequest.create(), proto.result || RaidSimResult.create());
	}

	static async makeNew(request: RaidSimRequest, result: RaidSimResult): Promise<SimResult> {
		const resultData = new SimResultData(request, result);
		const logs = await SimLog.parseAll(result);

		const raidPromise = RaidMetrics.makeNew(resultData, request.raid!, result.raidMetrics!, logs);
		const encounterPromise = EncounterMetrics.makeNew(resultData, request.encounter!, result.encounterMetrics!, logs);

		const raidMetrics = await raidPromise;
		const encounterMetrics = await encounterPromise;

		return new SimResult(request, result, raidMetrics, encounterMetrics, logs);
	}
}

export class RaidMetrics {
	private readonly raid: RaidProto;
	private readonly metrics: RaidMetricsProto;

	readonly dps: DistributionMetricsProto;
	readonly parties: Array<PartyMetrics>;

	private constructor(raid: RaidProto, metrics: RaidMetricsProto, parties: Array<PartyMetrics>) {
		this.raid = raid;
		this.metrics = metrics;
		this.dps = this.metrics.dps!;
		this.parties = parties;
	}

	static async makeNew(resultData: SimResultData, raid: RaidProto, metrics: RaidMetricsProto, logs: Array<SimLog>): Promise<RaidMetrics> {
		const numParties = Math.min(raid.parties.length, metrics.parties.length);

		const parties = await Promise.all(
			[...new Array(numParties).keys()]
				.map(i => PartyMetrics.makeNew(
					resultData,
					raid.parties[i],
					metrics.parties[i],
					i,
					logs)));

		return new RaidMetrics(raid, metrics, parties);
	}
}

export class PartyMetrics {
	private readonly party: PartyProto;
	private readonly metrics: PartyMetricsProto;

	readonly partyIndex: number;
	readonly dps: DistributionMetricsProto;
	readonly players: Array<UnitMetrics>;

	private constructor(party: PartyProto, metrics: PartyMetricsProto, partyIndex: number, players: Array<UnitMetrics>) {
		this.party = party;
		this.metrics = metrics;
		this.partyIndex = partyIndex;
		this.dps = this.metrics.dps!;
		this.players = players;
	}

	static async makeNew(resultData: SimResultData, party: PartyProto, metrics: PartyMetricsProto, partyIndex: number, logs: Array<SimLog>): Promise<PartyMetrics> {
		const numPlayers = Math.min(party.players.length, metrics.players.length);
		const players = await Promise.all(
			[...new Array(numPlayers).keys()]
				.filter(i => party.players[i].class != Class.ClassUnknown)
				.map(i => UnitMetrics.makeNewPlayer(
					resultData,
					party.players[i],
					metrics.players[i],
					partyIndex * 5 + i,
					false,
					logs)));

		return new PartyMetrics(party, metrics, partyIndex, players);
	}
}

export class UnitMetrics {
	// If this Unit is a pet, player is the owner. If it's a target, player is null.
	private readonly player: PlayerProto | null;
	private readonly target: TargetProto | null;
	private readonly metrics: UnitMetricsProto;

	readonly index: number;
	readonly name: string;
	readonly spec: Spec;
	readonly petActionId: ActionId | null;
	readonly iconUrl: string;
	readonly classColor: string;
	readonly dps: DistributionMetricsProto;
	readonly tps: DistributionMetricsProto;
	readonly actions: Array<ActionMetrics>;
	readonly auras: Array<AuraMetrics>;
	readonly resources: Array<ResourceMetrics>;
	readonly pets: Array<UnitMetrics>;
	private readonly iterations: number;
	private readonly duration: number;

	readonly logs: Array<SimLog>;
	readonly damageDealtLogs: Array<DamageDealtLog>;
	readonly groupedResourceLogs: Record<ResourceType, Array<ResourceChangedLogGroup>>;
	readonly dpsLogs: Array<DpsLog>;
	readonly auraUptimeLogs: Array<AuraUptimeLog>;
	readonly majorCooldownLogs: Array<MajorCooldownUsedLog>;
	readonly castLogs: Array<CastLog>;
	readonly threatLogs: Array<ThreatLogGroup>;

	// Aura uptime logs, filtered to include only auras that correspond to a
	// major cooldown.
	readonly majorCooldownAuraUptimeLogs: Array<AuraUptimeLog>;

	private constructor(
		player: PlayerProto | null,
		target: TargetProto | null,
		petActionId: ActionId | null,
		metrics: UnitMetricsProto,
		index: number,
		actions: Array<ActionMetrics>,
		auras: Array<AuraMetrics>,
		resources: Array<ResourceMetrics>,
		pets: Array<UnitMetrics>,
		logs: Array<SimLog>,
		resultData: SimResultData) {
		this.player = player;
		this.target = target;
		this.metrics = metrics;

		this.index = index;
		this.name = metrics.name;
		this.spec = player ? playerToSpec(player) : 0;
		this.petActionId = petActionId;
		this.iconUrl = player ? getTalentTreeIcon(this.spec, player.talentsString) : '';
		this.classColor = classColors[specToClass[this.spec]];
		this.dps = this.metrics.dps!;
		this.tps = this.metrics.threat!;
		this.actions = actions;
		this.auras = auras;
		this.resources = resources;
		this.pets = pets;
		this.logs = logs;
		this.iterations = resultData.iterations;
		this.duration = resultData.duration;

		this.damageDealtLogs = this.logs.filter((log): log is DamageDealtLog => log.isDamageDealt());
		this.dpsLogs = DpsLog.fromLogs(this.damageDealtLogs);
		this.castLogs = CastLog.fromLogs(this.logs);
		this.threatLogs = ThreatLogGroup.fromLogs(this.logs);

		this.auraUptimeLogs = AuraUptimeLog.fromLogs(this.logs, new Entity(this.name, '', this.index, false, this.isPet), resultData.firstIterationDuration);
		this.majorCooldownLogs = this.logs.filter((log): log is MajorCooldownUsedLog => log.isMajorCooldownUsed());

		this.groupedResourceLogs = ResourceChangedLogGroup.fromLogs(this.logs);
		AuraUptimeLog.populateActiveAuras(this.dpsLogs, this.auraUptimeLogs);
		AuraUptimeLog.populateActiveAuras(this.groupedResourceLogs[ResourceType.ResourceTypeMana], this.auraUptimeLogs);

		this.majorCooldownAuraUptimeLogs = this.auraUptimeLogs.filter(auraLog => this.majorCooldownLogs.find(mcdLog => mcdLog.actionId!.equals(auraLog.actionId!)));
	}

	get label() {
		if (this.target == null) {
			return `${this.name} (#${this.index + 1})`;
		} else {
			return this.name;
		}
	}

	get isPet() {
		return this.petActionId != null;
	}

	get maxThreat() {
		return this.threatLogs[this.threatLogs.length - 1]?.threatAfter || 0;
	}

	get secondsOomAvg() {
		return this.metrics.secondsOomAvg
	}

	get totalDamage() {
		return this.dps.avg * this.duration;
	}

	getPlayerAndPetActions(): Array<ActionMetrics> {
		return this.actions.concat(this.pets.map(pet => pet.getPlayerAndPetActions()).flat());
	}

	getMeleeActions(): Array<ActionMetrics> {
		return this.actions.filter(e => e.hitAttempts != 0 && e.isMeleeAction);
	}

	getSpellActions(): Array<ActionMetrics> {
		return this.actions.filter(e => e.hitAttempts != 0 && !e.isMeleeAction)
	}

	getResourceMetrics(resourceType: ResourceType): Array<ResourceMetrics> {
		return this.resources.filter(resource => resource.type == resourceType);
	}

	static async makeNewPlayer(resultData: SimResultData, player: PlayerProto, metrics: UnitMetricsProto, raidIndex: number, isPet: boolean, logs: Array<SimLog>): Promise<UnitMetrics> {
		const playerLogs = logs.filter(log => log.source && (!log.source.isTarget && (isPet == log.source.isPet) && log.source.index == raidIndex));

		const actionsPromise = Promise.all(metrics.actions.map(actionMetrics => ActionMetrics.makeNew(null, resultData, actionMetrics, raidIndex)));
		const aurasPromise = Promise.all(metrics.auras.map(auraMetrics => AuraMetrics.makeNew(null, resultData, auraMetrics, raidIndex)));
		const resourcesPromise = Promise.all(metrics.resources.map(resourceMetrics => ResourceMetrics.makeNew(null, resultData, resourceMetrics, raidIndex)));
		const petsPromise = Promise.all(metrics.pets.map(petMetrics => UnitMetrics.makeNewPlayer(resultData, player, petMetrics, raidIndex, true, playerLogs)));

		let petIdPromise: Promise<ActionId | null> = Promise.resolve(null);
		if (isPet) {
			petIdPromise = ActionId.fromPetName(metrics.name).fill(raidIndex);
		}

		const actions = await actionsPromise;
		const auras = await aurasPromise;
		const resources = await resourcesPromise;
		const pets = await petsPromise;
		const petActionId = await petIdPromise;

		const playerMetrics = new UnitMetrics(player, null, petActionId, metrics, raidIndex, actions, auras, resources, pets, playerLogs, resultData);
		actions.forEach(action => action.unit = playerMetrics);
		auras.forEach(aura => aura.unit = playerMetrics);
		resources.forEach(resource => resource.unit = playerMetrics);
		return playerMetrics;
	}

	static async makeNewTarget(resultData: SimResultData, target: TargetProto, metrics: UnitMetricsProto, index: number, logs: Array<SimLog>): Promise<UnitMetrics> {
		const targetLogs = logs.filter(log => log.source && (log.source.isTarget && log.source.index == index));
		const auras = await Promise.all(metrics.auras.map(auraMetrics => AuraMetrics.makeNew(null, resultData, auraMetrics)));
		return new UnitMetrics(null, target, null, metrics, index, [], auras, [], [], targetLogs, resultData);
	}
}

export class EncounterMetrics {
	private readonly encounter: EncounterProto;
	private readonly metrics: EncounterMetricsProto;

	readonly targets: Array<UnitMetrics>;

	private constructor(encounter: EncounterProto, metrics: EncounterMetricsProto, targets: Array<UnitMetrics>) {
		this.encounter = encounter;
		this.metrics = metrics;
		this.targets = targets;
	}

	static async makeNew(resultData: SimResultData, encounter: EncounterProto, metrics: EncounterMetricsProto, logs: Array<SimLog>): Promise<EncounterMetrics> {
		const numTargets = Math.min(encounter.targets.length, metrics.targets.length);
		const targets = await Promise.all(
			[...new Array(numTargets).keys()]
				.map(i => UnitMetrics.makeNewTarget(
					resultData,
					encounter.targets[i],
					metrics.targets[i],
					i,
					logs)));

		return new EncounterMetrics(encounter, metrics, targets);
	}

	get durationSeconds() {
		return this.encounter.duration;
	}
}

export class AuraMetrics {
	unit: UnitMetrics | null;
	readonly actionId: ActionId;
	readonly name: string;
	readonly iconUrl: string;
	private readonly resultData: SimResultData;
	private readonly iterations: number;
	private readonly duration: number;
	private readonly data: AuraMetricsProto;

	private constructor(unit: UnitMetrics | null, actionId: ActionId, data: AuraMetricsProto, resultData: SimResultData) {
		this.unit = unit;
		this.actionId = actionId;
		this.name = actionId.name;
		this.iconUrl = actionId.iconUrl;
		this.data = data;
		this.resultData = resultData;
		this.iterations = resultData.iterations;
		this.duration = resultData.duration;
	}

	get uptimePercent() {
		return this.data.uptimeSecondsAvg / this.duration * 100;
	}

	static async makeNew(unit: UnitMetrics | null, resultData: SimResultData, auraMetrics: AuraMetricsProto, playerIndex?: number): Promise<AuraMetrics> {
		const actionId = await ActionId.fromProto(auraMetrics.id!).fill(playerIndex);
		return new AuraMetrics(unit, actionId, auraMetrics, resultData);
	}

	// Merges an array of metrics into a single metrics.
	static merge(auras: Array<AuraMetrics>, removeTag?: boolean, actionIdOverride?: ActionId): AuraMetrics {
		const firstAura = auras[0];
		const unit = auras.every(aura => aura.unit == firstAura.unit) ? firstAura.unit : null;
		let actionId = actionIdOverride || firstAura.actionId;
		if (removeTag) {
			actionId = actionId.withoutTag();
		}
		return new AuraMetrics(
			unit,
			actionId,
			AuraMetricsProto.create({
				uptimeSecondsAvg: Math.max(...auras.map(a => a.data.uptimeSecondsAvg)),
			}),
			firstAura.resultData);
	}

	// Groups similar metrics, i.e. metrics with the same item/spell/other ID but
	// different tags, and returns them as separate arrays.
	static groupById(auras: Array<AuraMetrics>, useTag?: boolean): Array<Array<AuraMetrics>> {
		if (useTag) {
			return Object.values(bucket(auras, aura => aura.actionId.toString()));
		} else {
			return Object.values(bucket(auras, aura => aura.actionId.toStringIgnoringTag()));
		}
	}

	// Merges aura metrics that have the same name/ID, adding their stats together.
	static joinById(auras: Array<AuraMetrics>, useTag?: boolean): Array<AuraMetrics> {
		return AuraMetrics.groupById(auras, useTag).map(aurasToJoin => AuraMetrics.merge(aurasToJoin));
	}
};

export class ResourceMetrics {
	unit: UnitMetrics | null;
	readonly actionId: ActionId;
	readonly name: string;
	readonly iconUrl: string;
	readonly type: ResourceType;
	private readonly resultData: SimResultData;
	private readonly iterations: number;
	private readonly duration: number;
	private readonly data: ResourceMetricsProto;

	private constructor(unit: UnitMetrics | null, actionId: ActionId, data: ResourceMetricsProto, resultData: SimResultData) {
		this.unit = unit;
		this.actionId = actionId;
		this.name = actionId.name;
		this.iconUrl = actionId.iconUrl;
		this.type = data.type;
		this.resultData = resultData;
		this.iterations = resultData.iterations;
		this.duration = resultData.duration;
		this.data = data;
	}

	get events() {
		return this.data.events / this.iterations;
	}

	get gain() {
		return this.data.gain / this.iterations;
	}

	get gainPerSecond() {
		return this.data.gain / this.iterations / this.duration;
	}

	get avgGain() {
		return this.data.gain / this.data.events;
	}

	get wastedGain() {
		return (this.data.gain - this.data.actualGain) / this.iterations;
	}

	static async makeNew(unit: UnitMetrics | null, resultData: SimResultData, resourceMetrics: ResourceMetricsProto, playerIndex?: number): Promise<ResourceMetrics> {
		const actionId = await ActionId.fromProto(resourceMetrics.id!).fill(playerIndex);
		return new ResourceMetrics(unit, actionId, resourceMetrics, resultData);
	}

	// Merges an array of metrics into a single metrics.
	static merge(resources: Array<ResourceMetrics>, removeTag?: boolean, actionIdOverride?: ActionId): ResourceMetrics {
		const firstResource = resources[0];
		const unit = resources.every(resource => resource.unit == firstResource.unit) ? firstResource.unit : null;
		let actionId = actionIdOverride || firstResource.actionId;
		if (removeTag) {
			actionId = actionId.withoutTag();
		}
		return new ResourceMetrics(
			unit,
			actionId,
			ResourceMetricsProto.create({
				events: sum(resources.map(a => a.data.events)),
				gain: sum(resources.map(a => a.data.gain)),
				actualGain: sum(resources.map(a => a.data.actualGain)),
			}),
			firstResource.resultData);
	}

	// Groups similar metrics, i.e. metrics with the same item/spell/other ID but
	// different tags, and returns them as separate arrays.
	static groupById(resources: Array<ResourceMetrics>, useTag?: boolean): Array<Array<ResourceMetrics>> {
		if (useTag) {
			return Object.values(bucket(resources, resource => resource.actionId.toString()));
		} else {
			return Object.values(bucket(resources, resource => resource.actionId.toStringIgnoringTag()));
		}
	}

	// Merges resource metrics that have the same name/ID, adding their stats together.
	static joinById(resources: Array<ResourceMetrics>, useTag?: boolean): Array<ResourceMetrics> {
		return ResourceMetrics.groupById(resources, useTag).map(resourcesToJoin => ResourceMetrics.merge(resourcesToJoin));
	}
};

// Manages the metrics for a single unit action (e.g. Lightning Bolt).
export class ActionMetrics {
	unit: UnitMetrics | null;
	readonly actionId: ActionId;
	readonly name: string;
	readonly iconUrl: string;
	private readonly resultData: SimResultData;
	private readonly iterations: number;
	private readonly duration: number;
	private readonly data: ActionMetricsProto;

	private constructor(unit: UnitMetrics | null, actionId: ActionId, data: ActionMetricsProto, resultData: SimResultData) {
		this.unit = unit;
		this.actionId = actionId;
		this.name = actionId.name;
		this.iconUrl = actionId.iconUrl;
		this.resultData = resultData;
		this.iterations = resultData.iterations;
		this.duration = resultData.duration;
		this.data = data;
	}

	get isMeleeAction() {
		return this.data.isMelee;
	}

	get damage() {
		return this.data.damage;
	}

	get dps() {
		return this.data.damage / this.iterations / this.duration;
	}

	get tps() {
		return this.data.threat / this.iterations / this.duration;
	}

	get casts() {
		return this.data.casts / this.iterations;
	}

	get castsPerMinute() {
		return this.data.casts / this.iterations / (this.duration / 60);
	}

	get avgCast() {
		return this.data.damage / this.data.casts;
	}

	get avgCastThreat() {
		return this.data.threat / this.data.casts;
	}

	private get landedHitsRaw() {
		return this.data.hits + this.data.crits + this.data.blocks + this.data.glances;
	}
	get landedHits() {
		return this.landedHitsRaw / this.iterations;
	}

	get hitAttempts() {
		return this.data.misses
			+ this.data.dodges
			+ this.data.parries
			+ this.data.blocks
			+ this.data.glances
			+ this.data.crits
			+ this.data.hits;
	}

	get avgHit() {
		return this.data.damage / this.landedHitsRaw;
	}

	get avgHitThreat() {
		return this.data.threat / this.landedHitsRaw;
	}

	get critPercent() {
		return (this.data.crits / this.hitAttempts) * 100;
	}

	get misses() {
		return this.data.misses / this.iterations;
	}

	get missPercent() {
		return (this.data.misses / this.hitAttempts) * 100;
	}

	get dodges() {
		return this.data.dodges / this.iterations;
	}

	get dodgePercent() {
		return (this.data.dodges / this.hitAttempts) * 100;
	}

	get parries() {
		return this.data.parries / this.iterations;
	}

	get parryPercent() {
		return (this.data.parries / this.hitAttempts) * 100;
	}

	get blocks() {
		return this.data.blocks / this.iterations;
	}

	get blockPercent() {
		return (this.data.blocks / this.hitAttempts) * 100;
	}

	get glances() {
		return this.data.glances / this.iterations;
	}

	get glancePercent() {
		return (this.data.glances / this.hitAttempts) * 100;
	}

	static async makeNew(unit: UnitMetrics | null, resultData: SimResultData, actionMetrics: ActionMetricsProto, playerIndex?: number): Promise<ActionMetrics> {
		const actionId = await ActionId.fromProto(actionMetrics.id!).fill(playerIndex);
		return new ActionMetrics(unit, actionId, actionMetrics, resultData);
	}

	// Merges an array of metrics into a single metric.
	static merge(actions: Array<ActionMetrics>, removeTag?: boolean, actionIdOverride?: ActionId): ActionMetrics {
		const firstAction = actions[0];
		const unit = actions.every(action => action.unit == firstAction.unit) ? firstAction.unit : null;
		let actionId = actionIdOverride || firstAction.actionId;
		if (removeTag) {
			actionId = actionId.withoutTag();
		}
		return new ActionMetrics(
			unit,
			actionId,
			ActionMetricsProto.create({
				isMelee: firstAction.isMeleeAction,
				casts: sum(actions.map(a => a.data.casts)),
				hits: sum(actions.map(a => a.data.hits)),
				crits: sum(actions.map(a => a.data.crits)),
				misses: sum(actions.map(a => a.data.misses)),
				dodges: sum(actions.map(a => a.data.dodges)),
				parries: sum(actions.map(a => a.data.parries)),
				blocks: sum(actions.map(a => a.data.blocks)),
				glances: sum(actions.map(a => a.data.glances)),
				damage: sum(actions.map(a => a.data.damage)),
				threat: sum(actions.map(a => a.data.threat)),
			}),
			firstAction.resultData);
	}

	// Groups similar metrics, i.e. metrics with the same item/spell/other ID but
	// different tags, and returns them as separate arrays.
	static groupById(actions: Array<ActionMetrics>, useTag?: boolean): Array<Array<ActionMetrics>> {
		if (useTag) {
			return Object.values(bucket(actions, action => action.actionId.toString()));
		} else {
			return Object.values(bucket(actions, action => action.actionId.toStringIgnoringTag()));
		}
	}

	// Merges action metrics that have the same name/ID, adding their stats together.
	static joinById(actions: Array<ActionMetrics>, useTag?: boolean): Array<ActionMetrics> {
		return ActionMetrics.groupById(actions, useTag).map(actionsToJoin => ActionMetrics.merge(actionsToJoin));
	}
}

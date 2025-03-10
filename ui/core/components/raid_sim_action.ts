import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { Raid as RaidProto } from '/tbc/core/proto/api.js';
import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { SimResult } from '/tbc/core/proto_utils/sim_result.js';
import { SimUI } from '/tbc/core/sim_ui.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';

declare var tippy: any;

export function addRaidSimAction(simUI: SimUI): RaidSimResultsManager {
	simUI.addAction('DPS', 'dps-action', async () => {
		simUI.setResultsPending();
		try {
			const result = await simUI.sim.runRaidSim(TypedEvent.nextEventID());
		} catch (e) {
			simUI.hideAllResults();
			alert(e);
		}
	});

	const resultsManager = new RaidSimResultsManager(simUI);
	simUI.sim.simResultEmitter.on((eventID, simResult) => {
		resultsManager.setSimResult(eventID, simResult);
	});
	return resultsManager;
}

export type ReferenceData = {
	simResult: SimResult,
	settings: any,
	raidProto: RaidProto,
	encounterProto: EncounterProto,
};

export class RaidSimResultsManager {
	readonly currentChangeEmitter: TypedEvent<void> = new TypedEvent<void>();
	readonly referenceChangeEmitter: TypedEvent<void> = new TypedEvent<void>();

	readonly changeEmitter: TypedEvent<void> = new TypedEvent<void>();

	private readonly simUI: SimUI;

	private currentData: ReferenceData | null = null;
	private referenceData: ReferenceData | null = null;

  constructor(simUI: SimUI) {
		this.simUI = simUI;

		[
      this.currentChangeEmitter,
      this.referenceChangeEmitter,
    ].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
  }

  setSimResult(eventID: EventID, simResult: SimResult) {
		this.currentData = {
			simResult: simResult,
			settings: this.simUI.sim.toJson(),
			raidProto: RaidProto.clone(simResult.request.raid || RaidProto.create()),
			encounterProto: EncounterProto.clone(simResult.request.encounter || EncounterProto.create()),
		};
		this.currentChangeEmitter.emit(eventID);

		const dpsMetrics = simResult.raidMetrics.dps;
		this.simUI.setResultsContent(`
      <div class="results-sim">
				<div class="results-sim-dps">
					<span class="results-sim-dps-avg">${dpsMetrics.avg.toFixed(2)}</span>
					<span class="results-sim-dps-stdev">${dpsMetrics.stdev.toFixed(2)}</span>
				</div>
				<div class="results-sim-reference">
					<span class="results-sim-set-reference fa fa-bookmark"></span>
					<div class="results-sim-reference-bar">
						<span class="results-sim-reference-diff"></span>
						<span class="results-sim-reference-text"> vs. reference</span>
						<span class="results-sim-reference-swap fa fa-retweet"></span>
						<span class="results-sim-reference-delete fa fa-times"></span>
					</div>
				</div>
      </div>
    `);

    const simReferenceElem = this.simUI.resultsContentElem.getElementsByClassName('results-sim-reference')[0] as HTMLDivElement;
    const simReferenceDiffElem = this.simUI.resultsContentElem.getElementsByClassName('results-sim-reference-diff')[0] as HTMLSpanElement;

    const simReferenceSetButton = this.simUI.resultsContentElem.getElementsByClassName('results-sim-set-reference')[0] as HTMLSpanElement;
		simReferenceSetButton.addEventListener('click', event => {
			this.referenceData = this.currentData;
			this.referenceChangeEmitter.emit(TypedEvent.nextEventID());
			this.updateReference();
		});
		tippy(simReferenceSetButton, {
			'content': 'Use as reference',
			'allowHTML': true,
		});

    const simReferenceSwapButton = this.simUI.resultsContentElem.getElementsByClassName('results-sim-reference-swap')[0] as HTMLSpanElement;
		simReferenceSwapButton.addEventListener('click', event => {
			TypedEvent.freezeAllAndDo(() => {
				if (this.currentData && this.referenceData) {
					const swapEventID = TypedEvent.nextEventID();
					const tmpData = this.currentData;
					this.currentData = this.referenceData;
					this.referenceData = tmpData;

					this.simUI.sim.raid.fromProto(swapEventID, this.currentData.raidProto);
					this.simUI.sim.encounter.fromProto(swapEventID, this.currentData.encounterProto);
					this.setSimResult(swapEventID, this.currentData.simResult);

					this.referenceChangeEmitter.emit(swapEventID);
					this.updateReference();
				}
			});
		});
		tippy(simReferenceSwapButton, {
			'content': 'Swap reference with current',
			'allowHTML': true,
		});

    const simReferenceDeleteButton = this.simUI.resultsContentElem.getElementsByClassName('results-sim-reference-delete')[0] as HTMLSpanElement;
		simReferenceDeleteButton.addEventListener('click', event => {
			this.referenceData = null;
			this.referenceChangeEmitter.emit(TypedEvent.nextEventID());
			this.updateReference();
		});
		tippy(simReferenceDeleteButton, {
			'content': 'Remove reference',
			'allowHTML': true,
		});

		this.updateReference();
  }

	private updateReference() {
    const simReferenceElem = this.simUI.resultsContentElem.getElementsByClassName('results-sim-reference')[0] as HTMLDivElement;
    const simReferenceDiffElem = this.simUI.resultsContentElem.getElementsByClassName('results-sim-reference-diff')[0] as HTMLSpanElement;

		if (!this.referenceData || !this.currentData) {
			simReferenceElem.classList.remove('has-reference');
			return;
		}
		simReferenceElem.classList.add('has-reference');

		const currentDpsMetrics = this.currentData.simResult.raidMetrics.dps;
		const referenceDpsMetrics = this.referenceData.simResult.raidMetrics.dps;
		const delta = currentDpsMetrics.avg - referenceDpsMetrics.avg;
		const deltaStr = delta.toFixed(2);
		if (delta >= 0) {
			simReferenceDiffElem.textContent = '+' + deltaStr;
			simReferenceDiffElem.classList.remove('negative');
			simReferenceDiffElem.classList.add('positive');
		} else {
			simReferenceDiffElem.textContent = '' + deltaStr;
			simReferenceDiffElem.classList.remove('positive');
			simReferenceDiffElem.classList.add('negative');
		}
	}

	getCurrentData(): ReferenceData | null {
		if (this.currentData == null) {
			return null;
		}

		// Defensive copy.
		return {
			simResult: this.currentData.simResult,
			settings: JSON.parse(JSON.stringify(this.currentData.settings)),
			raidProto: this.currentData.raidProto,
			encounterProto: this.currentData.encounterProto,
		};
	}

	getReferenceData(): ReferenceData | null {
		if (this.referenceData == null) {
			return null;
		}

		// Defensive copy.
		return {
			simResult: this.referenceData.simResult,
			settings: JSON.parse(JSON.stringify(this.referenceData.settings)),
			raidProto: this.referenceData.raidProto,
			encounterProto: this.referenceData.encounterProto,
		};
	}
}

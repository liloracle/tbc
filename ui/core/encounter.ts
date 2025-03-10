import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { Target } from '/tbc/core/target.js';

import { Sim } from './sim.js';
import { EventID, TypedEvent } from './typed_event.js';

// Manages all the settings for an Encounter.
export class Encounter {
	private readonly sim: Sim;

  private duration: number = 300;
  private numTargets: number = 1;
  private executeProportion: number = 0.2;
	readonly primaryTarget: Target;

  readonly durationChangeEmitter = new TypedEvent<void>();
  readonly numTargetsChangeEmitter = new TypedEvent<void>();
  readonly executeProportionChangeEmitter = new TypedEvent<void>();

  // Emits when any of the above emitters emit.
  readonly changeEmitter = new TypedEvent<void>();

  constructor(sim: Sim) {
		this.sim = sim;
		this.primaryTarget = new Target(sim);

    [
      this.durationChangeEmitter,
      this.numTargetsChangeEmitter,
      this.executeProportionChangeEmitter,
      this.primaryTarget.changeEmitter,
    ].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
  }
  
  getDuration(): number {
    return this.duration;
  }
  setDuration(eventID: EventID, newDuration: number) {
    if (newDuration == this.duration)
			return;

		this.duration = newDuration;
		this.durationChangeEmitter.emit(eventID);
  }
  
  getExecuteProportion(): number {
    return this.executeProportion;
  }
  setExecuteProportion(eventID: EventID, newExecuteProportion: number) {
    if (newExecuteProportion == this.executeProportion)
			return;

		this.executeProportion = newExecuteProportion;
		this.executeProportionChangeEmitter.emit(eventID);
  }
  
  getNumTargets(): number {
    return this.numTargets;
  }
  setNumTargets(eventID: EventID, newNumTargets: number) {
    if (newNumTargets == this.numTargets)
			return;

		this.numTargets = newNumTargets;
		this.numTargetsChangeEmitter.emit(eventID);
  }

	toProto(): EncounterProto {
		const numTargets = Math.max(1, this.numTargets);
		const targetProtos = [];
		for (let i = 0; i < numTargets; i++) {
			targetProtos.push(this.primaryTarget.toProto());
		}

		return EncounterProto.create({
			duration: this.duration,
			executeProportion: this.executeProportion,
			targets: targetProtos,
		});
	}

	fromProto(eventID: EventID, proto: EncounterProto) {
		TypedEvent.freezeAllAndDo(() => {
			this.setDuration(eventID, proto.duration);
			this.setExecuteProportion(eventID, proto.executeProportion);
			this.setNumTargets(eventID, proto.targets.length);

			if (proto.targets.length > 0) {
				this.primaryTarget.fromProto(eventID, proto.targets[0]);
			}
		});
	}

  // Returns JSON representing all the current values.
  toJson(): Object {
    return {
			'duration': this.getDuration(),
			'executeProportion': this.getExecuteProportion(),
			'numTargets': this.getNumTargets(),
			'primaryTarget': this.primaryTarget.toJson(),
    };
  }

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(eventID: EventID, obj: any) {
		TypedEvent.freezeAllAndDo(() => {
			const parsedDuration = parseInt(obj['duration']);
			if (!isNaN(parsedDuration) && parsedDuration != 0) {
				this.setDuration(eventID, parsedDuration);
			}

			const parsedExecuteProportion = parseInt(obj['executeProportion']);
			if (!isNaN(parsedExecuteProportion) && parsedExecuteProportion != 0) {
				this.setExecuteProportion(eventID, parsedExecuteProportion);
			}

			const parsedNumTargets = parseInt(obj['numTargets']);
			if (!isNaN(parsedNumTargets) && parsedNumTargets != 0) {
				this.setNumTargets(eventID, parsedNumTargets);
			}

			try {
				this.primaryTarget.fromJson(eventID, obj['primaryTarget']);
			} catch (e) {
				console.warn('Failed to parse debuffs: ' + e);
			}
		});
  }
}

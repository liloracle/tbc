import { SimUI } from '/tbc/core/sim_ui.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';

import { Component } from './component.js';

export class LogRunner extends Component {
  constructor(parent: HTMLElement, simUI: SimUI) {
    super(parent, 'log-runner-root');

		const controlBar = document.createElement('div');
		controlBar.classList.add('log-runner-control-bar');
		this.rootElem.appendChild(controlBar);

    const simButton = document.createElement('button');
    simButton.classList.add('log-runner-button', 'sim-button');
    simButton.textContent = 'SIM 1 ITERATION';
    controlBar.appendChild(simButton);

		const logsDiv = document.createElement('div');
		logsDiv.classList.add('log-runner-logs');
		this.rootElem.appendChild(logsDiv);

    simButton.addEventListener('click', async () => {
			simUI.setResultsPending();
			try {
				const result = await simUI.sim.runRaidSimWithLogs(TypedEvent.nextEventID());
			} catch (e) {
				simUI.hideAllResults();
				alert(e);
			}
    });

		simUI.sim.simResultEmitter.on((eventID, simResult) => {
			const logs = simResult.getLogs();
			logsDiv.textContent = '';
			logs.forEach(log => {
				const lineElem = document.createElement('span');
				lineElem.textContent = log;
				logsDiv.appendChild(lineElem);
			});
		});
  }
}

/**
 * @module bugs-central-sk
 * @description <h2><code>bugs-central-sk</code></h2>
 *
 * <p>
 *   Displays a table with issue counts for client+source+queries. Also
 *   displays that information in charts.
 * </p>
 *
 */

import { define } from 'elements-sk/define';
import { html, TemplateResult } from 'lit-html';
import { errorMessage } from 'elements-sk/errorMessage';
import { jsonOrThrow } from 'common-sk/modules/jsonOrThrow';
import { stateReflector } from 'common-sk/modules/stateReflector';

import 'elements-sk/spinner-sk';
import '../bugs-chart-sk';

import { HintableObject } from 'common-sk/modules/hintable';
import { ElementSk } from '../../../infra-sk/modules/ElementSk';

const CLIENT_KEY_DELIMITER = ' > ';

const SKIA_SLO_DOC = 'https://docs.google.com/document/d/1OgpX1KDDq3YkHzRJjqRHSPJ9CJ8hH0RTvMAApKVxwm8/edit';

function getClientKey(c: string, s: string, q: string) {
  if (!c) {
    return '';
  } if (!s) {
    return `${c}`;
  } if (!q) {
    return `${c}${CLIENT_KEY_DELIMITER}${s}`;
  }
  return `${c}${CLIENT_KEY_DELIMITER}${s}${CLIENT_KEY_DELIMITER}${q}`;
}

function breakupClientKey(clientKey: string) {
  const ret = {
    client: '',
    source: '',
    query: '',
  };
  const tokens = clientKey.split(CLIENT_KEY_DELIMITER);
  if (tokens.length === 0) {
    // Leave all values blank.
  } else if (tokens.length === 1) {
    ret.client = tokens[0];
  } else if (tokens.length === 2) {
    ret.client = tokens[0];
    ret.source = tokens[1];
  } else if (tokens.length === 3) {
    ret.client = tokens[0];
    ret.source = tokens[1];
    ret.query = tokens[2];
  }
  return ret;
}

// TODO(rmistry): Generate this using go2ts.
declare interface CountsData {
  open_count: number,
  unassigned_count: number,
  untriaged_count: number,

  p0_count: number,
  p1_count: number,
  p2_count: number,
  p3_count: number,
  p4_count: number,
  p5_count: number,
  p6_count: number,

  p0_slo_count: number,
  p1_slo_count: number,
  p2_slo_count: number,
  p3_slo_count: number,

  query_link: string,
  untriaged_query_link: string,
}

// ChartsData is directly fed to @google-web-components/google-chart.
// Example format: '[["Month", "Days"], ["Jan", 31], ["Feb", 28], ["Mar", 31]]'
declare interface ChartsData {
  open_data: string,
  slo_data: string,
  untriaged_data: string,
}

// State is reflected to the URL via stateReflector.
declare interface State {
  client: string,
  source: string,
  query: string,
}

declare interface ClientsResponse {
  clients: Record<string, Record<string, Record<string, boolean>>>,
}

export class BugsCentralSk extends ElementSk {
  public state: State = {
    client: '',
    source: '',
    query: '',
  };

  private clients_to_counts: Record<string, CountsData> = {};

  private clients_map: Record<string, Record<string, Record<string, boolean>>> = {};

  private open_chart_data: string = '';

  private slo_chart_data: string = '';

  private untriaged_chart_data: string = '';

  private updatingData: boolean = true;

  constructor() {
    super(BugsCentralSk.template);
  }

  private static template = (el: BugsCentralSk) => html`
  <h2>${el.getTitle()}</h2>
  <spinner-sk ?active=${el.updatingData}></spinner-sk>
  <br/><br/>
  <div class="charts-container">
    <div class="chart-div">
      <bugs-chart-sk chart_type='open'
                     chart_title='Bug Count'
                     data=${el.open_chart_data}>
      </bugs-chart-sk>
    </div>
    <div class="chart-div">
      <bugs-chart-sk chart_type='slo'
                     chart_title='SLO Violations'
                     data=${el.slo_chart_data}>
      </bugs-chart-sk>
    </div>
    <div class="chart-div">
      <bugs-chart-sk chart_type='untriaged'
                     chart_title='Untriaged Bugs'
                     data=${el.untriaged_chart_data}>
      </bugs-chart-sk>
    </div>
  </div>
  <br/><br/>
  ${el.displayClientsTable()}
  `;


  async connectedCallback(): Promise<void> {
    super.connectedCallback();

    // Populate map of clients to sources to queries.
    await this.doImpl('/_/get_clients_sources_queries', {}, async (json: ClientsResponse) => {
      this.clients_map = json.clients;
    });

    // From this point on reflect the state to the URL.
    this.startStateReflector();

    this.updatingData = true;
    this._render();
    await this.populateDataAndRender();
    this.updatingData = false;
    this._render();
  }

  // Call this anytime something in private state is changed. Will be replaced
  // with the real function once stateReflector has been setup.
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  private stateHasChanged = () => {};

  private displayClientsTable(): TemplateResult {
    return html`
    <table class=client-counts>
      <colgroup>
        <col span="1" style="width: 50%">
        <col span="1" style="width: 10%">
        <col span="1" style="width: 10%">
        <col span="1" style="width: 10%">
        <col span="1" style="width: 10%">
        <col span="1" style="width: 10%">
      </colgroup>
      <tr>
        <th>Client</th>
        <th>P0/P1 <span class="small">[<a href="${SKIA_SLO_DOC}">SLO</a>]</span></th>
        <th>P2 <span class="small">[<a href="${SKIA_SLO_DOC}">SLO</a>]</span></th>
        <th>P3+ <span class="small">[<a href="${SKIA_SLO_DOC}">SLO</a>]</span></th>
        <th>Untriaged</th>
        <th>Total</th>
      </tr>
       ${this.displayClientsRows()}
    </table>
  `;
  }

  private getTitle(): TemplateResult {
    if (!this.state.client) {
      return html`Displaying all clients`;
    }
    const clientKey = getClientKey(this.state.client, this.state.source, this.state.query);
    const clientCounts = this.clients_to_counts[clientKey];
    if (clientCounts && clientCounts.query_link) {
      return html`
        ${clientKey}
        [
          <span class=query-link><a href="${clientCounts.query_link}" target=_blank>open issues</a></span>,
          <span class=query-link><a href="${clientCounts.untriaged_query_link}" target=_blank>untriaged issues</a></span>
        ]
      `;
    }
    return html`${clientKey}`;
  }

  private displayClientsRows(): TemplateResult[] {
    const rowsHTML = [];
    const clientKeys = Object.keys(this.clients_to_counts);
    clientKeys.sort();
    for (let i = 0; i < clientKeys.length; i++) {
      const clientKey = clientKeys[i];
      const clientCounts = this.clients_to_counts[clientKey];
      rowsHTML.push(html`
        <tr>
          <td @click=${() => this.clickClient(clientKey)}>
            <span class=client-link>${clientKey}</span>
          </td>
          <td>
            ${clientCounts.p0_count + clientCounts.p1_count}
            ${clientCounts.p0_slo_count + clientCounts.p1_slo_count > 0
    ? html`<span class="small"> [${clientCounts.p0_slo_count + clientCounts.p1_slo_count}]</span>`
    : ''}
          </td>
          <td>
            ${clientCounts.p2_count}
            ${clientCounts.p2_slo_count > 0
    ? html`<span class="small"> [${clientCounts.p2_slo_count}]</span>`
    : ''}
          </td>
          <td>
            ${clientCounts.p3_count + clientCounts.p4_count + clientCounts.p5_count + clientCounts.p6_count}
            ${clientCounts.p3_slo_count > 0
    ? html`<span class="small"> [${clientCounts.p3_slo_count}]</span>`
    : ''}
          </td>
          <td>
          ${clientCounts.untriaged_query_link
    ? html`<span class=query-link><a href="${clientCounts.untriaged_query_link}" target=_blank>${clientCounts.untriaged_count}</a></span>`
    : html`${clientCounts.untriaged_count}`}
          </td>
          <td>
          ${clientCounts.query_link
    ? html`<span class=query-link><a href="${clientCounts.query_link}" target=_blank>${clientCounts.open_count}</a></span>`
    : html`${clientCounts.open_count}`}
            ${clientCounts.p0_slo_count + clientCounts.p1_slo_count + clientCounts.p2_slo_count + clientCounts.p3_slo_count > 0
    ? html`<span class="small"> [${clientCounts.p0_slo_count + clientCounts.p1_slo_count + clientCounts.p2_slo_count + clientCounts.p3_slo_count}]</span>`
    : ''}
          </td>
        </tr>
      `);
    }
    return rowsHTML;
  }

  private clickClient(clientKey: string) {
    const tokens = breakupClientKey(clientKey);
    this.state.client = tokens.client ? tokens.client : '';
    this.state.source = tokens.source ? tokens.source : '';
    this.state.query = tokens.query ? tokens.query : '';
    this.stateHasChanged();
    this.populateDataAndRender();
  }

  // If client is specified and there is only one source then directly display
  // it's queries. If there is only one query available then directly set it
  // on the status. This saves unnecessary clicks for users.
  //
  // Eg: When a user clicks on 'Android' the UI would show 'Android>Buganizer'.
  // Clicking on that would then display 'Android>Buganizer>query'. Instead of these
  // unnecessary clicks, this function directly displays 'Android>Buganizer>query'
  // when 'Android' is clicked.
  private addExtraInformationToState(state: State): boolean {
    let stateUpdated = false;
    if (state.client && !state.source && !state.query) {
      const sources = Object.keys(this.clients_map[state.client as string]);
      if (sources.length === 1) {
        state.source = sources[0];
        stateUpdated = true;
        const queries = Object.keys(this.clients_map[state.client as string][state.source]);
        if (queries.length === 1) {
          state.query = queries[0];
          stateUpdated = true;
        }
      }
    }
    return stateUpdated;
  }

  private startStateReflector() {
    this.stateHasChanged = stateReflector(
      /* getState */() => {
        this.addExtraInformationToState(this.state);
        return (this.state as unknown) as HintableObject;
      },
      /* setState */(newState) => {
        this.state = (newState as unknown) as State;
        const stateUpdated = this.addExtraInformationToState(this.state);
        if (stateUpdated) {
          this.stateHasChanged();
        }
        this.populateDataAndRender();
      },
    );
  }

  // Common work done for all fetch requests.
  private async doImpl(url: string, detail: any, action: (json: any)=> void): Promise<void> {
    try {
      const resp = await fetch(url, {
        body: JSON.stringify(detail),
        headers: {
          'content-type': 'application/json',
        },
        credentials: 'include',
        method: 'POST',
      });
      const json = await jsonOrThrow(resp);
      action(json);
    } catch (msg) {
      errorMessage(msg);
    }
  }

  private async getCounts(client: string, source: string, query: string) {
    const detail = {
      client: client,
      source: source,
      query: query,
    };
    let countsData = {} as CountsData;
    await this.doImpl('/_/get_issue_counts', detail, (json: CountsData) => {
      countsData = json;
    });
    return countsData;
  }

  private async populateChartData() {
    const detail = {
      client: this.state.client,
      source: this.state.source,
      query: this.state.query,
    };
    await this.doImpl('/_/get_charts_data', detail, (json: ChartsData) => {
      this.open_chart_data = JSON.stringify(json.open_data);
      this.slo_chart_data = JSON.stringify(json.slo_data);
      this.untriaged_chart_data = JSON.stringify(json.untriaged_data);
    });
  }

  private async populateDataAndRender() {
    this.clients_to_counts = {};
    const c = this.state.client;
    const s = this.state.source;
    const q = this.state.query;

    if (!c) {
      await Promise.all(Object.keys(this.clients_map).map(async (client) => this.clients_to_counts[getClientKey(client, '', '')] = await this.getCounts(client, '', '')));
    } else if (!s) {
      await Promise.all(Object.keys(this.clients_map[c]).map(async (source) => this.clients_to_counts[getClientKey(c, source, '')] = await this.getCounts(c, source, '')));
    } else if (!q) {
      await Promise.all(Object.keys(this.clients_map[c][s]).map(async (query) => this.clients_to_counts[getClientKey(c, s, query)] = await this.getCounts(c, s, query)));
    } else {
      this.clients_to_counts[getClientKey(c, s, q)] = await this.getCounts(c, s, q);
    }
    // Render counts as soon we have them. Rendering charts will take longer.
    this._render();

    // Get chart data and render.
    await this.populateChartData();
    this._render();
  }
}

define('bugs-central-sk', BugsCentralSk);

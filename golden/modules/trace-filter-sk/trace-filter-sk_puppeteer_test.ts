import { expect } from 'chai';
import { takeScreenshot, TestBed } from '../../../puppeteer-tests/util';
import { loadGoldWebpack } from '../common_puppeteer_test/common_puppeteer_test';
import { TraceFilterSkPO } from './trace-filter-sk_po';

describe('trace-filter-sk', () => {
  let testBed: TestBed;
  let traceFilterSkPO: TraceFilterSkPO;

  before(async () => {
    testBed = await loadGoldWebpack();
  });

  beforeEach(async () => {
    await testBed.page.goto(`${testBed.baseUrl}/dist/trace-filter-sk.html`);
    traceFilterSkPO = new TraceFilterSkPO((await testBed.page.$('trace-filter-sk'))!);
  });

  it('should render the demo page', async () => {
    // Basic smoke test that things loaded.
    expect(await testBed.page.$$('trace-filter-sk')).to.have.length(1);
  });

  describe('empty selection', () => {
    beforeEach(async () => {
      await traceFilterSkPO.clickEditBtn();
      await traceFilterSkPO.setQueryDialogSkSelection({});
      await traceFilterSkPO.clickQueryDialogSkShowMatchesBtn();
    });

    it('shows the user input', async () => {
      await takeScreenshot(testBed.page, 'gold', 'trace-filter-sk');
    });

    it('opens the query dialog', async () => {
      await traceFilterSkPO.clickEditBtn();
      const queryDialogSkPO = await traceFilterSkPO.getQueryDialogSkPO();
      await queryDialogSkPO.clickKey('car make');
      await takeScreenshot(
        testBed.page, 'gold', 'trace-filter-sk_query-dialog-open');
    });
  });

  describe('non-empty selection', () => {
    it('shows the user input', async () => {
      await takeScreenshot(testBed.page, 'gold', 'trace-filter-sk_nonempty');
    });

    it('opens the query dialog', async () => {
      await traceFilterSkPO.clickEditBtn();
      const queryDialogSkPO = await traceFilterSkPO.getQueryDialogSkPO();
      await queryDialogSkPO.clickKey('car make');
      await takeScreenshot(
        testBed.page, 'gold', 'trace-filter-sk_nonempty_query-dialog-open');
    });
  })
});

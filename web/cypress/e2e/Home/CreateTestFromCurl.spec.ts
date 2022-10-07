import {SupportedPlugins} from '../../../src/constants/Common.constants';

describe('Create test from CURL Command', () => {
  beforeEach(() => cy.visit('/'));

  it('should create a basic GET test', () => {
    cy.inteceptHomeApiCall();
    const name = `Test - Pokemon - #${String(Date.now()).slice(-4)}`;
    cy.openTestCreationModal();
    cy.get(`[data-cy=${SupportedPlugins.CURL.toLowerCase()}-plugin]`).click();
    cy.fillCreateFormBasicStep(name, 'Create from Curl Command');

    cy.get('[data-cy=import-command-input] [contenteditable]')
      .first()
      .type(
        `curl -XPOST 'http://demo-pokemon-api.demo.svc.cluster.local/pokemon'
    -H "Content-type: application/json"
    --data '{"name":"meowth","type":"normal","imageUrl":"https://assets.pokemon.com/assets/cms2/img/pokedex/full/052.png","isFeatured":true}'
   `,
        {parseSpecialCharSequences: false}
      );
    cy.get('[data-cy=create-test-next-button]').last().click();

    cy.submitCreateTestForm();
    cy.matchTestRunPageUrl();
    cy.cancelOnBoarding();
    cy.get('[data-cy=test-details-name]').should('have.text', `${name} (v1)`);
    cy.deleteTest(true);
  });
});
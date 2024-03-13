import { check } from "k6";
import { textSummary } from "https://jslib.k6.io/k6-summary/0.0.2/index.js";
import { Http, Qualitytrace } from "k6/x/quality-trace";
import { sleep } from "k6";

export const options = {
  vus: 1,
  duration: "6s",
};

const qualitytrace = Qualitytrace({
  serverUrl: "http://localhost:11633",
});
const testId = "kc_MgKoVR";
const pokemonId = 6; // charizad
const http = new Http();
const url = "http://localhost:8081/pokemon/import";

export default function () {
  const payload = JSON.stringify({
    id: pokemonId,
  });
  const params = {
    qualitytrace: {
      testId,
    },
    headers: {
      "Content-Type": "application/json",
    },
  };

  const response = http.post(url, payload, params);

  check(response, {
    "is status 200": (r) => r.status === 200,
    "body matches de id": (r) => JSON.parse(r.body).id === pokemonId,
  });
  sleep(1);
}

// enable this to return a non-zero status code if a quality-trace test fails
export function teardown() {
  qualitytrace.validateResult();
}

export function handleSummary(data) {
  // combine the default summary with the quality-trace summary
  const qualitytraceSummary = qualitytrace.summary();
  const defaultSummary = textSummary(data);
  const summary = `
    ${defaultSummary}
    ${qualitytraceSummary}
  `;

  return {
    stdout: summary,
    "quality-trace.json": qualitytrace.json(),
  };
}

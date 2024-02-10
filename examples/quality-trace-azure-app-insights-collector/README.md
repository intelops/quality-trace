# Qualitytrace + OTel Collector + Azure Application Insights (using the OpenTelemetry Collector)

> <!--[Read the detailed recipe for setting up Qualitytrace + OTel Collector + Azure Application Insights (using the OpenTelemetry Collector) in our documentation.](https://docs.tracetest.io/examples-tutorials/recipes/running-tracetest-with-azure-app-insights-collector)-->

This repository objective is to show how you can configure your Qualitytrace instance using the OpenTelemetry collector to send telemetry data to both Azure App Insights and the Qualitytrace.

## Steps

1. [Install the quality-trace CLI]<!--(https://docs.tracetest.io/installing/)-->
2. Run `quality-trace configure --endpoint http://localhost:11633` on a terminal
3. Update the `.env` file adding a valid set the valid App Insights Instrumentation Key
4. Run the project by using docker-compose: `docker compose -f ./docker-compose.yaml -f ./quality-trace/docker-compose.yaml up -d`
5. Test if it works by running: `quality-trace run test -f tests/test.yaml`. This would trigger a test that will send spans to Azure Monitor API and directly to Qualitytrace that is running on your machine.

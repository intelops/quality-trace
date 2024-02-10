# Qualitytrace + Azure Application Insights (using Qualitytrace App Insights direct integration)

<!--> [Read the detailed recipe for setting up Qualitytrace + Azure Application Insights (using Qualitytrace App Insights direct integration) in our documentation.](https://docs.tracetest.io/examples-tutorials/recipes/running-tracetest-with-azure-app-insights)-->

This repository objective is to show how you can configure your Qualitytrace instance to connect to Azure App Insights and use it as its tracing backend.

## Steps

1. [Install the quality-trace CLI]<!--(https://docs.tracetest.io/installing/)-->
2. Run `quality-trace configure --endpoint http://localhost:11633` on a terminal
3. Update the `.env` file adding a valid set the valid App Insights Connection String
4. Update the `quality-trace.provision.yaml` file adding a valid set the Azure ARM Id and secret token
5. Run the project by using docker-compose: `docker compose -f ./docker-compose.yaml -f ./quality-trace/docker-compose.yaml up -d`
6. Test if it works by running: `quality-trace run test -f tests/test.yaml`. This would trigger a test that will send and retrieve spans from the Azure Monitor API instance that is running on your machine.

# Quick Start - .NET Core API with Jaeger, OpenTelemetry and Qualitytrace

> [Read the detailed recipe for setting up Jaeger with Tractest in our documentation.](https://docs.quality-trace.io/examples-tutorials/recipes/running-quality-trace-with-jaeger)

This is a simple quick start on how to configure a .NET Core API to use OpenTelemetry instrumentation with traces and Qualitytrace for enhancing your E2E and integration tests with trace-based testing. The infrastructure will use Jaeger as the trace data store, and OpenTelemetry Collector to receive traces from the API and send them to Jaeger.

## Steps

1. [Install the quality-trace CLI](https://github.com/intelops/quality-trace/blob/main/docs/installing.md#cli-installation)
2. Run `quality-trace configure --endpoint http://localhost:11633` on a terminal to configure the CLI to send all commands to that address
3. Run the project by using docker-compose: `docker-compose up -d` (Linux) or `docker compose up -d` (Mac)
4. Test if it works by running: `quality-trace run test -f tests/test.yaml`. This would execute a test against the .NET Core API that will send spans to Jaeger to be fetched from the Qualitytrace server.

Feel free to check out the [docs](https://docs.quality-trace.io/)<!-- and join our [Discord Community](https://discord.gg/8MtcMrQNbX) -->for more info!

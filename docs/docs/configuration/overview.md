# Configuration

There are several configuration options with Qualitytrace:

- [Server configuration](./server) to set database connection information needed to connect to required PostgreSQL instance.
- [Provisioning configuration](./provisioning) to 'preload' the Qualitytrace server with resources when first running the Qualitytrace server.

## Supported Trace Data Stores

Qualitytrace is designed to work with different trace data stores. To enable Qualitytrace to run end-to-end tests against trace data, you need to configure Qualitytrace to access trace data.

Currently, Qualitytrace supports the following data stores. Click on the respective data store to view configuration examples:

- [AWS X-Ray](./connecting-to-data-stores/awsxray)
- [Azure App Insights](./connecting-to-data-stores/azure-app-insights.md)
- [Datadog](./connecting-to-data-stores/datadog)
- [Dynatrace](./connecting-to-data-stores/dynatrace)
- [Elastic APM](./connecting-to-data-stores/elasticapm)
- [Grafana Tempo](./connecting-to-data-stores/tempo)
- [Honeycomb](./connecting-to-data-stores/honeycomb)
- [Jaeger](./connecting-to-data-stores/jaeger)
- [Lightstep](./connecting-to-data-stores/lightstep)
- [New Relic](./connecting-to-data-stores/new-relic)
- [OpenSearch](./connecting-to-data-stores/opensearch)
- [OpenTelemetry Collector](./connecting-to-data-stores/opentelemetry-collector)
- [SignalFX](./connecting-to-data-stores/signalfx)
- [Signoz](./connecting-to-data-stores/signoz)

## Using Qualitytrace without a Trace Data Store

Another option is to send traces directly to Qualitytrace using the OpenTelemetry Collector. And, you don't have to change your existing pipelines to do so.

View [configuration for OpenTelemetry Collector](./connecting-to-data-stores/opentelemetry-collector.md) for more details.

## Trace Data Store Configuration Examples

Examples of configuring Qualitytrace to access different data stores can be found in the [`examples` folder of the Qualitytrace GitHub repo](https://github.com/intelops/quality-trace/tree/main/examples). Check out the [**Recipes**](../examples-tutorials/recipes.md) for guided walkthroughs of sample use cases.

We will be adding new data stores over the next couple of months - [let us know](https://github.com/intelops/quality-trace/issues/new/choose) any additional data stores you would like to see us support.

## Configuring the Server

Qualitytrace has a configuration file to contain the minimal information needed to start the Qualitytrace server. See more at [Qualitytrace Server Configuration](./server).

You can also provision the server when it first starts, configuring most aspects of your Qualitytrace server environment. This is useful in a CI/CD environment to preload and configure the server. See more at [Provisioning a Qualitytrace Server](./provisioning).

Many of the server configuration settings can be set individually in the UI or via the CLI. See:

- [Trace Polling](./trace-polling)
- [Demo Applications](./demo)
- [Analytics](./analytics)

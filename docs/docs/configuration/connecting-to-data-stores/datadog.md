# Datadog

If you want to use [Datadog](https://www.datadoghq.com/) as the trace data store, you'll configure the OpenTelemetry Collector to receive traces from your system and then send them to both Qualitytrace and Datadog. And, you don't have to change your existing pipelines to do so.

:::tip
Examples of configuring Qualitytrace with Datadog can be found in the [`examples` folder of the Qualitytrace GitHub repo](https://github.com/intelops/quality-trace/tree/main/examples).
:::

## Configuring OpenTelemetry Collector to Send Traces to Both Datadog and Qualitytrace

In your OpenTelemetry Collector config file:

- Set the `exporter` to `otlp/quality-trace`
- Set the `endpoint` to your Qualitytrace instance on port `4317`

:::tip
If you are running Qualitytrace with Docker, and Qualitytrace's service name is `quality-trace`, then the endpoint might look like this `http://quality-trace:4317`
:::

Additionally, add another config:

- Set the `exporter` to `datadog`.
- Set the `api` pointing to your Datadog account.
- Set the `site` to Datadog API `datadoghq.com`.
- Set the `key` to your Datadog API key.

```yaml
# collector.config.yaml

# If you already have receivers declared, you can just ignore
# this one and still use yours instead.
receivers:
  otlp:
    protocols:
      http:
      grpc:

processors:
  # This configuration is needed to guarantee that the data is sent correctly to Datadog
  batch:
    send_batch_max_size: 100
    send_batch_size: 10
    timeout: 10s

exporters:
  # OTLP for Qualitytrace
  # Send traces to Qualitytrace.
  # Read more in docs here: https://docs.quality-trace.io/configuration/connecting-to-data-stores/opentelemetry-collector
  otlp/quality-trace:
    endpoint: quality-trace:4317
    tls:
      insecure: true

  # Datadog exporter
  # One example on how to set up a collector configuration for Datadog can be seen here:
  # https://docs.datadoghq.com/opentelemetry/otel_collector_datadog_exporter/?tab=onahost
  datadog:
    api:
      site: datadoghq.com
      key: ${DATADOG_API_KEY} # Add here you API key for Datadog

service:
  pipelines:
    traces/quality-trace:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp/quality-trace] # exporter sending traces to your Qualitytrace instance
    traces/datadog:
      receivers: [otlp]
      processors: [batch]
      exporters: [datadog] # exporter sending traces to directly to Datadog
```

## Configure Qualitytrace to Use Datadog as a Trace Data Store

Configure your Qualitytrace instance to expose an `otlp` endpoint to make it aware it will receive traces from the OpenTelemetry Collector. This will expose Qualitytrace's trace receiver on port `4317`.

## Connect Qualitytrace to Datadog with the Web UI

In the Web UI, (1) open Settings, and, on the (2) Configure Data Store tab, select (3) Datadog.

![Datadog](../img/configure-datadog-0.11.3.png)

## Connect Qualitytrace to Datadog with the CLI

Or, if you prefer using the CLI, you can use this file config.

```yaml
type: DataStore
spec:
  name: Datadog pipeline
  type: datadog
  default: true
```

Proceed to run this command in the terminal and specify the file above.

```bash
quality-trace apply datastore -f my/data-store/file/location.yaml
```

:::tip
To learn more, [read the recipe for running a sample app with Datadog and Qualitytrace](../../examples-tutorials/recipes/running-quality-trace-with-datadog.md).
:::

# Dynatrace

If you want to use [Dynatrace](https://www.dynatrace.com/) as the trace data store, you'll configure the OpenTelemetry Collector to receive traces from your system and then send them to both Qualitytrace and Dynatrace. And, you don't have to change your existing pipelines to do so.

:::tip
Examples of configuring Qualitytrace with Dynatrace can be found in the [`examples` folder of the Qualitytrace GitHub repo](https://github.com/intelops/quality-trace/tree/main/examples).
:::

## Configuring OpenTelemetry Collector to Send Traces to both Dynatrace and Qualitytrace

In your OpenTelemetry Collector config file:

- Set the `exporter` to `otlp/quality-trace`
- Set the `endpoint` to your Qualitytrace instance on port `4317`

:::tip
If you are running Qualitytrace with Docker, and Qualitytrace's service name is `quality-trace`, then the endpoint might look like this `http://quality-trace:4317`
:::

Additionally, add another config:

- Set the `exporter` to `otlphttp/dynatrace`
- Set the `endpoint` to your Dynatrace tenant and include the: `https://{your-environment-id}.live.dynatrace.com/api/v2/otlp`

```yaml
# collector.config.yaml

# If you already have receivers declared, you can just ignore
# this one and still use yours instead.
receivers:
  otlp:
    protocols:
      grpc:
      http:

processors:
  batch:
    timeout: 100ms

exporters:
  logging:
    verbosity: detailed
  # OTLP for Qualitytrace
  otlp/quality-trace:
    endpoint: quality-trace:4317 # Send traces to Qualitytrace. Read more in docs here:  https://docs.quality-trace.io/configuration/connecting-to-data-stores/opentelemetry-collector
    tls:
      insecure: true
  # OTLP for Dynatrace
  otlphttp/dynatrace:
    endpoint: https://abc123.live.dynatrace.com/api/v2/otlp # Send traces to Dynatrace. Read more in docs here: https://www.dynatrace.com/support/help/extend-dynatrace/opentelemetry/collector#configuration
    headers:
      Authorization: "Api-Token dt0c01.sample.secret" # Requires "openTelemetryTrace.ingest" permission
service:
  pipelines:
    traces/quality-trace: # Pipeline to send data to Qualitytrace
      receivers: [otlp]
      processors: [batch]
      exporters: [logging, otlp/quality-trace]
    traces/Dynatrace: # Pipeline to send data to Dynatrace
      receivers: [otlp]
      processors: [batch]
      exporters: [logging, otlphttp/dynatrace]
```

## Configure Qualitytrace to Use Dynatrace as a Trace Data Store

Configure your Qualitytrace instance to expose an `otlp` endpoint to make it aware it will receive traces from the OpenTelemetry Collector. This will expose Qualitytrace's trace receiver on port `4317`.

## Connect Qualitytrace to Dynatrace with the Web UI

In the Web UI, (1) open Settings, and, on the (2) Configure Data Store tab, select (3) Dynatrace.

<!-- TODO: create this image using the same standard as the other stores -->
![Dynatrace](../img/Dynatrace-settings.png)

## Connect Qualitytrace to Dynatrace with the CLI

Or, if you prefer using the CLI, you can use this file config.

```yaml
type: DataStore
spec:
  name: Dynatrace pipeline
  type: dynatrace
  default: true
```

Proceed to run this command in the terminal and specify the file above.

```bash
quality-trace apply datastore -f my/data-store/file/location.yaml
```

<!-- 
TODO: create a tutorial for Dynatrace
:::tip
To learn more, [read the recipe on running a sample app with Dynatrace and Qualitytrace](../../examples-tutorials/recipes/running-quality-trace-with-dynatrace.md).
::: 
-->

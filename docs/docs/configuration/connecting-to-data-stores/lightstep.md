# Lightstep

If you want to use [Lightstep](https://lightstep.com/) as the trace data store, you'll configure the OpenTelemetry Collector to receive traces from your system and then send them to both Qualitytrace and Lightstep. And, you don't have to change your existing pipelines to do so.

:::tip
Examples of configuring Qualitytrace with Lightstep can be found in the [`examples` folder of the Qualitytrace GitHub repo](https://github.com/intelops/quality-trace/tree/main/examples).
:::

## Configuring OpenTelemetry Collector to Send Traces to both Lightstep and Qualitytrace

In your OpenTelemetry Collector config file:

- Set the `exporter` to `otlp/quality-trace`
- Set the `endpoint` to your Qualitytrace instance on port `4317`

:::tip
If you are running Qualitytrace with Docker, and Qualitytrace's service name is `quality-trace`, then the endpoint might look like this `http://quality-trace:4317`
:::

Additionally, add another config:

- Set the `exporter` to `otlp/lightstep`
- Set the `endpoint` pointing to your Lightstep account and the Lightstep ingest API
- Set your Lightstep access token as a `header`

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
    logLevel: debug
  # OTLP for Qualitytrace
  otlp/quality-trace:
    endpoint: quality-trace:4317 # Send traces to Qualitytrace. Read more in docs here:  https://docs.quality-trace.io/configuration/connecting-to-data-stores/opentelemetry-collector
    tls:
      insecure: true
  # OTLP for Lightstep
  otlp/lightstep:
    endpoint: ingest.lightstep.com:443
    headers:
      "lightstep-access-token": "<lightstep_access_token>" # Send traces to Lightstep. Read more in docs here: https://docs.lightstep.com/otel/otel-quick-start

service:
  pipelines:
    # Your probably already have a traces pipeline, you don't have to change it.
    # Add these two to your configuration. Just make sure to not have two
    # pipelines with the same name
    traces/quality-trace:
      receivers: [otlp] # your receiver
      processors: [batch]
      exporters: [otlp/quality-trace] # your exporter pointing to your quality-trace instance
    traces/lightstep:
      receivers: [otlp] # your receiver
      processors: [batch]
      exporters: [logging, otlp/lightstep] # your exporter pointing to your lighstep account
```

## Configure Qualitytrace to Use Lightstep as a Trace Data Store

Configure your Qualitytrace instance to expose an `otlp` endpoint to make it aware it will receive traces from the OpenTelemetry Collector. This will expose Qualitytrace's trace receiver on port `4317`.

## Connect Qualitytrace to Lightstep with the Web UI

In the Web UI, (1) open Settings, and, on the (2) Configure Data Store tab, select (3) Lightstep.

![Lightstep](../img/Lightstep-settings.png)

<!---![](https://res.cloudinary.com/djwdcmwdz/image/upload/v1674643396/Blogposts/Docs/screely-1674643391899_w6k22s.png)-->

## Connect Qualitytrace to Lightstep with the CLI

Or, if you prefer using the CLI, you can use this file config.

```yaml
type: DataStore
spec:
  name: Lightstep pipeline
  type: lightstep
  default: true
```

Proceed to run this command in the terminal and specify the file above.

```bash
quality-trace apply datastore -f my/data-store/file/location.yaml
```

:::tip
To learn more, [read the recipe on running a sample app with Lightstep and Qualitytrace](../../examples-tutorials/recipes/running-quality-trace-with-lightstep.md).
:::

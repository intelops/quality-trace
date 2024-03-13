# OpenTelemetry Collector

Qualitytrace receives trace data on port `4317`. Qualitytrace's trace receiver endpoint might look like:

```
http://your-quality-trace-instance.com:4317
```

:::tip
Examples of configuring Qualitytrace can be found in the [`examples` folder of the Qualitytrace GitHub repo](https://github.com/intelops/quality-trace/tree/main/examples).
:::

## Configuring OpenTelemetry Collector to Send Traces to Qualitytrace

If you don't want to use a trace data store, you can send all traces directly to Qualitytrace using your OpenTelemetry Collector. And, you don't have to change your existing pipelines to do so.

In your OpenTelemetry Collector config file:

- Set the `exporter` to `otlp/1`
- Set the `endpoint` to your Qualitytrace instance on port `4317`

:::tip
If you are running Qualitytrace with Docker, and Qualitytrace's service name is `quality-trace`, then the endpoint might look like this `http://quality-trace:4317`
:::

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
  # This is the exporter that will send traces to Qualitytrace
  otlp/1:
    endpoint: http://your-quality-trace-instance.com:4317
    tls:
      insecure: true

service:
  pipelines:
    # your probably already have a traces pipeline, you don't have to change it.
    # just add this one to your configuration. Just make sure to not have two
    # pipelines with the same name
    traces/1:
      receivers: [otlp] # your receiver
      processors: [batch]
      exporters: [otlp/1] # your exporter pointing to your quality-trace instance
```

## Configure Qualitytrace to Use OpenTelemetry Collector

Configure your Qualitytrace instance to expose an `otlp` endpoint to make it aware it will receive traces from the OpenTelemetry Collector. This will expose Qualitytrace's trace receiver on port `4317`.

## Connect Qualitytrace to OpenTelemetry Collector with the Web UI

In the Web UI, (1) open Settings, and, on the (2) Configure Data Store tab, select (3) OpenTelemetry.

![OpenTelemetry](../img/open-telemetry-settings.png)

<!---![](https://res.cloudinary.com/djwdcmwdz/image/upload/v1674644190/Blogposts/Docs/screely-1674644186486_pahrds.png) -->

## Connect Qualitytrace to OpenTelemetry Collector with the CLI

Or, if you prefer using the CLI, you can use this file config.

```yaml
type: DataStore
spec:
  name: Opentelemetry Collector pipeline
  type: otlp
  default: true
```

Proceed to run this command in the terminal, and specify the file above.

```bash
quality-trace apply datastore -f my/data-store/file/location.yaml
```

:::tip
To learn more, [read the recipe on running a sample app with OpenTelemetry Collector and Qualitytrace](../../examples-tutorials/recipes/running-quality-trace-without-a-trace-data-store.md).
:::

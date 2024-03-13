# Telemetry

The Qualitytrace server generates internal observability trace data. You can use this data to track Qualitytrace test runs over time and gain observability of how the Qualitytrace server is behaving.

The Qualitytrace team uses an observability-driven development approach in developing the Qualitytrace server, capturing traces and then running Qualitytrace tests against it as part of the CI/CD process.<!-- You can read more about how we "eat our own dog food" in [this blog post](https://tracetestio/blog/integrating-tracetest-with-github-actions-in-a-ci-pipeline) about the CI/CD configuration where we test Qualitytrace with Qualitytrace.-->

## Configuring Qualitytrace Server Internal Telemetry

You can configure an exporter to send the trace data to an OpenTelemetry Collector and then store it safely in your trace data store for further historical analysis. View the [supported trace data stores](./overview#supported-trace-data-stores) for more guidance on setting them up.

In the `quality-trace-config.yaml` file, alongside the [configuration](./server.md) of connecting Qualitytrace to the Postgres instance, you can also define a `telemetry` and `server` section.

With these two additional sections, you define an exporter where the Qualitytrace server's internal telemetry will be routed to. In the `telemetry` section, you define the endpoint of the OpenTelemetry Collector. And, in the `server` section you define which exporter the Qualitytrace server will use.

```yaml
# quality-trace-config.yaml
postgres:
# [...]

telemetry:
  exporters:
    collector:
      serviceName: quality-trace
      sampling: 100 # 100%
      exporter:
        type: collector
        collector:
          endpoint: otel-collector:4317
          # Replace with your OpenTelemetry Collector endpoint

server:
  telemetry:
    exporter: collector
    applicationExporter: collector
```

:::note
Make sure to check what the service endpoint for the OpenTelemetry Collector in your infrastructure is. The example above is using `otel-collector` because that is the service name in Docker Compose. Your infrastructure might differ.
:::

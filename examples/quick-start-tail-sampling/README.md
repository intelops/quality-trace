# Quick Start - Trace-based Tests with Tail Sampling Configuration

> [Read the detailed recipe for setting up OpenTelemetry Collector with Qualitytrace in our documentation.](https://docs.quality-trace.io/examples-tutorials/recipes/running-quality-trace-without-a-trace-data-store)

This is a simple quick start example on how to set up `tail_sampling` into OTel Collector, allowing Qualitytrace to run tests in environments where we have a probabilistic sampler enabled and a percentage of the traces are not sent to the final data store.

## Scenario

In this scenario, we have [Go API](./simple-go-service/) that sends Trace data to one instance of [OTel Collector](https://opentelemetry.io/docs/collector/), which samples 50% of the traces and sends it to the [Jaeger](https://www.jaegertracing.io/) data store.

```mermaid
---
title: Initial observability architecture
---
flowchart LR
    GoAPI["Go API"]
    OTelCol["OTel Collector"]
    DataStore[("Jaeger")]

    GoAPI -- Send traces --> OTelCol
    OTelCol -- Forward 50% of Traces --> DataStore
``````

The collector configuration is defined by a `traces/jaeger` pipeline with one receiver, two processors (one to send traces in batch and the other set up the probabilistic sampling) and one exporter to Jaeger:

```yaml
receivers:
  otlp:
    protocols:
      grpc:
      http:

processors:
  batch:
    timeout: 100ms
  probabilistic_sampler:
    hash_seed: 22
    sampling_percentage: 50.0 # 50%

exporters:
  otlp/jaeger:
    endpoint: jaeger:4317 # Send traces to Jaeger
    tls:
      insecure: true

service:
  pipelines:
    traces/jaeger:
      receivers: [otlp]
      processors: [batch, probabilistic_sampler]
      exporters: [otlp/jaeger]

```

However, to run a Trace-based test on this scenario and validate if the API is ok, we need to have the testing trace collected by our OTelCollector. How can we do that?

Since Qualitytrace always starts a trace by adding a `quality-trace=true` key/value to [TraceState](https://opentelemetry.io/docs/specs/otel/trace/api/#tracestate) that is [propagated](https://opentelemetry.io/docs/instrumentation/js/propagation/), we can filter these traces on the collector and send them to Qualitytrace.

To do this we need to set up Qualitytrace to use an OTLP Data Store on its provision ([here](./quality-trace/quality-trace.provision.yaml)) and also set up another pipeline with a [Tail Sampling](https://opentelemetry.io/docs/concepts/sampling/#tail-sampling) processor on the OTel Collector, filtering by TraceState `quality-trace=true` value.

We can do that by adding the following configuration to the OTel Collector (the entire collector config YAML is [here](./quality-trace/collector.config.yaml)):
```yaml
#...

processors:
  #...
  tail_sampling:
    decision_wait: 10s
    num_traces: 100
    expected_new_traces_per_sec: 10
    policies:
      [
        {
          name: Accept only traces that started from quality-trace,
          type: trace_state,
          trace_state: {
            key: quality-trace,
            values: ["true"]
          }
        }
      ]

exporters:
  #...
  otlp/quality-trace:
    endpoint: quality-trace:4317 # Send traces to Qualitytrace.
    tls:
      insecure: true

service:
  pipelines:
    #...
    traces/quality-trace:
      receivers: [otlp]
      processors: [batch, tail_sampling]
      exporters: [otlp/quality-trace, logging]

```

With this configuration, we can run Trace-based tests, even with a sampling, into the main data store, having the following structure:

```mermaid
---
title: Observability architecture with Trace-based testing support
---
flowchart LR
    GoAPI["Go API"]
    OTelCol["OTel Collector"]
    DataStore[("Jaeger")]
    Qualitytrace

    GoAPI -- Send traces --> OTelCol
    OTelCol -- Forward 50% of Traces --> DataStore
    OTelCol -- Forward only `quality-trace=true` traces --> Qualitytrace
``````

## Running the example

If you want to run this example, execute `docker compose up` on this folder.

To execute a Trace-based test with Qualitytrace against this structure, run `quality-trace run test -f test-api-working.yaml`.

Feel free to check out the [docs](https://docs.quality-trace.io/), and join our [Discord Community](https://discord.gg/8MtcMrQNbX) for more info!

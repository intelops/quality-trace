# Server configuration

Qualitytrace can be configured using a config.yaml file placed on the same directory as its executable. It is useful to configure some aspects of how quality-trace should behave. This section is dedicated to explain the options we currently have available.

## Configuration file example
```yaml
# Connection string to the postgres instance
postgres:
  host: localhost
  user: postgres
  password: postgres

# Instance of jaeger that will be used to retrieve the trace of the service under test
tracingBackend:
  dataStore:
    type: jaeger
    jaeger:
      endpoint: localhost:16685
      tls:
        insecure: true

# Configure how traces should be pooled from the tracing storage.
poolingConfig:
    # How long quality-trace can wait for a trace to be complete? After this period, the pooling process will timeout
    # and the test will be marked as failed.
    maxWaitTimeForTrace: 90s

    # How much time quality-trace should wait before trying to fetch the trace since the last execution?
    retryDelay: 5s

# Server configuration
server:
  # Enables you to add a prefix to the server path. So, instead of running quality-trace on http://localhost:11633, it would run on http://localhost:11633/quality-trace instead.
  pathPrefix: /quality-trace
  httpPort: 11633

# Google analytics configuration
googleAnalytics:
  enabled: false
  measurementId: ""
  secretKey: ""

# How quality-trace should generate telemetry data.
telemetry:
  serviceName: quality-trace
  sampling: 100
  otelCollectorEndpoint: localhost:4317
```

## Providing a configuration when running a container
```cmd
docker run --volume "`pwd`/my-config-file.yaml:/app/config.yaml" intelops/quality-trace
```

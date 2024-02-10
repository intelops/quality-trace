# Jaeger

Qualitytrace fetches traces from [Jaeger's gRPC Protobuf/gRPC QueryService](https://www.jaegertracing.io/docs/1.42/deployment/#query-service--ui) on port `16685`.

:::tip
Examples of configuring Qualitytrace can be found in the [`examples` folder of the Qualitytrace GitHub repo](https://github.com/intelops/quality-trace/tree/main/examples).
:::

## Configure Qualitytrace to Use Jaeger as a Trace Data Store

Configure Qualitytrace to be aware that it has to fetch trace data from Jaeger.

Qualitytrace uses [Jaeger's gRPC Protobuf/gRPC QueryService](https://www.jaegertracing.io/docs/1.42/deployment/#query-service--ui) on port `16685` to fetch trace data.

:::tip
Need help configuring the OpenTelemetry Collector so send trace data from your application to Jaeger? Read more in [the reference page here](../opentelemetry-collector-configuration-file-reference)).
:::

## Connect Qualitytrace to Jaeger with the Web UI

In the Web UI, (1) open Settings, and, on the (2) Configure Data Store tab, (3) select Jaeger. If you are using Docker like in the screenshot below, use the service name as the hostname with port `16685` like this:

```
jaeger:16685
```

![Jaeger](../img/Jaeger-settings.png)

<!---![](https://res.cloudinary.com/djwdcmwdz/image/upload/v1674643178/Blogposts/Docs/screely-1674643170953_vazb9h.png)-->

## Connect Qualitytrace to Jaeger with the CLI

Or, if you prefer using the CLI, you can use this file config.

```yaml
type: DataStore
spec:
  name: jaeger
  type: jaeger
  default: true
  jaeger:
    endpoint: jaeger:16685
    tls:
      insecure: true
```

Proceed to run this command in the terminal, and specify the file above.

```bash
quality-trace apply datastore -f my/data-store/file/location.yaml
```

:::tip
To learn more, [read the recipe on running a sample app with Jaeger and Qualitytrace](../../examples-tutorials/recipes/running-quality-trace-with-jaeger.md).
:::

# SignalFx

Qualitytrace fetches traces from [SignalFx's realm and token](https://docs.splunk.com/Observability/references/organizations.html).

:::tip
Examples of configuring Qualitytrace can be found in the [`examples` folder of the Qualitytrace GitHub repo](https://github.com/intelops/quality-trace/tree/main/examples).
:::

## Configure Qualitytrace to Use SignalFx as a Trace Data Store

Configure Qualitytrace to be aware that it has to fetch trace data from SignalFx.

:::tip
Need help configuring the OpenTelemetry Collector so send trace data from your application to SignalFx? Read more in [the reference page here](../opentelemetry-collector-configuration-file-reference)).
:::

## Connect Qualitytrace to SignalFx with the Web UI

In the Web UI, (1) open Settings, and, on the (2) Configure Data Store tab, (3) select SignalFx.

You need your SignalFx:

- **Realm**
- **Token**

Follow this [guide](https://docs.splunk.com/Observability/references/organizations.html).

![SignalFX](../img/SignalFX-settings.png)

<!---![](https://res.cloudinary.com/djwdcmwdz/image/upload/v1674644337/Blogposts/Docs/screely-1674644332529_cks0lw.png)-->

## Connect Qualitytrace to SignalFx with the CLI

Or, if you prefer using the CLI, you can use this file config.

```yaml
type: DataStore
spec:
  name: SignalFX
  type: signalfx
  default: true
  signalfx:
    realm: us1
    token: mytoken
```

Proceed to run this command in the terminal, and specify the file above.

```bash
quality-trace apply datastore -f my/data-store/file/location.yaml
```

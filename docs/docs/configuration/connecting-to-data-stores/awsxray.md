# AWS X-Ray

If you want to use [AWS X-Ray](https://aws.amazon.com/xray/) as the trace data store, you can use the native connection from Qualitytrace to pull telemetry data directly from any region.

:::tip
Examples of configuring Qualitytrace with X-Ray can be found in the [`examples` folder of the Qualitytrace GitHub repo](https://github.com/intelops/quality-trace/tree/main/examples).
:::

## Configure Qualitytrace to Use X-Ray as a Trace Data Store

Configure Qualitytrace to be aware that it has to fetch trace data from X-Ray.

Qualitytrace uses the Golang [AWS-SDK](https://aws.amazon.com/sdk-for-go/) library to pull to fetch trace data.

:::tip
Need help configuring the OpenTelemetry Collector so send trace data from your application to AWS X-Ray? Read more in [the reference page here](../opentelemetry-collector-configuration-file-reference).
:::

## Connect Qualitytrace to X-Ray with the Web UI

In the Web UI, (1) open Settings, and, on the (2) Configure Data Store tab, select (3) AWS X-Ray.

![AWS X-Ray Settings](../img/configure-awsxray-0.11.3.png)

From the configuration page, add your AWS credentials. You can use temporary credentials by running `aws sts get-session-token`.
Lastly, you can ensure the connection is working as expected by pressing the `Test Connection` button.

## Connect Qualitytrace to AWS X-Ray with the CLI

Or, if you prefer using the CLI, you can use this file config.

```yaml
type: DataStore
spec:
  name: AWS X-Ray
  type: awsxray
  default: true
  awsxray:
    accessKeyId: <your-accessKeyId>
    secretAccessKey: <your-secretAccessKey>
    sessionToken: <your-session-token>
    region: "us-west-2"
```

Run this command in the terminal and specify the file above.

```bash
quality-trace apply datastore -f my/data-store/file/location.yaml
```

:::tip
To learn more, [read the recipe on running a sample app with AWS X-Ray and Qualitytrace](../../examples-tutorials/recipes/running-quality-trace-with-aws-x-ray.md).
:::

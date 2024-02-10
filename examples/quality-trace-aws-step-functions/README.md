# Quick Start - Qualitytrace with Terraform, XRay, Step Functions and .NET

This is a simple quick start on how to configure a .NET State Machine (AWS Step Functions) using XRay instrumentation and Qualitytrace for enhancing your E2E and integration tests with trace-based testing. The infrastructure will use Jaeger as the trace data store and Terraform to provision the required AWS infrastructure (Fargate/Lambda).

## Steps

1. [Install the quality-trace CLI](https://github.com/intelops/quality-trace/blob/main/docs/installing.md#cli-installation)
2. From the `infra` folder run `terraform init` and `terraform apply` and accept the changes
3. From the terraform outputs, grab the `quality-trace_url` and run `quality-trace configure --endpoint <tracetest_url>` on a terminal to configure the CLI to send all commands to that address
4. From the `src` folder run `sam build` and `sam deploy --guided`
5. Follow the instructions from the guided deployment
![functions](./assets/functions.png)
6. Grab the API gateway endpoint from the outputs, and update the `<your_api_endpoint>` section from `test/incident.yaml` and the `test/exam.yaml` files
7. Inject the Tests and Transactions definitions to the quality-trace server  using the following:

```bash
quality-trace run test -f tests/incident.yaml \
quality-trace run test -f tests/exam.yaml \
quality-trace run transaction -f tests/transaction.yaml
```

Feel free to check out the [docs](https://docs.quality-trace.io/), and join our [Discord Community](https://discord.gg/8MtcMrQNbX) for more info!

# Synthetic monitoring with Qualitytrace and GitHub Actions

> <!--[Read the detailed recipe for setting up OpenTelemetry Collector with Tractest in our documentation.](https://docs.tracetest.io/examples-tutorials/recipes/running-quality-trace-without-a-trace-data-store)-->

This is a simple quick start on how to configure Qualitytrace and GitHub Actions to emulate synthetic monitoring using your existing Qualitytrace tests. The idea behind this example
is to have a way of running a set of tests using Qualitytrace on a schedule and notify a Slack channel in case the test fails. This way, you can keep testing your application and
identifying issues constantly.

This example is based on the [Golang quick-start example](https://github.com/intelops/quality-trace/tree/main/examples/quick-start-go). All important information about
how to run synthetic monitoring using Qualitytrace can be found in the [.github/workflows/synthetic-monitoring.yaml](https://github.com/intelops/quality-trace/tree/main/examples/quality-trace-synthetic-monitoring/.github/workflows/synthetic-monitoring.yaml) file.

<!--Feel free to check out the [docs](https://docs.tracetest.io/), and join our [Discord Community](https://discord.gg/8MtcMrQNbX) for more info!-->

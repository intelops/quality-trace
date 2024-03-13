# Detailed Instructions on Installing Qualitytrace Using the CLI

This page contains detailed instructions for all installation options the Qualitytrace CLI provides.

Qualitytrace has a command line interface (CLI) which includes an **install wizard** that helps to install the Qualitytrace server into Docker or Kubernetes. The CLI can also be used to run tests, download or upload tests, and manage much of the capability of Qualitytrace.

This page provides a step-by-step guide to install Qualitytrace using the Qualitytrace CLI.

## Prerequisites

:::info
Make sure you have the Qualitytrace CLI installed. Read the installation reference [here](../cli/cli-installation-reference.md).
:::

:::info
Make sure you have [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/), or [Kubernetes](https://kubernetes.io/) installed.
:::

:::info
In this quick start, OpenTelemetry Collector is used to send traces directly to Qualitytrace. If you have an existing trace data source, [read here](../configuration/overview.md).
:::

## Docker Compose

This guide showcases using the Docker Compose option in the Qualitytrace CLI.

### 1. Run the `server install` Command

Once you've installed the CLI you can install a Qualitytrace server by running:

```bash
quality-trace server install
```
<!--
```text title="Expected output:"
████████ ██████   █████   ██████ ███████ ████████ ███████ ███████ ████████ 
   ██    ██   ██ ██   ██ ██      ██         ██    ██      ██         ██    
   ██    ██████  ███████ ██      █████      ██    █████   ███████    ██    
   ██    ██   ██ ██   ██ ██      ██         ██    ██           ██    ██    
   ██    ██   ██ ██   ██  ██████ ███████    ██    ███████ ███████    ██    
-->
Version: v0.10.1


Hi! Welcome to the Qualitytrace server installer. I'll help you set up your Qualitytrace server by asking you a few questions
and configuring your system with all the requirements, so you can start Qualitytracing right away!

To get more info about Qualitytrace, you can check our docs at https://intelops.github.io/quality-trace/

If you have any issues, please let us know by creating an issue (https://github.com/intelops/quality-trace/issues/new/choose)
<!--or reach us on Discord https://discord.gg/6zupCZFQbe-->


How do you want to run Qualitytrace? [type to search]: 
> Using Docker Compose
  Using Kubernetes
```

### 2. Select Docker Compose

This prompts you to select if you want to get either a Docker Compose or Kubernetes setup generated for you.

Select `Using Docker Compose`.

```text title="Expected output:"
...

How do you want to run Qualitytrace?: 
  > Using Docker Compose


-> Let's check if your system has everything we need

✔ docker already installed
✔ docker is ready
✔ docker compose already installed

-> Your system is ready! Now, let's configure Qualitytrace

Do you have OpenTelemetry based tracing already set up, or would you like us to install a demo tracing environment and app? [type to search]: 
> I have a tracing environment already. Just install Qualitytrace
  Just learning tracing! Install Qualitytrace, OpenTelemetry Collector and the sample app.
```

After choosing this option, the installer will check if your Docker installation is ok on your machine and will proceed to the next step.

### 3. Select a default installation or an installation with sample app

In this step, you can choose to install just Qualitytrace alone or install it with a sample app. By seeing the following options:

```text title="Expected output:"
Do you have OpenTelemetry based tracing already set up, or would you like us to install a demo tracing environment and app? [type to search]: 
  I have a tracing environment already - Just install Qualitytrace.
> Just learning tracing! Install Qualitytrace, OpenTelemetry Collector and the sample app.
```

By choosing any option, this installer will create a `quality-trace` directory in the current directory and will add a `docker-compose.yaml` file to it.
If you choose the first one, the `docker-compose.yaml` will have only Qualitytrace and its dependencies. By choosing the second, a sample app called [Pokeshop](../live-examples/pokeshop/overview.md) will be installed with Qualitytrace, allowing you to create some tests against it in the future.

For demonstration purposes, we will choose `Just learning tracing! Install Qualitytrace, OpenTelemetry Collector and the sample app.` option.

### 4. Finish the installation

Qualitytrace will proceed with the installation and show how to start it.

```text title="Expected output:"
-> Thanks! We are ready to install Qualitytrace now

 SUCCESS  Install successful!

To start quality-trace:

        docker compose -f quality-trace/docker-compose.yaml  up -d

Then, use your browser to navigate to:

  http://localhost:11633

Happy Qualitytracing =)
```

### 5. Start Docker Compose

```bash
docker compose -f quality-trace/docker-compose.yaml up -d
```

```bash title="Condensed expected output from the Qualitytrace container:"
Starting quality-trace ...
...
2022/11/28 18:24:09 HTTP Server started
...
```

### 6. Open the Qualitytrace Web UI

Open [`http://localhost:11633`](http://localhost:11633) in your browser.

Create a [test](../web-ui/creating-tests.md).

:::info
Running a test against `localhost` will resolve as `127.0.0.1` inside the Qualitytrace container. To run tests against apps running on your local machine, add them to the same network and use service name mapping instead. Example: Instead of running an app on `localhost:8080`, add it to your Docker Compose file, connect it to the same network as your Qualitytrace service, and use `service-name:8080` in the URL field when creating an app.

You can reach services running on your local machine using:

- Linux (docker version < 20.10.0): `172.17.0.1:8080`
- MacOS (docker version >= 18.03) and Linux (docker version >= 20.10.0): `host.docker.internal:8080`
:::

## Kubernetes

This guide showcases using the Kubernetes option in the Qualitytrace CLI.

### 1. Run the `server install` Command

Once you've installed the CLI you can install a Qualitytrace server by running:

```bash
quality-trace server install
```
<!--
```text title="Expected output:"
████████ ██████   █████   ██████ ███████ ████████ ███████ ███████ ████████ 
   ██    ██   ██ ██   ██ ██      ██         ██    ██      ██         ██    
   ██    ██████  ███████ ██      █████      ██    █████   ███████    ██    
   ██    ██   ██ ██   ██ ██      ██         ██    ██           ██    ██    
   ██    ██   ██ ██   ██  ██████ ███████    ██    ███████ ███████    ██    
-->
Version: v0.10.1


Hi! Welcome to the Qualitytrace server installer. I'll help you set up your Qualitytrace server by asking you a few questions
and configuring your system with all the requirements, so you can start TraceTesting right away!

To get more info about Qualitytrace, you can check our docs at https://intelops.github.io/quality-trace/

If you have any issues, please let us know by creating an issue (https://github.com/intelops/quality-trace/issues/new/choose)
<!--or reach us on Discord https://discord.gg/6zupCZFQbe-->


How do you want to run Qualitytrace? [type to search]: 
  Using Docker Compose
> Using Kubernetes
```

### 2. Select Kubernetes

This prompts you to select if you want to get either a Docker Compose or Kubernetes setup generated for you.

Select `Using Kubernetes`.

```text title="Expected output:"
...

How do you want to run Qualitytrace?:
  > Using Kubernetes


-> Let's check if your system has everything we need

✔ kubectl already installed
✔ helm already installed

-> Your system is ready! Now, let's configure Qualitytrace

Do you have OpenTelemetry based tracing already set up, or would you like us to install a demo tracing environment and app? [type to search]:
> I have a tracing environment already. Just install Qualitytrace
  Just learning tracing! Install Qualitytrace, OpenTelemetry Collector and the sample app.
```

After choosing this option, the installer will check if your kubectl and Helm installations are okay on your machine and will proceed to the next step.

### 3. Select a default installation or an installation with sample app

In this step, you can choose to install just Qualitytrace alone or install it with a sample app. By seeing the following options:

```text title="Expected output:"
Do you have OpenTelemetry based tracing already set up, or would you like us to install a demo tracing environment and app? [type to search]: 
  I have a tracing environment already. Just install Qualitytrace
> Just learning tracing! Install Qualitytrace, OpenTelemetry Collector and the sample app.
```

By choosing any option, this installer will create a `quality-trace` namespace with all Qualitytrace related containers.
If you choose the first one, the `quality-trace` namespace will have only Qualitytrace and its dependencies. By choosing the second, two things will be added. First, a sample app called [Pokeshop](../live-examples/pokeshop/overview.md) will be installed in a `demo` namespace, allowing you to create some tests against it. Second, an OpenTelemetry Collector will be installed in the `quality-trace` namespace to receive traces from the demo app and forward them to Qualitytrace.

For demonstration purposes, we will choose `Just learning tracing! Install Qualitytrace, OpenTelemetry Collector and the sample app.` option.

### 4. Finish the installation

Qualitytrace will proceed with the installation and show how to start it.

```text title="Expected output:"
SUCCESS  Install successful!

To access quality-trace:

	kubectl --kubeconfig <path-to-your-home>/.kube/config --context <your-cluster-context> --namespace quality-trace port-forward svc/quality-trace 11633

Then, use your browser to navigate to:

  http://localhost:11633

Happy Qualitytracing =)
```

### 5. Port forward the Qualitytrace service

```bash
kubectl --kubeconfig <path-to-your-home>/.kube/config --context <your-cluster-context> --namespace quality-trace port-forward svc/quality-trace 11633
```

```bash title="Expected output..."
Forwarding from 127.0.0.1:11633 -> 11633
Forwarding from [::1]:11633 -> 11633
```

### 6. Open the Qualitytrace Web UI

Open [`http://localhost:11633`](http://localhost:11633) in your browser.

Create a [test](../web-ui/creating-tests.md).

:::info
Running a test against `localhost` will resolve as `127.0.0.1` inside the Qualitytrace container. To run tests against apps running on your local machine, add them to the same network and use service name mapping instead. Example: Instead of running an app on `localhost:8080`, add it to your Docker Compose file, connect it to the same network as your Qualitytrace service, and use `service-name:8080` in the URL field when creating an app.

You can reach services running on your local machine using:

- Linux (docker version < 20.10.0): `172.17.0.1:8080`
- MacOS (docker version >= 18.03) and Linux (docker version >= 20.10.0): `host.docker.internal:8080`
:::

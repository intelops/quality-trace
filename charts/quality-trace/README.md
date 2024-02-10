# Qualitytrace

This is the Helm chart for [Qualitytrace](https://github.com/intelops/quality-trace) installation.

## Installation

### Chart installation

Add repo:

```sh
helm repo add quality-trace https://intelops.github.io/quality-trace
helm repo update

```

```sh
helm install quality-trace quality-trace/quality-trace
```

## Uninstall

```sh
helm delete quality-trace
```

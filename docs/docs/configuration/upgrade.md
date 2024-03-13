# Upgrade Qualitytrace Version

This page explains how to upgrade the version of your Qualitytrace Server and CLI.

If you've ever seen this error, you've come to the right page:

```text
✖️ Error: Version Mismatch
The CLI version and the server version are not compatible. To fix this, you'll need to make sure that both your CLI and server are using compatible versions.
We recommend upgrading both of them to the latest available version. 
Thank you for using Qualitytrace! We apologize for any inconvenience caused.
```

This means your Qualitytrace CLI and Server versions must be the same.

```sh
quality-trace version
```

```text title="Expected output"
CLI: v0.11.9
Server: v0.11.9
✔️ Version match
```

## Upgrade Qualitytrace CLI Version

### Linux/MacOS

Run the Qualitytrace CLI install script to upgrade to the latest version of the CLI:

```sh
curl -L https://raw.githubusercontent.com/intelops/quality-trace/main/install-cli.sh | bash
```

### Homebrew

```sh
brew upgrade
brew update
brew install intelops/quality-trace/quality-trace
```

### APT

```sh
sudo apt-get update
sudo apt-get install quality-trace
```

### YUM

```sh
sudo yum update
sudo yum install quality-trace --refresh
```

### Windows

```sh
choco source add --name=intelops_repo --source=https://chocolatey.intelops.io/chocolatey ; choco upgrade quality-trace
```

## Upgrade Qualitytrace Server Version

Make sure to match the CLI version you have installed to the Server version.

```sh
intelops/quality-trace:vX.X.X
```

If you are using version `v0.11.9` of the CLI, make sure to use the same version of the server.

```yaml
image: intelops/quality-trace:v0.11.9
```

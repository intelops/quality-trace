# CLI Installation Reference

This page contains a reference of all options for installing Qualitytrace CLI.

Qualitytrace has a command line interface (CLI) which includes an **install wizard** that helps to install the Qualitytrace server into Docker or Kubernetes. The CLI can also be used to run tests, download or upload tests, and manage much of the capability of Qualitytrace.

:::tip Want more info?
Read more about the installation guide [here](../getting-started/installation.mdx).
:::

## Installing the Qualitytrace CLI in different operating systems

Every time we release a new version of Qualitytrace, we generate binaries for Linux, MacOS, and Windows. Supporting both amd64, and ARM64 architectures, in `tar.gz`, `deb`, `rpm` and `exe` formats.

You can find the latest version [here](https://github.com/intelops/quality-trace/releases/latest).

### Linux/MacOS

Qualitytrace CLI can be installed automatically using the following script:

```sh
curl -L https://raw.githubusercontent.com/intelops/quality-trace/main/install-cli.sh | bash
```

It works for systems with Homebrew, `apt-get`, `dpkg`, `yum`, `rpm` installed, and if no package manager is available, it will try to download the build and install it manually.

You can also manually install it with one of the following methods.

#### Homebrew

```sh
brew install intelops/quality-trace/quality-trace
```

#### APT

```sh
# requirements for our deb repo to work
sudo apt-get update && sudo apt-get install -y apt-transport-https ca-certificates

# add repo
echo "deb [trusted=yes] https://apt.fury.io/quality-trace/ /" | sudo tee /etc/apt/sources.list.d/fury.list

# update and install
sudo apt-get update
sudo apt-get install quality-trace
```

#### YUM

```sh
# add repository
cat <<EOF | sudo tee /etc/yum.repos.d/quality-trace.repo
[quality-trace]
name=Qualitytrace
baseurl=https://yum.fury.io/quality-trace/
enabled=1
gpgcheck=0
EOF

# install
sudo yum install quality-trace --refresh
```

### Windows

#### Chocolatey

```bash
choco source add --name=intelops_repo --source=https://chocolatey.intelops.io/chocolatey ; choco install quality-trace
```

#### From source

Download one of the files from the latest tag, extract to your machine, and then [add the quality-trace binary to your PATH variable](https://stackoverflow.com/a/41895179).

## Installing a specific version of the Qualitytrace CLI

You can request to install a specific version by appending `-s -- [version]` to the installation script.

```bash
curl -L https://raw.githubusercontent.com/intelops/quality-trace/main/install-cli.sh | bash -s -- [version]
```

If you would want version `v0.13.0` you would run this command:

```bash
curl -L https://raw.githubusercontent.com/intelops/quality-trace/main/install-cli.sh | bash -s -- v0.12.1
```

This method skips package managers and directly downloads the build from the release page.

Due to this, if the Qualitytrace CLI was previously installed using a package manager, that version will still exist in the system, so depending on the `$PATH` environment variable configuration, that version might take precedence over the specific version you installed with the install script.

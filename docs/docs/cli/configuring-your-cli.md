# Configuring your CLI

Our web interface makes it easier to visualize your traces and add assertions, but sometimes a CLI is needed for automation. The CLI was developed for users creating tests and executing them each time a change is made in the system, so Qualitytrace can detect regressions and check your Service Level Objectives (SLOs).

## Available Commands

Here is a list of all available commands and how to use them:

### Configure

Configure your CLI to connect to your Qualitytrace server.

**How to Use**:

```sh
quality-trace configure
```

If you want to set values without having to answer questions from a prompt, you can provide the flag `--endpoint` to define the server endpoint.

```sh
quality-trace configure --endpoint http://my-quality-trace-server:11633
```

### Test List

Allows you to list all tests.

**How to Use**:

```sh
quality-trace list test
```

### Run a Test

Allows you to run a test by referencing a [test definition file](./creating-tests).

> Note: If the definition file contains the field `id`, this command will not create a new test. Instead, it will update the test with that ID. If that test doesn't exist, a new one will be created with that ID on the server.

Every time the test is run, changes are detected and, if any change is introduced, we use Tractest's [versioning](../concepts/versioning) mechanism to ensure that it will not cause problems with previous test runs.

**How to Use**:

```sh
quality-trace run test --file <file-path>
```

### Running Qualitytrace CLI from Docker

There are times when it is easier to directly execute the Qualitytrace CLI from a Docker image rather than installing the CLI on your local machine. This can be convenient when you wish to execute the CLI in a CI/CD environment.

**How to Use**:

Use the command below, substituting the following placeholders:

- `your-quality-trace-server-url` - The URL to the running Qualitytrace server you wish to execute the test on. Example: `http://localhost:11633/`
- `file-path` - The path to the saved Qualitytrace test. Example: `./mytest.yaml`

```bash wordWrap=true
docker run --rm -it -v$(pwd):$(pwd) -w $(pwd) --network host --entrypoint quality-trace intelops/quality-trace:latest -s <your-quality-trace-server-url> run test --file <file-path>
```

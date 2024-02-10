#/bin/bash

set -e

export TAG=${TAG:-"latest"}
export TEST_ENV=${TEST_ENV:-"local"}

if [ $TEST_ENV = "local" ]; then
  export QUALITYTRACE_ENDPOINT="localhost:11633"
  export QUALITYTRACE_CLI_COMMAND=$QUALITYTRACE_CLI
else
  export QUALITYTRACE_ENDPOINT="host.docker.internal:11633"
  export QUALITYTRACE_CLI_COMMAND="docker run --volume $PWD/tests:/app/tests --entrypoint quality-trace --add-host=host.docker.internal:host-gateway intelops/quality-trace:$TAG"
fi

echo "Preparing to run CLI tests..."
echo ""

echo "Environment variables considered on this run:"
echo "TAG:                   $TAG"
echo "TEST_ENV:              $TEST_ENV"
echo "QUALITYTRACE_ENDPOINT:    $QUALITYTRACE_ENDPOINT"
echo "QUALITYTRACE_CLI_COMMAND: $QUALITYTRACE_CLI_COMMAND"
echo ""

echo "Setting up quality-trace CLI configuration..."
cat << EOF > tests/config.yml
scheme: http
endpoint: $QUALITYTRACE_ENDPOINT
analyticsEnabled: false
EOF
echo "quality-trace CLI set up."
echo ""

echo "Setting up test helpers..."

run_cli_command() {
  args=$1

  $QUALITYTRACE_CLI_COMMAND --config ./tests/config.yml $args
  return $?
}

echo "Test helpers set."
echo ""

echo "Starting tests..."

EXIT_STATUS=0

run_cli_command '--help' || EXIT_STATUS=$?
run_cli_command 'version' || EXIT_STATUS=$?
run_cli_command 'run test --file ./tests/simple-test.yaml' || EXIT_STATUS=$?

echo ""
echo "Tests done! Exit code: $EXIT_STATUS"

exit $EXIT_STATUS

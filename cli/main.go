package main

import (
	"github.com/kubeshop/tracetest/cli/cmd"
	logger "github.com/sirupsen/logrus"
)

func main() {
	logger.Info("Executing command..")
	cmd.Execute()
}

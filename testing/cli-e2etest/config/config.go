package config

import "os"

type EnvironmentVars struct {
	EnableCLIDebug   bool
	QualitytraceCommand string
	TestEnvironment  string
}

var instance *EnvironmentVars

func GetConfigAsEnvVars() *EnvironmentVars {
	if instance != nil {
		return instance
	}

	enableCLIDebug := (os.Getenv("ENABLE_CLI_DEBUG") == "true")

	qualitytraceCommand := os.Getenv("QUALITYTRACE_CLI")
	if qualitytraceCommand == "" {
		qualitytraceCommand = "quality-trace"
	}

	testEnvironment := os.Getenv("TEST_ENVIRONMENT")
	if testEnvironment == "" {
		testEnvironment = "jaeger"
	}

	return &EnvironmentVars{
		EnableCLIDebug:   enableCLIDebug,
		QualitytraceCommand: qualitytraceCommand,
		TestEnvironment:  testEnvironment,
	}
}

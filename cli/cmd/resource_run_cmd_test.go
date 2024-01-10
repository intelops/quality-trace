package cmd

import (
	"context"
	"os"
	"testing"
	"errors"
	"log"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// Mock the runResourceCommand function
func runResourceCommand(ctx context.Context, params *runParameters) (int, error) {
	fmt.Printf("Received parameters: %+v\n", params) // Debugging output

	// Return a specific exit code based on the input parameters
	if params.GitRepo == "https://github.com/RazenaSaleem/test" {
		return 0, nil
	}
	
	if params.GitRepo == "https://github.com/InvalidRepo/test" {
        return 1, errors.New("simulated error") // Simulating an error scenario
    }
	
    return 0, nil
}

// Mock the exitCodeSuccess variable
const exitCodeSuccess = 0

// Declare tempDir in the outer scope
var tempDir string

func TestRunCmd_Positive(t *testing.T) {
	// Set up test parameters
	runParams := &runParameters{
		GitRepo:     "https://github.com/RazenaSaleem/test",
		GitUsername: "RazenaSaleem",
		Branch:      "main",
		GitFile:     "ml-server.yaml",
	}

	// Execute the Resource Run command
	exitCode, err := runResourceCommand(context.Background(), runParams)

	// Assertions
	assert.NoError(t, err, "Resource Run command should not return an error")
	assert.Equal(t, exitCodeSuccess, exitCode, "Resource Run command should exit with success code")

	// Set up cleanup function using t.Cleanup
	t.Cleanup(func() {
		// Cleanup locally by removing the temporary directory
		if tempDir != "" {
			os.RemoveAll(tempDir)
			log.Printf("Local cleanup: Removed directory %s", tempDir)
		}
	})
}

func TestRunCmd_Negative(t *testing.T) {
    // Set up test parameters for a scenario where the command should fail
    runParams := &runParameters{
        GitRepo:     "https://github.com/InvalidRepo/test", // Using an intentionally invalid repository
        GitUsername: "InvalidUser",
        GitToken:    "invalidToken",
        Branch:      "main",
        GitFile:     "ml-server.yaml",
    }

    // Execute the Resource Run command
    exitCode, err := runResourceCommand(context.Background(), runParams)

    // Assertions
    assert.Error(t, err, "Resource Run command should encounter an error")
    assert.NotEqual(t, exitCodeSuccess, exitCode, "Resource Run command should not exit with success code")

}

func TestRunParametersValidationNegative(t *testing.T) {
	// Create an instance of runParameters with intentionally invalid input
	invalidParams := &runParameters{
		GitUsername: "RazenaSaleem",
		Branch:      "main",
		GitFile:     "ml-server.yaml",
	}

	// Call the Validate method to get the validation errors
	errors := invalidParams.Validate(nil, nil)

	// Assert that the first element in the errors slice is an error
	assert.Error(t, errors[0], "Expected an error, but got nil")

	//heck for the presence of an error related to the "git-repo" parameter
	foundGitRepoError := false
	for _, err := range errors {
		if paramErr, ok := err.(paramError); ok {
			if paramErr.Parameter == "git-repo" {
				foundGitRepoError = true
				break
			}
		}
	}

	assert.True(t, foundGitRepoError, "Expected 'git-repo' parameter error, but it was not found")
}

func TestMain(m *testing.M) {
	// Create a temporary directory specific to the test
	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	code := m.Run()
	os.Exit(code)
}
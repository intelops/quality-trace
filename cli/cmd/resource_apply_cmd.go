package cmd

import (
	"context"
	"fmt"

	"github.com/kubeshop/tracetest/cli/pkg/fileutil"
	"github.com/kubeshop/tracetest/cli/pkg/resourcemanager"
	"github.com/spf13/cobra"
)

var (
	applyParams = &applyParameters{}
	applyCmd    *cobra.Command
)

func init() {
	applyCmd = &cobra.Command{
		GroupID: cmdGroupResources.ID,
		Use:     "apply " + resourceList(),
		Short:   "Apply resources",
		Long:    "Apply (create/update) resources to your qt server",
		PreRun:  setupCommand(),
		Run: WithResourceMiddleware(func(_ *cobra.Command, args []string) (string, error) {
			resourceType := resourceParams.ResourceName
			ctx := context.Background()

			resourceClient, err := resources.Get(resourceType)
			if err != nil {
				return "", err
			}

			resultFormat, err := resourcemanager.Formats.GetWithFallback(output, "yaml")
			if err != nil {
				return "", err
			}

			inputFile, err := fileutil.Read(applyParams.DefinitionFile)
			if err != nil {
				return "", fmt.Errorf("cannot read file %s: %w", applyParams.DefinitionFile, err)
			}

			result, err := resourceClient.Apply(ctx, inputFile, resultFormat)
			if err != nil {
				return "", err
			}

			return result, nil
		}, applyParams),
		PostRun: teardownCommand,
	}

	applyCmd.Flags().StringVarP(&applyParams.DefinitionFile, "file", "f", "", "path to the definition file")
	rootCmd.AddCommand(applyCmd)
}

type applyParameters struct {
	DefinitionFile string
	GitRepo        string
	GitUsername    string
	GitToken       string
	RepoName       string
	Branch         string
	GitFile        string
}

func (p applyParameters) Validate(cmd *cobra.Command, args []string) []error {
	errors := make([]error, 0)

	if p.DefinitionFile == "" && (p.GitRepo == "" || p.GitUsername == "" || p.GitToken == "" || p.RepoName == "" || p.Branch == "" || p.GitFile == "") {
		errors = append(errors, paramError{
			Parameter: "file/git parameters",
			Message:   "Either a definition file or Git parameters (GitRepo, GitUsername, GitToken, RepoName, Branch, GitFile) must be provided",
		})
	}

	return errors
}

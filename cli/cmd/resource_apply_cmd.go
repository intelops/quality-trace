package cmd

import (
	"context"
	"fmt"
	"encoding/json"

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

			// Debug: Print the value of the "gitrepo" flag
			fmt.Println("Git Repo Flag:", applyParams.GitRepo)

			resourceClient, err := resources.Get(resourceType)
			if err != nil {
				return "", err
			}

			resultFormat, err := resourcemanager.Formats.GetWithFallback(output, "yaml")
			if err != nil {
				return "", err
			}
			
			// Check if a definition file is provided
			if applyParams.DefinitionFile != "" {
			// If a file name is provided, read its contents
        		inputFile, err := fileutil.Read(applyParams.DefinitionFile)
				if err != nil {
					return "", fmt.Errorf("cannot read file %s: %w", applyParams.DefinitionFile, err)
				}

				result, err := resourceClient.Apply(ctx, inputFile, resultFormat)
				if err != nil {
					return "", err
				}

				return result, nil
			} else {
				// If no file name is provided, use Git parameters as JSON body
				gitParams, err := resourceClient.ApplyWithGitParameters(ctx, resultFormat)
                if err != nil {
                    return "", err
                }
				// Convert the map to JSON
				jsonBody, err := json.Marshal(gitParams)
				if err != nil {
					return "", fmt.Errorf("error creating JSON body: %w", err)
				}
		
				return string(jsonBody), nil
			}
		}, applyParams),
		PostRun: teardownCommand,
	}

	applyCmd.Flags().StringVarP(&applyParams.DefinitionFile, "file", "f", "", "path to the definition file")
	applyCmd.Flags().StringVarP(&applyParams.GitRepo, "gitrepo", "", "", "Git repository name")
	applyCmd.Flags().StringVarP(&applyParams.GitUsername, "gitusername", "", "", "Git username")
	applyCmd.Flags().StringVarP(&applyParams.GitToken, "gittoken", "", "", "Git token")
	applyCmd.Flags().StringVarP(&applyParams.RepoName, "reponame", "", "", "Repository name")
	applyCmd.Flags().StringVarP(&applyParams.Branch, "branch", "", "", "Branch name")
	applyCmd.Flags().StringVarP(&applyParams.GitFile, "gitfile", "", "", "Git file name")

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

    if p.DefinitionFile == "" {
        gitErrors := p.validateGitParameters()
        errors = append(errors, gitErrors...)
    }

    return errors
}

func (p applyParameters) validateGitParameters() []error {
    gitErrors := make([]error, 0)

    // Add specific validation checks for Git parameters
    if p.GitRepo == "" {
        gitErrors = append(gitErrors, paramError{
            Parameter: "git-repo",
            Message:   "Git repository is required",
        })
    }
	if p.GitUsername == "" {
        gitErrors = append(gitErrors, paramError{
            Parameter: "gitusername",
            Message:   "Git username is required",
        })
    }

    if p.GitToken == "" {
        gitErrors = append(gitErrors, paramError{
            Parameter: "gittoken",
            Message:   "Git token is required",
        })
    }

    if p.RepoName == "" {
        gitErrors = append(gitErrors, paramError{
            Parameter: "reponame",
            Message:   "Repository name is required",
        })
    }

    if p.Branch == "" {
        gitErrors = append(gitErrors, paramError{
            Parameter: "branch",
            Message:   "Branch name is required",
        })
    }

    if p.GitFile == "" {
        gitErrors = append(gitErrors, paramError{
            Parameter: "gitfile",
            Message:   "Git file name is required",
        })
   }
	return gitErrors
}
package cmd

import (
	"context"
	"errors"
	//"fmt"
	//"encoding/json"

	"github.com/kubeshop/tracetest/cli/pkg/resourcemanager"
	"github.com/spf13/cobra"
)

var (
	getParams = &resourceIDParameters{}
	getCmd    *cobra.Command
)

func init() {
	getCmd = &cobra.Command{
		GroupID: cmdGroupResources.ID,
		Use:     "get " + resourceList(),
		Short:   "Get resource",
		Long:    "Get a resource from your QT server",
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
			
			// Check if Git parameters are provided
			if getParams.GitRepo != "" || getParams.Username != "" || getParams.Token != "" || getParams.RepoName != "" || getParams.Branch != "" || getParams.GitFile != "" {
				// Construct Git parameters JSON
				gitParams := resourcemanager.GitParams{
					GitRepo:     getParams.GitRepo,
					Username:    getParams.Username,
					Token:       getParams.Token,
					RepoName:    getParams.RepoName,
					Branch:      getParams.Branch,
					GitFile:     getParams.GitFile,
				}
				// // Marshal GitParams to JSON
				// jsonBody, err := json.Marshal(gitParams)
				// if err != nil {
				// 	return "", fmt.Errorf("error creating JSON body: %w", err)
				// }

				// Call the Get method with GitParams
				result, err := resourceClient.GetWithGit(ctx, getParams.ResourceID, resultFormat, gitParams)
				if errors.Is(err, resourcemanager.ErrNotFound) {
					return result, nil
				}
				if err != nil {
					return "", err
				}

				return result, nil
			}
		
			// Call the Get method without GitParams
			result, err := resourceClient.Get(ctx, getParams.ResourceID, resultFormat)
			if errors.Is(err, resourcemanager.ErrNotFound) {
				return result, nil
			}
			if err != nil {
				return "", err
			}

			return result, nil
			}, getParams),
		PostRun: teardownCommand,
	}

	getCmd.Flags().StringVar(&getParams.ResourceID, "id", "", "id of the resource to get")
	getCmd.Flags().StringVar(&getParams.GitRepo, "gitrepo", "", "Git repository name")
	getCmd.Flags().StringVar(&getParams.Username, "gitusername", "", "Git username")
	getCmd.Flags().StringVar(&getParams.Token, "gittoken", "", "Git token")
	getCmd.Flags().StringVar(&getParams.RepoName, "reponame", "", "Repository name")
	getCmd.Flags().StringVar(&getParams.Branch, "branch", "", "Branch name")
	getCmd.Flags().StringVar(&getParams.GitFile, "gitfile", "", "Git file name")

	rootCmd.AddCommand(getCmd)
}

type resourceIDParameters struct {
	ResourceID string
	GitRepo    string
	Username string
	Token    string
	RepoName    string
	Branch      string
	GitFile     string
}

func (p resourceIDParameters) Validate(cmd *cobra.Command, args []string) []error {
	errors := make([]error, 0)

	if p.ResourceID == "" {
		errors = append(errors, paramError{
			Parameter: "id",
			Message:   "resource id must be provided",
		})
	}

	return errors
}

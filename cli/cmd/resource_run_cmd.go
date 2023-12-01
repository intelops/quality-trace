package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/kubeshop/tracetest/cli/openapi"
	"github.com/kubeshop/tracetest/cli/runner"
	"github.com/kubeshop/tracetest/cli/utils"
	"github.com/spf13/cobra"
)

var (
	runParams = &runParameters{}
	runCmd    *cobra.Command
)

func init() {
	runCmd = &cobra.Command{
		GroupID: cmdGroupResources.ID,
		Use:     "run " + runnableResourceList(),
		Short:   "run resources",
		Long:    "run resources",
		PreRun:  setupCommand(),
		Run: WithResourceMiddleware(func(_ *cobra.Command, args []string) (string, error) {
			resourceType := resourceParams.ResourceName
			ctx := context.Background()

			r, err := runnerRegistry.Get(resourceType)
			if err != nil {
				return "", fmt.Errorf("resource type '%s' cannot be run", resourceType)
			}

			orchestrator := runner.Orchestrator(
				cliLogger,
				utils.GetAPIClient(cliConfig),
				variableSetClient,
			)

			if runParams.EnvID != "" {
				runParams.VarsID = runParams.EnvID
			}

			runParams := runner.RunOptions{
				ID:              runParams.ID,
				DefinitionFile:  runParams.DefinitionFile,
				VarsID:          runParams.VarsID,
				SkipResultWait:  runParams.SkipResultWait,
				JUnitOuptutFile: runParams.JUnitOuptutFile,
				RequiredGates:   runParams.RequriedGates,
				GitRepo:         runParams.GitRepo,
   				GitUsername:     runParams.GitUsername,
   				GitToken:        runParams.GitToken,
   				RepoName:        runParams.RepoName,
   				Branch:          runParams.Branch,
				GitFile:         runParams.GitFile,
			}

			exitCode, err := orchestrator.Run(ctx, r, runParams, output)
			if err != nil {
				return "", err
			}

			ExitCLI(exitCode)

			// ExitCLI will exit the process, so this return is just to satisfy the compiler
			return "", nil

		}, runParams),
		PostRun: teardownCommand,
	}

	runCmd.Flags().StringVarP(&runParams.DefinitionFile, "file", "f", "", "path to the definition file")
	runCmd.Flags().StringVar(&runParams.ID, "id", "", "id of the resource to run")
	runCmd.Flags().StringVarP(&runParams.VarsID, "vars", "", "", "variable set file or ID to be used")
	runCmd.Flags().BoolVarP(&runParams.SkipResultWait, "skip-result-wait", "W", false, "do not wait for results. exit immediately after test run started")
	runCmd.Flags().StringVarP(&runParams.JUnitOuptutFile, "junit", "j", "", "file path to save test results in junit format")
	runCmd.Flags().StringSliceVar(&runParams.RequriedGates, "required-gates", []string{}, "override default required gate. "+validRequiredGatesMsg())

	runCmd.Flags().StringVarP(&runParams.GitRepo, "gitrepo", "", "", "Git repository name")
	runCmd.Flags().StringVarP(&runParams.GitUsername, "gitusername", "", "", "Git username")
	runCmd.Flags().StringVarP(&runParams.GitToken, "gittoken", "", "", "Git token")
	runCmd.Flags().StringVarP(&runParams.RepoName, "reponame", "", "", "Repository name")
	runCmd.Flags().StringVarP(&runParams.Branch, "branch", "", "", "Branch name")
	runCmd.Flags().StringVarP(&runParams.GitFile, "gitfile", "", "", "Git file name")

	//deprecated
	runCmd.Flags().StringVarP(&runParams.EnvID, "environment", "e", "", "environment file or ID to be used")
	runCmd.Flags().MarkDeprecated("environment", "use --vars instead")
	runCmd.Flags().MarkShorthandDeprecated("e", "use --vars instead")

	rootCmd.AddCommand(runCmd)
}

func validRequiredGatesMsg() string {
	opts := make([]string, 0, len(openapi.AllowedSupportedGatesEnumValues))
	for _, v := range openapi.AllowedSupportedGatesEnumValues {
		opts = append(opts, string(v))
	}

	return "valid options: " + strings.Join(opts, ", ")
}

type runParameters struct {
	ID              string
	DefinitionFile  string
	VarsID          string
	EnvID           string
	SkipResultWait  bool
	JUnitOuptutFile string
	RequriedGates   []string
	GitRepo         string
   	GitUsername     string
   	GitToken        string
   	RepoName        string
   	Branch          string
	GitFile         string
}

func (p runParameters) Validate(cmd *cobra.Command, args []string) []error {
	errs := []error{}
	if p.DefinitionFile == "" && p.ID == "" {
		errs = append(errs, paramError{
			Parameter: "resource",
			Message:   "you must specify a definition file or resource ID",
		})
	}

	if p.DefinitionFile != "" && p.ID != "" {
		errs = append(errs, paramError{
			Parameter: "resource",
			Message:   "you cannot specify both a definition file and resource ID",
		})
	}

	if p.JUnitOuptutFile != "" && p.SkipResultWait {
		errs = append(errs, paramError{
			Parameter: "junit",
			Message:   "--junit option is incompatible with --skip-result-wait option",
		})
	}

	// New checks for Git parameters
	if p.GitRepo == "" {
		errs = append(errs, paramError{
			Parameter: "git-repo",
			Message:   "you must specify a Git repository",
		})
	}

	if p.GitUsername == "" {
		errs = append(errs, paramError{
			Parameter: "git-username",
			Message:   "you must specify a Git username",
		})
	}
	if p.GitToken == "" {
		errs = append(errs, paramError{
			Parameter: "git-token",
			Message:   "you must specify a Git token",
		})
	}

	if p.RepoName == "" {
		errs = append(errs, paramError{
			Parameter: "repo-name",
			Message:   "you must specify a repository name",
		})
	}

	if p.Branch == "" {
		errs = append(errs, paramError{
			Parameter: "branch",
			Message:   "you must specify a branch name",
		})
	}

	if p.GitFile == "" {
		errs = append(errs, paramError{
			Parameter: "git-file",
			Message:   "you must specify a file name",
		})
	}

	for _, rg := range p.RequriedGates {
		_, err := openapi.NewSupportedGatesFromValue(rg)
		if err != nil {
			errs = append(errs, paramError{
				Parameter: "required-gates",
				Message:   fmt.Sprintf("invalid option '%s'. "+validRequiredGatesMsg(), rg),
			})
		}
	}

	return errs
}

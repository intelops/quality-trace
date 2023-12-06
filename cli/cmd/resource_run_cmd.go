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

			params := runner.RunOptions{
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

			exitCode, err := orchestrator.Run(ctx, r, params, output)
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

	// Print Git parameters for debugging
    fmt.Println("Git Repo:", p.GitRepo)
    fmt.Println("Git Username:", p.GitUsername)
    fmt.Println("Git Token:", p.GitToken)
    fmt.Println("Repo Name:", p.RepoName)
    fmt.Println("Branch:", p.Branch)
    fmt.Println("Git File:", p.GitFile)

	errs := []error{}
	// Check for incompatibility between JUnit and SkipResultWait options
    if p.JUnitOuptutFile != "" && p.SkipResultWait {
        errs = append(errs, paramError{
            Parameter: "junit",
            Message:   "--junit option is incompatible with --skip-result-wait option",
        })
    }

	if p.GitRepo != "" || p.GitUsername != "" || p.GitToken != "" || p.RepoName != "" || p.Branch != "" || p.GitFile != "" {
		// Call the validateGitParameters function
		gitErrors := p.validateGitParameters()
		errs = append(errs, gitErrors...)

	} else if p.DefinitionFile == "" && p.ID == "" {
		// Check for either DefinitionFile or ID
        errs = append(errs, paramError{
            Parameter: "resource",
            Message:   "you must specify a definition file or resource ID",
        })
		
    } else if p.DefinitionFile != "" && p.ID != "" {
        errs = append(errs, paramError{
            Parameter: "resource",
            Message:   "you cannot specify both a definition file and resource ID",
        })
    }
	

    // Validate required gates
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

func (p runParameters) validateGitParameters() []error {
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
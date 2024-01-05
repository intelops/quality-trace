package gitutil

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	logger "github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	http "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// RunParameters defines the parameters for the Git-related operations.
type RunParameters struct {
	GitRepo     string
	GitUsername string
	GitToken    string
	RepoName    string
	Branch      string
	GitFile     string
}

// CloneGitFile clones a file from a Git repository using the specified Git parameters.
func CloneGitFile(ctx context.Context, gitParams RunParameters) ([]byte, error) {

	logger.Info("Cloning Git repository...")

	// Set up Basic Authentication using custom BasicAuth struct
	auth := &http.BasicAuth{
		Username: gitParams.GitUsername,
		Password:    gitParams.GitToken,
	}

	repo, err := git.PlainCloneContext(ctx, "/tmp/new_01", false, &git.CloneOptions{
		URL: gitParams.GitRepo,
		Auth: auth,
	})
	if err != nil {
		logger.Error("Error during cloning:", err)
		return nil, fmt.Errorf("failed to clone Git repository: %v", err)
	}

	// Check out the specified branch
	wt, err := repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("failed to get worktree: %v", err)
	}
	err = wt.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName("refs/heads/" + gitParams.Branch),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to checkout branch: %v", err)
	}

	filePath := filepath.Join("/tmp/new_01", gitParams.GitFile)

	logger.Info("Git repository cloned successfully.")
	return []byte(filePath), nil
}

// CleanupClonedRepo deletes the cloned repository file.
func CleanupClonedRepo(filePath string) error {
	// Get the directory containing the file
    dirPath := filepath.Dir(filePath)

    // Remove the directory containing the file
    if err := os.RemoveAll(dirPath); err != nil {
        logger.Info("Failed to delete directory:", err)
        return err
    }

    logger.Info("Cloned repository file and its directory deleted successfully.")
    return nil
}
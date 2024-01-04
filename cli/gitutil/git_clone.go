package gitutil

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	logger "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
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

// BasicAuth is a basic authentication structure.
type BasicAuth struct {
	Username string
	Token    string
}

// NewBasicAuth creates a new BasicAuth instance with the provided username and password.
func NewBasicAuth(username, token string) *BasicAuth {
	return &BasicAuth{
		Username: username,
		Token:    token,
	}
}

// SetRequest sets the BasicAuth information in the request.
func (a *BasicAuth) SetRequest(req *http.Request) {
	authString := a.String()
	req.Header.Set("Authorization", authString)
}

// Name returns the name of the authentication method.
func (a *BasicAuth) Name() string {
	return "Basic"
}

// String returns the string representation of the authentication method.
func (a *BasicAuth) String() string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(a.Username+":"+a.Token))
}

// CloneGitFile clones a file from a Git repository using the specified Git parameters.
func CloneGitFile(ctx context.Context, gitParams RunParameters) ([]byte, error) {

	logger.Info("Cloning Git repository...")

	// Set up Basic Authentication using custom BasicAuth struct
	auth := &BasicAuth{
		Username: gitParams.GitUsername,
		Token:    gitParams.GitToken,
	}

	// Log authentication information for debugging

	logger.Info("Authentication String:", auth.String())
	// Before cloning
	logger.Info("Before cloning...")
	logger.Info("git repo ", gitParams.GitRepo)

	// Clone the Git repository
	repoURL := fmt.Sprintf("%s", gitParams.GitRepo)
	logger.Info("Repository URL:", repoURL)
	repo, err := git.PlainCloneContext(ctx, "/tmp/new_01", false, &git.CloneOptions{
		URL: repoURL,
		//Auth: auth,
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

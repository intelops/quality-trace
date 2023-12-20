package resourcemanager

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"encoding/json"
	"bytes"

	"github.com/Jeffail/gabs/v2"
	"github.com/kubeshop/tracetest/cli/pkg/fileutil"
	"go.uber.org/zap"
)

const VerbApply Verb = "apply"

type applyPreProcessorFn func(context.Context, fileutil.File) (fileutil.File, error)

func (c Client) validType(inputFile fileutil.File) error {
	c.logger.Debug("Validating resource type", zap.String("inputFile", inputFile.AbsPath()))

	var raw any
	err := (yamlFormat{}).Unmarshal(inputFile.Contents(), &raw)
	if err != nil {
		return fmt.Errorf("cannot unmarshal yaml: %w", err)
	}
	c.logger.Debug("Unmarshaled yaml", zap.Any("raw", raw))

	parsed := gabs.Wrap(raw)
	rawType := parsed.Path("type").Data()
	if rawType == nil {
		return errors.New("cannot find type in yaml")
	}
	c.logger.Debug("Found type", zap.String("type", fmt.Sprintf("%T", rawType)))
	t, ok := rawType.(string)
	if !ok {
		return fmt.Errorf("cannot parse type from yaml: %w", err)
	}
	c.logger.Debug("Parsed type", zap.String("type", t))

	if t != c.resourceType() && t != c.options.deprecatedAlias {
		return fmt.Errorf("cannot apply %s to %s resource", t, c.resourceType())
	}

	c.logger.Debug("resource type is valid")

	return nil

}

func (c Client) Apply(ctx context.Context, inputFile fileutil.File, requestedFormat Format) (string, error) {
	fmt.Println("Entering apply code")

	// A file is provided, follow the existing procedure
	if err := c.validType(inputFile); err != nil {
		println("Validation error:", err) //debug
		return "", err
	}

	c.logger.Debug("Applying resource",
		zap.String("format", requestedFormat.String()),
		zap.String("resource", c.resourceName),
		zap.String("inputFile", inputFile.AbsPath()),
		zap.String("contents", string(inputFile.Contents())),
	)

	if c.options.applyPreProcessor != nil {
		var err error
		inputFile, err = c.options.applyPreProcessor(ctx, inputFile)
		if err != nil {

			println("Preprocess error:", err) //debug
			return "", fmt.Errorf("cannot preprocess Apply request: %w", err)
		}
	}

	c.logger.Debug("preprocessed",
		zap.String("inputFile", inputFile.AbsPath()),
		zap.String("contents", string(inputFile.Contents())),
	)

	url := c.client.url(c.resourceNamePlural)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), inputFile.Reader())
	if err != nil {

		println("Build request error:", err) //debug
		return "", fmt.Errorf("cannot build Apply request: %w", err)
	}

	// we want the response inthe user's requested format
	err = requestedFormat.BuildRequest(req, VerbApply)
	if err != nil {

		println("Build request error:", err) //debug
		return "", fmt.Errorf("cannot build Apply request: %w", err)
	}

	// the files must be in yaml format, so we can safely force the content type,
	// even if it doesn't matcht he user's requested format
	req.Header.Set("Content-Type", (yamlFormat{}).ContentType())

	// final request looks like this:
		// PUT {server}/{resourceNamePlural}
		// Content-Type: text/yaml
		// Accept: {requestedFormat.contentType}
		//
		// {yamlFileContent}
		//
		// This means that we'll send the request body as YAML (read from the user provided file)
		// and we'll get the reponse in the users's requrested format.
	d, _ := httputil.DumpRequestOut(req, true)
	c.logger.Debug("apply request",
		zap.String("request", string(d)),
	)

	resp, err := c.client.do(req)
	if err != nil {

		println("Execute request error:", err) //debug
		return "", fmt.Errorf("cannot execute Apply request: %w", err)
	}
	defer resp.Body.Close()

	d, _ = httputil.DumpResponse(resp, true)
	c.logger.Debug("apply response",
		zap.String("response", string(d)),
	)

	if !isSuccessResponse(resp) {
		err := parseRequestError(resp, requestedFormat)

		println("Could not Apply resource:", err) //debug
		return "", fmt.Errorf("could not Apply resource: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		println("Read response error:", err) //debug
		return "", fmt.Errorf("cannot read Apply response: %w", err)
	}

	c.logger.Debug("file has id?", zap.Bool("hasID", inputFile.HasID()))
		// if the original file doesn't have an ID, we need to get the server generated ID from the response
		// and write it to the original file
	if !inputFile.HasID() {
		jsonBody, err := requestedFormat.ToJSON(body)
		if err != nil {

			println("Convert response body to JSON error:", err) //debug
			return "", fmt.Errorf("cannot convert response body to JSON format: %w", err)
		}

		parsed, err := gabs.ParseJSON(jsonBody)
		if err != nil {

			println("Parse Apply response error:", err) //debug
			return "", fmt.Errorf("cannot parse Apply response: %w", err)
		}

		id, ok := parsed.Path("spec.id").Data().(string)
		if !ok {
			println("Get ID from Apply response error:", err) //debug
			return "", fmt.Errorf("cannot get ID from Apply response")
		}
		c.logger.Debug("New ID", zap.String("id", id))

		inputFile, err = inputFile.SetID(id)
		if err != nil {
			println("Set ID on input file error:", err) //debug
			return "", fmt.Errorf("cannot set ID on input file: %w", err)
		}

		_, err = inputFile.Write()
		if err != nil {
			println("Write updated input file error:", err) //debug
			return "", fmt.Errorf("cannot write updated input file: %w", err)
		}
	}
	return requestedFormat.Format(string(body), c.options.tableConfig)
}

func (c *Client) ApplyWithGitParameters(ctx context.Context, requestedFormat Format) (string, error) {
	
	fmt.Println("Entering apply code")

	// Function to get Git parameters
	getGitParameters := func() (string, string, string, string, string, string) {
		// Check if Git parameters are provided through CLI flags
		gitRepo := c.options.cmd.Flags().Lookup("gitrepo").Value.String()
		gitUsername := c.options.cmd.Flags().Lookup("gitusername").Value.String()
		gitToken := c.options.cmd.Flags().Lookup("gittoken").Value.String()
		repoName := c.options.cmd.Flags().Lookup("reponame").Value.String()
		branch := c.options.cmd.Flags().Lookup("branch").Value.String()
		gitFile := c.options.cmd.Flags().Lookup("gitfile").Value.String()

		// Log Git parameters for debugging
		c.logger.Debug("Git Repo:", zap.String("gitRepo", gitRepo))
		c.logger.Debug("Git Username:", zap.String("gitUsername", gitUsername))
		c.logger.Debug("Git Token:", zap.String("gitToken", gitToken))
		c.logger.Debug("Repo Name:", zap.String("repoName", repoName))
		c.logger.Debug("Branch:", zap.String("branch", branch))
		c.logger.Debug("Git File:", zap.String("gitFile", gitFile))

		return gitRepo, gitUsername, gitToken, repoName, branch, gitFile
	}

	// Get Git parameters
	gitRepo, gitUsername, gitToken, repoName, branch, gitFile := getGitParameters()

	// construct JSON body with Git parameters and send it to the server
	c.logger.Debug("Applying resource with Git parameters",
		zap.String("format", requestedFormat.String()),
		zap.String("resource", c.resourceName),
		zap.String("gitRepo", gitRepo),
		zap.String("gitUsername", gitUsername),
		zap.String("gitToken", gitToken),
		zap.String("repoName", repoName),
		zap.String("branch", branch),
		zap.String("gitfile", gitFile),
	)

	// Construct JSON body
	gitParams := map[string]string{
		"gitRepo":     gitRepo,
		"gitUsername": gitUsername,
		"gitToken":    gitToken,
		"repoName":    repoName,
		"branch":      branch,
		"gitfile":     gitFile,
	}

	jsonBody, err := json.Marshal(gitParams)
	if err != nil {
		// Log JSON marshal error
		c.logger.Error("JSON marshal error", zap.Error(err))
		return "", fmt.Errorf("cannot marshal Git parameters to JSON: %w", err)
	}

	url := c.client.url(c.resourceNamePlural)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), bytes.NewReader(jsonBody))
	if err != nil {
		// Log build request error
		c.logger.Error("Build request error", zap.Error(err))
		return "", fmt.Errorf("cannot build Apply request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Execute the HTTP request
	resp, err := c.client.do(req)
	if err != nil {
		// Log execute request error
		c.logger.Error("Execute request error", zap.Error(err))
		return "", fmt.Errorf("cannot execute Apply request: %w", err)
	}
	defer resp.Body.Close()

	// Process the response
	if !isSuccessResponse(resp) {
		err := parseRequestError(resp, requestedFormat)
		// Log apply failure
		c.logger.Error("Could not Apply resource", zap.Error(err))
		return "", fmt.Errorf("could not Apply resource: %w", err)
	}

	// Returning the JSON body containing Git parameters
	return string(jsonBody), nil
}

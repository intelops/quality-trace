package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var (
	Version = "dev"
	Env     = "dev"
)

type Config struct {
	Scheme     string  `yaml:"scheme"`
	Endpoint   string  `yaml:"endpoint"`
	ServerPath *string `yaml:"serverPath,omitempty"`
}

func (c Config) URL() string {
	if c.Scheme == "" || c.Endpoint == "" {
		return ""
	}

	return fmt.Sprintf("%s://%s", c.Scheme, strings.TrimSuffix(c.Endpoint, "/"))
}

func (c Config) IsEmpty() bool {
	thisConfigJson, _ := json.Marshal(c)
	emptyConfigJson, _ := json.Marshal(Config{})

	return string(thisConfigJson) == string(emptyConfigJson)
}

func LoadConfig(configFile string) (Config, error) {
	fmt.Println("Loading configuration from:", configFile) // Debug line
	
	config, err := loadConfig(configFile)
	if err != nil {
		return Config{}, err
	}

	if !config.IsEmpty() {
		return config, nil
	}

	homePath, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("could not get user home path")
	}

	globalConfigPath := filepath.Join(homePath, ".tracetest/config.yml")
	return loadConfig(globalConfigPath)
}

func loadConfig(configFile string) (Config, error) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("tracetest")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	fmt.Println("Reading configuration from:", configFile) // Debug line
	if err := viper.ReadInConfig(); err != nil {
		// No configuration file found, return empty config
		return Config{}, nil
	}

	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("could not unmarshal config: %w", err)
	}

	return config, nil
}

func ValidateServerURL(serverURL string) error {
	fmt.Println("Validating server URL:", serverURL) // Debug line

	if !strings.HasPrefix(serverURL, "http://") && !strings.HasPrefix(serverURL, "https://") {
		return fmt.Errorf(`the server URL must start with the scheme, either "http://" or "https://"`)
	}

	return nil
}

func ParseServerURL(serverURL string) (scheme, endpoint string, err error) {
	fmt.Println("Parsing server URL:", serverURL) // Debug line

	urlParts := strings.Split(serverURL, "://")
	if len(urlParts) != 2 {
		return "", "", fmt.Errorf("invalid server url")
	}

	scheme = urlParts[0]
	endpoint = strings.TrimSuffix(urlParts[1], "/")

	return scheme, endpoint, nil
}

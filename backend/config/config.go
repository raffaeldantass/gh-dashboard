package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GithubAppID          string
	GithubPrivateKey     []byte
	GithubInstallationID string
	ServerPort           string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	privateKeyPath := os.Getenv("GITHUB_PRIVATE_KEY_PATH")
	privateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("error reading private key: %w", err)
	}

	config := &Config{
		GithubAppID:          os.Getenv("GITHUB_APP_ID"),
		GithubPrivateKey:     privateKey,
		GithubInstallationID: os.Getenv("GITHUB_INSTALLATION_ID"),
		ServerPort:           os.Getenv("SERVER_PORT"),
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfig(config *Config) error {
	if config.GithubAppID == "" {
		return fmt.Errorf("GITHUB_APP_ID is required")
	}
	if len(config.GithubPrivateKey) == 0 {
		return fmt.Errorf("GITHUB_PRIVATE_KEY is required")
	}
	if config.GithubInstallationID == "" {
		return fmt.Errorf("GITHUB_INSTALLATION_ID is required")
	}
	if config.ServerPort == "" {
		config.ServerPort = "8080" // default port
	}
	return nil
}

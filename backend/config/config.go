package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	GitHubAppID      string
	GitHubPrivateKey string
	Organizations    []string
	PersonalAccounts []string
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	log.Printf("Loading .env file")
	if err != nil {
		log.Println("Error loading .env file, using system environment")
	}

	return &Config{
		GitHubAppID:      os.Getenv("GITHUB_APP_ID"),
		GitHubPrivateKey: os.Getenv("GITHUB_PRIVATE_KEY"),
		Organizations:    strings.Split(os.Getenv("GITHUB_ORGANIZATIONS"), ","),
		PersonalAccounts: strings.Split(os.Getenv("GITHUB_PERSONAL_ACCOUNTS"), ","),
	}
}

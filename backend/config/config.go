package config

import (
	"golang.org/x/oauth2"
	"log"
	"os"
)

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	OAuth2Config *oauth2.Config
	Env          string
}

func Load() *Config {
	cfg := &Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GITHUB_REDIRECT_URL"),
		Env:          os.Getenv("APP_ENV"),
	}

	if cfg.Env == "" {
		cfg.Env = "development"
	}

	if cfg.ClientID == "" || cfg.ClientSecret == "" || cfg.RedirectURL == "" {
		log.Fatal("Missing required environment variables")
	}

	cfg.OAuth2Config = &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes: []string{
			"repo",
			"read:org",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	return cfg
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raffaeldantass/gh-dashboard/api/handlers"
	"github.com/raffaeldantass/gh-dashboard/api/routes"
	"github.com/raffaeldantass/gh-dashboard/config"
	"github.com/raffaeldantass/gh-dashboard/internal/github"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Create GitHub client
	githubClient := github.NewClient(
		cfg.GithubAppID,
		cfg.GithubInstallationID,
		cfg.GithubPrivateKey,
	)

	// Create handlers
	repoHandler := handlers.NewRepoHandler(githubClient)

	// Setup routes
	router := routes.SetupRoutes(repoHandler)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server starting on port %s", cfg.ServerPort)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

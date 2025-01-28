package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/raffaeldantass/gh-dashboard/internal/github"
)

type RepoHandler struct {
	githubClient *github.Client
}

func NewRepoHandler(githubClient *github.Client) *RepoHandler {
	return &RepoHandler{githubClient: githubClient}
}

func (h *RepoHandler) GetRepositories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse query parameters
	query := r.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	perPage, _ := strconv.Atoi(query.Get("per_page"))
	searchQuery := query.Get("query")
	org := query.Get("organization")

	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	// Get repositories
	repos, err := h.githubClient.GetRepositories(r.Context(), page, perPage, searchQuery, org)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write response
	if err := json.NewEncoder(w).Encode(repos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

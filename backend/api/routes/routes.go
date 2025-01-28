package routes

import (
	"net/http"

	"github.com/raffaeldantass/gh-dashboard/api/handlers"
)

func SetupRoutes(repoHandler *handlers.RepoHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/repositories", repoHandler.GetRepositories)

	// Add middleware for CORS
	return corsMiddleware(mux)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

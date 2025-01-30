package main

import (
	"github.com/google/go-github/v68/github"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

type Repository struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	LastUpdate  time.Time `json:"last_update"`
	IsPrivate   bool      `json:"is_private"`
	Owner       string    `json:"owner"`
}

var (
	config       Config
	oauth2Config *oauth2.Config
)

func init() {
	// Load configuration from environment variables
	config = Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GITHUB_REDIRECT_URL"),
	}

	// Debug print configuration (remove in production)
	log.Printf("ClientID: %s", config.ClientID)
	log.Printf("ClientSecret: %s", config.ClientSecret)
	log.Printf("RedirectURL: %s", config.RedirectURL)

	if config.ClientID == "" || config.ClientSecret == "" || config.RedirectURL == "" {
		log.Fatal("Missing required environment variables")
	}

	// Configure OAuth2
	oauth2Config = &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		Scopes: []string{
			"repo",
			"read:org",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
}

func main() {
	router := gin.Default()

	// Routes
	router.GET("/login", handleLogin)
	router.GET("/callback", handleCallback)
	router.GET("/repositories", authenticateToken(), getRepositories)

	log.Fatal(router.Run(":8080"))
}

// handleLogin initiates the GitHub OAuth flow
func handleLogin(c *gin.Context) {
	state := "random-state-string" // In production, generate this randomly
	url := oauth2Config.AuthCodeURL(state)
	log.Printf("Redirecting to GitHub: %s", url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// handleCallback processes the OAuth callback from GitHub
func handleCallback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	errorCallback := c.Query("error")
	errorDescription := c.Query("error_description")

	log.Printf("Callback received - Code: %s, State: %s", code, state)

	if errorCallback != "" {
		log.Printf("GitHub returned error: %s - %s", errorCallback, errorDescription)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":       errorCallback,
			"description": errorDescription,
		})
		return
	}

	if code == "" {
		log.Printf("No code received from GitHub")
		c.JSON(http.StatusBadRequest, gin.H{"error": "No code received"})
		return
	}

	token, err := oauth2Config.Exchange(c, code)
	if err != nil {
		log.Printf("Token exchange error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to exchange token",
			"details": err.Error(),
		})
		return
	}

	c.SetCookie("github_token", token.AccessToken, 3600, "/", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/repositories")
}

func authenticateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check header first
		token := c.GetHeader("Authorization")

		// If no header, check cookie
		if token == "" {
			cookieToken, err := c.Cookie("github_token")
			if err == nil {
				token = "Bearer " + cookieToken
			}
		}

		log.Printf("Token: %s", token)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token provided"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix if present
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		c.Set("access_token", token)
		c.Next()
	}
}

// getRepositories fetches repositories from GitHub
func getRepositories(c *gin.Context) {
	token := c.MustGet("access_token").(string)

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(c, ts)
	client := github.NewClient(tc)

	repos, _, err := client.Repositories.ListByAuthenticatedUser(c, &github.RepositoryListByAuthenticatedUserOptions{
		Type: "private",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch repositories"})
		return
	}

	orgs, _, err := client.Organizations.List(c, "", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch organizations"})
		return
	}

	log.Printf("Orgs: %s", orgs)
	var allRepos []Repository

	for _, repo := range repos {
		if repo.Private != nil && *repo.Private {
			allRepos = append(allRepos, Repository{
				Name:        *repo.Name,
				Description: getDescription(repo.Description),
				LastUpdate:  repo.UpdatedAt.Time,
				IsPrivate:   true,
				Owner:       *repo.Owner.Login,
			})
		}
	}

	for _, org := range orgs {
		orgRepos, _, err := client.Repositories.ListByOrg(c, *org.Login, &github.RepositoryListByOrgOptions{
			Type: "private",
		})
		if err != nil {
			continue
		}

		for _, repo := range orgRepos {
			if repo.Private != nil && *repo.Private {
				allRepos = append(allRepos, Repository{
					Name:        *repo.Name,
					Description: getDescription(repo.Description),
					LastUpdate:  repo.UpdatedAt.Time,
					IsPrivate:   true,
					Owner:       *org.Login,
				})
			}
		}
	}

	c.JSON(http.StatusOK, allRepos)
}

func getDescription(desc *string) string {
	if desc == nil {
		return ""
	}
	return *desc
}

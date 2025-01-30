package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/raffaeldantass/gh-dashboard/config"
	"log"
	"net/http"
)

func HandleLogin(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		state := "random-state-string" // In production, generate this randomly
		url := cfg.OAuth2Config.AuthCodeURL(state)
		log.Printf("Redirecting to GitHub: %s", url)
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func HandleCallback(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		errorCallback := c.Query("error")
		errorDescription := c.Query("error_description")

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

		token, err := cfg.OAuth2Config.Exchange(c, code)
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
}

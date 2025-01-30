package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/raffaeldantass/gh-dashboard/config"
	"log"
	"net/http"
)

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func HandleLogin(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate state"})
			return
		}

		c.SetCookie("oauth_state", state, 3600, "/", "", false, true)

		url := cfg.OAuth2Config.AuthCodeURL(state)
		log.Printf("Redirecting to GitHub: %s", url)
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func HandleCallback(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		returnedState := c.Query("state")
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

		savedState, err := c.Cookie("oauth_state")
		if err != nil || savedState != returnedState {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state parameter"})
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

		c.SetCookie("oauth_state", "", -1, "/", "", false, true)
		c.SetCookie("github_token", token.AccessToken, 3600, "/", "", false, true)
		c.Redirect(http.StatusTemporaryRedirect, "/repositories")
	}
}

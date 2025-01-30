package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/raffaeldantass/gh-dashboard/services"
	"net/http"
)

func GetRepositories() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.MustGet("access_token").(string)

		githubService := services.NewGitHubService(token)
		repos, err := githubService.GetAllRepositories(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, repos)
	}
}

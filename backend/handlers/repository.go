package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raffaeldantass/gh-dashboard/services"
)

func GetRepositories() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.MustGet("access_token").(string)

		page := 1
		perPage := 10

		if pageStr := c.Query("page"); pageStr != "" {
			if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
				page = p
			}
		}

		if perPageStr := c.Query("per_page"); perPageStr != "" {
			if pp, err := strconv.Atoi(perPageStr); err == nil && pp > 0 && pp <= 100 {
				perPage = pp
			}
		}

		githubService := services.NewGitHubService(token)
		result, err := githubService.GetAllRepositories(c, page, perPage)

		if err != nil {
			log.Printf("Failed to get repositories: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

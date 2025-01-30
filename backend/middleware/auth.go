package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthenticateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			cookieToken, err := c.Cookie("github_token")
			if err == nil {
				token = "Bearer " + cookieToken
			}
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token provided"})
			c.Abort()
			return
		}

		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		c.Set("access_token", token)
		c.Next()
	}
}

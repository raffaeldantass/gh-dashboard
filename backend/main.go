package main

import (
	"log"

	"github.com/gin-contrib/cors" // Change this import
	"github.com/gin-gonic/gin"
	"github.com/raffaeldantass/gh-dashboard/config"
	"github.com/raffaeldantass/gh-dashboard/handlers"
	"github.com/raffaeldantass/gh-dashboard/middleware"
)

func main() {
	cfg := config.Load()
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Trust only local network proxies
	router.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	// Use Gin's CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.FrontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // 12 hours
	}))

	router.GET("/api/login", handlers.HandleLogin(cfg))
	router.GET("/callback", handlers.HandleCallback(cfg))
	router.GET("/api/get-repositories", middleware.AuthenticateToken(), handlers.GetRepositories())

	log.Fatal(router.Run(":8080"))
}

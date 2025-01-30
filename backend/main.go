package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raffaeldantass/gh-dashboard/config"
	"github.com/raffaeldantass/gh-dashboard/handlers"
	"github.com/raffaeldantass/gh-dashboard/middleware"
	"log"
)

func main() {
	cfg := config.Load()
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.GET("/login", handlers.HandleLogin(cfg))
	router.GET("/callback", handlers.HandleCallback(cfg))
	router.GET("/repositories", middleware.AuthenticateToken(), handlers.GetRepositories())

	log.Fatal(router.Run(":8080"))
}

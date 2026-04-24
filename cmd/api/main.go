package main

import (
	"github.com/elisio-tech/jwt-autentication-go.git/internal/app/handlers"
	"github.com/elisio-tech/jwt-autentication-go.git/internal/database"
	"github.com/elisio-tech/jwt-autentication-go.git/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()

	r.POST("/register", handlers.Login)
	r.POST("/login", handlers.Login)

	api := r.Group("/api")
	api.Use(middleware.ValidateToken())
	{
		api.GET("/profile", handlers.Profile)
	}
	r.Run()
}

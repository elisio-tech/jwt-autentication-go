package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jwt-autentication-go.git/helper"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to JWT")
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}

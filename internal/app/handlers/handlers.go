package handlers

import (
	"net/http"

	"github.com/elisio-tech/jwt-autentication-go.git/internal/app/dto"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)

	ctx.JSON(200, gin.H{"message": "login"})
}

func Register(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "login"})
}

func Profile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	c.JSON(http.StatusOK, gin.H{"message": "Perfil acessado", "user_id": userID})
}

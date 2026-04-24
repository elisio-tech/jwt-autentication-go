package handlers

import (
	"net/http"

	"github.com/elisio-tech/jwt-autentication-go.git/internal/app/dto"
	"github.com/elisio-tech/jwt-autentication-go.git/internal/database"
	"github.com/elisio-tech/jwt-autentication-go.git/internal/domain/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)

	user := models.User{
		Username: req.UserName,
		Password: string(hashedPassword),
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Usuário criado com sucesso",
		"user":    user,
	})
}

func Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "login"})
}

func Profile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	c.JSON(http.StatusOK, gin.H{"message": "Perfil acessado", "user_id": userID})
}

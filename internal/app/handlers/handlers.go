package handlers

import (
	"net/http"
	"time"

	"github.com/elisio-tech/jwt-autentication-go.git/internal/app/dto"
	"github.com/elisio-tech/jwt-autentication-go.git/internal/database"
	"github.com/elisio-tech/jwt-autentication-go.git/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("minha-chave-secreta-super-segura")

func Register(ctx *gin.Context) {
	var request dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar senha"})
		return
	}

	user := models.User{
		UserName: request.UserName,
		Password: string(hashedPassword),
		Email:    request.Email,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Usuário ou email já cadastrado"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"user":    gin.H{"id": user.ID, "username": user.UserName, "email": user.Email},
	})
}

func Login(ctx *gin.Context) {
	var request dto.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	claims := jwt.MapClaims{
		"sub":  user.ID,
		"name": user.UserName,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Profile(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Perfil acessado", "user_id": userID})
}

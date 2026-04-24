package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("minha-chave-secreta-super-segura")

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Erro ao ler claims"})
			ctx.Abort()
			return
		}

		userID, exists := claims["user_id"]
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token sem user_id"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", userID)
		ctx.Next()
	}
}

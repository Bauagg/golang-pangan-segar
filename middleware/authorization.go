package middleware

import (
	"fmt"
	"net/http"
	"pangan-segar/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := utils.ExtractTokenFromHeader(ctx)

		if err != nil || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Token is missing or invalid",
			})

			ctx.Abort()
			return
		}

		claims, err := utils.VerifyToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Invalid token",
			})

			ctx.Abort()
			return
		}

		userID := fmt.Sprintf("%d", claims.ID)
		userRole := claims.Role
		userPhone := claims.Phone

		ctx.Set("userID", userID)
		ctx.Set("userRole", userRole)
		ctx.Set("userPhone", userPhone)

		ctx.Next()
	}
}

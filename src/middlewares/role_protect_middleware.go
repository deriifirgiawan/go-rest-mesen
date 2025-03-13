package middlewares

import (
	"net/http"
	"rest-app-pos/src/config"
	"rest-app-pos/src/dto"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RoleProtectMiddleware(role_id uint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			response := dto.ErrorResponse(http.StatusUnauthorized, "Authorization header required")
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 || tokenString[0] != "Bearer" {
			response := dto.ErrorResponse(http.StatusUnauthorized, "Invalid token format!!")
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}

		secret := config.AppConfig.JWT.Secret
		if secret == "" {
			secret = "defaultsecret"
		}

		token, err := jwt.Parse(tokenString[1], func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			response := dto.ErrorResponse(http.StatusUnauthorized, "Invalid or expired token")
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		roleFloat := claims["role"].(float64)

		roleID := uint(roleFloat)

		if roleID != role_id {
			response := dto.ErrorResponse(http.StatusUnauthorized, "Invalid or expired token")
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
package utils

import (
	"rest-app-pos/src/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type DecodeTokenType struct {
	User_ID uint
	ROLE_ID uint
}

func DecodeToken(c *gin.Context) DecodeTokenType {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWT.Secret), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	userID, _ := claims["user_id"].(float64)
	roleID, _ := claims["role"].(float64)

	return DecodeTokenType{
		User_ID: uint(userID),
		ROLE_ID: uint(roleID),
	}
}
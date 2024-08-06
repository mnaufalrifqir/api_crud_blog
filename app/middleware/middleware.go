package middleware

import (
	"api_crud_blog/app/model"
	"api_crud_blog/app/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "You are not login",
				Error:   "Token is empty",
			})
			return
		}

		fields := strings.Fields(tokenString)
		token := fields[1]

		id, role, err := ValidateToken(token, utils.GetConfig("ACCESS_TOKEN_SECRET"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "Failed to validate token",
				Error:   err.Error(),
			})
			return
		}

		c.Set("userId", id)
		c.Set("role", role)
		c.Next()
	}
}
 
func ValidateToken(token string, signedJWTKey string) (string, string, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})

	if err != nil {
		return "", "", fmt.Errorf("invalidate token: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return "", "", fmt.Errorf("invalid token claim")
	}

	return claims["sub"].(string), claims["role"].(string), nil
}
package utils

import (
	"api_crud_blog/app/model"
	"fmt"
	"time"
 
	"github.com/golang-jwt/jwt"
)
 
func GenerateAccessToken(payload model.User, secretJWTKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
 
	now := time.Now().UTC()
 
	expirationTime := time.Now().Add(time.Hour * 24 * 3)
 
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = payload.ID
	claims["exp"] = expirationTime.Unix()
	claims["role"] = payload.Role
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
 
	tokenString, err := token.SignedString([]byte(secretJWTKey))
 
	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}
 
	return tokenString, nil
}
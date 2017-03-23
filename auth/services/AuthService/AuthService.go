package AuthService

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fschr/go/auth/config"
	"github.com/fschr/go/auth/models"
)

func IssueToken(requestUser *models.User, password string) (newToken string, err error) {
	if requestUser.VerifyPassword(password) {
		newToken, err = GenerateJWTToken(requestUser.Username)
	} else {
		err = errors.New("Invalid password")
	}
	return newToken, err
}

func GenerateJWTToken(email string) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(72)).Unix()
	claims["username"] = email
	token.Claims = claims

	mConfig := config.DevConfig
	tokenString, err = token.SignedString([]byte(mConfig.Signing.SecretKey))
	return tokenString, err
}

func VerifyToken(tokenString string) bool {
	mConfig := config.DevConfig

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(mConfig.Signing.SecretKey), nil
	})
	return (err == nil && token.Valid)
}

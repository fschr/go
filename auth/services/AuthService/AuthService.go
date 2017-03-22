package AuthService

import (
	"../../models"
	"time"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	SecretKey = []byte("MasterOfNone")
)

func IssueToken(requestUser *models.User, password string) (newToken string, err error) {
	if requestUser.VerifyPassword(password) {
		newToken, err = GenerateJWTToken(requestUser.Username)
	}else{
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
	
	tokenString, err = token.SignedString(SecretKey)
	return tokenString, err
}

func VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
		return []byte(SecretKey), nil
	})
	return (err == nil && token.Valid)
}
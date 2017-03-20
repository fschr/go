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

func IssueToken(requestUser *models.User, password string) (string, error) {
	if requestUser.VerifyPassword(password) {
		newToken, err := GenerateJWTToken(requestUser.Username)
		if err != nil {
			return "", err
		}
		return newToken, nil
	}else{
		return "", errors.New("Invalid password")
	}
}

func GenerateJWTToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(72)).Unix()
	claims["id"] = email
	token.Claims = claims
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		panic(err)
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
		return []byte(SecretKey), nil
	})
	return (err == nil && token.Valid)
}
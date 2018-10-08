package services

import (
	"github.com/spf13/viper"
	jwt "github.com/dgrijalva/jwt-go"
)

type UserJWT struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

func CryptJWT(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &UserJWT{
		ID: id,
	})
	tokenString, err := token.SignedString([]byte(viper.GetString("jwt.secret_key")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func DecryptJWT(tokenString string) (*UserJWT, bool) {
	user := &UserJWT{}
	token, _ := jwt.ParseWithClaims(tokenString, user, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.secret_key")), nil
	})
	return user, token.Valid
}
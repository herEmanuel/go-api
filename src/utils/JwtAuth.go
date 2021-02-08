package utils

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CreateToken creates a new jwt token
func CreateToken(userID uint, secretKey string) (string, error) {

	claims := jwt.MapClaims{}
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	preToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := preToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.New("Could not sign the new token")
	}

	return token, nil
}

//VerifyToken checks wheter a token is valid or not
func VerifyToken(userToken string, secretKey string) (string, error) {

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(userToken, claims, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return "", errors.New(err.Error())
	}

	return claims["id"].(string), nil
}

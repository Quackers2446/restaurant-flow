package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const jwtExpMins = 30

var privateKey string = ""

func GenerateJWT(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(jwtExpMins * time.Minute)
	claims["iat"] = time.Now()
	claims["iss"] = "restaurantFlowAuth"
	claims["aud"] = "restaurantFlow"
	claims["sub"] = userId

	if privateKey == "" {
		readKey, err := os.ReadFile("../../jwtRS256-sample.key")

		if err != nil {
			return "", err
		}

		privateKey = string(readKey)
	}

	tokenString, err := token.SignedString(string(privateKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

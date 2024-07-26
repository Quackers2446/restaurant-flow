package util

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var publicKey []byte = nil

type Claims interface {
	jwt.Claims
	jwt.RegisteredClaims
}

func ValidateJWT(tokenString string) (*jwt.Token, *jwt.RegisteredClaims, error) {
	if publicKey == nil {
		readKey, err := os.ReadFile("./jwtRS256-sample.key.pub")

		if err != nil {
			return nil, nil, err
		}

		publicKey = readKey
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM(publicKey)
	}, jwt.WithValidMethods([]string{"RS256"}))

	if !token.Valid || err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)

	if !ok {
		return nil, nil, errors.New("could not get claims")
	}

	return token, claims, nil
}

// ValidateTokenHeader validates a JWT in the form "Bearer TOKEN"
func ValidateTokenHeader(headers *http.Header) (*jwt.Token, *jwt.RegisteredClaims, error) {
	tokenHeader, ok := (*headers)["Authorization"]

	if !ok {
		return nil, nil, errors.New("no token")
	}

	if len(tokenHeader) == 0 {
		return nil, nil, errors.New("no token")
	}

	words := strings.Fields(tokenHeader[0]) // Split by whitespace

	if words[0] != "Bearer" || len(words) != 2 {
		return nil, nil, errors.New("bad token")
	}

	return ValidateJWT(words[1])
}

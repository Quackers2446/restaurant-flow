package util

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const jwtExpMins = 30

var privateKey []byte = nil
var publicKey []byte = nil

type Claims interface {
	jwt.Claims
	jwt.RegisteredClaims
}

func GenerateJWT(userId string) (string, error) {
	if privateKey == nil {
		readKey, err := os.ReadFile("./jwtRS256-sample.key")

		if err != nil {
			return "", err
		}

		privateKey = readKey
	}

	token := jwt.New(jwt.SigningMethodRS256)

	token.Claims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpMins * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "restaurantFlowAuth",
		Audience:  []string{"restaurantFlow"},
		Subject:   userId,
	}

	paredKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)

	if err != nil {
		return "", errors.New("key error")
	}

	tokenString, err := token.SignedString(paredKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, *jwt.RegisteredClaims, error) {
	if publicKey == nil {
		readKey, err := os.ReadFile("./jwtRS256-sample.key.pub")

		if err != nil {
			return nil, nil, err
		}

		publicKey = readKey
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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
func ValidateTokenHeader(tokenHeader []string) (*jwt.Token, *jwt.RegisteredClaims, error) {
	if len(tokenHeader) == 0 {
		return nil, nil, errors.New("no token")
	}

	words := strings.Fields(tokenHeader[0]) // Split by whitespace

	if words[0] != "Bearer" || len(words) != 2 {
		return nil, nil, errors.New("bad token")
	}

	return ValidateJWT(words[1])
}

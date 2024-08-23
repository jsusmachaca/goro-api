package utils

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenstring string) error {
	key, fileErr := os.ReadFile("jwt.key.pub")

	if fileErr != nil {
		panic(fileErr)
	}
	pubKey, keyErr := jwt.ParseRSAPublicKeyFromPEM(key)

	if keyErr != nil {
		panic(keyErr)
	}

	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

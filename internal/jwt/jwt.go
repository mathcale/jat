package jwt

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
)

var hmacSecret = make([]byte, 0)

func Parse(receivedToken string) string {
	token, parseErr := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Fatalf("Parse error: %+v\n", parseErr)
	}

	val, claimsMarshalErr := json.MarshalIndent(claims, "", "  ")

	if claimsMarshalErr != nil {
		log.Fatalf("Claims parse error: %+v\n", claimsMarshalErr)
	}

	return string(val)
}

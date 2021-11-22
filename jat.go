package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

func main() {
	args := os.Args
	hmacSecret := make([]byte, 0)

	if len(args) < 2 {
		fmt.Println("No token provided!")
		os.Exit(1)
	}

	tokenString := args[1]

	if tokenString == "" {
		fmt.Println("No token provided!")
		os.Exit(2)
	}

	token, parseErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		val, marshalErr := json.MarshalIndent(claims, "", "  ")

		if marshalErr != nil {
			panic(marshalErr)
		}

		fmt.Println(string(val))
	} else {
		fmt.Printf("Parse error: %+v", parseErr)
		os.Exit(3)
	}
}

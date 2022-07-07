package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mathcale/jat/internal/jwt"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatalf("Not enough arguments.\n")
	}

	tokenString := args[1]

	if tokenString == "" {
		log.Fatalf("No token provided!\n")
		os.Exit(2)
	}

	parseResult := jwt.Parse(tokenString)
	fmt.Println(parseResult)
}

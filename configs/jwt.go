package configs

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
}

var JWT_KEY = []byte("cCM7rrsrmle6")

type JWTClaim struct {
	ID    uint
	Name  string
	Email string
	jwt.RegisteredClaims
}

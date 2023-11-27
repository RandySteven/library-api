package configs

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte(os.Getenv("JWT_TOKEN"))

type JWTClaim struct {
	UserId uint
	jwt.RegisteredClaims
}

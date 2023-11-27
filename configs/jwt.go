package configs

import (
	"os"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/response"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte(os.Getenv("JWT_TOKEN"))

type JWTClaim struct {
	response.UserLogin
	jwt.RegisteredClaims
}

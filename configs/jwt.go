package configs

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte("cCM7rrsrmle6")

type JWTClaim struct {
	User *models.User
	jwt.RegisteredClaims
}

package middleware

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/configs"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func validateToken(c *gin.Context) *configs.JWTClaim {
	// tokenString, err := c.Cookie("token")
	// if err != nil {
	// 	return nil
	// }
	tokenStringValidate := c.GetHeader("Authorization")

	if tokenStringValidate == "" {
		return nil
	}

	claims := &configs.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenStringValidate, claims, func(t *jwt.Token) (interface{}, error) {
		return configs.JWT_KEY, nil
	})

	if err != nil || !token.Valid {
		return nil
	}

	return claims
}

func AuthMiddleware(c *gin.Context) {
	claims := validateToken(c)

	if claims == nil {
		resp := response.Response{
			Errors: []string{"Unauthorized. Invalid token"},
		}
		// utils.ResponseHandler(c.Writer, http.StatusUnauthorized, resp)
		// c.Abort()
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	c.Set("user", claims.UserLogin)
	c.Next()
}

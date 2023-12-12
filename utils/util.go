package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/configs"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CheckOldAndNew(old, new string) string {
	if new == "" {
		return old
	}
	return new
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", apperror.NewErrPasswordTooLong()
	}
	return string(bytes), nil
}

func Validate(obj any) []string {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		var errs = make([]string, 0)
		for _, currErr := range err.(validator.ValidationErrors) {
			errMsg := fmt.Sprintf("%s field is %s", currErr.Field(), currErr.ActualTag())
			errs = append(errs, errMsg)
		}
		return errs
	}
	return nil
}

func ResponseHandler(res http.ResponseWriter, httpStatus int, resp response.Response) {
	res.Header().Set("content-type", "application/json")
	res.WriteHeader(httpStatus)
	json.NewEncoder(res).Encode(resp)
}

func ResponseEncoder(c *gin.Context, statusCode int, resp response.Response) {
	c.JSON(statusCode, resp)
}

func ValidateToken(tokenString string) *configs.JWTClaim {

	if tokenString == "" {
		return nil
	}

	claims := &configs.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return configs.JWT_KEY, nil
	})

	if err != nil || !token.Valid {
		return nil
	}

	return claims
}

func Encode(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	return string(jsonData)
}

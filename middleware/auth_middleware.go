package middleware

import (
	"net/http"
	"os"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	if os.Getenv("ENV_MODE") == "testing" {
		c.Next()
		return
	}

	tokenStringValidate := c.GetHeader("Authorization")[len("Bearer "):]

	claims := utils.ValidateToken(tokenStringValidate)

	if claims == nil {
		resp := response.Response{
			Errors: []string{apperror.NewErrUnauthorized().Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusUnauthorized, resp)
		c.Abort()
		return
	}

	c.Set("id", claims.ID)
	c.Set("name", claims.Name)
	c.Set("email", claims.Email)
	c.Next()
}

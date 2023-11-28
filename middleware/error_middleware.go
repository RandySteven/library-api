package middleware

import (
	"context"
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			if errors.Is(ginErr, context.DeadlineExceeded) {
				c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{"errors": ginErr.Err.Error()})
				return
			}
			apperror.ErrorChecker(c, ginErr.Err)
			return
		}
	}
}

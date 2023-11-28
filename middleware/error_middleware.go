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

		err := c.Errors.Last()
		if errors.Is(err, context.DeadlineExceeded) {
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{"errors": err.Error()})
			return
		}

		for _, ginErr := range c.Errors {
			apperror.ErrorChecker(c, ginErr.Err)
			return
		}
	}
}

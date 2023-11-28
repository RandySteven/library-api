package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func WithTimeOut(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)
	c.Next()
	// err := c.Request.Context().Err()
	// if errors.Is(err, context.DeadlineExceeded) {
	// 	c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{"errors": "request timeout"})
	// 	return
	// }
}

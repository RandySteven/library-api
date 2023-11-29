package middleware

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/logger"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			logger.Log.Info("Success")
			return
		}

		args := make(map[string]interface{})
		args["method"] = c.Request.Method
		args["uri"] = c.Request.RequestURI
		args["status"] = c.Writer.Status()

		logger.Log.WithFields(args).Error(err)

	}
}

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"octcarp/sustech/cs328/a2/api/grpc/logclient"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		latency := time.Since(start)
		logclient.Info(fmt.Sprintf("[%s] %s %s %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			latency,
		), "default_id")

		c.Next()
	}
}

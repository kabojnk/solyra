package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger middleware logs request details
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		// Format log message
		logMessage := fmt.Sprintf("[API] %s | %3d | %13v | %15s | %s %s",
			time.Now().Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)

		// Log based on status code
		switch {
		case statusCode >= 500:
			fmt.Println(logMessage, "| Error:", c.Errors.String())
		case statusCode >= 400:
			fmt.Println(logMessage, "| Warning:", c.Errors.String())
		default:
			fmt.Println(logMessage)
		}
	}
}

package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)
		log.Printf("Request: %s %s | Status: %d | Duration: %v", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	}
}

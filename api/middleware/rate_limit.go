package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimitMiddleware() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(1*time.Second), 10)
	return func(c *gin.Context) {
		if err := limiter.Wait(context.Background()); err != nil {
			c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests"})
			return
		}
		c.Next()
	}
}

type Rate struct {
	limit    int
	interval time.Duration
}

func NewRate(limit int, interval time.Duration) *Rate {
	return &Rate{
		limit:    limit,
		interval: interval,
	}
}

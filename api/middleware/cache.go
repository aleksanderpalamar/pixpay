package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	cache       = make(map[string]cacheEntry)
	cacheMutex  sync.RWMutex
	cacheExpiry = time.Duration(60 * time.Second)
)

type cacheEntry struct {
	response    []byte
	contentType string
	expiry      time.Time
}

func CacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cacheKey := c.Request.URL.Path

		cacheMutex.RLock()
		entry, found := cache[cacheKey]
		cacheMutex.RUnlock()

		if found && entry.expiry.After(time.Now()) {
			c.Data(http.StatusOK, entry.contentType, entry.response)
			c.Abort()
			return
		}

		responseWriter := &responseCapture{ResponseWriter: c.Writer, body: &[]byte{}}
		c.Writer = responseWriter

		c.Next()

		cacheMutex.Lock()
		cache[cacheKey] = cacheEntry{
			response:    *responseWriter.body,
			contentType: responseWriter.Header().Get("Content-Type"),
			expiry:      time.Now().Add(cacheExpiry),
		}
		cacheMutex.Unlock()
	}
}

type responseCapture struct {
	gin.ResponseWriter
	body *[]byte
}

func (r *responseCapture) Write(b []byte) (int, error) {
	*r.body = append(*r.body, b...)
	return r.ResponseWriter.Write(b)
}

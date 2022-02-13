package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type bucket struct {
	tokens uint
	time   time.Time
}

var maxTokens = uint(500)
var refillDuration = time.Minute
var refillAmount = maxTokens / 2

func ThrottleMiddleware() gin.HandlerFunc {
	buckets := map[string]*bucket{}

	return func(c *gin.Context) {
		user := c.ClientIP()
		b := buckets[user]

		if b == nil {
			buckets[user] = &bucket{tokens: maxTokens - 1, time: time.Now()}
			c.Next()
			return
		}

		refillInterval := uint(time.Since(b.time) / refillDuration)
		tokensAdded := refillAmount * refillInterval
		currentTokens := b.tokens + tokensAdded

		if currentTokens < 1 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			c.Abort()
			return
		}
		if currentTokens > maxTokens {
			b.time = time.Now()
			b.tokens = maxTokens - 1
		} else {
			deltaTokens := currentTokens - b.tokens
			deltaRefills := deltaTokens / refillAmount
			deltaTime := time.Duration(deltaRefills) * refillDuration

			b.time = b.time.Add(deltaTime)
			b.tokens = currentTokens - 1
		}
	}
}

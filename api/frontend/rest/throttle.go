package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	cache "github.com/hashicorp/golang-lru"
)

type bucket struct {
	tokens uint
	time   time.Time
}

var maxTokens = uint(500)
var refillDuration = time.Minute
var refillAmount = maxTokens / 2

// ThrottleMiddleware sets a throttle on how many requests a user can make in certain duration
func ThrottleMiddleware() gin.HandlerFunc {
	bucketCache, _ := cache.New(500)

	return func(c *gin.Context) {
		user := c.ClientIP()
		bucketResult, ok := bucketCache.Get(user)

		if !ok {
			bucketCache.Add(user, &bucket{tokens: maxTokens - 1, time: time.Now()})
			c.Next()
			return
		}

		b := bucketResult.(*bucket)

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

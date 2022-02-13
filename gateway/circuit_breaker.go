package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type circuit struct {
	mu         sync.RWMutex
	errorCount uint
	time       time.Time
}

var circuitErrorThreshold = uint(25)
var circuitTimeThreshold = 20 * time.Second

func (c *circuit) ShouldBreak() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.errorCount > circuitErrorThreshold
}

func (c *circuit) RecordError() {
	c.mu.Lock()
	defer c.mu.Unlock()

	passedIntervals := uint(time.Since(c.time) / circuitTimeThreshold)
	removedErrorCount := circuitErrorThreshold * passedIntervals

	if removedErrorCount > c.errorCount {
		c.time = time.Now()
		c.errorCount = 1
	} else {
		c.errorCount = c.errorCount - removedErrorCount + 1
	}
}

func CircuitBreakerMiddleware() gin.HandlerFunc {
	breaker := circuit{errorCount: 0, time: time.Now()}
	return func(c *gin.Context) {
		if breaker.ShouldBreak() {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "service unavailable"})
			c.Abort()
			return
		}
		c.Next()
		if c.Writer.Status() >= 500 {
			breaker.RecordError()
		}
	}
}

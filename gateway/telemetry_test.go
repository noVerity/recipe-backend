package main

import "github.com/gin-gonic/gin"

type NoopTelemetry struct{}

func NewNoopTelemetryManager() Telemetry {
	return NoopTelemetry{}
}

func (m NoopTelemetry) TracerMiddleware(endpoint string) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}

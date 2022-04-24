package main

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

func SetupUserService(router *gin.Engine, userService *url.URL, telemetry Telemetry) {
	breaker := CircuitBreakerMiddleware()
	router.Any("/user", breaker, telemetry.TracerMiddleware("/user"), ReverseProxy(userService))
	router.Any("/login", breaker, telemetry.TracerMiddleware("/login"), ReverseProxy(userService))
}

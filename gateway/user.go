package main

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

func SetupUserService(router *gin.Engine, userService *url.URL) {
	breaker := CircuitBreakerMiddleware()
	router.Any("/user", breaker, ReverseProxy(userService))
	router.Any("/login", breaker, ReverseProxy(userService))
}

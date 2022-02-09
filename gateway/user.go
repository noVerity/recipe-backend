package main

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

func SetupUserService(router *gin.Engine, userService *url.URL) {
	router.Any("/user", ReverseProxy(userService))
	router.Any("/login", ReverseProxy(userService))
}

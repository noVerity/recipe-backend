package main

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

func SetupUserService(router *gin.Engine, userService *url.URL) {
	router.Any("/user", ReverseProxy(userService), func(c *gin.Context) {
		fmt.Printf("HERE: %v\n", c.Request.URL.Host)
	})
	router.Any("/login", ReverseProxy(userService))
}

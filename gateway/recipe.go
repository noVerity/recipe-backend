package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"

	"github.com/gin-gonic/gin"
)

type RecipeRoute struct {
	Id string `uri:"id" binding:"required"`
}

type RecipeController struct {
	ShardMap map[string]Shard
}

func SetupRecipeService(router *gin.Engine, manager *AuthManager, shardMap ShardMap) {
	mapping := make(map[string]Shard)

	for _, shard := range shardMap.Map {
		mapping[shard.Name] = shard
	}

	controller := RecipeController{mapping}
	router.GET("/recipe/:id", controller.HandleIdRedirectRoute)
	router.POST("/recipe/:id", controller.HandleIdRedirectRoute)
	router.PUT("/recipe/:id", controller.HandleIdRedirectRoute)
	router.DELETE("/recipe/:id", controller.HandleIdRedirectRoute)
	router.GET("/recipe", manager.AuthMiddleware(), controller.HandleAuthRedirectRoute)
	router.POST("/recipe", manager.AuthMiddleware(), controller.HandleAuthRedirectRoute)
}

var shardIDParser = regexp.MustCompile("^([^_]+)_")

func (controller RecipeController) HandleIdRedirectRoute(c *gin.Context) {
	var uriElement RecipeRoute
	if err := c.BindUri(&uriElement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	match := shardIDParser.FindStringSubmatch(uriElement.Id)

	if len(match) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	shard, prs := controller.ShardMap[match[1]]
	if !prs {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	shardUrl, err := url.ParseRequestURI(shard.URL)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	director := func(req *http.Request) {
		shardUrl.Path = req.URL.Path
		req.URL = shardUrl
		req.Host = shardUrl.Host
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func (controller RecipeController) HandleAuthRedirectRoute(c *gin.Context) {
	shard, prs := controller.ShardMap[c.GetString("userShard")]

	if !prs {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	shardUrl, err := url.ParseRequestURI(shard.URL)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	director := func(req *http.Request) {
		shardUrl.Path = req.URL.Path
		req.URL = shardUrl
		req.Host = shardUrl.Host
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)
}

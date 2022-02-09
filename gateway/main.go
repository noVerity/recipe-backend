package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

type ShardMap struct {
	Map []Shard `json:"shards"`
}

type Shard struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	r := gin.Default()

	userService, err := url.ParseRequestURI(getenv("APP_USER_SERVICE", "localhost:9901"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid user service url (%v)\n", err)
		return
	}

	var recipeShards ShardMap
	shard := getenv("APP_RECIPE_SHARDS", "{\"shards\":[{\"name\":\"test\",\"url\":\"localhost:9902\"}]}")
	err = json.Unmarshal([]byte(shard), &recipeShards)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid recipe shard definition (%v)\n", err)
		return
	}

	manager := NewAuthManager(getenv("JWT_SECRET", "NON_SECRET_DEFAULT"))

	SetupUserService(r, userService)
	SetupRecipeService(r, manager, recipeShards)

	r.Run()
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

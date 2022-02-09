package main

import "encoding/json"

type ShardMap struct {
	Map []Shard `json:"shards"`
}

type Shard struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetShardMap() (ShardMap, error) {
	var recipeShards ShardMap
	shard := getenv("APP_RECIPE_SHARDS", "{\"shards\":[{\"name\":\"test\",\"url\":\"localhost:9902\"}]}")
	err := json.Unmarshal([]byte(shard), &recipeShards)
	return recipeShards, err
}

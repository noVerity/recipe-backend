package main

import (
	"adomeit.xyz/shared"
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRecipeService(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	manager := shared.NewAuthManager("TEST")

	// Create tokens for users of various shards
	tokenOne, _ := manager.GetToken("user1", "one")
	tokenTwo, _ := manager.GetToken("user2", "two")
	tokenUnknown, _ := manager.GetToken("user3", "three")

	// Create two fake shards
	backendOne, requestsOne, requesterOne := ProxyTester(t, router)
	defer backendOne.Close()
	backendUrl, _ := url.ParseRequestURI(backendOne.URL)
	backendTwo, requestsTwo, _ := ProxyTester(t, router)
	defer backendTwo.Close()
	backendUrlTwo, _ := url.ParseRequestURI(backendTwo.URL)
	SetupRecipeService(router, manager, ShardMap{[]Shard{{"one", backendUrl.String()}, {"two", backendUrlTwo.String()}}})

	// Requests with a valid id will go to the relevant shard
	resp := requesterOne(http.MethodGet, "/recipe/one_123456", tokenUnknown)
	req := <-requestsOne

	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/recipe/one_123456", req.URL.Path)
	assert.Equal(t, http.StatusOK, resp.Code)

	// List and creation requests withput id, will go to a service based on the user token
	resp = requesterOne(http.MethodPost, "/recipe", tokenTwo)
	req = <-requestsTwo

	assert.Equal(t, http.MethodPost, req.Method)
	assert.Equal(t, "/recipe", req.URL.Path)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Unknown shards in ids give a 404
	resp = requesterOne(http.MethodGet, "/recipe/three_123456", tokenOne)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

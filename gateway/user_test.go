package main

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupUserService(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	backend, requests, requester := ProxyTester(t, router)
	defer backend.Close()
	backendUrl, _ := url.ParseRequestURI(backend.URL)
	SetupUserService(router, backendUrl, NewNoopTelemetryManager())

	resp := requester(http.MethodGet, "/user", "")
	req := <-requests

	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/user", req.URL.Path)
	assert.Equal(t, http.StatusOK, resp.Code)

	resp = requester(http.MethodGet, "/login", "")
	req = <-requests

	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/login", req.URL.Path)
	assert.Equal(t, http.StatusOK, resp.Code)
}

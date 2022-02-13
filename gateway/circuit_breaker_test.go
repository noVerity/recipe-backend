package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestCircuitBreakerMiddleware(t *testing.T) {
	backendResponse := "I am the backend"
	count := 0
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(502)
		count = count + 1
		w.Write([]byte(backendResponse))
	}))
	backendUrl, _ := url.ParseRequestURI(backend.URL)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/test", CircuitBreakerMiddleware(), ReverseProxy(backendUrl))
	var res *closeNotifyingRecorder

	for i := 0; i <= 50; i++ {
		body := bytes.NewBufferString("")
		req, _ := http.NewRequest(http.MethodGet, "/test", body)
		req.Header.Set("Content-Type", "application/json")
		res = newCloseNotifyingRecorder()
		router.ServeHTTP(res, req)
	}

	assert.Equal(t, http.StatusServiceUnavailable, res.Code)
	assert.Equal(t, 26, count)
}

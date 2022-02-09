package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type closeNotifyingRecorder struct {
	*httptest.ResponseRecorder
	closed chan bool
}

func newCloseNotifyingRecorder() *closeNotifyingRecorder {
	return &closeNotifyingRecorder{
		httptest.NewRecorder(),
		make(chan bool, 1),
	}
}

func (c *closeNotifyingRecorder) close() {
	c.closed <- true
}

func (c *closeNotifyingRecorder) CloseNotify() <-chan bool {
	return c.closed
}

func ProxyTester(t *testing.T, router *gin.Engine) (*httptest.Server, chan *http.Request, func(method string, srcPath string, token string) *closeNotifyingRecorder) {
	backendResponse := "I am the backend"
	requests := make(chan *http.Request)

	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(backendResponse))
		requestCopy := r
		fmt.Printf("Arrived at %v\n", requestCopy.URL.Path)
		go func() {
			requests <- requestCopy
		}()
	}))

	requester := func(method string, srcPath string, token string) *closeNotifyingRecorder {
		fmt.Printf("Sending: %v\n", srcPath)
		body := bytes.NewBufferString("")
		req, _ := http.NewRequest(method, srcPath, body)
		req.Header.Set("Content-Type", "application/json")
		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}
		res := newCloseNotifyingRecorder()
		router.ServeHTTP(res, req)
		return res
	}
	return backend, requests, requester
}

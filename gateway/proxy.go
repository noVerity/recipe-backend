package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(target *url.URL) gin.HandlerFunc {
	rewriter := func(path string) string { return path }
	return ReverseRewriteProxy(target, rewriter)
}

func ReverseRewriteProxy(target *url.URL, rewritePath func(string) string) gin.HandlerFunc {
	return func(c *gin.Context) {
		director := func(req *http.Request) {
			//r := c.Request
			//req = r
			copy := target
			copy.Path = rewritePath(req.URL.Path)
			req.URL = copy
			// req.Header["my-header"] = []string{r.Header.Get("my-header")}
			// Golang camelcases headers
			//delete(req.Header, "My-Header")
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
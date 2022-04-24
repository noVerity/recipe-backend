package main

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
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
			targetCopy := target
			targetCopy.Path = rewritePath(req.URL.Path)
			req.URL = targetCopy
			req.Host = targetCopy.Host
			otel.GetTextMapPropagator().Inject(c.Request.Context(), propagation.HeaderCarrier(req.Header))
		}
		proxy := &httputil.ReverseProxy{
			Director: director,
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

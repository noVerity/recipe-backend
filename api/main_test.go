package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"adomeit.xyz/recipe/ent"
	"adomeit.xyz/recipe/ent/enttest"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func SetupTestORM(t *testing.T) (*ent.Client, func(method string, endpoint string, payload string) *httptest.ResponseRecorder) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	gin.SetMode(gin.ReleaseMode)
	router := SetupRouter(client, gin.New())
	requestTester := GetJSONRequestTester(router)

	return client, requestTester
}

func GetJSONRequestTester(router *gin.Engine) func(method string, endpoint string, payload string) *httptest.ResponseRecorder {
	return func(method string, endpoint string, payload string) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		var jsonStr = []byte(payload)
		req, _ := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		return w
	}
}

package main

import (
	"adomeit.xyz/shared"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"adomeit.xyz/user/ent"
	"adomeit.xyz/user/ent/enttest"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func SetupTestORM(t *testing.T) (*ent.Client, func(method string, endpoint string, validToken bool, payload string) *httptest.ResponseRecorder) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	gin.SetMode(gin.ReleaseMode)
	auth := shared.NewAuthManager("TEST_SECRET")
	r := gin.New()
	shardMap := ShardMap{[]Shard{{"one", "http://localhost"}, {"two", "http://localhost"}}}
	NewUserController(r, client, auth, &shardMap)
	requestTester := GetJSONRequestTester(r, auth)

	return client, requestTester
}

func GetJSONRequestTester(router *gin.Engine, auth *shared.AuthManager) func(method string, endpoint string, validToken bool, payload string) *httptest.ResponseRecorder {
	testToken, _ := auth.GetToken("Frodo", "one")

	return func(method string, endpoint string, validToken bool, payload string) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		var jsonStr = []byte(payload)
		req, _ := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		if validToken {
			req.Header.Set("Authorization", "Bearer "+testToken)
		}
		router.ServeHTTP(w, req)
		return w
	}
}

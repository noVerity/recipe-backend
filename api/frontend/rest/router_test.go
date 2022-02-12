package rest

import (
	"adomeit.xyz/recipe/core"
	"adomeit.xyz/recipe/ent"
	"adomeit.xyz/recipe/ent/enttest"
	"bytes"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupTestORM(t *testing.T) (*ent.Client, func(method string, endpoint string, payload string) *httptest.ResponseRecorder) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	gin.SetMode(gin.ReleaseMode)
	auth := NewAuthManager("TEST_SECRET")
	recipeCore := core.NewRecipeCore(client, nil)
	ingredientCore := core.NewIngredientCore(client, nil)
	router := gin.New()
	SetupRouter(router, auth, recipeCore, ingredientCore)
	requestTester := GetJSONRequestTester(router, auth)

	return client, requestTester
}

func GetJSONRequestTester(router *gin.Engine, auth *AuthManager) func(method string, endpoint string, payload string) *httptest.ResponseRecorder {
	testToken, _ := auth.GetToken("TestUser", "one")

	return func(method string, endpoint string, payload string) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		var jsonStr = []byte(payload)
		req, _ := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+testToken)
		router.ServeHTTP(w, req)
		return w
	}
}

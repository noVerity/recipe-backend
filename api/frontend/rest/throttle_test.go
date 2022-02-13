package rest

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestThrottleMiddleware(t *testing.T) {
	client, requestTester := SetupTestORM(t)
	defer client.Close()

	ingredientRoute := "/ingredient"

	for range [500]int{} {
		requestTester(
			http.MethodPost,
			ingredientRoute,
			`{}`,
		)
	}

	// Invalid payload: Invalid JSON
	w := requestTester(
		http.MethodPost,
		ingredientRoute,
		`{
			"name""Potatoes",
			"calories":0,
			"fat":0,
			"carbohydrates":0,
			"protein":0
		}`, // Missing : between name and Potatoes
	)
	assert.Equal(t, http.StatusTooManyRequests, w.Code)
}

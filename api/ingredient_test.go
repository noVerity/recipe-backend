package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupIngredientRoutes(t *testing.T) {
	client, requestTester := SetupTestORM(t)
	defer client.Close()

	ingredientRoute := "/ingredient"

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

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"invalid character '\"' after object key"}`, w.Body.String())

	// Create new ingredient
	w = requestTester(
		http.MethodPost,
		ingredientRoute,
		`{
			"name":"Potatoes",
			"calories":1,
			"fat":2,
			"carbohydrates":3.4,
			"protein":5
		}`,
	)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, `{"name":"Potatoes","calories":1,"fat":2,"carbohydrates":3.4,"protein":5}`, w.Body.String())

	// Try to create a duplicate ingredient
	w = requestTester(
		http.MethodPost,
		ingredientRoute,
		`{
			"name":"Potatoes",
			"calories":1,
			"fat":2,
			"carbohydrates":3.4,
			"protein":5
		}`,
	)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, `{"error":"ingredient with this name already exists"}`, w.Body.String())

	// Update on invalid URL
	w = requestTester(
		http.MethodPut,
		ingredientRoute+"/",
		`{
			"name":"Potatoes",
			"calories":2,
			"fat":2,
			"carbohydrates":3.4,
			"protein":5
		}`,
	)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"Key: 'URIElement.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`, w.Body.String())

	// Update with wrong types
	w = requestTester(
		http.MethodPut,
		ingredientRoute+"/pOtAtoEs",
		`{
			"name":"Potatoes",
			"calories":"2",
			"fat":2,
			"carbohydrates":3.4,
			"protein":5
		}`, // calories is a string instead of a number here
	)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"json: cannot unmarshal string into Go struct field Ingredient.calories of type float32"}`, w.Body.String())

	// Update non existing ingredient
	w = requestTester(
		http.MethodPut,
		ingredientRoute+"/cake", // The cake is a lie
		`{
			"name":"Potatoes",
			"calories":2,
			"fat":2,
			"carbohydrates":3.4,
			"protein":5
		}`,
	)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, `{"error":"ingredient does not exist"}`, w.Body.String())

	// Update our ingredient
	w = requestTester(
		http.MethodPut,
		ingredientRoute+"/pOtAtoEs",
		`{
			"name":"potatoeS",
			"calories":2,
			"fat":2,
			"carbohydrates":3.4,
			"protein":5
		}`, // case in the url should not matter
	)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"name":"potatoeS","calories":2,"fat":2,"carbohydrates":3.4,"protein":5}`, w.Body.String())

	// Retrieve our ingredient
	w = requestTester(http.MethodGet, ingredientRoute+"/pOtAtoEs", ``)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"name":"potatoeS","calories":2,"fat":2,"carbohydrates":3.4,"protein":5}`, w.Body.String())

	// Get all ingredients
	w = requestTester(http.MethodGet, ingredientRoute, ``)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"pagination":{"count":1,"offset":0},"data":[{"name":"potatoeS","calories":2,"fat":2,"carbohydrates":3.4,"protein":5}]}`, w.Body.String())

	// Page too far when getting all ingredients
	w = requestTester(http.MethodGet, ingredientRoute+"?offset=1000&limit=1", ``)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"pagination":{"count":1,"offset":1000},"data":[]}`, w.Body.String())

	// Retrieve non-existent ingredient
	w = requestTester(http.MethodGet, ingredientRoute+"/pizza", ``)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, `{"error":"ingredient not found"}`, w.Body.String())

	// Delete our ingredient
	w = requestTester(http.MethodDelete, ingredientRoute+"/pOtAtoEs", ``)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"name":"potatoeS","calories":2,"fat":2,"carbohydrates":3.4,"protein":5}`, w.Body.String())

	// Can't find the now non-existing ingredients
	w = requestTester(http.MethodDelete, ingredientRoute+"/pOtAtoEs", ``)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, `{"error":"ingredient not found"}`, w.Body.String())

	// Can't delete on root path
	w = requestTester(http.MethodDelete, ingredientRoute+"/", ``)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"Key: 'URIElement.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`, w.Body.String())

}

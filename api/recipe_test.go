package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupRecipeRoutes(t *testing.T) {
	client, requestTester := SetupTestORM(t)
	defer client.Close()

	recipeRoute := "/recipe"

	// Add some ingredients to the backend

	client.Ingredient.Create().
		SetName("Hope").
		SetCalories(999).
		SetFat(1).
		SetCarbohydrates(2).
		SetProtein(3).
		SaveX(context.Background())

	client.Ingredient.Create().
		SetName("Banana").
		SetCalories(123.4).
		SetFat(3).
		SetCarbohydrates(2).
		SetProtein(1).
		SaveX(context.Background())

	// Invalid payload: Invalid JSON
	w := requestTester(
		http.MethodPost,
		recipeRoute,
		`{
			"name""The Cake",
			"ingredientslist": "1 dose Hope\n1 tbsp Banana",
			"servings": 1,
			"instructions": "Throw everything together and hope for the best"
		}`, // Missing : between name and the cake
	)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"invalid character '\"' after object key"}`, w.Body.String())

	// Valid payload
	w = requestTester(
		http.MethodPost,
		recipeRoute,
		`{
			"name": "The Cake",
			"ingredientslist": "1 dose Hope\n1 tbsp Banana",
			"servings": 1,
			"instructions": "Throw everything together and hope for the best"
		}`,
	)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, `{"name":"The Cake","ingredientslist":"1 dose Hope\n1 tbsp Banana","instructions":"Throw everything together and hope for the best","nutrition":"Throw everything together and hope for the best","servings":1}`, w.Body.String())

	// Retrieve the created recipe
	w = requestTester(
		http.MethodGet,
		recipeRoute+"/the-cake",
		``,
	)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"name":"The Cake","ingredientslist":"1 dose Hope\n1 tbsp Banana","instructions":"Throw everything together and hope for the best","nutrition":"Throw everything together and hope for the best","servings":1,"ingredients":[{"name":"Hope","calories":999,"fat":1,"carbohydrates":2,"protein":3},{"name":"Banana","calories":123.4,"fat":3,"carbohydrates":2,"protein":1}]}`, w.Body.String())

	// Retrieve the all recipes
	w = requestTester(
		http.MethodGet,
		recipeRoute,
		``,
	)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"pagination":{"count":1,"offset":0},"data":[{"name":"The Cake","ingredientslist":"1 dose Hope\n1 tbsp Banana","instructions":"Throw everything together and hope for the best","nutrition":"Throw everything together and hope for the best","servings":1}]}`, w.Body.String())

	// Update recipe
	w = requestTester(
		http.MethodPut,
		recipeRoute+"/the-cake",
		`{
			"name": "The Cake",
			"ingredientslist": "1 dose Hope\n1 tbsp Banana",
			"servings": 4,
			"instructions": "Throw everything together and hope for the best"
		}`,
	)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"name":"The Cake","ingredientslist":"1 dose Hope\n1 tbsp Banana","instructions":"Throw everything together and hope for the best","nutrition":"Throw everything together and hope for the best","servings":4}`, w.Body.String())

}

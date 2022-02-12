package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupRecipeRoutes(t *testing.T) {
	os.Setenv("SHARD", "TEST")
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

	var recipe Recipe
	json.Unmarshal(w.Body.Bytes(), &recipe)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "The Cake", recipe.Name)
	assert.Equal(t, "1 dose Hope\n1 tbsp Banana", recipe.IngredientsList)
	assert.Equal(t, 1, recipe.Servings)
	assert.Equal(t, "Throw everything together and hope for the best", recipe.Instructions)

	// Retrieve the created recipe
	w = requestTester(
		http.MethodGet,
		recipeRoute+"/"+recipe.Id,
		``,
	)

	json.Unmarshal(w.Body.Bytes(), &recipe)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "The Cake", recipe.Name)
	assert.Equal(t, "1 dose Hope\n1 tbsp Banana", recipe.IngredientsList)
	assert.Equal(t, 1, recipe.Servings)
	assert.Equal(t, "Throw everything together and hope for the best", recipe.Instructions)

	// Retrieve the all recipes
	w = requestTester(
		http.MethodGet,
		recipeRoute,
		``,
	)

	var result RecipeResult
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, result.Pagination.Count)
	assert.Equal(t, 0, result.Pagination.Offset)
	assert.Equal(t, 1, len(result.Data))

	// Update recipe
	w = requestTester(
		http.MethodPut,
		recipeRoute+"/"+recipe.Id,
		`{
			"name": "The Cake",
			"ingredientslist": "1 dose Hope\n1 tbsp Banana",
			"servings": 4,
			"instructions": "Throw everything together and hope for the best"
		}`,
	)

	json.Unmarshal(w.Body.Bytes(), &recipe)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 4, recipe.Servings)

	// Delete recipe
	w = requestTester(
		http.MethodDelete,
		recipeRoute+"/"+recipe.Id,
		``,
	)

	json.Unmarshal(w.Body.Bytes(), &recipe)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "The Cake", recipe.Name)
	assert.Equal(t, "1 dose Hope\n1 tbsp Banana", recipe.IngredientsList)
	assert.Equal(t, 4, recipe.Servings)
	assert.Equal(t, "Throw everything together and hope for the best", recipe.Instructions)

}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIngredientList(t *testing.T) {
	ingredientList := ParseIngredientList(`
	1 tbsp Sugar
	2 tsp Purified Soul
	3.14 tsp Apple Pie
	# Do not fail on unknown lines
	`)

	assert.Equal(t, 3, len(ingredientList))

	assert.Equal(t, float32(1), ingredientList[0].Amount)
	assert.Equal(t, "Tablespoon", ingredientList[0].Measure)
	assert.Equal(t, "Sugar", ingredientList[0].Name)

	assert.Equal(t, float32(2), ingredientList[1].Amount)
	assert.Equal(t, "Teaspoon", ingredientList[1].Measure)
	assert.Equal(t, "Purified Soul", ingredientList[1].Name)

	assert.Equal(t, float32(3.14), ingredientList[2].Amount)
	assert.Equal(t, "Teaspoon", ingredientList[2].Measure)
	assert.Equal(t, "Apple Pie", ingredientList[2].Name)
}

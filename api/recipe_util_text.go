package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIngredientList(t *testing.T) {
	ingredientList := ParseIngredientList(`
	1 tbsp Sugar
	2 gallons Purified Soul
	3.14 pies Apple Pie
	# Do not fail on unknown lines
	`)

	assert.Equal(t, 3, len(ingredientList))

	assert.Equal(t, 1, ingredientList[0].Amount)
	assert.Equal(t, "tbsp", ingredientList[0].Measure)
	assert.Equal(t, "Sugar", ingredientList[0].Name)

	assert.Equal(t, 2, ingredientList[1].Amount)
	assert.Equal(t, "gallons", ingredientList[1].Measure)
	assert.Equal(t, "Purified Soul", ingredientList[1].Name)

	assert.Equal(t, 3.14, ingredientList[2].Amount)
	assert.Equal(t, "pies", ingredientList[2].Measure)
	assert.Equal(t, "Apple Pie", ingredientList[2].Name)
}

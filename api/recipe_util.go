package main

import (
	"strings"

	"github.com/noVerity/gofromto"
)

type IngredientEntry struct {
	Name    string
	Amount  float32
	Measure string
}

// ParseIngredientList takes a string and splits it into individual ingredients with their amounts
func ParseIngredientList(ingredientList string) []IngredientEntry {
	ingredients := strings.Split(ingredientList, "\n")
	var entryList []IngredientEntry
	for _, line := range ingredients {
		measure, err := gofromto.ParseMeasure(line)

		if err != nil {
			continue
		}

		entryList = append(entryList, IngredientEntry{
			Name:    measure.Name,
			Amount:  float32(measure.Amount),
			Measure: measure.Unit.String(),
		})
	}
	return entryList
}

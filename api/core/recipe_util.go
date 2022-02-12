package core

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
func ParseIngredientList(ingredientList string) []gofromto.Measure {
	ingredients := strings.Split(ingredientList, "\n")
	var entryList []gofromto.Measure
	for _, line := range ingredients {
		measure, err := gofromto.ParseMeasure(line)

		if err != nil {
			continue
		}

		entryList = append(entryList, measure)
	}
	return entryList
}

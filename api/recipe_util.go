package main

import (
	"regexp"
	"strconv"
	"strings"
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
	ingredientMatcher, _ := regexp.Compile(`^(\d+(?:\.\d+)?)\s+(\w+)\s+(.+)$`)
	for _, line := range ingredients {
		matched := ingredientMatcher.FindStringSubmatch(strings.TrimSpace(line))
		amount, err := strconv.ParseFloat(matched[1], 32)

		if err != nil || len(matched) < 4 {
			continue
		}

		entryList = append(entryList, IngredientEntry{
			Name:    matched[3],
			Amount:  float32(amount),
			Measure: matched[2],
		})
	}
	return entryList
}

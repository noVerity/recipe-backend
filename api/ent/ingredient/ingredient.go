// Code generated by entc, DO NOT EDIT.

package ingredient

const (
	// Label holds the string label denoting the ingredient type in the database.
	Label = "ingredient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCalories holds the string denoting the calories field in the database.
	FieldCalories = "calories"
	// FieldFat holds the string denoting the fat field in the database.
	FieldFat = "fat"
	// FieldCarbohydrates holds the string denoting the carbohydrates field in the database.
	FieldCarbohydrates = "carbohydrates"
	// FieldProtein holds the string denoting the protein field in the database.
	FieldProtein = "protein"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// EdgeRecipe holds the string denoting the recipe edge name in mutations.
	EdgeRecipe = "recipe"
	// Table holds the table name of the ingredient in the database.
	Table = "ingredients"
	// RecipeTable is the table that holds the recipe relation/edge. The primary key declared below.
	RecipeTable = "recipe_ingredients"
	// RecipeInverseTable is the table name for the Recipe entity.
	// It exists in this package in order to avoid circular dependency with the "recipe" package.
	RecipeInverseTable = "recipes"
)

// Columns holds all SQL columns for ingredient fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCalories,
	FieldFat,
	FieldCarbohydrates,
	FieldProtein,
	FieldSource,
}

var (
	// RecipePrimaryKey and RecipeColumn2 are the table columns denoting the
	// primary key for the recipe relation (M2M).
	RecipePrimaryKey = []string{"recipe_id", "ingredient_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

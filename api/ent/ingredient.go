// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"adomeit.xyz/recipe/ent/ingredient"
	"entgo.io/ent/dialect/sql"
)

// Ingredient is the model entity for the Ingredient schema.
type Ingredient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Calories holds the value of the "calories" field.
	Calories float32 `json:"calories,omitempty"`
	// Fat holds the value of the "fat" field.
	Fat float32 `json:"fat,omitempty"`
	// Carbohydrates holds the value of the "carbohydrates" field.
	Carbohydrates float32 `json:"carbohydrates,omitempty"`
	// Protein holds the value of the "protein" field.
	Protein float32 `json:"protein,omitempty"`
	// Source holds the value of the "source" field.
	Source *string `json:"source,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IngredientQuery when eager-loading is set.
	Edges IngredientEdges `json:"edges"`
}

// IngredientEdges holds the relations/edges for other nodes in the graph.
type IngredientEdges struct {
	// Recipe holds the value of the recipe edge.
	Recipe []*Recipe `json:"recipe,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RecipeOrErr returns the Recipe value or an error if the edge
// was not loaded in eager-loading.
func (e IngredientEdges) RecipeOrErr() ([]*Recipe, error) {
	if e.loadedTypes[0] {
		return e.Recipe, nil
	}
	return nil, &NotLoadedError{edge: "recipe"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ingredient) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case ingredient.FieldCalories, ingredient.FieldFat, ingredient.FieldCarbohydrates, ingredient.FieldProtein:
			values[i] = new(sql.NullFloat64)
		case ingredient.FieldID:
			values[i] = new(sql.NullInt64)
		case ingredient.FieldName, ingredient.FieldSource:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Ingredient", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ingredient fields.
func (i *Ingredient) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case ingredient.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case ingredient.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case ingredient.FieldCalories:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field calories", values[j])
			} else if value.Valid {
				i.Calories = float32(value.Float64)
			}
		case ingredient.FieldFat:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field fat", values[j])
			} else if value.Valid {
				i.Fat = float32(value.Float64)
			}
		case ingredient.FieldCarbohydrates:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field carbohydrates", values[j])
			} else if value.Valid {
				i.Carbohydrates = float32(value.Float64)
			}
		case ingredient.FieldProtein:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field protein", values[j])
			} else if value.Valid {
				i.Protein = float32(value.Float64)
			}
		case ingredient.FieldSource:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[j])
			} else if value.Valid {
				i.Source = new(string)
				*i.Source = value.String
			}
		}
	}
	return nil
}

// QueryRecipe queries the "recipe" edge of the Ingredient entity.
func (i *Ingredient) QueryRecipe() *RecipeQuery {
	return (&IngredientClient{config: i.config}).QueryRecipe(i)
}

// Update returns a builder for updating this Ingredient.
// Note that you need to call Ingredient.Unwrap() before calling this method if this Ingredient
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Ingredient) Update() *IngredientUpdateOne {
	return (&IngredientClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Ingredient entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Ingredient) Unwrap() *Ingredient {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ingredient is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Ingredient) String() string {
	var builder strings.Builder
	builder.WriteString("Ingredient(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", name=")
	builder.WriteString(i.Name)
	builder.WriteString(", calories=")
	builder.WriteString(fmt.Sprintf("%v", i.Calories))
	builder.WriteString(", fat=")
	builder.WriteString(fmt.Sprintf("%v", i.Fat))
	builder.WriteString(", carbohydrates=")
	builder.WriteString(fmt.Sprintf("%v", i.Carbohydrates))
	builder.WriteString(", protein=")
	builder.WriteString(fmt.Sprintf("%v", i.Protein))
	if v := i.Source; v != nil {
		builder.WriteString(", source=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Ingredients is a parsable slice of Ingredient.
type Ingredients []*Ingredient

func (i Ingredients) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}

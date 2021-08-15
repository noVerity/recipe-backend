package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Recipe struct {
	ent.Schema
}

// Fields of the Recipe.
func (Recipe) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			StructTag(`json:"id,omitempty"`),
		field.String("slug").
			Unique().
			NotEmpty(),
		field.String("name").
			NotEmpty(),
		field.String("ingredientslist"),
		field.String("instructions"),
		field.String("nutrition"),
		field.Int("servings").
			Min(1),
	}
}

// Edges of the Recipe.
func (Recipe) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ingredients", Ingredient.Type),
	}
}

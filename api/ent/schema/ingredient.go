package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Ingredient struct {
	ent.Schema
}

// Fields of the Ingredient.
func (Ingredient) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			StructTag(`json:"id,omitempty"`),
		field.String("name").
			Unique().
			NotEmpty(),
		field.Float32("calories"),
		field.Float32("fat"),
		field.Float32("carbohydrates"),
		field.Float32("protein"),
	}
}

// Edges of the Ingredient.
func (Ingredient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipe", Recipe.Type).
			Ref("ingredients"),
	}
}

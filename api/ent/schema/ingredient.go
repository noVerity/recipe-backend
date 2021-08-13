package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Ingredient struct {
	ent.Schema
}

// Fields of the User.
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

// Edges of the User.
func (Ingredient) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"os"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Recipe struct {
	ent.Schema
}

// Fields of the Recipe.
func (Recipe) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			DefaultFunc(func() string {
				shard_prefix := os.Getenv("SHARD")
				return shard_prefix + "_" + gonanoid.Must(22)
			}).
			StructTag(`json:"id,omitempty"`),
		field.String("slug").
			NotEmpty(),
		field.String("name").
			NotEmpty(),
		field.String("ingredientslist"),
		field.String("instructions"),
		field.String("nutrition"),
		field.String("user"),
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

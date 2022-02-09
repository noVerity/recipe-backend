package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

var validEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			StructTag(`json:"id,omitempty"`),
		field.String("username").
			Unique().
			NotEmpty(),
		field.String("email").
			Unique().
			NotEmpty().
			Match(validEmail),
		field.String("recipeShard"),
		field.String("password").
			NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

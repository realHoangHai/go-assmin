package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),

		field.String("email").Optional(),

		field.String("password").Optional(),

		field.String("phone").Optional(),

		field.Int("status").Optional(),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Common{},
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

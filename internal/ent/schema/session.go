package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}).Optional(),

		field.String("refresh_token").Optional().NotEmpty(),

		field.String("user_agent").Optional(),

		field.String("client_ip").Optional(),

		field.Bool("is_blocked").Optional(),

		field.Time("expire_time").Optional(),
	}
}

// Mixin of the Session.
func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Common{},
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return nil
}

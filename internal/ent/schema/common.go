package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"time"
)

// Common struct.
type Common struct {
	mixin.Schema
}

// Fields of the BaseMixin.
func (Common) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).
			Default(uuid.New).
			Unique(),

		field.Time("create_time").
			Immutable().
			Default(time.Now),

		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now).
			Immutable(),
	}
}

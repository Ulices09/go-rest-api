package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// -------------------------------------------------
// Mixin definition

// TimeMixin implements the ent.Mixin for sharing
// time fields with package schemas.
type TimeMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").
			Immutable().
			Default(time.Now),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

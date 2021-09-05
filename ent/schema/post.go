package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("text"),
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("posts").
			Unique(),
	}
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "post"},
	}
}

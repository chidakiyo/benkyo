package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Hoge holds the schema definition for the Hoge entity.
type Hoge struct {
	ent.Schema
}

// Fields of the Hoge.
func (Hoge) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
	}
}

// Edges of the Hoge.
func (Hoge) Edges() []ent.Edge {
	return nil
}

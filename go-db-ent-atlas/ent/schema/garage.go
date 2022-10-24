package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Garage holds the schema definition for the Car entity.
type Garage struct {
	ent.Schema
}

// Fields of the Car.
func (Garage) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int64("capacity"),
	}
}

// Edges of the Car.
func (Garage) Edges() []ent.Edge {
	return nil
}

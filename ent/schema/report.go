package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Report holds the schema definition for the Report entity.
type Report struct {
	ent.Schema
}

// Fields of the Report.
func (Report) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
	}
}

// Edges of the Report.
func (Report) Edges() []ent.Edge {
	return nil
}

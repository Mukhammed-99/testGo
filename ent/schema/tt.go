package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// Tt holds the schema definition for the Tt entity.
type Tt struct {
	ent.Schema
}

// Fields of the Tt.
func (Tt) Fields() []ent.Field {
	return nil
}

// Edges of the Tt.
func (Tt) Edges() []ent.Edge {
	return []ent.Edge{
	edge.To("id_tt", Pet.Type).
		StorageKey(edge.Column("id_tt")),
	}
}

package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// User schema.
type User struct {
	ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_at").Default(time.Now),
	}
}

// Edges of the user.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type).
			// Set the column name in the "pets" table for O2M relationship.
			StorageKey(edge.Column("user_id")),
			
	}
}

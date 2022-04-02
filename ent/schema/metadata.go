package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Metadata holds the schema definition for the Metadata entity.
type Metadata struct {
	ent.Schema
}

// Fields of the Metadata.
func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("key"),
		field.String("value"),
	}
}

// Edges of the Metadata.
func (Metadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("artifact", Artifact.Type).
			Ref("metadata").Unique(),
	}
}

func (Metadata) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key").Edges("artifact").Unique(),
	}
}

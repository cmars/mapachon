package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Artifact holds the schema definition for the Artifact entity.
type Artifact struct {
	ent.Schema
}

// Fields of the Artifact.
func (Artifact) Fields() []ent.Field {
	return []ent.Field{
		field.String("file_path"),
		field.String("archive_path").Optional(),
		field.String("file_type"),
		field.String("content").SchemaType(map[string]string{
			dialect.SQLite: "TEXT",
		}),
	}
	return nil
}

// Edges of the Artifact.
func (Artifact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata", Metadata.Type),
	}
}

func (Artifact) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("file_path", "archive_path"),
	}
}

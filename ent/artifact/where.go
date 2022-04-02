// Code generated by entc, DO NOT EDIT.

package artifact

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cmars/mapachon/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FilePath applies equality check predicate on the "file_path" field. It's identical to FilePathEQ.
func FilePath(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFilePath), v))
	})
}

// ArchivePath applies equality check predicate on the "archive_path" field. It's identical to ArchivePathEQ.
func ArchivePath(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArchivePath), v))
	})
}

// FileType applies equality check predicate on the "file_type" field. It's identical to FileTypeEQ.
func FileType(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFileType), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// FilePathEQ applies the EQ predicate on the "file_path" field.
func FilePathEQ(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFilePath), v))
	})
}

// FilePathNEQ applies the NEQ predicate on the "file_path" field.
func FilePathNEQ(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFilePath), v))
	})
}

// FilePathIn applies the In predicate on the "file_path" field.
func FilePathIn(vs ...string) predicate.Artifact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artifact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFilePath), v...))
	})
}

// FilePathNotIn applies the NotIn predicate on the "file_path" field.
func FilePathNotIn(vs ...string) predicate.Artifact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artifact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFilePath), v...))
	})
}

// FilePathGT applies the GT predicate on the "file_path" field.
func FilePathGT(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFilePath), v))
	})
}

// FilePathGTE applies the GTE predicate on the "file_path" field.
func FilePathGTE(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFilePath), v))
	})
}

// FilePathLT applies the LT predicate on the "file_path" field.
func FilePathLT(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFilePath), v))
	})
}

// FilePathLTE applies the LTE predicate on the "file_path" field.
func FilePathLTE(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFilePath), v))
	})
}

// FilePathContains applies the Contains predicate on the "file_path" field.
func FilePathContains(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFilePath), v))
	})
}

// FilePathHasPrefix applies the HasPrefix predicate on the "file_path" field.
func FilePathHasPrefix(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFilePath), v))
	})
}

// FilePathHasSuffix applies the HasSuffix predicate on the "file_path" field.
func FilePathHasSuffix(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFilePath), v))
	})
}

// FilePathEqualFold applies the EqualFold predicate on the "file_path" field.
func FilePathEqualFold(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFilePath), v))
	})
}

// FilePathContainsFold applies the ContainsFold predicate on the "file_path" field.
func FilePathContainsFold(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFilePath), v))
	})
}

// ArchivePathEQ applies the EQ predicate on the "archive_path" field.
func ArchivePathEQ(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArchivePath), v))
	})
}

// ArchivePathNEQ applies the NEQ predicate on the "archive_path" field.
func ArchivePathNEQ(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldArchivePath), v))
	})
}

// ArchivePathIn applies the In predicate on the "archive_path" field.
func ArchivePathIn(vs ...string) predicate.Artifact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artifact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldArchivePath), v...))
	})
}

// ArchivePathNotIn applies the NotIn predicate on the "archive_path" field.
func ArchivePathNotIn(vs ...string) predicate.Artifact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artifact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldArchivePath), v...))
	})
}

// ArchivePathGT applies the GT predicate on the "archive_path" field.
func ArchivePathGT(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldArchivePath), v))
	})
}

// ArchivePathGTE applies the GTE predicate on the "archive_path" field.
func ArchivePathGTE(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldArchivePath), v))
	})
}

// ArchivePathLT applies the LT predicate on the "archive_path" field.
func ArchivePathLT(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldArchivePath), v))
	})
}

// ArchivePathLTE applies the LTE predicate on the "archive_path" field.
func ArchivePathLTE(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldArchivePath), v))
	})
}

// ArchivePathContains applies the Contains predicate on the "archive_path" field.
func ArchivePathContains(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldArchivePath), v))
	})
}

// ArchivePathHasPrefix applies the HasPrefix predicate on the "archive_path" field.
func ArchivePathHasPrefix(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldArchivePath), v))
	})
}

// ArchivePathHasSuffix applies the HasSuffix predicate on the "archive_path" field.
func ArchivePathHasSuffix(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldArchivePath), v))
	})
}

// ArchivePathIsNil applies the IsNil predicate on the "archive_path" field.
func ArchivePathIsNil() predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldArchivePath)))
	})
}

// ArchivePathNotNil applies the NotNil predicate on the "archive_path" field.
func ArchivePathNotNil() predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldArchivePath)))
	})
}

// ArchivePathEqualFold applies the EqualFold predicate on the "archive_path" field.
func ArchivePathEqualFold(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldArchivePath), v))
	})
}

// ArchivePathContainsFold applies the ContainsFold predicate on the "archive_path" field.
func ArchivePathContainsFold(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldArchivePath), v))
	})
}

// FileTypeEQ applies the EQ predicate on the "file_type" field.
func FileTypeEQ(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFileType), v))
	})
}

// FileTypeNEQ applies the NEQ predicate on the "file_type" field.
func FileTypeNEQ(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFileType), v))
	})
}

// FileTypeIn applies the In predicate on the "file_type" field.
func FileTypeIn(vs ...string) predicate.Artifact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artifact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFileType), v...))
	})
}

// FileTypeNotIn applies the NotIn predicate on the "file_type" field.
func FileTypeNotIn(vs ...string) predicate.Artifact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artifact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFileType), v...))
	})
}

// FileTypeGT applies the GT predicate on the "file_type" field.
func FileTypeGT(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFileType), v))
	})
}

// FileTypeGTE applies the GTE predicate on the "file_type" field.
func FileTypeGTE(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFileType), v))
	})
}

// FileTypeLT applies the LT predicate on the "file_type" field.
func FileTypeLT(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFileType), v))
	})
}

// FileTypeLTE applies the LTE predicate on the "file_type" field.
func FileTypeLTE(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFileType), v))
	})
}

// FileTypeContains applies the Contains predicate on the "file_type" field.
func FileTypeContains(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFileType), v))
	})
}

// FileTypeHasPrefix applies the HasPrefix predicate on the "file_type" field.
func FileTypeHasPrefix(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFileType), v))
	})
}

// FileTypeHasSuffix applies the HasSuffix predicate on the "file_type" field.
func FileTypeHasSuffix(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFileType), v))
	})
}

// FileTypeEqualFold applies the EqualFold predicate on the "file_type" field.
func FileTypeEqualFold(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFileType), v))
	})
}

// FileTypeContainsFold applies the ContainsFold predicate on the "file_type" field.
func FileTypeContainsFold(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFileType), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Artifact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artifact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Artifact {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artifact(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// HasMetadata applies the HasEdge predicate on the "metadata" edge.
func HasMetadata() predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetadataTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetadataWith applies the HasEdge predicate on the "metadata" edge with a given conditions (other predicates).
func HasMetadataWith(preds ...predicate.Metadata) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MetadataInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Artifact) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Artifact) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Artifact) predicate.Artifact {
	return predicate.Artifact(func(s *sql.Selector) {
		p(s.Not())
	})
}

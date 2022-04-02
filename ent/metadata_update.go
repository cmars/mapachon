// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cmars/mapachon/ent/artifact"
	"github.com/cmars/mapachon/ent/metadata"
	"github.com/cmars/mapachon/ent/predicate"
)

// MetadataUpdate is the builder for updating Metadata entities.
type MetadataUpdate struct {
	config
	hooks    []Hook
	mutation *MetadataMutation
}

// Where appends a list predicates to the MetadataUpdate builder.
func (mu *MetadataUpdate) Where(ps ...predicate.Metadata) *MetadataUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetKey sets the "key" field.
func (mu *MetadataUpdate) SetKey(s string) *MetadataUpdate {
	mu.mutation.SetKey(s)
	return mu
}

// SetValue sets the "value" field.
func (mu *MetadataUpdate) SetValue(s string) *MetadataUpdate {
	mu.mutation.SetValue(s)
	return mu
}

// SetArtifactID sets the "artifact" edge to the Artifact entity by ID.
func (mu *MetadataUpdate) SetArtifactID(id int) *MetadataUpdate {
	mu.mutation.SetArtifactID(id)
	return mu
}

// SetNillableArtifactID sets the "artifact" edge to the Artifact entity by ID if the given value is not nil.
func (mu *MetadataUpdate) SetNillableArtifactID(id *int) *MetadataUpdate {
	if id != nil {
		mu = mu.SetArtifactID(*id)
	}
	return mu
}

// SetArtifact sets the "artifact" edge to the Artifact entity.
func (mu *MetadataUpdate) SetArtifact(a *Artifact) *MetadataUpdate {
	return mu.SetArtifactID(a.ID)
}

// Mutation returns the MetadataMutation object of the builder.
func (mu *MetadataUpdate) Mutation() *MetadataMutation {
	return mu.mutation
}

// ClearArtifact clears the "artifact" edge to the Artifact entity.
func (mu *MetadataUpdate) ClearArtifact() *MetadataUpdate {
	mu.mutation.ClearArtifact()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MetadataUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetadataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MetadataUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MetadataUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MetadataUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MetadataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metadata.Table,
			Columns: metadata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metadata.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metadata.FieldKey,
		})
	}
	if value, ok := mu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metadata.FieldValue,
		})
	}
	if mu.mutation.ArtifactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metadata.ArtifactTable,
			Columns: []string{metadata.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artifact.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ArtifactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metadata.ArtifactTable,
			Columns: []string{metadata.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artifact.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MetadataUpdateOne is the builder for updating a single Metadata entity.
type MetadataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MetadataMutation
}

// SetKey sets the "key" field.
func (muo *MetadataUpdateOne) SetKey(s string) *MetadataUpdateOne {
	muo.mutation.SetKey(s)
	return muo
}

// SetValue sets the "value" field.
func (muo *MetadataUpdateOne) SetValue(s string) *MetadataUpdateOne {
	muo.mutation.SetValue(s)
	return muo
}

// SetArtifactID sets the "artifact" edge to the Artifact entity by ID.
func (muo *MetadataUpdateOne) SetArtifactID(id int) *MetadataUpdateOne {
	muo.mutation.SetArtifactID(id)
	return muo
}

// SetNillableArtifactID sets the "artifact" edge to the Artifact entity by ID if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableArtifactID(id *int) *MetadataUpdateOne {
	if id != nil {
		muo = muo.SetArtifactID(*id)
	}
	return muo
}

// SetArtifact sets the "artifact" edge to the Artifact entity.
func (muo *MetadataUpdateOne) SetArtifact(a *Artifact) *MetadataUpdateOne {
	return muo.SetArtifactID(a.ID)
}

// Mutation returns the MetadataMutation object of the builder.
func (muo *MetadataUpdateOne) Mutation() *MetadataMutation {
	return muo.mutation
}

// ClearArtifact clears the "artifact" edge to the Artifact entity.
func (muo *MetadataUpdateOne) ClearArtifact() *MetadataUpdateOne {
	muo.mutation.ClearArtifact()
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MetadataUpdateOne) Select(field string, fields ...string) *MetadataUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Metadata entity.
func (muo *MetadataUpdateOne) Save(ctx context.Context) (*Metadata, error) {
	var (
		err  error
		node *Metadata
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetadataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MetadataUpdateOne) SaveX(ctx context.Context) *Metadata {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MetadataUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MetadataUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MetadataUpdateOne) sqlSave(ctx context.Context) (_node *Metadata, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metadata.Table,
			Columns: metadata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metadata.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Metadata.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metadata.FieldID)
		for _, f := range fields {
			if !metadata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != metadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metadata.FieldKey,
		})
	}
	if value, ok := muo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metadata.FieldValue,
		})
	}
	if muo.mutation.ArtifactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metadata.ArtifactTable,
			Columns: []string{metadata.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artifact.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ArtifactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metadata.ArtifactTable,
			Columns: []string{metadata.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artifact.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Metadata{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
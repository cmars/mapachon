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
)

// MetadataCreate is the builder for creating a Metadata entity.
type MetadataCreate struct {
	config
	mutation *MetadataMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetKey sets the "key" field.
func (mc *MetadataCreate) SetKey(s string) *MetadataCreate {
	mc.mutation.SetKey(s)
	return mc
}

// SetValue sets the "value" field.
func (mc *MetadataCreate) SetValue(s string) *MetadataCreate {
	mc.mutation.SetValue(s)
	return mc
}

// SetArtifactID sets the "artifact" edge to the Artifact entity by ID.
func (mc *MetadataCreate) SetArtifactID(id int) *MetadataCreate {
	mc.mutation.SetArtifactID(id)
	return mc
}

// SetNillableArtifactID sets the "artifact" edge to the Artifact entity by ID if the given value is not nil.
func (mc *MetadataCreate) SetNillableArtifactID(id *int) *MetadataCreate {
	if id != nil {
		mc = mc.SetArtifactID(*id)
	}
	return mc
}

// SetArtifact sets the "artifact" edge to the Artifact entity.
func (mc *MetadataCreate) SetArtifact(a *Artifact) *MetadataCreate {
	return mc.SetArtifactID(a.ID)
}

// Mutation returns the MetadataMutation object of the builder.
func (mc *MetadataCreate) Mutation() *MetadataMutation {
	return mc.mutation
}

// Save creates the Metadata in the database.
func (mc *MetadataCreate) Save(ctx context.Context) (*Metadata, error) {
	var (
		err  error
		node *Metadata
	)
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetadataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MetadataCreate) SaveX(ctx context.Context) *Metadata {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MetadataCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MetadataCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MetadataCreate) check() error {
	if _, ok := mc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "Metadata.key"`)}
	}
	if _, ok := mc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Metadata.value"`)}
	}
	return nil
}

func (mc *MetadataCreate) sqlSave(ctx context.Context) (*Metadata, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MetadataCreate) createSpec() (*Metadata, *sqlgraph.CreateSpec) {
	var (
		_node = &Metadata{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: metadata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metadata.FieldID,
			},
		}
	)
	_spec.OnConflict = mc.conflict
	if value, ok := mc.mutation.Key(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metadata.FieldKey,
		})
		_node.Key = value
	}
	if value, ok := mc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metadata.FieldValue,
		})
		_node.Value = value
	}
	if nodes := mc.mutation.ArtifactIDs(); len(nodes) > 0 {
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
		_node.artifact_metadata = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Metadata.Create().
//		SetKey(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MetadataUpsert) {
//			SetKey(v+v).
//		}).
//		Exec(ctx)
//
func (mc *MetadataCreate) OnConflict(opts ...sql.ConflictOption) *MetadataUpsertOne {
	mc.conflict = opts
	return &MetadataUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Metadata.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mc *MetadataCreate) OnConflictColumns(columns ...string) *MetadataUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MetadataUpsertOne{
		create: mc,
	}
}

type (
	// MetadataUpsertOne is the builder for "upsert"-ing
	//  one Metadata node.
	MetadataUpsertOne struct {
		create *MetadataCreate
	}

	// MetadataUpsert is the "OnConflict" setter.
	MetadataUpsert struct {
		*sql.UpdateSet
	}
)

// SetKey sets the "key" field.
func (u *MetadataUpsert) SetKey(v string) *MetadataUpsert {
	u.Set(metadata.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *MetadataUpsert) UpdateKey() *MetadataUpsert {
	u.SetExcluded(metadata.FieldKey)
	return u
}

// SetValue sets the "value" field.
func (u *MetadataUpsert) SetValue(v string) *MetadataUpsert {
	u.Set(metadata.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *MetadataUpsert) UpdateValue() *MetadataUpsert {
	u.SetExcluded(metadata.FieldValue)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Metadata.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *MetadataUpsertOne) UpdateNewValues() *MetadataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Metadata.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *MetadataUpsertOne) Ignore() *MetadataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MetadataUpsertOne) DoNothing() *MetadataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MetadataCreate.OnConflict
// documentation for more info.
func (u *MetadataUpsertOne) Update(set func(*MetadataUpsert)) *MetadataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MetadataUpsert{UpdateSet: update})
	}))
	return u
}

// SetKey sets the "key" field.
func (u *MetadataUpsertOne) SetKey(v string) *MetadataUpsertOne {
	return u.Update(func(s *MetadataUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *MetadataUpsertOne) UpdateKey() *MetadataUpsertOne {
	return u.Update(func(s *MetadataUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *MetadataUpsertOne) SetValue(v string) *MetadataUpsertOne {
	return u.Update(func(s *MetadataUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *MetadataUpsertOne) UpdateValue() *MetadataUpsertOne {
	return u.Update(func(s *MetadataUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *MetadataUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MetadataCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MetadataUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MetadataUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MetadataUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MetadataCreateBulk is the builder for creating many Metadata entities in bulk.
type MetadataCreateBulk struct {
	config
	builders []*MetadataCreate
	conflict []sql.ConflictOption
}

// Save creates the Metadata entities in the database.
func (mcb *MetadataCreateBulk) Save(ctx context.Context) ([]*Metadata, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Metadata, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetadataMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MetadataCreateBulk) SaveX(ctx context.Context) []*Metadata {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MetadataCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MetadataCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Metadata.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MetadataUpsert) {
//			SetKey(v+v).
//		}).
//		Exec(ctx)
//
func (mcb *MetadataCreateBulk) OnConflict(opts ...sql.ConflictOption) *MetadataUpsertBulk {
	mcb.conflict = opts
	return &MetadataUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Metadata.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mcb *MetadataCreateBulk) OnConflictColumns(columns ...string) *MetadataUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MetadataUpsertBulk{
		create: mcb,
	}
}

// MetadataUpsertBulk is the builder for "upsert"-ing
// a bulk of Metadata nodes.
type MetadataUpsertBulk struct {
	create *MetadataCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Metadata.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *MetadataUpsertBulk) UpdateNewValues() *MetadataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Metadata.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *MetadataUpsertBulk) Ignore() *MetadataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MetadataUpsertBulk) DoNothing() *MetadataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MetadataCreateBulk.OnConflict
// documentation for more info.
func (u *MetadataUpsertBulk) Update(set func(*MetadataUpsert)) *MetadataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MetadataUpsert{UpdateSet: update})
	}))
	return u
}

// SetKey sets the "key" field.
func (u *MetadataUpsertBulk) SetKey(v string) *MetadataUpsertBulk {
	return u.Update(func(s *MetadataUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *MetadataUpsertBulk) UpdateKey() *MetadataUpsertBulk {
	return u.Update(func(s *MetadataUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *MetadataUpsertBulk) SetValue(v string) *MetadataUpsertBulk {
	return u.Update(func(s *MetadataUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *MetadataUpsertBulk) UpdateValue() *MetadataUpsertBulk {
	return u.Update(func(s *MetadataUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *MetadataUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MetadataCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MetadataCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MetadataUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

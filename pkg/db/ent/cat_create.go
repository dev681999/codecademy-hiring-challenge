// Code generated by ent, DO NOT EDIT.

package ent

import (
	"catinator-backend/pkg/db/ent/cat"
	"catinator-backend/pkg/db/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CatCreate is the builder for creating a Cat entity.
type CatCreate struct {
	config
	mutation *CatMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (cc *CatCreate) SetCreateTime(t time.Time) *CatCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *CatCreate) SetNillableCreateTime(t *time.Time) *CatCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *CatCreate) SetUpdateTime(t time.Time) *CatCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *CatCreate) SetNillableUpdateTime(t *time.Time) *CatCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CatCreate) SetName(s string) *CatCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *CatCreate) SetNillableName(s *string) *CatCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetDescription sets the "description" field.
func (cc *CatCreate) SetDescription(s string) *CatCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *CatCreate) SetNillableDescription(s *string) *CatCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetImageID sets the "image_id" field.
func (cc *CatCreate) SetImageID(s string) *CatCreate {
	cc.mutation.SetImageID(s)
	return cc
}

// SetNillableImageID sets the "image_id" field if the given value is not nil.
func (cc *CatCreate) SetNillableImageID(s *string) *CatCreate {
	if s != nil {
		cc.SetImageID(*s)
	}
	return cc
}

// SetOwnerID sets the "owner_id" field.
func (cc *CatCreate) SetOwnerID(s string) *CatCreate {
	cc.mutation.SetOwnerID(s)
	return cc
}

// SetTags sets the "tags" field.
func (cc *CatCreate) SetTags(s []string) *CatCreate {
	cc.mutation.SetTags(s)
	return cc
}

// SetID sets the "id" field.
func (cc *CatCreate) SetID(s string) *CatCreate {
	cc.mutation.SetID(s)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CatCreate) SetNillableID(s *string) *CatCreate {
	if s != nil {
		cc.SetID(*s)
	}
	return cc
}

// SetOwner sets the "owner" edge to the User entity.
func (cc *CatCreate) SetOwner(u *User) *CatCreate {
	return cc.SetOwnerID(u.ID)
}

// Mutation returns the CatMutation object of the builder.
func (cc *CatCreate) Mutation() *CatMutation {
	return cc.mutation
}

// Save creates the Cat in the database.
func (cc *CatCreate) Save(ctx context.Context) (*Cat, error) {
	var (
		err  error
		node *Cat
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Cat)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CatMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CatCreate) SaveX(ctx context.Context) *Cat {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CatCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CatCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CatCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := cat.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := cat.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
	if _, ok := cc.mutation.Name(); !ok {
		v := cat.DefaultName
		cc.mutation.SetName(v)
	}
	if _, ok := cc.mutation.Description(); !ok {
		v := cat.DefaultDescription
		cc.mutation.SetDescription(v)
	}
	if _, ok := cc.mutation.ImageID(); !ok {
		v := cat.DefaultImageID
		cc.mutation.SetImageID(v)
	}
	if _, ok := cc.mutation.Tags(); !ok {
		v := cat.DefaultTags
		cc.mutation.SetTags(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := cat.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CatCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Cat.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Cat.update_time"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Cat.name"`)}
	}
	if _, ok := cc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Cat.description"`)}
	}
	if _, ok := cc.mutation.ImageID(); !ok {
		return &ValidationError{Name: "image_id", err: errors.New(`ent: missing required field "Cat.image_id"`)}
	}
	if _, ok := cc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "Cat.owner_id"`)}
	}
	if _, ok := cc.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New(`ent: missing required field "Cat.tags"`)}
	}
	if _, ok := cc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Cat.owner"`)}
	}
	return nil
}

func (cc *CatCreate) sqlSave(ctx context.Context) (*Cat, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Cat.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (cc *CatCreate) createSpec() (*Cat, *sqlgraph.CreateSpec) {
	var (
		_node = &Cat{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cat.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: cat.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cat.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cat.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cat.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cat.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := cc.mutation.ImageID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cat.FieldImageID,
		})
		_node.ImageID = value
	}
	if value, ok := cc.mutation.Tags(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: cat.FieldTags,
		})
		_node.Tags = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cat.OwnerTable,
			Columns: []string{cat.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Cat.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CatUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (cc *CatCreate) OnConflict(opts ...sql.ConflictOption) *CatUpsertOne {
	cc.conflict = opts
	return &CatUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Cat.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *CatCreate) OnConflictColumns(columns ...string) *CatUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CatUpsertOne{
		create: cc,
	}
}

type (
	// CatUpsertOne is the builder for "upsert"-ing
	//  one Cat node.
	CatUpsertOne struct {
		create *CatCreate
	}

	// CatUpsert is the "OnConflict" setter.
	CatUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateTime sets the "create_time" field.
func (u *CatUpsert) SetCreateTime(v time.Time) *CatUpsert {
	u.Set(cat.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CatUpsert) UpdateCreateTime() *CatUpsert {
	u.SetExcluded(cat.FieldCreateTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *CatUpsert) SetUpdateTime(v time.Time) *CatUpsert {
	u.Set(cat.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CatUpsert) UpdateUpdateTime() *CatUpsert {
	u.SetExcluded(cat.FieldUpdateTime)
	return u
}

// SetName sets the "name" field.
func (u *CatUpsert) SetName(v string) *CatUpsert {
	u.Set(cat.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CatUpsert) UpdateName() *CatUpsert {
	u.SetExcluded(cat.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *CatUpsert) SetDescription(v string) *CatUpsert {
	u.Set(cat.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CatUpsert) UpdateDescription() *CatUpsert {
	u.SetExcluded(cat.FieldDescription)
	return u
}

// SetImageID sets the "image_id" field.
func (u *CatUpsert) SetImageID(v string) *CatUpsert {
	u.Set(cat.FieldImageID, v)
	return u
}

// UpdateImageID sets the "image_id" field to the value that was provided on create.
func (u *CatUpsert) UpdateImageID() *CatUpsert {
	u.SetExcluded(cat.FieldImageID)
	return u
}

// SetOwnerID sets the "owner_id" field.
func (u *CatUpsert) SetOwnerID(v string) *CatUpsert {
	u.Set(cat.FieldOwnerID, v)
	return u
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *CatUpsert) UpdateOwnerID() *CatUpsert {
	u.SetExcluded(cat.FieldOwnerID)
	return u
}

// SetTags sets the "tags" field.
func (u *CatUpsert) SetTags(v []string) *CatUpsert {
	u.Set(cat.FieldTags, v)
	return u
}

// UpdateTags sets the "tags" field to the value that was provided on create.
func (u *CatUpsert) UpdateTags() *CatUpsert {
	u.SetExcluded(cat.FieldTags)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Cat.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(cat.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CatUpsertOne) UpdateNewValues() *CatUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(cat.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(cat.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Cat.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CatUpsertOne) Ignore() *CatUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CatUpsertOne) DoNothing() *CatUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CatCreate.OnConflict
// documentation for more info.
func (u *CatUpsertOne) Update(set func(*CatUpsert)) *CatUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CatUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *CatUpsertOne) SetCreateTime(v time.Time) *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CatUpsertOne) UpdateCreateTime() *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *CatUpsertOne) SetUpdateTime(v time.Time) *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CatUpsertOne) UpdateUpdateTime() *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *CatUpsertOne) SetName(v string) *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CatUpsertOne) UpdateName() *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *CatUpsertOne) SetDescription(v string) *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CatUpsertOne) UpdateDescription() *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.UpdateDescription()
	})
}

// SetImageID sets the "image_id" field.
func (u *CatUpsertOne) SetImageID(v string) *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.SetImageID(v)
	})
}

// UpdateImageID sets the "image_id" field to the value that was provided on create.
func (u *CatUpsertOne) UpdateImageID() *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.UpdateImageID()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *CatUpsertOne) SetOwnerID(v string) *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.SetOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *CatUpsertOne) UpdateOwnerID() *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.UpdateOwnerID()
	})
}

// SetTags sets the "tags" field.
func (u *CatUpsertOne) SetTags(v []string) *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.SetTags(v)
	})
}

// UpdateTags sets the "tags" field to the value that was provided on create.
func (u *CatUpsertOne) UpdateTags() *CatUpsertOne {
	return u.Update(func(s *CatUpsert) {
		s.UpdateTags()
	})
}

// Exec executes the query.
func (u *CatUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CatCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CatUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CatUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CatUpsertOne.ID is not supported by MySQL driver. Use CatUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CatUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CatCreateBulk is the builder for creating many Cat entities in bulk.
type CatCreateBulk struct {
	config
	builders []*CatCreate
	conflict []sql.ConflictOption
}

// Save creates the Cat entities in the database.
func (ccb *CatCreateBulk) Save(ctx context.Context) ([]*Cat, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Cat, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CatMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CatCreateBulk) SaveX(ctx context.Context) []*Cat {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CatCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CatCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Cat.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CatUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *CatCreateBulk) OnConflict(opts ...sql.ConflictOption) *CatUpsertBulk {
	ccb.conflict = opts
	return &CatUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Cat.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *CatCreateBulk) OnConflictColumns(columns ...string) *CatUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CatUpsertBulk{
		create: ccb,
	}
}

// CatUpsertBulk is the builder for "upsert"-ing
// a bulk of Cat nodes.
type CatUpsertBulk struct {
	create *CatCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Cat.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(cat.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CatUpsertBulk) UpdateNewValues() *CatUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(cat.FieldID)
				return
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(cat.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Cat.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CatUpsertBulk) Ignore() *CatUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CatUpsertBulk) DoNothing() *CatUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CatCreateBulk.OnConflict
// documentation for more info.
func (u *CatUpsertBulk) Update(set func(*CatUpsert)) *CatUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CatUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *CatUpsertBulk) SetCreateTime(v time.Time) *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CatUpsertBulk) UpdateCreateTime() *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *CatUpsertBulk) SetUpdateTime(v time.Time) *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CatUpsertBulk) UpdateUpdateTime() *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *CatUpsertBulk) SetName(v string) *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CatUpsertBulk) UpdateName() *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *CatUpsertBulk) SetDescription(v string) *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CatUpsertBulk) UpdateDescription() *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.UpdateDescription()
	})
}

// SetImageID sets the "image_id" field.
func (u *CatUpsertBulk) SetImageID(v string) *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.SetImageID(v)
	})
}

// UpdateImageID sets the "image_id" field to the value that was provided on create.
func (u *CatUpsertBulk) UpdateImageID() *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.UpdateImageID()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *CatUpsertBulk) SetOwnerID(v string) *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.SetOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *CatUpsertBulk) UpdateOwnerID() *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.UpdateOwnerID()
	})
}

// SetTags sets the "tags" field.
func (u *CatUpsertBulk) SetTags(v []string) *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.SetTags(v)
	})
}

// UpdateTags sets the "tags" field to the value that was provided on create.
func (u *CatUpsertBulk) UpdateTags() *CatUpsertBulk {
	return u.Update(func(s *CatUpsert) {
		s.UpdateTags()
	})
}

// Exec executes the query.
func (u *CatUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CatCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CatCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CatUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

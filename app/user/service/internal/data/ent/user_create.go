// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-ent-example/app/user/service/internal/data/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (uc *UserCreate) SetCreateTime(i int64) *UserCreate {
	uc.mutation.SetCreateTime(i)
	return uc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateTime(i *int64) *UserCreate {
	if i != nil {
		uc.SetCreateTime(*i)
	}
	return uc
}

// SetUpdateTime sets the "update_time" field.
func (uc *UserCreate) SetUpdateTime(i int64) *UserCreate {
	uc.mutation.SetUpdateTime(i)
	return uc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdateTime(i *int64) *UserCreate {
	if i != nil {
		uc.SetUpdateTime(*i)
	}
	return uc
}

// SetDeleteTime sets the "delete_time" field.
func (uc *UserCreate) SetDeleteTime(i int64) *UserCreate {
	uc.mutation.SetDeleteTime(i)
	return uc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeleteTime(i *int64) *UserCreate {
	if i != nil {
		uc.SetDeleteTime(*i)
	}
	return uc
}

// SetUserName sets the "user_name" field.
func (uc *UserCreate) SetUserName(s string) *UserCreate {
	uc.mutation.SetUserName(s)
	return uc
}

// SetNillableUserName sets the "user_name" field if the given value is not nil.
func (uc *UserCreate) SetNillableUserName(s *string) *UserCreate {
	if s != nil {
		uc.SetUserName(*s)
	}
	return uc
}

// SetNickName sets the "nick_name" field.
func (uc *UserCreate) SetNickName(s string) *UserCreate {
	uc.mutation.SetNickName(s)
	return uc
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (uc *UserCreate) SetNillableNickName(s *string) *UserCreate {
	if s != nil {
		uc.SetNickName(*s)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uc *UserCreate) SetNillablePassword(s *string) *UserCreate {
	if s != nil {
		uc.SetPassword(*s)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uint32) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreateTime(); !ok {
		v := user.DefaultCreateTime()
		uc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if v, ok := uc.mutation.UserName(); ok {
		if err := user.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`ent: validator failed for field "User.user_name": %w`, err)}
		}
	}
	if v, ok := uc.mutation.NickName(); ok {
		if err := user.NickNameValidator(v); err != nil {
			return &ValidationError{Name: "nick_name", err: fmt.Errorf(`ent: validator failed for field "User.nick_name": %w`, err)}
		}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uc.mutation.ID(); ok {
		if err := user.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "User.id": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = uc.conflict
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.CreateTime(); ok {
		_spec.SetField(user.FieldCreateTime, field.TypeInt64, value)
		_node.CreateTime = &value
	}
	if value, ok := uc.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeInt64, value)
		_node.UpdateTime = &value
	}
	if value, ok := uc.mutation.DeleteTime(); ok {
		_spec.SetField(user.FieldDeleteTime, field.TypeInt64, value)
		_node.DeleteTime = &value
	}
	if value, ok := uc.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
		_node.UserName = &value
	}
	if value, ok := uc.mutation.NickName(); ok {
		_spec.SetField(user.FieldNickName, field.TypeString, value)
		_node.NickName = &value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = &value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *UserUpsert) SetUpdateTime(v int64) *UserUpsert {
	u.Set(user.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserUpsert) UpdateUpdateTime() *UserUpsert {
	u.SetExcluded(user.FieldUpdateTime)
	return u
}

// AddUpdateTime adds v to the "update_time" field.
func (u *UserUpsert) AddUpdateTime(v int64) *UserUpsert {
	u.Add(user.FieldUpdateTime, v)
	return u
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *UserUpsert) ClearUpdateTime() *UserUpsert {
	u.SetNull(user.FieldUpdateTime)
	return u
}

// SetDeleteTime sets the "delete_time" field.
func (u *UserUpsert) SetDeleteTime(v int64) *UserUpsert {
	u.Set(user.FieldDeleteTime, v)
	return u
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *UserUpsert) UpdateDeleteTime() *UserUpsert {
	u.SetExcluded(user.FieldDeleteTime)
	return u
}

// AddDeleteTime adds v to the "delete_time" field.
func (u *UserUpsert) AddDeleteTime(v int64) *UserUpsert {
	u.Add(user.FieldDeleteTime, v)
	return u
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *UserUpsert) ClearDeleteTime() *UserUpsert {
	u.SetNull(user.FieldDeleteTime)
	return u
}

// SetNickName sets the "nick_name" field.
func (u *UserUpsert) SetNickName(v string) *UserUpsert {
	u.Set(user.FieldNickName, v)
	return u
}

// UpdateNickName sets the "nick_name" field to the value that was provided on create.
func (u *UserUpsert) UpdateNickName() *UserUpsert {
	u.SetExcluded(user.FieldNickName)
	return u
}

// ClearNickName clears the value of the "nick_name" field.
func (u *UserUpsert) ClearNickName() *UserUpsert {
	u.SetNull(user.FieldNickName)
	return u
}

// SetPassword sets the "password" field.
func (u *UserUpsert) SetPassword(v string) *UserUpsert {
	u.Set(user.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsert) UpdatePassword() *UserUpsert {
	u.SetExcluded(user.FieldPassword)
	return u
}

// ClearPassword clears the value of the "password" field.
func (u *UserUpsert) ClearPassword() *UserUpsert {
	u.SetNull(user.FieldPassword)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(user.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(user.FieldCreateTime)
		}
		if _, exists := u.create.mutation.UserName(); exists {
			s.SetIgnore(user.FieldUserName)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserUpsertOne) SetUpdateTime(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdateTime(v)
	})
}

// AddUpdateTime adds v to the "update_time" field.
func (u *UserUpsertOne) AddUpdateTime(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUpdateTime() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *UserUpsertOne) ClearUpdateTime() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *UserUpsertOne) SetDeleteTime(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetDeleteTime(v)
	})
}

// AddDeleteTime adds v to the "delete_time" field.
func (u *UserUpsertOne) AddDeleteTime(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateDeleteTime() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *UserUpsertOne) ClearDeleteTime() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearDeleteTime()
	})
}

// SetNickName sets the "nick_name" field.
func (u *UserUpsertOne) SetNickName(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetNickName(v)
	})
}

// UpdateNickName sets the "nick_name" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateNickName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateNickName()
	})
}

// ClearNickName clears the value of the "nick_name" field.
func (u *UserUpsertOne) ClearNickName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearNickName()
	})
}

// SetPassword sets the "password" field.
func (u *UserUpsertOne) SetPassword(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsertOne) UpdatePassword() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePassword()
	})
}

// ClearPassword clears the value of the "password" field.
func (u *UserUpsertOne) ClearPassword() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearPassword()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(user.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(user.FieldCreateTime)
			}
			if _, exists := b.mutation.UserName(); exists {
				s.SetIgnore(user.FieldUserName)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserUpsertBulk) SetUpdateTime(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdateTime(v)
	})
}

// AddUpdateTime adds v to the "update_time" field.
func (u *UserUpsertBulk) AddUpdateTime(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUpdateTime() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *UserUpsertBulk) ClearUpdateTime() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *UserUpsertBulk) SetDeleteTime(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetDeleteTime(v)
	})
}

// AddDeleteTime adds v to the "delete_time" field.
func (u *UserUpsertBulk) AddDeleteTime(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateDeleteTime() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *UserUpsertBulk) ClearDeleteTime() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearDeleteTime()
	})
}

// SetNickName sets the "nick_name" field.
func (u *UserUpsertBulk) SetNickName(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetNickName(v)
	})
}

// UpdateNickName sets the "nick_name" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateNickName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateNickName()
	})
}

// ClearNickName clears the value of the "nick_name" field.
func (u *UserUpsertBulk) ClearNickName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearNickName()
	})
}

// SetPassword sets the "password" field.
func (u *UserUpsertBulk) SetPassword(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdatePassword() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePassword()
	})
}

// ClearPassword clears the value of the "password" field.
func (u *UserUpsertBulk) ClearPassword() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearPassword()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

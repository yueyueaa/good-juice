// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"juice/app/public/ent/predicate"
	"juice/app/public/ent/userbaseinfo"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserBaseInfoDelete is the builder for deleting a UserBaseInfo entity.
type UserBaseInfoDelete struct {
	config
	hooks    []Hook
	mutation *UserBaseInfoMutation
}

// Where appends a list predicates to the UserBaseInfoDelete builder.
func (ubid *UserBaseInfoDelete) Where(ps ...predicate.UserBaseInfo) *UserBaseInfoDelete {
	ubid.mutation.Where(ps...)
	return ubid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ubid *UserBaseInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ubid.sqlExec, ubid.mutation, ubid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ubid *UserBaseInfoDelete) ExecX(ctx context.Context) int {
	n, err := ubid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ubid *UserBaseInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userbaseinfo.Table, sqlgraph.NewFieldSpec(userbaseinfo.FieldID, field.TypeUint64))
	if ps := ubid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ubid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ubid.mutation.done = true
	return affected, err
}

// UserBaseInfoDeleteOne is the builder for deleting a single UserBaseInfo entity.
type UserBaseInfoDeleteOne struct {
	ubid *UserBaseInfoDelete
}

// Where appends a list predicates to the UserBaseInfoDelete builder.
func (ubido *UserBaseInfoDeleteOne) Where(ps ...predicate.UserBaseInfo) *UserBaseInfoDeleteOne {
	ubido.ubid.mutation.Where(ps...)
	return ubido
}

// Exec executes the deletion query.
func (ubido *UserBaseInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := ubido.ubid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userbaseinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ubido *UserBaseInfoDeleteOne) ExecX(ctx context.Context) {
	if err := ubido.Exec(ctx); err != nil {
		panic(err)
	}
}

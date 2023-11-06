// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"juice/app/public/ent/predicate"
	"juice/app/public/ent/userfollowinfo"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserFollowInfoDelete is the builder for deleting a UserFollowInfo entity.
type UserFollowInfoDelete struct {
	config
	hooks    []Hook
	mutation *UserFollowInfoMutation
}

// Where appends a list predicates to the UserFollowInfoDelete builder.
func (ufid *UserFollowInfoDelete) Where(ps ...predicate.UserFollowInfo) *UserFollowInfoDelete {
	ufid.mutation.Where(ps...)
	return ufid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ufid *UserFollowInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ufid.sqlExec, ufid.mutation, ufid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ufid *UserFollowInfoDelete) ExecX(ctx context.Context) int {
	n, err := ufid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ufid *UserFollowInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userfollowinfo.Table, sqlgraph.NewFieldSpec(userfollowinfo.FieldID, field.TypeUint64))
	if ps := ufid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ufid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ufid.mutation.done = true
	return affected, err
}

// UserFollowInfoDeleteOne is the builder for deleting a single UserFollowInfo entity.
type UserFollowInfoDeleteOne struct {
	ufid *UserFollowInfoDelete
}

// Where appends a list predicates to the UserFollowInfoDelete builder.
func (ufido *UserFollowInfoDeleteOne) Where(ps ...predicate.UserFollowInfo) *UserFollowInfoDeleteOne {
	ufido.ufid.mutation.Where(ps...)
	return ufido
}

// Exec executes the deletion query.
func (ufido *UserFollowInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := ufido.ufid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userfollowinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ufido *UserFollowInfoDeleteOne) ExecX(ctx context.Context) {
	if err := ufido.Exec(ctx); err != nil {
		panic(err)
	}
}

// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	ent2 "juice/app/public/ent"
)

// The UserBaseInfoFunc type is an adapter to allow the use of ordinary
// function as UserBaseInfo mutator.
type UserBaseInfoFunc func(context.Context, *ent2.UserBaseInfoMutation) (ent2.Value, error)

// Mutate calls f(ctx, m).
func (f UserBaseInfoFunc) Mutate(ctx context.Context, m ent2.Mutation) (ent2.Value, error) {
	if mv, ok := m.(*ent2.UserBaseInfoMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserBaseInfoMutation", m)
}

// The UserFollowInfoFunc type is an adapter to allow the use of ordinary
// function as UserFollowInfo mutator.
type UserFollowInfoFunc func(context.Context, *ent2.UserFollowInfoMutation) (ent2.Value, error)

// Mutate calls f(ctx, m).
func (f UserFollowInfoFunc) Mutate(ctx context.Context, m ent2.Mutation) (ent2.Value, error) {
	if mv, ok := m.(*ent2.UserFollowInfoMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserFollowInfoMutation", m)
}

// The UserPasswordFunc type is an adapter to allow the use of ordinary
// function as UserPassword mutator.
type UserPasswordFunc func(context.Context, *ent2.UserPasswordMutation) (ent2.Value, error)

// Mutate calls f(ctx, m).
func (f UserPasswordFunc) Mutate(ctx context.Context, m ent2.Mutation) (ent2.Value, error) {
	if mv, ok := m.(*ent2.UserPasswordMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserPasswordMutation", m)
}

// The VideoCollectionFunc type is an adapter to allow the use of ordinary
// function as VideoCollection mutator.
type VideoCollectionFunc func(context.Context, *ent2.VideoCollectionMutation) (ent2.Value, error)

// Mutate calls f(ctx, m).
func (f VideoCollectionFunc) Mutate(ctx context.Context, m ent2.Mutation) (ent2.Value, error) {
	if mv, ok := m.(*ent2.VideoCollectionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VideoCollectionMutation", m)
}

// The VideoCommentFunc type is an adapter to allow the use of ordinary
// function as VideoComment mutator.
type VideoCommentFunc func(context.Context, *ent2.VideoCommentMutation) (ent2.Value, error)

// Mutate calls f(ctx, m).
func (f VideoCommentFunc) Mutate(ctx context.Context, m ent2.Mutation) (ent2.Value, error) {
	if mv, ok := m.(*ent2.VideoCommentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VideoCommentMutation", m)
}

// The VideoLikeFunc type is an adapter to allow the use of ordinary
// function as VideoLike mutator.
type VideoLikeFunc func(context.Context, *ent2.VideoLikeMutation) (ent2.Value, error)

// Mutate calls f(ctx, m).
func (f VideoLikeFunc) Mutate(ctx context.Context, m ent2.Mutation) (ent2.Value, error) {
	if mv, ok := m.(*ent2.VideoLikeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VideoLikeMutation", m)
}

// The VideoMetadatumFunc type is an adapter to allow the use of ordinary
// function as VideoMetadatum mutator.
type VideoMetadatumFunc func(context.Context, *ent2.VideoMetadatumMutation) (ent2.Value, error)

// Mutate calls f(ctx, m).
func (f VideoMetadatumFunc) Mutate(ctx context.Context, m ent2.Mutation) (ent2.Value, error) {
	if mv, ok := m.(*ent2.VideoMetadatumMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VideoMetadatumMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent2.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent2.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent2.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent2.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent2.Op) Condition {
	return func(_ context.Context, m ent2.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent2.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent2.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent2.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent2.Hook, cond Condition) ent2.Hook {
	return func(next ent2.Mutator) ent2.Mutator {
		return ent2.MutateFunc(func(ctx context.Context, m ent2.Mutation) (ent2.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent2.Hook, op ent2.Op) ent2.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent2.Hook, op ent2.Op) ent2.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent2.Hook {
	return func(ent2.Mutator) ent2.Mutator {
		return ent2.MutateFunc(func(context.Context, ent2.Mutation) (ent2.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent2.Op) ent2.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent2.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent2.Hook) Chain {
	return Chain{append([]ent2.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent2.Hook {
	return func(mutator ent2.Mutator) ent2.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent2.Hook) Chain {
	newHooks := make([]ent2.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
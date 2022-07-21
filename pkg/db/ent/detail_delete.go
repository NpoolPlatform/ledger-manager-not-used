// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/detail"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/predicate"
)

// DetailDelete is the builder for deleting a Detail entity.
type DetailDelete struct {
	config
	hooks    []Hook
	mutation *DetailMutation
}

// Where appends a list predicates to the DetailDelete builder.
func (dd *DetailDelete) Where(ps ...predicate.Detail) *DetailDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DetailDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dd.hooks) == 0 {
		affected, err = dd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dd.mutation = mutation
			affected, err = dd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dd.hooks) - 1; i >= 0; i-- {
			if dd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DetailDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: detail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: detail.FieldID,
			},
		},
	}
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
}

// DetailDeleteOne is the builder for deleting a single Detail entity.
type DetailDeleteOne struct {
	dd *DetailDelete
}

// Exec executes the deletion query.
func (ddo *DetailDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{detail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DetailDeleteOne) ExecX(ctx context.Context) {
	ddo.dd.ExecX(ctx)
}
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/products/internal/pkg/ent/plan"
	"github.com/pepeunlimited/products/internal/pkg/ent/predicate"
)

// PlanDelete is the builder for deleting a Plan entity.
type PlanDelete struct {
	config
	predicates []predicate.Plan
}

// Where adds a new predicate to the delete builder.
func (pd *PlanDelete) Where(ps ...predicate.Plan) *PlanDelete {
	pd.predicates = append(pd.predicates, ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PlanDelete) Exec(ctx context.Context) (int, error) {
	return pd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PlanDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PlanDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: plan.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: plan.FieldID,
			},
		},
	}
	if ps := pd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
}

// PlanDeleteOne is the builder for deleting a single Plan entity.
type PlanDeleteOne struct {
	pd *PlanDelete
}

// Exec executes the deletion query.
func (pdo *PlanDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{plan.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PlanDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/products/internal/pkg/ent/plan"
	"github.com/pepeunlimited/products/internal/pkg/ent/predicate"
	"github.com/pepeunlimited/products/internal/pkg/ent/price"
	"github.com/pepeunlimited/products/internal/pkg/ent/product"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	sku           *string
	prices        map[int]struct{}
	plans         map[int]struct{}
	removedPrices map[int]struct{}
	removedPlans  map[int]struct{}
	predicates    []predicate.Product
}

// Where adds a new predicate for the builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetSku sets the sku field.
func (pu *ProductUpdate) SetSku(s string) *ProductUpdate {
	pu.sku = &s
	return pu
}

// AddPriceIDs adds the prices edge to Price by ids.
func (pu *ProductUpdate) AddPriceIDs(ids ...int) *ProductUpdate {
	if pu.prices == nil {
		pu.prices = make(map[int]struct{})
	}
	for i := range ids {
		pu.prices[ids[i]] = struct{}{}
	}
	return pu
}

// AddPrices adds the prices edges to Price.
func (pu *ProductUpdate) AddPrices(p ...*Price) *ProductUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPriceIDs(ids...)
}

// AddPlanIDs adds the plans edge to Plan by ids.
func (pu *ProductUpdate) AddPlanIDs(ids ...int) *ProductUpdate {
	if pu.plans == nil {
		pu.plans = make(map[int]struct{})
	}
	for i := range ids {
		pu.plans[ids[i]] = struct{}{}
	}
	return pu
}

// AddPlans adds the plans edges to Plan.
func (pu *ProductUpdate) AddPlans(p ...*Plan) *ProductUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPlanIDs(ids...)
}

// RemovePriceIDs removes the prices edge to Price by ids.
func (pu *ProductUpdate) RemovePriceIDs(ids ...int) *ProductUpdate {
	if pu.removedPrices == nil {
		pu.removedPrices = make(map[int]struct{})
	}
	for i := range ids {
		pu.removedPrices[ids[i]] = struct{}{}
	}
	return pu
}

// RemovePrices removes prices edges to Price.
func (pu *ProductUpdate) RemovePrices(p ...*Price) *ProductUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePriceIDs(ids...)
}

// RemovePlanIDs removes the plans edge to Plan by ids.
func (pu *ProductUpdate) RemovePlanIDs(ids ...int) *ProductUpdate {
	if pu.removedPlans == nil {
		pu.removedPlans = make(map[int]struct{})
	}
	for i := range ids {
		pu.removedPlans[ids[i]] = struct{}{}
	}
	return pu
}

// RemovePlans removes plans edges to Plan.
func (pu *ProductUpdate) RemovePlans(p ...*Plan) *ProductUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePlanIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	if pu.sku != nil {
		if err := product.SkuValidator(*pu.sku); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"sku\": %v", err)
		}
	}
	return pu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := pu.sku; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: product.FieldSku,
		})
	}
	if nodes := pu.removedPrices; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.PricesTable,
			Columns: []string{product.PricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: price.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.prices; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.PricesTable,
			Columns: []string{product.PricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: price.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.removedPlans; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.PlansTable,
			Columns: []string{product.PlansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.plans; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.PlansTable,
			Columns: []string{product.PlansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	id            int
	sku           *string
	prices        map[int]struct{}
	plans         map[int]struct{}
	removedPrices map[int]struct{}
	removedPlans  map[int]struct{}
}

// SetSku sets the sku field.
func (puo *ProductUpdateOne) SetSku(s string) *ProductUpdateOne {
	puo.sku = &s
	return puo
}

// AddPriceIDs adds the prices edge to Price by ids.
func (puo *ProductUpdateOne) AddPriceIDs(ids ...int) *ProductUpdateOne {
	if puo.prices == nil {
		puo.prices = make(map[int]struct{})
	}
	for i := range ids {
		puo.prices[ids[i]] = struct{}{}
	}
	return puo
}

// AddPrices adds the prices edges to Price.
func (puo *ProductUpdateOne) AddPrices(p ...*Price) *ProductUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPriceIDs(ids...)
}

// AddPlanIDs adds the plans edge to Plan by ids.
func (puo *ProductUpdateOne) AddPlanIDs(ids ...int) *ProductUpdateOne {
	if puo.plans == nil {
		puo.plans = make(map[int]struct{})
	}
	for i := range ids {
		puo.plans[ids[i]] = struct{}{}
	}
	return puo
}

// AddPlans adds the plans edges to Plan.
func (puo *ProductUpdateOne) AddPlans(p ...*Plan) *ProductUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPlanIDs(ids...)
}

// RemovePriceIDs removes the prices edge to Price by ids.
func (puo *ProductUpdateOne) RemovePriceIDs(ids ...int) *ProductUpdateOne {
	if puo.removedPrices == nil {
		puo.removedPrices = make(map[int]struct{})
	}
	for i := range ids {
		puo.removedPrices[ids[i]] = struct{}{}
	}
	return puo
}

// RemovePrices removes prices edges to Price.
func (puo *ProductUpdateOne) RemovePrices(p ...*Price) *ProductUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePriceIDs(ids...)
}

// RemovePlanIDs removes the plans edge to Plan by ids.
func (puo *ProductUpdateOne) RemovePlanIDs(ids ...int) *ProductUpdateOne {
	if puo.removedPlans == nil {
		puo.removedPlans = make(map[int]struct{})
	}
	for i := range ids {
		puo.removedPlans[ids[i]] = struct{}{}
	}
	return puo
}

// RemovePlans removes plans edges to Plan.
func (puo *ProductUpdateOne) RemovePlans(p ...*Plan) *ProductUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePlanIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	if puo.sku != nil {
		if err := product.SkuValidator(*puo.sku); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"sku\": %v", err)
		}
	}
	return puo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (pr *Product, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  puo.id,
				Type:   field.TypeInt,
				Column: product.FieldID,
			},
		},
	}
	if value := puo.sku; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: product.FieldSku,
		})
	}
	if nodes := puo.removedPrices; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.PricesTable,
			Columns: []string{product.PricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: price.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.prices; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.PricesTable,
			Columns: []string{product.PricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: price.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.removedPlans; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.PlansTable,
			Columns: []string{product.PlansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.plans; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.PlansTable,
			Columns: []string{product.PlansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Product{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}

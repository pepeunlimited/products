// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/products/internal/pkg/ent/plan"
	"github.com/pepeunlimited/products/internal/pkg/ent/price"
	"github.com/pepeunlimited/products/internal/pkg/ent/product"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	sku    *string
	prices map[int]struct{}
	plans  map[int]struct{}
}

// SetSku sets the sku field.
func (pc *ProductCreate) SetSku(s string) *ProductCreate {
	pc.sku = &s
	return pc
}

// AddPriceIDs adds the prices edge to Price by ids.
func (pc *ProductCreate) AddPriceIDs(ids ...int) *ProductCreate {
	if pc.prices == nil {
		pc.prices = make(map[int]struct{})
	}
	for i := range ids {
		pc.prices[ids[i]] = struct{}{}
	}
	return pc
}

// AddPrices adds the prices edges to Price.
func (pc *ProductCreate) AddPrices(p ...*Price) *ProductCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPriceIDs(ids...)
}

// AddPlanIDs adds the plans edge to Plan by ids.
func (pc *ProductCreate) AddPlanIDs(ids ...int) *ProductCreate {
	if pc.plans == nil {
		pc.plans = make(map[int]struct{})
	}
	for i := range ids {
		pc.plans[ids[i]] = struct{}{}
	}
	return pc
}

// AddPlans adds the plans edges to Plan.
func (pc *ProductCreate) AddPlans(p ...*Plan) *ProductCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPlanIDs(ids...)
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	if pc.sku == nil {
		return nil, errors.New("ent: missing required field \"sku\"")
	}
	if err := product.SkuValidator(*pc.sku); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"sku\": %v", err)
	}
	return pc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	var (
		pr    = &Product{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: product.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product.FieldID,
			},
		}
	)
	if value := pc.sku; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: product.FieldSku,
		})
		pr.Sku = *value
	}
	if nodes := pc.prices; len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.plans; len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pr.ID = int(id)
	return pr, nil
}

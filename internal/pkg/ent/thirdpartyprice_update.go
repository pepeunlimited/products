// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/products/internal/pkg/ent/plan"
	"github.com/pepeunlimited/products/internal/pkg/ent/predicate"
	"github.com/pepeunlimited/products/internal/pkg/ent/price"
	"github.com/pepeunlimited/products/internal/pkg/ent/thirdpartyprice"
)

// ThirdPartyPriceUpdate is the builder for updating ThirdPartyPrice entities.
type ThirdPartyPriceUpdate struct {
	config
	in_app_purchase_sku             *string
	google_billing_service_sku      *string
	cleargoogle_billing_service_sku bool
	start_at                        *time.Time
	end_at                          *time.Time
	_type                           *string
	prices                          map[int]struct{}
	plans                           map[int]struct{}
	removedPrices                   map[int]struct{}
	removedPlans                    map[int]struct{}
	predicates                      []predicate.ThirdPartyPrice
}

// Where adds a new predicate for the builder.
func (tppu *ThirdPartyPriceUpdate) Where(ps ...predicate.ThirdPartyPrice) *ThirdPartyPriceUpdate {
	tppu.predicates = append(tppu.predicates, ps...)
	return tppu
}

// SetInAppPurchaseSku sets the in_app_purchase_sku field.
func (tppu *ThirdPartyPriceUpdate) SetInAppPurchaseSku(s string) *ThirdPartyPriceUpdate {
	tppu.in_app_purchase_sku = &s
	return tppu
}

// SetGoogleBillingServiceSku sets the google_billing_service_sku field.
func (tppu *ThirdPartyPriceUpdate) SetGoogleBillingServiceSku(s string) *ThirdPartyPriceUpdate {
	tppu.google_billing_service_sku = &s
	return tppu
}

// SetNillableGoogleBillingServiceSku sets the google_billing_service_sku field if the given value is not nil.
func (tppu *ThirdPartyPriceUpdate) SetNillableGoogleBillingServiceSku(s *string) *ThirdPartyPriceUpdate {
	if s != nil {
		tppu.SetGoogleBillingServiceSku(*s)
	}
	return tppu
}

// ClearGoogleBillingServiceSku clears the value of google_billing_service_sku.
func (tppu *ThirdPartyPriceUpdate) ClearGoogleBillingServiceSku() *ThirdPartyPriceUpdate {
	tppu.google_billing_service_sku = nil
	tppu.cleargoogle_billing_service_sku = true
	return tppu
}

// SetStartAt sets the start_at field.
func (tppu *ThirdPartyPriceUpdate) SetStartAt(t time.Time) *ThirdPartyPriceUpdate {
	tppu.start_at = &t
	return tppu
}

// SetEndAt sets the end_at field.
func (tppu *ThirdPartyPriceUpdate) SetEndAt(t time.Time) *ThirdPartyPriceUpdate {
	tppu.end_at = &t
	return tppu
}

// SetType sets the type field.
func (tppu *ThirdPartyPriceUpdate) SetType(s string) *ThirdPartyPriceUpdate {
	tppu._type = &s
	return tppu
}

// AddPriceIDs adds the prices edge to Price by ids.
func (tppu *ThirdPartyPriceUpdate) AddPriceIDs(ids ...int) *ThirdPartyPriceUpdate {
	if tppu.prices == nil {
		tppu.prices = make(map[int]struct{})
	}
	for i := range ids {
		tppu.prices[ids[i]] = struct{}{}
	}
	return tppu
}

// AddPrices adds the prices edges to Price.
func (tppu *ThirdPartyPriceUpdate) AddPrices(p ...*Price) *ThirdPartyPriceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tppu.AddPriceIDs(ids...)
}

// AddPlanIDs adds the plans edge to Plan by ids.
func (tppu *ThirdPartyPriceUpdate) AddPlanIDs(ids ...int) *ThirdPartyPriceUpdate {
	if tppu.plans == nil {
		tppu.plans = make(map[int]struct{})
	}
	for i := range ids {
		tppu.plans[ids[i]] = struct{}{}
	}
	return tppu
}

// AddPlans adds the plans edges to Plan.
func (tppu *ThirdPartyPriceUpdate) AddPlans(p ...*Plan) *ThirdPartyPriceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tppu.AddPlanIDs(ids...)
}

// RemovePriceIDs removes the prices edge to Price by ids.
func (tppu *ThirdPartyPriceUpdate) RemovePriceIDs(ids ...int) *ThirdPartyPriceUpdate {
	if tppu.removedPrices == nil {
		tppu.removedPrices = make(map[int]struct{})
	}
	for i := range ids {
		tppu.removedPrices[ids[i]] = struct{}{}
	}
	return tppu
}

// RemovePrices removes prices edges to Price.
func (tppu *ThirdPartyPriceUpdate) RemovePrices(p ...*Price) *ThirdPartyPriceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tppu.RemovePriceIDs(ids...)
}

// RemovePlanIDs removes the plans edge to Plan by ids.
func (tppu *ThirdPartyPriceUpdate) RemovePlanIDs(ids ...int) *ThirdPartyPriceUpdate {
	if tppu.removedPlans == nil {
		tppu.removedPlans = make(map[int]struct{})
	}
	for i := range ids {
		tppu.removedPlans[ids[i]] = struct{}{}
	}
	return tppu
}

// RemovePlans removes plans edges to Plan.
func (tppu *ThirdPartyPriceUpdate) RemovePlans(p ...*Plan) *ThirdPartyPriceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tppu.RemovePlanIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tppu *ThirdPartyPriceUpdate) Save(ctx context.Context) (int, error) {
	if tppu.in_app_purchase_sku != nil {
		if err := thirdpartyprice.InAppPurchaseSkuValidator(*tppu.in_app_purchase_sku); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"in_app_purchase_sku\": %v", err)
		}
	}
	if tppu.google_billing_service_sku != nil {
		if err := thirdpartyprice.GoogleBillingServiceSkuValidator(*tppu.google_billing_service_sku); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"google_billing_service_sku\": %v", err)
		}
	}
	if tppu._type != nil {
		if err := thirdpartyprice.TypeValidator(*tppu._type); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"type\": %v", err)
		}
	}
	return tppu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tppu *ThirdPartyPriceUpdate) SaveX(ctx context.Context) int {
	affected, err := tppu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tppu *ThirdPartyPriceUpdate) Exec(ctx context.Context) error {
	_, err := tppu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tppu *ThirdPartyPriceUpdate) ExecX(ctx context.Context) {
	if err := tppu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tppu *ThirdPartyPriceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   thirdpartyprice.Table,
			Columns: thirdpartyprice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: thirdpartyprice.FieldID,
			},
		},
	}
	if ps := tppu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := tppu.in_app_purchase_sku; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: thirdpartyprice.FieldInAppPurchaseSku,
		})
	}
	if value := tppu.google_billing_service_sku; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: thirdpartyprice.FieldGoogleBillingServiceSku,
		})
	}
	if tppu.cleargoogle_billing_service_sku {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: thirdpartyprice.FieldGoogleBillingServiceSku,
		})
	}
	if value := tppu.start_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: thirdpartyprice.FieldStartAt,
		})
	}
	if value := tppu.end_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: thirdpartyprice.FieldEndAt,
		})
	}
	if value := tppu._type; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: thirdpartyprice.FieldType,
		})
	}
	if nodes := tppu.removedPrices; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thirdpartyprice.PricesTable,
			Columns: []string{thirdpartyprice.PricesColumn},
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
	if nodes := tppu.prices; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thirdpartyprice.PricesTable,
			Columns: []string{thirdpartyprice.PricesColumn},
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
	if nodes := tppu.removedPlans; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thirdpartyprice.PlansTable,
			Columns: []string{thirdpartyprice.PlansColumn},
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
	if nodes := tppu.plans; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thirdpartyprice.PlansTable,
			Columns: []string{thirdpartyprice.PlansColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tppu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ThirdPartyPriceUpdateOne is the builder for updating a single ThirdPartyPrice entity.
type ThirdPartyPriceUpdateOne struct {
	config
	id                              int
	in_app_purchase_sku             *string
	google_billing_service_sku      *string
	cleargoogle_billing_service_sku bool
	start_at                        *time.Time
	end_at                          *time.Time
	_type                           *string
	prices                          map[int]struct{}
	plans                           map[int]struct{}
	removedPrices                   map[int]struct{}
	removedPlans                    map[int]struct{}
}

// SetInAppPurchaseSku sets the in_app_purchase_sku field.
func (tppuo *ThirdPartyPriceUpdateOne) SetInAppPurchaseSku(s string) *ThirdPartyPriceUpdateOne {
	tppuo.in_app_purchase_sku = &s
	return tppuo
}

// SetGoogleBillingServiceSku sets the google_billing_service_sku field.
func (tppuo *ThirdPartyPriceUpdateOne) SetGoogleBillingServiceSku(s string) *ThirdPartyPriceUpdateOne {
	tppuo.google_billing_service_sku = &s
	return tppuo
}

// SetNillableGoogleBillingServiceSku sets the google_billing_service_sku field if the given value is not nil.
func (tppuo *ThirdPartyPriceUpdateOne) SetNillableGoogleBillingServiceSku(s *string) *ThirdPartyPriceUpdateOne {
	if s != nil {
		tppuo.SetGoogleBillingServiceSku(*s)
	}
	return tppuo
}

// ClearGoogleBillingServiceSku clears the value of google_billing_service_sku.
func (tppuo *ThirdPartyPriceUpdateOne) ClearGoogleBillingServiceSku() *ThirdPartyPriceUpdateOne {
	tppuo.google_billing_service_sku = nil
	tppuo.cleargoogle_billing_service_sku = true
	return tppuo
}

// SetStartAt sets the start_at field.
func (tppuo *ThirdPartyPriceUpdateOne) SetStartAt(t time.Time) *ThirdPartyPriceUpdateOne {
	tppuo.start_at = &t
	return tppuo
}

// SetEndAt sets the end_at field.
func (tppuo *ThirdPartyPriceUpdateOne) SetEndAt(t time.Time) *ThirdPartyPriceUpdateOne {
	tppuo.end_at = &t
	return tppuo
}

// SetType sets the type field.
func (tppuo *ThirdPartyPriceUpdateOne) SetType(s string) *ThirdPartyPriceUpdateOne {
	tppuo._type = &s
	return tppuo
}

// AddPriceIDs adds the prices edge to Price by ids.
func (tppuo *ThirdPartyPriceUpdateOne) AddPriceIDs(ids ...int) *ThirdPartyPriceUpdateOne {
	if tppuo.prices == nil {
		tppuo.prices = make(map[int]struct{})
	}
	for i := range ids {
		tppuo.prices[ids[i]] = struct{}{}
	}
	return tppuo
}

// AddPrices adds the prices edges to Price.
func (tppuo *ThirdPartyPriceUpdateOne) AddPrices(p ...*Price) *ThirdPartyPriceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tppuo.AddPriceIDs(ids...)
}

// AddPlanIDs adds the plans edge to Plan by ids.
func (tppuo *ThirdPartyPriceUpdateOne) AddPlanIDs(ids ...int) *ThirdPartyPriceUpdateOne {
	if tppuo.plans == nil {
		tppuo.plans = make(map[int]struct{})
	}
	for i := range ids {
		tppuo.plans[ids[i]] = struct{}{}
	}
	return tppuo
}

// AddPlans adds the plans edges to Plan.
func (tppuo *ThirdPartyPriceUpdateOne) AddPlans(p ...*Plan) *ThirdPartyPriceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tppuo.AddPlanIDs(ids...)
}

// RemovePriceIDs removes the prices edge to Price by ids.
func (tppuo *ThirdPartyPriceUpdateOne) RemovePriceIDs(ids ...int) *ThirdPartyPriceUpdateOne {
	if tppuo.removedPrices == nil {
		tppuo.removedPrices = make(map[int]struct{})
	}
	for i := range ids {
		tppuo.removedPrices[ids[i]] = struct{}{}
	}
	return tppuo
}

// RemovePrices removes prices edges to Price.
func (tppuo *ThirdPartyPriceUpdateOne) RemovePrices(p ...*Price) *ThirdPartyPriceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tppuo.RemovePriceIDs(ids...)
}

// RemovePlanIDs removes the plans edge to Plan by ids.
func (tppuo *ThirdPartyPriceUpdateOne) RemovePlanIDs(ids ...int) *ThirdPartyPriceUpdateOne {
	if tppuo.removedPlans == nil {
		tppuo.removedPlans = make(map[int]struct{})
	}
	for i := range ids {
		tppuo.removedPlans[ids[i]] = struct{}{}
	}
	return tppuo
}

// RemovePlans removes plans edges to Plan.
func (tppuo *ThirdPartyPriceUpdateOne) RemovePlans(p ...*Plan) *ThirdPartyPriceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tppuo.RemovePlanIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (tppuo *ThirdPartyPriceUpdateOne) Save(ctx context.Context) (*ThirdPartyPrice, error) {
	if tppuo.in_app_purchase_sku != nil {
		if err := thirdpartyprice.InAppPurchaseSkuValidator(*tppuo.in_app_purchase_sku); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"in_app_purchase_sku\": %v", err)
		}
	}
	if tppuo.google_billing_service_sku != nil {
		if err := thirdpartyprice.GoogleBillingServiceSkuValidator(*tppuo.google_billing_service_sku); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"google_billing_service_sku\": %v", err)
		}
	}
	if tppuo._type != nil {
		if err := thirdpartyprice.TypeValidator(*tppuo._type); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"type\": %v", err)
		}
	}
	return tppuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tppuo *ThirdPartyPriceUpdateOne) SaveX(ctx context.Context) *ThirdPartyPrice {
	tpp, err := tppuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return tpp
}

// Exec executes the query on the entity.
func (tppuo *ThirdPartyPriceUpdateOne) Exec(ctx context.Context) error {
	_, err := tppuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tppuo *ThirdPartyPriceUpdateOne) ExecX(ctx context.Context) {
	if err := tppuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tppuo *ThirdPartyPriceUpdateOne) sqlSave(ctx context.Context) (tpp *ThirdPartyPrice, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   thirdpartyprice.Table,
			Columns: thirdpartyprice.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  tppuo.id,
				Type:   field.TypeInt,
				Column: thirdpartyprice.FieldID,
			},
		},
	}
	if value := tppuo.in_app_purchase_sku; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: thirdpartyprice.FieldInAppPurchaseSku,
		})
	}
	if value := tppuo.google_billing_service_sku; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: thirdpartyprice.FieldGoogleBillingServiceSku,
		})
	}
	if tppuo.cleargoogle_billing_service_sku {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: thirdpartyprice.FieldGoogleBillingServiceSku,
		})
	}
	if value := tppuo.start_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: thirdpartyprice.FieldStartAt,
		})
	}
	if value := tppuo.end_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: thirdpartyprice.FieldEndAt,
		})
	}
	if value := tppuo._type; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: thirdpartyprice.FieldType,
		})
	}
	if nodes := tppuo.removedPrices; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thirdpartyprice.PricesTable,
			Columns: []string{thirdpartyprice.PricesColumn},
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
	if nodes := tppuo.prices; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thirdpartyprice.PricesTable,
			Columns: []string{thirdpartyprice.PricesColumn},
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
	if nodes := tppuo.removedPlans; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thirdpartyprice.PlansTable,
			Columns: []string{thirdpartyprice.PlansColumn},
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
	if nodes := tppuo.plans; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thirdpartyprice.PlansTable,
			Columns: []string{thirdpartyprice.PlansColumn},
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
	tpp = &ThirdPartyPrice{config: tppuo.config}
	_spec.Assign = tpp.assignValues
	_spec.ScanValues = tpp.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tppuo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return tpp, nil
}

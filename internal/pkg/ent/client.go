// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/pepeunlimited/prices/internal/pkg/ent/migrate"

	"github.com/pepeunlimited/prices/internal/pkg/ent/plan"
	"github.com/pepeunlimited/prices/internal/pkg/ent/price"
	"github.com/pepeunlimited/prices/internal/pkg/ent/product"
	"github.com/pepeunlimited/prices/internal/pkg/ent/subscription"
	"github.com/pepeunlimited/prices/internal/pkg/ent/thirdparty"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Plan is the client for interacting with the Plan builders.
	Plan *PlanClient
	// Price is the client for interacting with the Price builders.
	Price *PriceClient
	// Product is the client for interacting with the Product builders.
	Product *ProductClient
	// Subscription is the client for interacting with the Subscription builders.
	Subscription *SubscriptionClient
	// ThirdParty is the client for interacting with the ThirdParty builders.
	ThirdParty *ThirdPartyClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config:       c,
		Schema:       migrate.NewSchema(c.driver),
		Plan:         NewPlanClient(c),
		Price:        NewPriceClient(c),
		Product:      NewProductClient(c),
		Subscription: NewSubscriptionClient(c),
		ThirdParty:   NewThirdPartyClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
		config:       cfg,
		Plan:         NewPlanClient(cfg),
		Price:        NewPriceClient(cfg),
		Product:      NewProductClient(cfg),
		Subscription: NewSubscriptionClient(cfg),
		ThirdParty:   NewThirdPartyClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Plan.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config:       cfg,
		Schema:       migrate.NewSchema(cfg.driver),
		Plan:         NewPlanClient(cfg),
		Price:        NewPriceClient(cfg),
		Product:      NewProductClient(cfg),
		Subscription: NewSubscriptionClient(cfg),
		ThirdParty:   NewThirdPartyClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// PlanClient is a client for the Plan schema.
type PlanClient struct {
	config
}

// NewPlanClient returns a client for the Plan from the given config.
func NewPlanClient(c config) *PlanClient {
	return &PlanClient{config: c}
}

// Create returns a create builder for Plan.
func (c *PlanClient) Create() *PlanCreate {
	return &PlanCreate{config: c.config}
}

// Update returns an update builder for Plan.
func (c *PlanClient) Update() *PlanUpdate {
	return &PlanUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlanClient) UpdateOne(pl *Plan) *PlanUpdateOne {
	return c.UpdateOneID(pl.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *PlanClient) UpdateOneID(id int) *PlanUpdateOne {
	return &PlanUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Plan.
func (c *PlanClient) Delete() *PlanDelete {
	return &PlanDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PlanClient) DeleteOne(pl *Plan) *PlanDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PlanClient) DeleteOneID(id int) *PlanDeleteOne {
	return &PlanDeleteOne{c.Delete().Where(plan.ID(id))}
}

// Create returns a query builder for Plan.
func (c *PlanClient) Query() *PlanQuery {
	return &PlanQuery{config: c.config}
}

// Get returns a Plan entity by its id.
func (c *PlanClient) Get(ctx context.Context, id int) (*Plan, error) {
	return c.Query().Where(plan.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlanClient) GetX(ctx context.Context, id int) *Plan {
	pl, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pl
}

// QuerySubscriptions queries the subscriptions edge of a Plan.
func (c *PlanClient) QuerySubscriptions(pl *Plan) *SubscriptionQuery {
	query := &SubscriptionQuery{config: c.config}
	id := pl.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(plan.Table, plan.FieldID, id),
		sqlgraph.To(subscription.Table, subscription.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, plan.SubscriptionsTable, plan.SubscriptionsColumn),
	)
	query.sql = sqlgraph.Neighbors(pl.driver.Dialect(), step)

	return query
}

// QueryPrices queries the prices edge of a Plan.
func (c *PlanClient) QueryPrices(pl *Plan) *PriceQuery {
	query := &PriceQuery{config: c.config}
	id := pl.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(plan.Table, plan.FieldID, id),
		sqlgraph.To(price.Table, price.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, plan.PricesTable, plan.PricesColumn),
	)
	query.sql = sqlgraph.Neighbors(pl.driver.Dialect(), step)

	return query
}

// PriceClient is a client for the Price schema.
type PriceClient struct {
	config
}

// NewPriceClient returns a client for the Price from the given config.
func NewPriceClient(c config) *PriceClient {
	return &PriceClient{config: c}
}

// Create returns a create builder for Price.
func (c *PriceClient) Create() *PriceCreate {
	return &PriceCreate{config: c.config}
}

// Update returns an update builder for Price.
func (c *PriceClient) Update() *PriceUpdate {
	return &PriceUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *PriceClient) UpdateOne(pr *Price) *PriceUpdateOne {
	return c.UpdateOneID(pr.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *PriceClient) UpdateOneID(id int) *PriceUpdateOne {
	return &PriceUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Price.
func (c *PriceClient) Delete() *PriceDelete {
	return &PriceDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PriceClient) DeleteOne(pr *Price) *PriceDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PriceClient) DeleteOneID(id int) *PriceDeleteOne {
	return &PriceDeleteOne{c.Delete().Where(price.ID(id))}
}

// Create returns a query builder for Price.
func (c *PriceClient) Query() *PriceQuery {
	return &PriceQuery{config: c.config}
}

// Get returns a Price entity by its id.
func (c *PriceClient) Get(ctx context.Context, id int) (*Price, error) {
	return c.Query().Where(price.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PriceClient) GetX(ctx context.Context, id int) *Price {
	pr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pr
}

// QueryProducts queries the products edge of a Price.
func (c *PriceClient) QueryProducts(pr *Price) *ProductQuery {
	query := &ProductQuery{config: c.config}
	id := pr.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(price.Table, price.FieldID, id),
		sqlgraph.To(product.Table, product.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, price.ProductsTable, price.ProductsColumn),
	)
	query.sql = sqlgraph.Neighbors(pr.driver.Dialect(), step)

	return query
}

// QueryThirdParties queries the third_parties edge of a Price.
func (c *PriceClient) QueryThirdParties(pr *Price) *ThirdPartyQuery {
	query := &ThirdPartyQuery{config: c.config}
	id := pr.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(price.Table, price.FieldID, id),
		sqlgraph.To(thirdparty.Table, thirdparty.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, price.ThirdPartiesTable, price.ThirdPartiesColumn),
	)
	query.sql = sqlgraph.Neighbors(pr.driver.Dialect(), step)

	return query
}

// QueryPlans queries the plans edge of a Price.
func (c *PriceClient) QueryPlans(pr *Price) *PlanQuery {
	query := &PlanQuery{config: c.config}
	id := pr.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(price.Table, price.FieldID, id),
		sqlgraph.To(plan.Table, plan.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, price.PlansTable, price.PlansColumn),
	)
	query.sql = sqlgraph.Neighbors(pr.driver.Dialect(), step)

	return query
}

// ProductClient is a client for the Product schema.
type ProductClient struct {
	config
}

// NewProductClient returns a client for the Product from the given config.
func NewProductClient(c config) *ProductClient {
	return &ProductClient{config: c}
}

// Create returns a create builder for Product.
func (c *ProductClient) Create() *ProductCreate {
	return &ProductCreate{config: c.config}
}

// Update returns an update builder for Product.
func (c *ProductClient) Update() *ProductUpdate {
	return &ProductUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProductClient) UpdateOne(pr *Product) *ProductUpdateOne {
	return c.UpdateOneID(pr.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *ProductClient) UpdateOneID(id int) *ProductUpdateOne {
	return &ProductUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Product.
func (c *ProductClient) Delete() *ProductDelete {
	return &ProductDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProductClient) DeleteOne(pr *Product) *ProductDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProductClient) DeleteOneID(id int) *ProductDeleteOne {
	return &ProductDeleteOne{c.Delete().Where(product.ID(id))}
}

// Create returns a query builder for Product.
func (c *ProductClient) Query() *ProductQuery {
	return &ProductQuery{config: c.config}
}

// Get returns a Product entity by its id.
func (c *ProductClient) Get(ctx context.Context, id int) (*Product, error) {
	return c.Query().Where(product.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProductClient) GetX(ctx context.Context, id int) *Product {
	pr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pr
}

// QueryPrices queries the prices edge of a Product.
func (c *ProductClient) QueryPrices(pr *Product) *PriceQuery {
	query := &PriceQuery{config: c.config}
	id := pr.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(product.Table, product.FieldID, id),
		sqlgraph.To(price.Table, price.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, product.PricesTable, product.PricesColumn),
	)
	query.sql = sqlgraph.Neighbors(pr.driver.Dialect(), step)

	return query
}

// SubscriptionClient is a client for the Subscription schema.
type SubscriptionClient struct {
	config
}

// NewSubscriptionClient returns a client for the Subscription from the given config.
func NewSubscriptionClient(c config) *SubscriptionClient {
	return &SubscriptionClient{config: c}
}

// Create returns a create builder for Subscription.
func (c *SubscriptionClient) Create() *SubscriptionCreate {
	return &SubscriptionCreate{config: c.config}
}

// Update returns an update builder for Subscription.
func (c *SubscriptionClient) Update() *SubscriptionUpdate {
	return &SubscriptionUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubscriptionClient) UpdateOne(s *Subscription) *SubscriptionUpdateOne {
	return c.UpdateOneID(s.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *SubscriptionClient) UpdateOneID(id int) *SubscriptionUpdateOne {
	return &SubscriptionUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Subscription.
func (c *SubscriptionClient) Delete() *SubscriptionDelete {
	return &SubscriptionDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SubscriptionClient) DeleteOne(s *Subscription) *SubscriptionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SubscriptionClient) DeleteOneID(id int) *SubscriptionDeleteOne {
	return &SubscriptionDeleteOne{c.Delete().Where(subscription.ID(id))}
}

// Create returns a query builder for Subscription.
func (c *SubscriptionClient) Query() *SubscriptionQuery {
	return &SubscriptionQuery{config: c.config}
}

// Get returns a Subscription entity by its id.
func (c *SubscriptionClient) Get(ctx context.Context, id int) (*Subscription, error) {
	return c.Query().Where(subscription.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubscriptionClient) GetX(ctx context.Context, id int) *Subscription {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// QueryPlans queries the plans edge of a Subscription.
func (c *SubscriptionClient) QueryPlans(s *Subscription) *PlanQuery {
	query := &PlanQuery{config: c.config}
	id := s.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(subscription.Table, subscription.FieldID, id),
		sqlgraph.To(plan.Table, plan.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, subscription.PlansTable, subscription.PlansColumn),
	)
	query.sql = sqlgraph.Neighbors(s.driver.Dialect(), step)

	return query
}

// ThirdPartyClient is a client for the ThirdParty schema.
type ThirdPartyClient struct {
	config
}

// NewThirdPartyClient returns a client for the ThirdParty from the given config.
func NewThirdPartyClient(c config) *ThirdPartyClient {
	return &ThirdPartyClient{config: c}
}

// Create returns a create builder for ThirdParty.
func (c *ThirdPartyClient) Create() *ThirdPartyCreate {
	return &ThirdPartyCreate{config: c.config}
}

// Update returns an update builder for ThirdParty.
func (c *ThirdPartyClient) Update() *ThirdPartyUpdate {
	return &ThirdPartyUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *ThirdPartyClient) UpdateOne(tp *ThirdParty) *ThirdPartyUpdateOne {
	return c.UpdateOneID(tp.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *ThirdPartyClient) UpdateOneID(id int) *ThirdPartyUpdateOne {
	return &ThirdPartyUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for ThirdParty.
func (c *ThirdPartyClient) Delete() *ThirdPartyDelete {
	return &ThirdPartyDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ThirdPartyClient) DeleteOne(tp *ThirdParty) *ThirdPartyDeleteOne {
	return c.DeleteOneID(tp.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ThirdPartyClient) DeleteOneID(id int) *ThirdPartyDeleteOne {
	return &ThirdPartyDeleteOne{c.Delete().Where(thirdparty.ID(id))}
}

// Create returns a query builder for ThirdParty.
func (c *ThirdPartyClient) Query() *ThirdPartyQuery {
	return &ThirdPartyQuery{config: c.config}
}

// Get returns a ThirdParty entity by its id.
func (c *ThirdPartyClient) Get(ctx context.Context, id int) (*ThirdParty, error) {
	return c.Query().Where(thirdparty.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ThirdPartyClient) GetX(ctx context.Context, id int) *ThirdParty {
	tp, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return tp
}

// QueryPrices queries the prices edge of a ThirdParty.
func (c *ThirdPartyClient) QueryPrices(tp *ThirdParty) *PriceQuery {
	query := &PriceQuery{config: c.config}
	id := tp.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(thirdparty.Table, thirdparty.FieldID, id),
		sqlgraph.To(price.Table, price.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, thirdparty.PricesTable, thirdparty.PricesColumn),
	)
	query.sql = sqlgraph.Neighbors(tp.driver.Dialect(), step)

	return query
}
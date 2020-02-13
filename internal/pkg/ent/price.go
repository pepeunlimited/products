// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pepeunlimited/prices/internal/pkg/ent/iapsource"
	"github.com/pepeunlimited/prices/internal/pkg/ent/plan"
	"github.com/pepeunlimited/prices/internal/pkg/ent/price"
	"github.com/pepeunlimited/prices/internal/pkg/ent/product"
)

// Price is the model entity for the Price schema.
type Price struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt time.Time `json:"start_at,omitempty"`
	// EndAt holds the value of the "end_at" field.
	EndAt time.Time `json:"end_at,omitempty"`
	// Price holds the value of the "price" field.
	Price uint16 `json:"price,omitempty"`
	// Discount holds the value of the "discount" field.
	Discount uint16 `json:"discount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PriceQuery when eager-loading is set.
	Edges             PriceEdges `json:"edges"`
	iap_source_prices *int
	plan_prices       *int
	product_prices    *int
}

// PriceEdges holds the relations/edges for other nodes in the graph.
type PriceEdges struct {
	// Products holds the value of the products edge.
	Products *Product
	// IapSource holds the value of the iap_source edge.
	IapSource *IapSource
	// Plans holds the value of the plans edge.
	Plans *Plan
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PriceEdges) ProductsOrErr() (*Product, error) {
	if e.loadedTypes[0] {
		if e.Products == nil {
			// The edge products was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// IapSourceOrErr returns the IapSource value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PriceEdges) IapSourceOrErr() (*IapSource, error) {
	if e.loadedTypes[1] {
		if e.IapSource == nil {
			// The edge iap_source was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: iapsource.Label}
		}
		return e.IapSource, nil
	}
	return nil, &NotLoadedError{edge: "iap_source"}
}

// PlansOrErr returns the Plans value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PriceEdges) PlansOrErr() (*Plan, error) {
	if e.loadedTypes[2] {
		if e.Plans == nil {
			// The edge plans was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: plan.Label}
		}
		return e.Plans, nil
	}
	return nil, &NotLoadedError{edge: "plans"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Price) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // start_at
		&sql.NullTime{},  // end_at
		&sql.NullInt64{}, // price
		&sql.NullInt64{}, // discount
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Price) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // iap_source_prices
		&sql.NullInt64{}, // plan_prices
		&sql.NullInt64{}, // product_prices
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Price fields.
func (pr *Price) assignValues(values ...interface{}) error {
	if m, n := len(values), len(price.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field start_at", values[0])
	} else if value.Valid {
		pr.StartAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field end_at", values[1])
	} else if value.Valid {
		pr.EndAt = value.Time
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field price", values[2])
	} else if value.Valid {
		pr.Price = uint16(value.Int64)
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field discount", values[3])
	} else if value.Valid {
		pr.Discount = uint16(value.Int64)
	}
	values = values[4:]
	if len(values) == len(price.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field iap_source_prices", value)
		} else if value.Valid {
			pr.iap_source_prices = new(int)
			*pr.iap_source_prices = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field plan_prices", value)
		} else if value.Valid {
			pr.plan_prices = new(int)
			*pr.plan_prices = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field product_prices", value)
		} else if value.Valid {
			pr.product_prices = new(int)
			*pr.product_prices = int(value.Int64)
		}
	}
	return nil
}

// QueryProducts queries the products edge of the Price.
func (pr *Price) QueryProducts() *ProductQuery {
	return (&PriceClient{pr.config}).QueryProducts(pr)
}

// QueryIapSource queries the iap_source edge of the Price.
func (pr *Price) QueryIapSource() *IapSourceQuery {
	return (&PriceClient{pr.config}).QueryIapSource(pr)
}

// QueryPlans queries the plans edge of the Price.
func (pr *Price) QueryPlans() *PlanQuery {
	return (&PriceClient{pr.config}).QueryPlans(pr)
}

// Update returns a builder for updating this Price.
// Note that, you need to call Price.Unwrap() before calling this method, if this Price
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Price) Update() *PriceUpdateOne {
	return (&PriceClient{pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pr *Price) Unwrap() *Price {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Price is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Price) String() string {
	var builder strings.Builder
	builder.WriteString("Price(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", start_at=")
	builder.WriteString(pr.StartAt.Format(time.ANSIC))
	builder.WriteString(", end_at=")
	builder.WriteString(pr.EndAt.Format(time.ANSIC))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", pr.Price))
	builder.WriteString(", discount=")
	builder.WriteString(fmt.Sprintf("%v", pr.Discount))
	builder.WriteByte(')')
	return builder.String()
}

// Prices is a parsable slice of Price.
type Prices []*Price

func (pr Prices) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}

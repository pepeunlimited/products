// Code generated by entc, DO NOT EDIT.

package price

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/pepeunlimited/products/internal/pkg/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// Discount applies equality check predicate on the "discount" field. It's identical to DiscountEQ.
func Discount(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscount), v))
	})
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartAt), v))
	})
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...time.Time) predicate.Price {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Price(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartAt), v...))
	})
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...time.Time) predicate.Price {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Price(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartAt), v...))
	})
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartAt), v))
	})
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartAt), v))
	})
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartAt), v))
	})
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartAt), v))
	})
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndAt), v))
	})
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.Price {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Price(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndAt), v...))
	})
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.Price {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Price(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndAt), v...))
	})
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndAt), v))
	})
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndAt), v))
	})
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndAt), v))
	})
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndAt), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int64) predicate.Price {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Price(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int64) predicate.Price {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Price(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// DiscountEQ applies the EQ predicate on the "discount" field.
func DiscountEQ(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscount), v))
	})
}

// DiscountNEQ applies the NEQ predicate on the "discount" field.
func DiscountNEQ(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiscount), v))
	})
}

// DiscountIn applies the In predicate on the "discount" field.
func DiscountIn(vs ...int64) predicate.Price {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Price(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiscount), v...))
	})
}

// DiscountNotIn applies the NotIn predicate on the "discount" field.
func DiscountNotIn(vs ...int64) predicate.Price {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Price(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiscount), v...))
	})
}

// DiscountGT applies the GT predicate on the "discount" field.
func DiscountGT(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiscount), v))
	})
}

// DiscountGTE applies the GTE predicate on the "discount" field.
func DiscountGTE(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiscount), v))
	})
}

// DiscountLT applies the LT predicate on the "discount" field.
func DiscountLT(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiscount), v))
	})
}

// DiscountLTE applies the LTE predicate on the "discount" field.
func DiscountLTE(v int64) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiscount), v))
	})
}

// HasProducts applies the HasEdge predicate on the "products" edge.
func HasProducts() predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductsTable, ProductsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductsWith applies the HasEdge predicate on the "products" edge with a given conditions (other predicates).
func HasProductsWith(preds ...predicate.Product) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductsTable, ProductsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasThirdPartyPrices applies the HasEdge predicate on the "third_party_prices" edge.
func HasThirdPartyPrices() predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ThirdPartyPricesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ThirdPartyPricesTable, ThirdPartyPricesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasThirdPartyPricesWith applies the HasEdge predicate on the "third_party_prices" edge with a given conditions (other predicates).
func HasThirdPartyPricesWith(preds ...predicate.ThirdPartyPrice) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ThirdPartyPricesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ThirdPartyPricesTable, ThirdPartyPricesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Price) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Price) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Price) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		p(s.Not())
	})
}

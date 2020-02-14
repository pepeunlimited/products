// Code generated by entc, DO NOT EDIT.

package plan

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/pepeunlimited/prices/internal/pkg/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TitleI18nID applies equality check predicate on the "title_i18n_id" field. It's identical to TitleI18nIDEQ.
func TitleI18nID(v int64) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitleI18nID), v))
	})
}

// Length applies equality check predicate on the "length" field. It's identical to LengthEQ.
func Length(v uint8) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLength), v))
	})
}

// Unit applies equality check predicate on the "unit" field. It's identical to UnitEQ.
func Unit(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnit), v))
	})
}

// TitleI18nIDEQ applies the EQ predicate on the "title_i18n_id" field.
func TitleI18nIDEQ(v int64) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitleI18nID), v))
	})
}

// TitleI18nIDNEQ applies the NEQ predicate on the "title_i18n_id" field.
func TitleI18nIDNEQ(v int64) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitleI18nID), v))
	})
}

// TitleI18nIDIn applies the In predicate on the "title_i18n_id" field.
func TitleI18nIDIn(vs ...int64) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitleI18nID), v...))
	})
}

// TitleI18nIDNotIn applies the NotIn predicate on the "title_i18n_id" field.
func TitleI18nIDNotIn(vs ...int64) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitleI18nID), v...))
	})
}

// TitleI18nIDGT applies the GT predicate on the "title_i18n_id" field.
func TitleI18nIDGT(v int64) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitleI18nID), v))
	})
}

// TitleI18nIDGTE applies the GTE predicate on the "title_i18n_id" field.
func TitleI18nIDGTE(v int64) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitleI18nID), v))
	})
}

// TitleI18nIDLT applies the LT predicate on the "title_i18n_id" field.
func TitleI18nIDLT(v int64) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitleI18nID), v))
	})
}

// TitleI18nIDLTE applies the LTE predicate on the "title_i18n_id" field.
func TitleI18nIDLTE(v int64) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitleI18nID), v))
	})
}

// LengthEQ applies the EQ predicate on the "length" field.
func LengthEQ(v uint8) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLength), v))
	})
}

// LengthNEQ applies the NEQ predicate on the "length" field.
func LengthNEQ(v uint8) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLength), v))
	})
}

// LengthIn applies the In predicate on the "length" field.
func LengthIn(vs ...uint8) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLength), v...))
	})
}

// LengthNotIn applies the NotIn predicate on the "length" field.
func LengthNotIn(vs ...uint8) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLength), v...))
	})
}

// LengthGT applies the GT predicate on the "length" field.
func LengthGT(v uint8) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLength), v))
	})
}

// LengthGTE applies the GTE predicate on the "length" field.
func LengthGTE(v uint8) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLength), v))
	})
}

// LengthLT applies the LT predicate on the "length" field.
func LengthLT(v uint8) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLength), v))
	})
}

// LengthLTE applies the LTE predicate on the "length" field.
func LengthLTE(v uint8) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLength), v))
	})
}

// UnitEQ applies the EQ predicate on the "unit" field.
func UnitEQ(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnit), v))
	})
}

// UnitNEQ applies the NEQ predicate on the "unit" field.
func UnitNEQ(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnit), v))
	})
}

// UnitIn applies the In predicate on the "unit" field.
func UnitIn(vs ...string) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUnit), v...))
	})
}

// UnitNotIn applies the NotIn predicate on the "unit" field.
func UnitNotIn(vs ...string) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUnit), v...))
	})
}

// UnitGT applies the GT predicate on the "unit" field.
func UnitGT(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnit), v))
	})
}

// UnitGTE applies the GTE predicate on the "unit" field.
func UnitGTE(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnit), v))
	})
}

// UnitLT applies the LT predicate on the "unit" field.
func UnitLT(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnit), v))
	})
}

// UnitLTE applies the LTE predicate on the "unit" field.
func UnitLTE(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnit), v))
	})
}

// UnitContains applies the Contains predicate on the "unit" field.
func UnitContains(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUnit), v))
	})
}

// UnitHasPrefix applies the HasPrefix predicate on the "unit" field.
func UnitHasPrefix(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUnit), v))
	})
}

// UnitHasSuffix applies the HasSuffix predicate on the "unit" field.
func UnitHasSuffix(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUnit), v))
	})
}

// UnitEqualFold applies the EqualFold predicate on the "unit" field.
func UnitEqualFold(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUnit), v))
	})
}

// UnitContainsFold applies the ContainsFold predicate on the "unit" field.
func UnitContainsFold(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUnit), v))
	})
}

// HasSubscriptions applies the HasEdge predicate on the "subscriptions" edge.
func HasSubscriptions() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubscriptionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubscriptionsTable, SubscriptionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubscriptionsWith applies the HasEdge predicate on the "subscriptions" edge with a given conditions (other predicates).
func HasSubscriptionsWith(preds ...predicate.Subscription) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubscriptionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubscriptionsTable, SubscriptionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrices applies the HasEdge predicate on the "prices" edge.
func HasPrices() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PricesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PricesTable, PricesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPricesWith applies the HasEdge predicate on the "prices" edge with a given conditions (other predicates).
func HasPricesWith(preds ...predicate.Price) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PricesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PricesTable, PricesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Plan) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Plan) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
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
func Not(p predicate.Plan) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		p(s.Not())
	})
}

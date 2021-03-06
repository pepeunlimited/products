package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Product struct {
	ent.Schema
}

func (Product) Config() ent.Config {
	return ent.Config{Table:"products"}
}

func (Product) Fields() []ent.Field {
	return []ent.Field {
		field.String("sku").Unique().MaxLen(32),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("prices", Price.Type),     // one-to-many
		edge.To("plans", Plan.Type),     // one-to-many
	}
}
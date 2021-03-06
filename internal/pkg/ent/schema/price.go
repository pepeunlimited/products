package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Price struct {
	ent.Schema
}

func (Price) Config() ent.Config {
	return ent.Config{Table:"prices"}
}

func (Price) Fields() []ent.Field {
	return []ent.Field {
		field.Time("start_at"),
		field.Time("end_at"),
		field.Int64("price"),
		field.Int64("discount"),
	}
}

func (Price) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("products", Product.Type).Ref("prices").Unique(), // many-to-one
		edge.From("third_party_prices", ThirdPartyPrice.Type).Ref("prices").Unique(), // many-to-one
	}
}
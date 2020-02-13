package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type InAppPurchase struct {
	ent.Schema
}

func (InAppPurchase) Config() ent.Config {
	return ent.Config{Table:"in_app_purchases"}
}

func (InAppPurchase) Fields() []ent.Field {
	return []ent.Field {
		field.Int64("apple_sku").Optional().Unique(),
		field.Int64("google_sku").Optional().Unique(),
		field.Time("start_at"),
		field.Time("end_at"),
	}
}

func (InAppPurchase) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("prices", Price.Type), // one-to-many
	}
}
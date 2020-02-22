package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type ThirdPartyPrice struct {
	ent.Schema
}

func (ThirdPartyPrice) Config() ent.Config {
	return ent.Config{Table:"third_party_prices"}
}

func (ThirdPartyPrice) Fields() []ent.Field {
	return []ent.Field {
		field.String("in_app_purchase_sku").MaxLen(32).Unique(),
		field.String("google_billing_service_sku").Optional().MaxLen(32).Unique(),
		field.Time("start_at"),
		field.Time("end_at"),
		field.String("type").MaxLen(28).Unique(),
	}
}

func (ThirdPartyPrice) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("prices", Price.Type), // one-to-many
		edge.To("plans", Plan.Type), // one-to-many
	}
}
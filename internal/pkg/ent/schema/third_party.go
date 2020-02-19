package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type ThirdParty struct {
	ent.Schema
}

func (ThirdParty) Config() ent.Config {
	return ent.Config{Table:"third_parties"}
}

func (ThirdParty) Fields() []ent.Field {
	return []ent.Field {
		field.String("in_app_purchase_sku").MaxLen(32).Unique(),
		field.String("google_billing_service_sku").Optional().MaxLen(32).Unique(),
		field.Time("start_at"),
		field.Time("end_at"),
	}
}

func (ThirdParty) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("prices", Price.Type), // one-to-many
	}
}
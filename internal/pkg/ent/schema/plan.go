package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// => `plans`
type Plan struct {
	ent.Schema
}

func (Plan) Config() ent.Config {
	return ent.Config{Table:"plans"}
}

func (Plan) Fields() []ent.Field {
	return []ent.Field {
		field.Int64("title_i18n_id"),
		field.Uint8("length"),
		field.Time("start_at"),
		field.Time("end_at"),
		field.Uint16("price"),
		field.Uint16("discount"),
		field.String("unit").MaxLen(7),
	}
}

func (Plan) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("subscriptions", Subscription.Type), // one-to-many
		edge.From("products", Product.Type).Ref("plans").Unique(), // many-to-one
		edge.From("third_party_prices", ThirdPartyPrice.Type).Ref("plans").Unique(), // many-to-one
	}
}
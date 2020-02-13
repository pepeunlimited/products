package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// => `subscriptions`
type Subscription struct {
	ent.Schema
}

func (Subscription) Config() ent.Config {
	return ent.Config{Table:"subscriptions"}
}

func (Subscription) Fields() []ent.Field {
	return []ent.Field {
		field.Int64("user_id"),
		field.Time("start_at"),
		field.Time("end_at"),
	}
}

func (Subscription) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("plans", Plan.Type).Ref("subscriptions").Unique(), // many-to-one
	}
}

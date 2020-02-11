package schema

import (
	"github.com/facebookincubator/ent"
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
		field.Uint16("cost"),
		field.Uint16("discount"),
	}
}

func (Price) Edges() []ent.Edge {
	return []ent.Edge {}
}
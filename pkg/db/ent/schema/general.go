package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/mixin"
	"github.com/google/uuid"
)

// General holds the schema definition for the General entity.
type General struct {
	ent.Schema
}

func (General) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the General.
func (General) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("app_id", uuid.UUID{}).Optional().Default(func() uuid.UUID {
			return uuid.UUID{}
		}),
		field.UUID("user_id", uuid.UUID{}).Optional().Default(func() uuid.UUID {
			return uuid.UUID{}
		}),
		field.UUID("coin_type_id", uuid.UUID{}).Optional().Default(func() uuid.UUID {
			return uuid.UUID{}
		}),
		field.Uint64("incoming").Optional().Default(0),
		field.Uint64("locked").Optional().Default(0),
		field.Uint64("outcoming").Optional().Default(0),
		field.Uint64("spendable").Optional().Default(0),
	}
}

// Edges of the General.
func (General) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/mixin"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/message/npool/ledgermgr/detail"
)

// Detail holds the schema definition for the Detail entity.
type Detail struct {
	ent.Schema
}

func (Detail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Detail.
func (Detail) Fields() []ent.Field {
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
		field.String("io_type").Optional().Default(detail.IOType_DefaultType.String()),
		field.String("io_sub_type").Optional().Default(detail.IOSubType_DefaultSubType.String()),
		field.Uint64("amount").Optional().Default(0),
		field.Uint32("amount_precision").Optional().Default(6),
		field.UUID("from_coin_type_id", uuid.UUID{}).Optional().Default(func() uuid.UUID {
			return uuid.UUID{}
		}),
		field.Uint64("coin_usd_currency").Optional().Default(0),
		field.String("io_extra").Optional().Default(""),
		field.UUID("from_old_id", uuid.UUID{}).Optional().Default(func() uuid.UUID {
			return uuid.UUID{}
		}),
	}
}

// Edges of the Detail.
func (Detail) Edges() []ent.Edge {
	return nil
}

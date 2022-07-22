// Code generated by entc, DO NOT EDIT.

package general

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// Incoming applies equality check predicate on the "incoming" field. It's identical to IncomingEQ.
func Incoming(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIncoming), v))
	})
}

// Locked applies equality check predicate on the "locked" field. It's identical to LockedEQ.
func Locked(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocked), v))
	})
}

// Outcoming applies equality check predicate on the "outcoming" field. It's identical to OutcomingEQ.
func Outcoming(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcoming), v))
	})
}

// Spendable applies equality check predicate on the "spendable" field. It's identical to SpendableEQ.
func Spendable(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpendable), v))
	})
}

// Precision applies equality check predicate on the "precision" field. It's identical to PrecisionEQ.
func Precision(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrecision), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIsNil applies the IsNil predicate on the "coin_type_id" field.
func CoinTypeIDIsNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoinTypeID)))
	})
}

// CoinTypeIDNotNil applies the NotNil predicate on the "coin_type_id" field.
func CoinTypeIDNotNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoinTypeID)))
	})
}

// IncomingEQ applies the EQ predicate on the "incoming" field.
func IncomingEQ(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIncoming), v))
	})
}

// IncomingNEQ applies the NEQ predicate on the "incoming" field.
func IncomingNEQ(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIncoming), v))
	})
}

// IncomingIn applies the In predicate on the "incoming" field.
func IncomingIn(vs ...uint64) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIncoming), v...))
	})
}

// IncomingNotIn applies the NotIn predicate on the "incoming" field.
func IncomingNotIn(vs ...uint64) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIncoming), v...))
	})
}

// IncomingGT applies the GT predicate on the "incoming" field.
func IncomingGT(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIncoming), v))
	})
}

// IncomingGTE applies the GTE predicate on the "incoming" field.
func IncomingGTE(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIncoming), v))
	})
}

// IncomingLT applies the LT predicate on the "incoming" field.
func IncomingLT(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIncoming), v))
	})
}

// IncomingLTE applies the LTE predicate on the "incoming" field.
func IncomingLTE(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIncoming), v))
	})
}

// IncomingIsNil applies the IsNil predicate on the "incoming" field.
func IncomingIsNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIncoming)))
	})
}

// IncomingNotNil applies the NotNil predicate on the "incoming" field.
func IncomingNotNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIncoming)))
	})
}

// LockedEQ applies the EQ predicate on the "locked" field.
func LockedEQ(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocked), v))
	})
}

// LockedNEQ applies the NEQ predicate on the "locked" field.
func LockedNEQ(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocked), v))
	})
}

// LockedIn applies the In predicate on the "locked" field.
func LockedIn(vs ...uint64) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLocked), v...))
	})
}

// LockedNotIn applies the NotIn predicate on the "locked" field.
func LockedNotIn(vs ...uint64) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLocked), v...))
	})
}

// LockedGT applies the GT predicate on the "locked" field.
func LockedGT(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocked), v))
	})
}

// LockedGTE applies the GTE predicate on the "locked" field.
func LockedGTE(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocked), v))
	})
}

// LockedLT applies the LT predicate on the "locked" field.
func LockedLT(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocked), v))
	})
}

// LockedLTE applies the LTE predicate on the "locked" field.
func LockedLTE(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocked), v))
	})
}

// LockedIsNil applies the IsNil predicate on the "locked" field.
func LockedIsNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLocked)))
	})
}

// LockedNotNil applies the NotNil predicate on the "locked" field.
func LockedNotNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLocked)))
	})
}

// OutcomingEQ applies the EQ predicate on the "outcoming" field.
func OutcomingEQ(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutcoming), v))
	})
}

// OutcomingNEQ applies the NEQ predicate on the "outcoming" field.
func OutcomingNEQ(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutcoming), v))
	})
}

// OutcomingIn applies the In predicate on the "outcoming" field.
func OutcomingIn(vs ...uint64) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutcoming), v...))
	})
}

// OutcomingNotIn applies the NotIn predicate on the "outcoming" field.
func OutcomingNotIn(vs ...uint64) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutcoming), v...))
	})
}

// OutcomingGT applies the GT predicate on the "outcoming" field.
func OutcomingGT(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutcoming), v))
	})
}

// OutcomingGTE applies the GTE predicate on the "outcoming" field.
func OutcomingGTE(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutcoming), v))
	})
}

// OutcomingLT applies the LT predicate on the "outcoming" field.
func OutcomingLT(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutcoming), v))
	})
}

// OutcomingLTE applies the LTE predicate on the "outcoming" field.
func OutcomingLTE(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutcoming), v))
	})
}

// OutcomingIsNil applies the IsNil predicate on the "outcoming" field.
func OutcomingIsNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOutcoming)))
	})
}

// OutcomingNotNil applies the NotNil predicate on the "outcoming" field.
func OutcomingNotNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOutcoming)))
	})
}

// SpendableEQ applies the EQ predicate on the "spendable" field.
func SpendableEQ(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpendable), v))
	})
}

// SpendableNEQ applies the NEQ predicate on the "spendable" field.
func SpendableNEQ(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSpendable), v))
	})
}

// SpendableIn applies the In predicate on the "spendable" field.
func SpendableIn(vs ...uint64) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSpendable), v...))
	})
}

// SpendableNotIn applies the NotIn predicate on the "spendable" field.
func SpendableNotIn(vs ...uint64) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSpendable), v...))
	})
}

// SpendableGT applies the GT predicate on the "spendable" field.
func SpendableGT(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSpendable), v))
	})
}

// SpendableGTE applies the GTE predicate on the "spendable" field.
func SpendableGTE(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSpendable), v))
	})
}

// SpendableLT applies the LT predicate on the "spendable" field.
func SpendableLT(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSpendable), v))
	})
}

// SpendableLTE applies the LTE predicate on the "spendable" field.
func SpendableLTE(v uint64) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSpendable), v))
	})
}

// SpendableIsNil applies the IsNil predicate on the "spendable" field.
func SpendableIsNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSpendable)))
	})
}

// SpendableNotNil applies the NotNil predicate on the "spendable" field.
func SpendableNotNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSpendable)))
	})
}

// PrecisionEQ applies the EQ predicate on the "precision" field.
func PrecisionEQ(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrecision), v))
	})
}

// PrecisionNEQ applies the NEQ predicate on the "precision" field.
func PrecisionNEQ(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrecision), v))
	})
}

// PrecisionIn applies the In predicate on the "precision" field.
func PrecisionIn(vs ...uint32) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrecision), v...))
	})
}

// PrecisionNotIn applies the NotIn predicate on the "precision" field.
func PrecisionNotIn(vs ...uint32) predicate.General {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.General(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrecision), v...))
	})
}

// PrecisionGT applies the GT predicate on the "precision" field.
func PrecisionGT(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrecision), v))
	})
}

// PrecisionGTE applies the GTE predicate on the "precision" field.
func PrecisionGTE(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrecision), v))
	})
}

// PrecisionLT applies the LT predicate on the "precision" field.
func PrecisionLT(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrecision), v))
	})
}

// PrecisionLTE applies the LTE predicate on the "precision" field.
func PrecisionLTE(v uint32) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrecision), v))
	})
}

// PrecisionIsNil applies the IsNil predicate on the "precision" field.
func PrecisionIsNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPrecision)))
	})
}

// PrecisionNotNil applies the NotNil predicate on the "precision" field.
func PrecisionNotNil() predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPrecision)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.General) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.General) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.General) predicate.General {
	return predicate.General(func(s *sql.Selector) {
		p(s.Not())
	})
}

// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/detail"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/general"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/schema"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	detailMixin := schema.Detail{}.Mixin()
	detail.Policy = privacy.NewPolicies(detailMixin[0], schema.Detail{})
	detail.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := detail.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	detailMixinFields0 := detailMixin[0].Fields()
	_ = detailMixinFields0
	detailFields := schema.Detail{}.Fields()
	_ = detailFields
	// detailDescCreatedAt is the schema descriptor for created_at field.
	detailDescCreatedAt := detailMixinFields0[0].Descriptor()
	// detail.DefaultCreatedAt holds the default value on creation for the created_at field.
	detail.DefaultCreatedAt = detailDescCreatedAt.Default.(func() uint32)
	// detailDescUpdatedAt is the schema descriptor for updated_at field.
	detailDescUpdatedAt := detailMixinFields0[1].Descriptor()
	// detail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	detail.DefaultUpdatedAt = detailDescUpdatedAt.Default.(func() uint32)
	// detail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	detail.UpdateDefaultUpdatedAt = detailDescUpdatedAt.UpdateDefault.(func() uint32)
	// detailDescDeletedAt is the schema descriptor for deleted_at field.
	detailDescDeletedAt := detailMixinFields0[2].Descriptor()
	// detail.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	detail.DefaultDeletedAt = detailDescDeletedAt.Default.(func() uint32)
	// detailDescAppID is the schema descriptor for app_id field.
	detailDescAppID := detailFields[1].Descriptor()
	// detail.DefaultAppID holds the default value on creation for the app_id field.
	detail.DefaultAppID = detailDescAppID.Default.(func() uuid.UUID)
	// detailDescUserID is the schema descriptor for user_id field.
	detailDescUserID := detailFields[2].Descriptor()
	// detail.DefaultUserID holds the default value on creation for the user_id field.
	detail.DefaultUserID = detailDescUserID.Default.(func() uuid.UUID)
	// detailDescCoinTypeID is the schema descriptor for coin_type_id field.
	detailDescCoinTypeID := detailFields[3].Descriptor()
	// detail.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	detail.DefaultCoinTypeID = detailDescCoinTypeID.Default.(func() uuid.UUID)
	// detailDescIoType is the schema descriptor for io_type field.
	detailDescIoType := detailFields[4].Descriptor()
	// detail.DefaultIoType holds the default value on creation for the io_type field.
	detail.DefaultIoType = detailDescIoType.Default.(string)
	// detailDescIoSubType is the schema descriptor for io_sub_type field.
	detailDescIoSubType := detailFields[5].Descriptor()
	// detail.DefaultIoSubType holds the default value on creation for the io_sub_type field.
	detail.DefaultIoSubType = detailDescIoSubType.Default.(string)
	// detailDescAmount is the schema descriptor for amount field.
	detailDescAmount := detailFields[6].Descriptor()
	// detail.DefaultAmount holds the default value on creation for the amount field.
	detail.DefaultAmount = detailDescAmount.Default.(uint64)
	// detailDescFromCoinTypeID is the schema descriptor for from_coin_type_id field.
	detailDescFromCoinTypeID := detailFields[7].Descriptor()
	// detail.DefaultFromCoinTypeID holds the default value on creation for the from_coin_type_id field.
	detail.DefaultFromCoinTypeID = detailDescFromCoinTypeID.Default.(func() uuid.UUID)
	// detailDescCoinUsdCurrency is the schema descriptor for coin_usd_currency field.
	detailDescCoinUsdCurrency := detailFields[8].Descriptor()
	// detail.DefaultCoinUsdCurrency holds the default value on creation for the coin_usd_currency field.
	detail.DefaultCoinUsdCurrency = detailDescCoinUsdCurrency.Default.(uint64)
	// detailDescIoExtra is the schema descriptor for io_extra field.
	detailDescIoExtra := detailFields[9].Descriptor()
	// detail.DefaultIoExtra holds the default value on creation for the io_extra field.
	detail.DefaultIoExtra = detailDescIoExtra.Default.(string)
	// detailDescID is the schema descriptor for id field.
	detailDescID := detailFields[0].Descriptor()
	// detail.DefaultID holds the default value on creation for the id field.
	detail.DefaultID = detailDescID.Default.(func() uuid.UUID)
	generalMixin := schema.General{}.Mixin()
	general.Policy = privacy.NewPolicies(generalMixin[0], schema.General{})
	general.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := general.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	generalMixinFields0 := generalMixin[0].Fields()
	_ = generalMixinFields0
	generalFields := schema.General{}.Fields()
	_ = generalFields
	// generalDescCreatedAt is the schema descriptor for created_at field.
	generalDescCreatedAt := generalMixinFields0[0].Descriptor()
	// general.DefaultCreatedAt holds the default value on creation for the created_at field.
	general.DefaultCreatedAt = generalDescCreatedAt.Default.(func() uint32)
	// generalDescUpdatedAt is the schema descriptor for updated_at field.
	generalDescUpdatedAt := generalMixinFields0[1].Descriptor()
	// general.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	general.DefaultUpdatedAt = generalDescUpdatedAt.Default.(func() uint32)
	// general.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	general.UpdateDefaultUpdatedAt = generalDescUpdatedAt.UpdateDefault.(func() uint32)
	// generalDescDeletedAt is the schema descriptor for deleted_at field.
	generalDescDeletedAt := generalMixinFields0[2].Descriptor()
	// general.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	general.DefaultDeletedAt = generalDescDeletedAt.Default.(func() uint32)
	// generalDescAppID is the schema descriptor for app_id field.
	generalDescAppID := generalFields[1].Descriptor()
	// general.DefaultAppID holds the default value on creation for the app_id field.
	general.DefaultAppID = generalDescAppID.Default.(func() uuid.UUID)
	// generalDescUserID is the schema descriptor for user_id field.
	generalDescUserID := generalFields[2].Descriptor()
	// general.DefaultUserID holds the default value on creation for the user_id field.
	general.DefaultUserID = generalDescUserID.Default.(func() uuid.UUID)
	// generalDescCoinTypeID is the schema descriptor for coin_type_id field.
	generalDescCoinTypeID := generalFields[3].Descriptor()
	// general.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	general.DefaultCoinTypeID = generalDescCoinTypeID.Default.(func() uuid.UUID)
	// generalDescIncoming is the schema descriptor for incoming field.
	generalDescIncoming := generalFields[4].Descriptor()
	// general.DefaultIncoming holds the default value on creation for the incoming field.
	general.DefaultIncoming = generalDescIncoming.Default.(uint64)
	// generalDescLocked is the schema descriptor for locked field.
	generalDescLocked := generalFields[5].Descriptor()
	// general.DefaultLocked holds the default value on creation for the locked field.
	general.DefaultLocked = generalDescLocked.Default.(uint64)
	// generalDescOutcoming is the schema descriptor for outcoming field.
	generalDescOutcoming := generalFields[6].Descriptor()
	// general.DefaultOutcoming holds the default value on creation for the outcoming field.
	general.DefaultOutcoming = generalDescOutcoming.Default.(uint64)
	// generalDescSpendable is the schema descriptor for spendable field.
	generalDescSpendable := generalFields[7].Descriptor()
	// general.DefaultSpendable holds the default value on creation for the spendable field.
	general.DefaultSpendable = generalDescSpendable.Default.(uint64)
	// generalDescID is the schema descriptor for id field.
	generalDescID := generalFields[0].Descriptor()
	// general.DefaultID holds the default value on creation for the id field.
	general.DefaultID = generalDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.10.1"                                         // Version of ent codegen.
	Sum     = "h1:dM5h4Zk6yHGIgw4dCqVzGw3nWgpGYJiV4/kyHEF6PFo=" // Sum of ent codegen.
)

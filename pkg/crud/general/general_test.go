package general

/*
import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/ledger-manager/pkg/testinit"
	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/ledgermgr/general"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var entity = ent.General{
	ID:         uuid.New(),
	AppID:      uuid.New(),
	UserID:     uuid.New(),
	CoinTypeID: uuid.New(),
	// Incoming:   10.0,
	// Locked:     10.0,
	// Outcoming:  10.0,
	// Spendable:  10.0,
}

var (
	id         = entity.ID.String()
	appID      = entity.AppID.String()
	userID     = entity.UserID.String()
	coinTypeID = entity.CoinTypeID.String()
	req        = npool.GeneralReq{
		ID:         &id,
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		Incoming:   &entity.Incoming,
		Locked:     &entity.Locked,
		Outcoming:  &entity.Outcoming,
		Spendable:  &entity.Spendable,
	}
)

var info *ent.General

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &entity)
	}
}

func createBulk(t *testing.T) {
	entities := []ent.General{
		{
			ID:         uuid.New(),
			AppID:      uuid.New(),
			UserID:     uuid.New(),
			CoinTypeID: uuid.New(),
			// Incoming:   10.0,
			// Locked:     10.0,
			// Outcoming:  10.0,
			// Spendable:  10.0,
		},
		{
			ID:         uuid.New(),
			AppID:      uuid.New(),
			UserID:     uuid.New(),
			CoinTypeID: uuid.New(),
			// Incoming:   10.0,
			// Locked:     10.0,
			// Outcoming:  10.0,
			// Spendable:  10.0,
		},
	}

	reqs := []*npool.GeneralReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_appID := _entity.AppID.String()
		_userID := _entity.UserID.String()
		_coinTypeID := _entity.CoinTypeID.String()

		reqs = append(reqs, &npool.GeneralReq{
			ID:         &_id,
			AppID:      &_appID,
			UserID:     &_userID,
			CoinTypeID: &_coinTypeID,
			Incoming:   &_entity.Incoming,
			Locked:     &_entity.Locked,
			Outcoming:  &_entity.Outcoming,
			Spendable:  &_entity.Spendable,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.Equal(t, infos, entities)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &entity)
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, infos[0], &entity)
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &entity)
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, 1)
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &entity)
	}
}

func TestGeneral(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("delete", deleteA)
	t.Run("count", count)
}
*/

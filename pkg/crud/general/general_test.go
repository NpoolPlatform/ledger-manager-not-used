package general

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/ledger-manager/pkg/testinit"
	val "github.com/NpoolPlatform/message/npool"
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

var entGeneral = ent.General{
	ID:   uuid.New(),
	Name: uuid.New().String(),
	Age:  10,
}

var (
	id          = entGeneral.ID.String()
	generalInfo = npool.GeneralReq{
		ID:   &id,
		Name: &entGeneral.Name,
		Age:  &entGeneral.Age,
	}
)

var info *ent.General

func rowToObject(row *ent.General) *ent.General {
	return &ent.General{
		ID:   row.ID,
		Name: row.Name,
		Age:  row.Age,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &generalInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entGeneral.ID = info.ID
			entGeneral.CreatedAt = info.CreatedAt
		}
		assert.Equal(t, rowToObject(info), &entGeneral)
	}
}

func createBulk(t *testing.T) {
	entGeneral := []ent.General{
		{
			ID:   uuid.New(),
			Name: uuid.New().String(),
			Age:  10,
		},
		{
			ID:   uuid.New(),
			Name: uuid.New().String(),
			Age:  10,
		},
	}

	generals := []*npool.GeneralReq{}
	for key := range entGeneral {
		id := entGeneral[key].ID.String()
		generals = append(generals, &npool.GeneralReq{
			ID:   &id,
			Name: &entGeneral[key].Name,
			Age:  &entGeneral[key].Age,
		})
	}
	infos, err := CreateBulk(context.Background(), generals)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &generalInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entGeneral)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entGeneral)
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, rowToObject(infos[0]), &entGeneral)
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entGeneral)
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, count)
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entGeneral)
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
	t.Run("update", update)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("delete", deleteA)
	t.Run("count", count)
}

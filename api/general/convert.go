//nolint:nolintlint,dupl
package general

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-manager/pkg/crud/general"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/ledger-manager/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/ledgermgr/general"

	price "github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"
)

func ent2grpc(row *ent.General) *npool.General {
	return &npool.General{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		UserID:     row.UserID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		Incoming:   price.DBPriceToVisualPrice(row.Incoming),
		Locked:     price.DBPriceToVisualPrice(row.Locked),
		Outcoming:  price.DBPriceToVisualPrice(row.Outcoming),
		Spendable:  price.DBPriceToVisualPrice(row.Spendable),
	}
}

func ent2grpcMany(rows []*ent.General) []*npool.General {
	infos := []*npool.General{}
	for _, row := range rows {
		infos = append(infos, ent2grpc(row))
	}
	return infos
}

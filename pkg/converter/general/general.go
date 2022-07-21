//nolint:nolintlint,dupl
package general

import (
	price "github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledgermgr/general"
)

func Ent2Grpc(row *ent.General) *npool.General {
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

func Ent2GrpcMany(rows []*ent.General) []*npool.General {
	infos := []*npool.General{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}

package detail

import (
	price "github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledgermgr/detail"
)

func Ent2Grpc(row *ent.Detail) *npool.Detail {
	return &npool.Detail{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		UserID:          row.UserID.String(),
		CoinTypeID:      row.CoinTypeID.String(),
		IOType:          npool.IOType(npool.IOType_value[row.IoType]),
		IOSubType:       npool.IOSubType(npool.IOSubType_value[row.IoSubType]),
		Amount:          price.DBPriceToVisualPrice(row.Amount),
		FromCoinTypeID:  row.FromCoinTypeID.String(),
		CoinUSDCurrency: price.DBPriceToVisualPrice(row.CoinUsdCurrency),
		IOExtra:         row.IoExtra,
		FromOldID:       row.FromOldID.String(),
	}
}

func Ent2GrpcMany(rows []*ent.Detail) []*npool.Detail {
	infos := []*npool.Detail{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}

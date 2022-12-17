package general

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/general"

	"github.com/google/uuid"
)

func validate(info *npool.GeneralReq) error {
	if info.GoodID == nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GoodID)
		return status.Error(codes.InvalidArgument, "GoodID is empty")
	}

	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GoodID, "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("GoodID is invalid: %v", err))
	}

	if info.CoinTypeID == nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.CoinTypeID)
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.CoinTypeID, "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinTypeID is invalid: %v", err))
	}

	return nil
}

func duplicate(infos []*npool.GeneralReq) error {
	keys := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Infos has invalid element %v", err))
		}

		key := fmt.Sprintf("%v:%v", info.GetGoodID(), info.GetCoinTypeID())
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate GoodID:CoinTypeID")
		}

		keys[key] = struct{}{}
	}

	return nil
}

package detail

import (
	"fmt"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/detail"

	"github.com/google/uuid"
)

func validate(info *npool.DetailReq) error {
	if info.GoodID == nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GoodID)
		return status.Error(codes.InvalidArgument, "GoodID is empty")
	}

	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GetGoodID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("GoodID is invalid: %v", err))
	}

	if info.CoinTypeID == nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.CoinTypeID)
		return status.Error(codes.InvalidArgument, "CoinTypeID is empty")
	}

	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.GetCoinTypeID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinTypeID is invalid: %v", err))
	}

	if info.Amount != nil {
		amount, err := decimal.NewFromString(info.GetAmount())
		if err != nil {
			logger.Sugar().Errorw("validate", "Amount", info.GetAmount(), "error", err)
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Amount is invalid: %v", err))
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("validate", "Amount", info.GetAmount(), "error", "less than 0")
			return status.Error(codes.InvalidArgument, "Amount is less than 0")
		}
	}

	if info.BenefitDate == nil || info.GetBenefitDate() == 0 {
		logger.Sugar().Errorw("validate", "BenefitDate", info.BenefitDate)
		return status.Error(codes.InvalidArgument, "BenefitDate is invalid")
	}

	return nil
}

func duplicate(infos []*npool.DetailReq) error {
	for _, info := range infos {
		if err := validate(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Infos has invalid element %v", err))
		}
	}

	return nil
}

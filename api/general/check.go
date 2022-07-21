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

	"github.com/google/uuid"
)

func validate(info *npool.GeneralReq) error {
	if info.AppID == nil {
		logger.Sugar().Error("AppID is empty")
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if info.UserID == nil {
		logger.Sugar().Error("UserID is empty")
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if info.CoinTypeID == nil {
		logger.Sugar().Error("CoinTypeID is empty")
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if info.Incoming != nil && info.GetIncoming() < 0 {
		logger.Sugar().Error("Incoming is less than 0")
		return status.Error(codes.InvalidArgument, "Incoming is less than 0")
	}

	if info.Outcoming != nil && info.GetOutcoming() < 0 {
		logger.Sugar().Error("Outcoming is less than 0")
		return status.Error(codes.InvalidArgument, "Outcoming is less than 0")
	}

	return nil
}

func duplicate(infos []*npool.GeneralReq) error {
	keys := map[string]struct{}{}
	apps := map[string]struct{}{}

	for _, info := range infos {
		if err := check(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Errorf("Infos has invalid element %v", err))
		}

		key := fmt.Sprintf("%v:%v:%v", info.AppID, info.UserID, info.CoinTypeID)
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate AppID:UserID:CoinTypeID")
		}
	}

	if len(apps) > 1 {
		return status.Error(codes.InvalidArgument, "Infos has different AppID")
	}

	return nil
}
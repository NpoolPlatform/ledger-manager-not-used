package version

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/version"
	npool "github.com/NpoolPlatform/message/npool"
)

func Version() (*npool.VersionResponse, error) {
	info, err := version.GetVersion()
	if err != nil {
		logger.Sugar().Errorw("Version", "error", err)
		return nil, err
	}
	return &npool.VersionResponse{
		Info: info,
	}, nil
}

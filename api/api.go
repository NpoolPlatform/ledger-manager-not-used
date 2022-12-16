package api

import (
	"context"

	ledgermgr "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger"

	"github.com/NpoolPlatform/ledger-manager/api/detail"
	"github.com/NpoolPlatform/ledger-manager/api/general"
	mdetail "github.com/NpoolPlatform/ledger-manager/api/mining/detail"
	mgeneral "github.com/NpoolPlatform/ledger-manager/api/mining/general"
	"github.com/NpoolPlatform/ledger-manager/api/mining/unsold"
	"github.com/NpoolPlatform/ledger-manager/api/profit"
	"github.com/NpoolPlatform/ledger-manager/api/withdraw"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ledgermgr.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	ledgermgr.RegisterManagerServer(server, &Server{})
	general.Register(server)
	detail.Register(server)
	mgeneral.Register(server)
	mdetail.Register(server)
	unsold.Register(server)
	profit.Register(server)
	withdraw.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := ledgermgr.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

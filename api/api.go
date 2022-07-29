package api

import (
	"context"

	ledgermgr "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger"

	"github.com/NpoolPlatform/ledger-manager/api/detail"
	"github.com/NpoolPlatform/ledger-manager/api/general"

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
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := ledgermgr.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := general.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := detail.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

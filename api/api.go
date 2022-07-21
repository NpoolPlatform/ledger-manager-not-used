package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/ledgermgr"

	"github.com/NpoolPlatform/ledger-manager/api/detail"
	"github.com/NpoolPlatform/ledger-manager/api/general"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ledgermgr.UnimplementedLedgerManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	ledgermgr.RegisterLedgerManagerServer(server, &Server{})
	general.Register(server)
	detail.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := ledgermgr.RegisterLedgerManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
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

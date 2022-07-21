package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/ledgermgr"
	"github.com/NpoolPlatform/message/npool/ledgermgr/general"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ledgermgr.UnimplementedLedgerManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	ledgermgr.RegisterLedgerManagerServer(server, &Server{})
	general.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := ledgermgr.RegisterLedgerManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := general.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

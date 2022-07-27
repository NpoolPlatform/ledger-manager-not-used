package detail

import (
	"github.com/NpoolPlatform/message/npool/ledger/mgr/v1/detail"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	detail.UnimplementedLedgerDetailServer
}

func Register(server grpc.ServiceRegistrar) {
	detail.RegisterLedgerDetailServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}

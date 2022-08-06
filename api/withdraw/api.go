package withdraw

import (
	ledgerwithdraw "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/withdraw"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ledgerwithdraw.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	ledgerwithdraw.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}

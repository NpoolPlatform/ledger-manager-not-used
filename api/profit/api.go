package profit

import (
	ledgerprofit "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/profit"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ledgerprofit.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	ledgerprofit.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}

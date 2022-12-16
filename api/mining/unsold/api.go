package unsold

import (
	"github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/unsold"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	unsold.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	unsold.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}

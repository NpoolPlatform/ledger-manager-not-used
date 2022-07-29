package general

import (
	ledgergeneral "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/general"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ledgergeneral.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	ledgergeneral.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}

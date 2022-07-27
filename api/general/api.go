package general

import (
	"github.com/NpoolPlatform/message/npool/ledger/mgr/v1/general"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	general.UnimplementedLedgerGeneralServer
}

func Register(server grpc.ServiceRegistrar) {
	general.RegisterLedgerGeneralServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}

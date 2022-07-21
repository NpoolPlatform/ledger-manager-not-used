//nolint:nolintlint,dupl
package general

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-manager/pkg/crud/general"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/ledger-manager/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/ledgermgr/general"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Server struct {
	general.UnimplementedGeneralServer
}

func Register(server grpc.ServiceRegistrar) {
	general.RegisterGeneralServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return general.RegisterGeneralHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

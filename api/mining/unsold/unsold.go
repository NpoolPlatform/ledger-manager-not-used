//nolint:nolintlint,dupl
package unsold

import (
	"context"

	converter "github.com/NpoolPlatform/ledger-manager/pkg/converter/mining/unsold"
	crud "github.com/NpoolPlatform/ledger-manager/pkg/crud/mining/unsold"
	commontracer "github.com/NpoolPlatform/ledger-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/ledger-manager/pkg/tracer/mining/unsold"

	constant "github.com/NpoolPlatform/ledger-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/unsold"

	"github.com/google/uuid"
)

func (s *Server) CreateUnsold(ctx context.Context, in *npool.CreateUnsoldRequest) (*npool.CreateUnsoldResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateUnsold")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = validate(in.GetInfo())
	if err != nil {
		return &npool.CreateUnsoldResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "unsold", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create unsold: %v", err.Error())
		return &npool.CreateUnsoldResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateUnsoldResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateUnsolds(ctx context.Context, in *npool.CreateUnsoldsRequest) (*npool.CreateUnsoldsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateUnsolds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateUnsoldsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = duplicate(in.GetInfos())
	if err != nil {
		return &npool.CreateUnsoldsResponse{}, err
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "unsold", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create unsolds: %v", err)
		return &npool.CreateUnsoldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateUnsoldsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) GetUnsold(ctx context.Context, in *npool.GetUnsoldRequest) (*npool.GetUnsoldResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUnsold")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetUnsoldResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "unsold", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get unsold: %v", err)
		return &npool.GetUnsoldResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetUnsoldResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetUnsoldOnly(ctx context.Context, in *npool.GetUnsoldOnlyRequest) (*npool.GetUnsoldOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUnsoldOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "unsold", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get unsolds: %v", err)
		return &npool.GetUnsoldOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetUnsoldOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetUnsolds(ctx context.Context, in *npool.GetUnsoldsRequest) (*npool.GetUnsoldsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUnsolds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "unsold", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get unsolds: %v", err)
		return &npool.GetUnsoldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetUnsoldsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistUnsold(ctx context.Context, in *npool.ExistUnsoldRequest) (*npool.ExistUnsoldResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistUnsold")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistUnsoldResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "unsold", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check unsold: %v", err)
		return &npool.ExistUnsoldResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistUnsoldResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistUnsoldConds(ctx context.Context,
	in *npool.ExistUnsoldCondsRequest) (*npool.ExistUnsoldCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistUnsoldConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "unsold", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check unsold: %v", err)
		return &npool.ExistUnsoldCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistUnsoldCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountUnsolds(ctx context.Context, in *npool.CountUnsoldsRequest) (*npool.CountUnsoldsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountUnsolds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "unsold", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count unsolds: %v", err)
		return &npool.CountUnsoldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountUnsoldsResponse{
		Info: total,
	}, nil
}

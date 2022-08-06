//nolint:nolintlint,dupl
package withdraw

import (
	"context"

	converter "github.com/NpoolPlatform/ledger-manager/pkg/converter/withdraw"
	crud "github.com/NpoolPlatform/ledger-manager/pkg/crud/withdraw"
	commontracer "github.com/NpoolPlatform/ledger-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/ledger-manager/pkg/tracer/withdraw"

	constant "github.com/NpoolPlatform/ledger-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/withdraw"

	"github.com/google/uuid"
)

func (s *Server) CreateWithdraw(ctx context.Context, in *npool.CreateWithdrawRequest) (*npool.CreateWithdrawResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateWithdraw")
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
		return &npool.CreateWithdrawResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "withdraw", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateWithdraw", "error", err)
		return &npool.CreateWithdrawResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateWithdrawResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateWithdraws(ctx context.Context, in *npool.CreateWithdrawsRequest) (*npool.CreateWithdrawsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateWithdraws")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateWithdrawsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = duplicate(in.GetInfos())
	if err != nil {
		return &npool.CreateWithdrawsResponse{}, err
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "withdraw", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateWithdraws", "error", err)
		return &npool.CreateWithdrawsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateWithdrawsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) AddWithdraw(ctx context.Context, in *npool.AddWithdrawRequest) (*npool.AddWithdrawResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetWithdraw")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetInfo().GetID())

	_, err = uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return &npool.AddWithdrawResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "withdraw", "crud", "AddFields")

	info, err := crud.AddFields(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("AddWithdraw", "error", err)
		return &npool.AddWithdrawResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.AddWithdrawResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetWithdraw(ctx context.Context, in *npool.GetWithdrawRequest) (*npool.GetWithdrawResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetWithdraw")
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
		return &npool.GetWithdrawResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "withdraw", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetWithdraw", "error", err)
		return &npool.GetWithdrawResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetWithdrawResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetWithdrawOnly(ctx context.Context, in *npool.GetWithdrawOnlyRequest) (*npool.GetWithdrawOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetWithdrawOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "withdraw", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetWithdrawOnly", "error", err)
		return &npool.GetWithdrawOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetWithdrawOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetWithdraws(ctx context.Context, in *npool.GetWithdrawsRequest) (*npool.GetWithdrawsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetWithdraws")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "withdraw", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("fail get withdraws: %v", err)
		return &npool.GetWithdrawsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetWithdrawsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistWithdraw(ctx context.Context, in *npool.ExistWithdrawRequest) (*npool.ExistWithdrawResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistWithdraw")
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
		return &npool.ExistWithdrawResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "withdraw", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetWithdraws", "error", err)
		return &npool.ExistWithdrawResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistWithdrawResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistWithdrawConds(ctx context.Context,
	in *npool.ExistWithdrawCondsRequest) (*npool.ExistWithdrawCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistWithdrawConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "withdraw", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistWithdraw", "error", err)
		return &npool.ExistWithdrawCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistWithdrawCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountWithdraws(ctx context.Context, in *npool.CountWithdrawsRequest) (*npool.CountWithdrawsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountWithdraws")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "withdraw", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountWithdraws", "error", err)
		return &npool.CountWithdrawsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountWithdrawsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteWithdraw(ctx context.Context, in *npool.DeleteWithdrawRequest) (*npool.DeleteWithdrawResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteWithdraw")
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
		return &npool.DeleteWithdrawResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "withdraw", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteWithdraw", "error", err)
		return &npool.DeleteWithdrawResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteWithdrawResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

//nolint:nolintlint,dupl
package profit

import (
	"context"

	converter "github.com/NpoolPlatform/ledger-manager/pkg/converter/profit"
	crud "github.com/NpoolPlatform/ledger-manager/pkg/crud/profit"
	commontracer "github.com/NpoolPlatform/ledger-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/ledger-manager/pkg/tracer/profit"

	constant "github.com/NpoolPlatform/ledger-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/profit"

	"github.com/google/uuid"
)

func (s *Server) CreateProfit(ctx context.Context, in *npool.CreateProfitRequest) (*npool.CreateProfitResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateProfit")
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
		return &npool.CreateProfitResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "profit", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateProfit", "error", err)
		return &npool.CreateProfitResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateProfitResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateProfits(ctx context.Context, in *npool.CreateProfitsRequest) (*npool.CreateProfitsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateProfits")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateProfitsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = duplicate(in.GetInfos())
	if err != nil {
		return &npool.CreateProfitsResponse{}, err
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "profit", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateProfits", "error", err)
		return &npool.CreateProfitsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateProfitsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) AddProfit(ctx context.Context, in *npool.AddProfitRequest) (*npool.AddProfitResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetProfit")
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
		return &npool.AddProfitResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "profit", "crud", "AddFields")

	info, err := crud.AddFields(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("AddProfit", "error", err)
		return &npool.AddProfitResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.AddProfitResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetProfit(ctx context.Context, in *npool.GetProfitRequest) (*npool.GetProfitResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetProfit")
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
		return &npool.GetProfitResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "profit", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetProfit", "error", err)
		return &npool.GetProfitResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetProfitResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetProfitOnly(ctx context.Context, in *npool.GetProfitOnlyRequest) (*npool.GetProfitOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetProfitOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "profit", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetProfitOnly", "error", err)
		return &npool.GetProfitOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetProfitOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetProfits(ctx context.Context, in *npool.GetProfitsRequest) (*npool.GetProfitsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetProfits")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "profit", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetProfits", "error", err)
		return &npool.GetProfitsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetProfitsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistProfit(ctx context.Context, in *npool.ExistProfitRequest) (*npool.ExistProfitResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistProfit")
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
		return &npool.ExistProfitResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "profit", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistProfit", "error", err)
		return &npool.ExistProfitResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistProfitResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistProfitConds(ctx context.Context,
	in *npool.ExistProfitCondsRequest) (*npool.ExistProfitCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistProfitConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "profit", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistProfitConds", "error", err)
		return &npool.ExistProfitCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistProfitCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountProfits(ctx context.Context, in *npool.CountProfitsRequest) (*npool.CountProfitsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountProfits")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "profit", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountProfits", "error", err)
		return &npool.CountProfitsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountProfitsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteProfit(ctx context.Context, in *npool.DeleteProfitRequest) (*npool.DeleteProfitResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteProfit")
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
		return &npool.DeleteProfitResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "profit", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteProfit", "error", err)
		return &npool.DeleteProfitResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteProfitResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

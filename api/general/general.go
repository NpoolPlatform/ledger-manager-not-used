//nolint:nolintlint,dupl
package general

import (
	"context"
	"fmt"

	converter "github.com/NpoolPlatform/ledger-manager/pkg/converter/general"
	crud "github.com/NpoolPlatform/ledger-manager/pkg/crud/general"
	commontracer "github.com/NpoolPlatform/ledger-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/ledger-manager/pkg/tracer/general"

	constant "github.com/NpoolPlatform/ledger-manager/pkg/message/const"

	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/ledgermgr/general"

	"github.com/google/uuid"
)

func (s *GeneralServer) CreateGeneral(ctx context.Context, in *npool.CreateGeneralRequest) (*npool.CreateGeneralResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGeneral")
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
		return &npool.CreateGeneralResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create general: %v", err)
		return &npool.CreateGeneralResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateGeneralResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *GeneralServer) CreateGenerals(ctx context.Context, in *npool.CreateGeneralsRequest) (*npool.CreateGeneralsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGenerals")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateGeneralsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = duplicate(ino.GetInfos())
	if err != nil {
		return &npool.CreateGeneralsResponse{}, err
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create generals: %v", err)
		return &npool.CreateGeneralsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateGeneralsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *GeneralServer) GetGeneral(ctx context.Context, in *npool.GetGeneralRequest) (*npool.GetGeneralResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetGeneral")
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
		return &npool.GetGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get general: %v", err)
		return &npool.GetGeneralResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGeneralResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *GeneralServer) GetGeneralOnly(ctx context.Context, in *npool.GetGeneralOnlyRequest) (*npool.GetGeneralOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetGeneralOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = crud.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get generals: %v", err)
		return &npool.GetGeneralOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGeneralOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *GeneralServer) GetGenerals(ctx context.Context, in *npool.GetGeneralsRequest) (*npool.GetGeneralsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetGenerals")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commintracer.TraceInvoker(span, "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get generals: %v", err)
		return &npool.GetGeneralsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGeneralsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *GeneralServer) ExistGeneral(ctx context.Context, in *npool.ExistGeneralRequest) (*npool.ExistGeneralResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistGeneral")
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
		return &npool.ExistGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commintracer.TraceInvoker(span, "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check general: %v", err)
		return &npool.ExistGeneralResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistGeneralResponse{
		Info: exist,
	}, nil
}

func (s *GeneralServer) ExistGeneralConds(ctx context.Context,
	in *npool.ExistGeneralCondsRequest) (*npool.ExistGeneralCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistGeneralConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commintracer.TraceInvoker(span, "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check general: %v", err)
		return &npool.ExistGeneralCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistGeneralCondsResponse{
		Info: exist,
	}, nil
}

func (s *GeneralServer) CountGenerals(ctx context.Context, in *npool.CountGeneralsRequest) (*npool.CountGeneralsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountGenerals")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commintracer.TraceInvoker(span, "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count generals: %v", err)
		return &npool.CountGeneralsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountGeneralsResponse{
		Info: total,
	}, nil
}

func (s *GeneralServer) DeleteGeneral(ctx context.Context, in *npool.DeleteGeneralRequest) (*npool.DeleteGeneralResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteGeneral")
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
		return &npool.DeleteGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commintracer.TraceInvoker(span, "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete general: %v", err)
		return &npool.DeleteGeneralResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteGeneralResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

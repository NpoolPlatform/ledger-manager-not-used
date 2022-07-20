//nolint:nolintlint,dupl
package api

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
	npool "github.com/NpoolPlatform/message/npool/ledgermgr/general"

	"github.com/google/uuid"
)

func checkGeneralInfo(info *npool.GeneralReq) error {
	if info.Name == nil {
		logger.Sugar().Error("Name is empty")
		return status.Error(codes.InvalidArgument, "Name is empty")
	}

	if info.Age == nil {
		logger.Sugar().Error("Age is empty")
		return status.Error(codes.InvalidArgument, "Age is empty")
	}

	return nil
}

func generalRowToObject(row *ent.General) *npool.General {
	return &npool.General{
		ID:   row.ID.String(),
		Name: row.Name,
		Age:  row.Age,
	}
}

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

	span = crud.GeneralSpanAttributes(span, in.GetInfo())
	err = checkGeneralInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateGeneralResponse{}, err
	}

	span.AddEvent("call crud Create")
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create general: %v", err)
		return &npool.CreateGeneralResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateGeneralResponse{
		Info: generalRowToObject(info),
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
		return &npool.CreateGeneralsResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}

	dup := make(map[string]struct{})
	for key, info := range in.GetInfos() {
		span.SetAttributes(
			attribute.String(fmt.Sprintf("ID.%v", key), info.GetID()),
			attribute.String(fmt.Sprintf("Name.%v", key), info.GetName()),
			attribute.String(fmt.Sprintf("Age.%v", key), info.GetName()),
		)
		err := checkGeneralInfo(info)
		if err != nil {
			return &npool.CreateGeneralsResponse{}, err
		}

		if _, ok := dup[info.GetName()]; ok {
			return &npool.CreateGeneralsResponse{},
				status.Errorf(codes.AlreadyExists,
					"Name: %v duplicate create",
					info.GetName(),
				)
		}

		dup[info.GetName()] = struct{}{}
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create generals: %v", err)
		return &npool.CreateGeneralsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.General, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, generalRowToObject(val))
	}

	return &npool.CreateGeneralsResponse{
		Infos: infos,
	}, nil
}

func (s *GeneralServer) UpdateGeneral(ctx context.Context, in *npool.UpdateGeneralRequest) (*npool.UpdateGeneralResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateGeneral")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = crud.GeneralSpanAttributes(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("general id is invalid")
		return &npool.UpdateGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call crud Update")
	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update general: %v", err)
		return &npool.UpdateGeneralResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateGeneralResponse{
		Info: generalRowToObject(info),
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

	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call crud Row")
	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get general: %v", err)
		return &npool.GetGeneralResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGeneralResponse{
		Info: generalRowToObject(info),
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

	span = crud.GeneralCondsSpanAttributes(span, in.GetConds())

	span.AddEvent("call crud RowOnly")
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get generals: %v", err)
		return &npool.GetGeneralOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGeneralOnlyResponse{
		Info: generalRowToObject(info),
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

	span = crud.GeneralCondsSpanAttributes(span, in.GetConds())
	span.SetAttributes(
		attribute.Int("Offset", int(in.GetOffset())),
		attribute.Int("Limit", int(in.GetLimit())),
	)

	span.AddEvent("call crud Rows")
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get generals: %v", err)
		return &npool.GetGeneralsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.General, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, generalRowToObject(val))
	}

	return &npool.GetGeneralsResponse{
		Infos: infos,
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

	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call crud Exist")
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

	span = crud.GeneralCondsSpanAttributes(span, in.GetConds())

	span.AddEvent("call crud ExistConds")
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
	span = crud.GeneralCondsSpanAttributes(span, in.GetConds())

	span.AddEvent("call crud Count")
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

	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call crud Delete")
	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete general: %v", err)
		return &npool.DeleteGeneralResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteGeneralResponse{
		Info: generalRowToObject(info),
	}, nil
}

package general

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/ledger-manager/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/NpoolPlatform/ledger-manager/pkg/db"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/general"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledgermgr/general"
	"github.com/google/uuid"
)

func Trace(span trace.Span, in *npool.GeneralReq, index int) trace.Span {
	span.SetAttributes(
		attribute.String("ID", in.GetID()),
		attribute.String("AppID", in.GetAppID()),
		attribute.String("UserID", in.GetUserID()),
		attribute.String("CoinTypeID", in.GetCoinTypeID()),
		attribute.Float64("Incoming", in.GetIncoming()),
		attribute.Float64("Locked", in.GetLocked()),
		attribute.Float64("Outcoming", in.GetOutcoming()),
		attribute.Float64("Spendable", in.GetSpendable()),
	)
	return span
}

func TraceConds(span trace.Span, in *npool.Conds) trace.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("UserID.Op", in.GetUserID().GetOp()),
		attribute.String("UserID.Value", in.GetUserID().GetValue()),
		attribute.String("CoinTypeID.Op", in.GetCoinTypeID().GetOp()),
		attribute.String("CoinTypeID.Value", in.GetCoinTypeID().GetValue()),
		attribute.String("Incoming.Op", in.GetIncoming().GetOp()),
		attribute.Float64("Incoming.Value", in.GetIncoming().GetValue()),
		attribute.String("Locked.Op", in.GetLocked().GetOp()),
		attribute.Float64("Locked.Value", in.GetLocked().GetValue()),
		attribute.String("Outcoming.Op", in.GetOutcoming().GetOp()),
		attribute.Float64("Outcoming.Value", in.GetOutcoming().GetValue()),
		attribute.String("Spendable.Op", in.GetSpendable().GetOp()),
		attribute.Float64("Spendable.Value", in.GetSpendable().GetValue()),
	)
	return span
}

func TraceMany(span trace.Span, infos []*npool.GeneralReq) trace.Span {
	for index, info := range infos {
		Trace(span, info, index)
	}
}

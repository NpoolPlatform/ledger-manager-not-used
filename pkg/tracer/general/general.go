package general

import (
	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/ledgermgr/general"
)

func trace(span trace1.Span, in *npool.GeneralReq, index int) trace1.Span {
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

func Trace(span trace1.Span, in *npool.GeneralReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
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

func TraceMany(span trace1.Span, infos []*npool.GeneralReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}

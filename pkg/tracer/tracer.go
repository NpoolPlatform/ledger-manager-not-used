package tracer

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/ledger-manager/pkg/message/const"
	servicename "github.com/NpoolPlatform/ledger-manager/pkg/servicename"

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

func TraceID(span trace.Span, id uuid.UUID) trace.Span {
	span.SetAttributes(attribute.String("ID", id))
	return span
}

func TraceInvoker(span trace.Span, module, invokeName string) trace.Span {
	span.AddEvent(fmt.Sprintf("%v.%v.%v", servicename.ServiceName, module, invokeName))
	return span
}

func TraceOffsetLimit(span trace.Span, offset, limit int) trace.Span {
	span.SetAttributes(
		attribute.Int("Offset", int(in.GetOffset())),
		attribute.Int("Limit", int(in.GetLimit())),
	)
	return span
}

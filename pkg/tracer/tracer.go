package tracer

import (
	"fmt"

	constant "github.com/NpoolPlatform/ledger-manager/pkg/message/const"
	servicename "github.com/NpoolPlatform/ledger-manager/pkg/servicename"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/google/uuid"
)

func TraceID(span trace.Span, id string) trace.Span {
	span.SetAttributes(attribute.String("ID", id))
	return span
}

func TraceInvoker(span trace.Span, module, invokeName string) trace.Span {
	span.AddEvent(fmt.Sprintf("%v.%v.%v", servicename.ServiceName, module, invokeName))
	return span
}

func TraceOffsetLimit(span trace.Span, offset, limit int) trace.Span {
	span.SetAttributes(
		attribute.Int("Offset", offset),
		attribute.Int("Limit", limit),
	)
	return span
}

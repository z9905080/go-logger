package otel

import (
	"context"

	logger "github.com/z9905080/go-logger"
	"go.opentelemetry.io/otel/trace"
)

// Enricher returns a logger.EnricherFunc that extracts the OpenTelemetry
// trace ID and span ID from the context and appends them as log fields.
// Fields are only appended when a valid span is present.
func Enricher() logger.EnricherFunc {
	return func(ctx context.Context) []any {
		sc := trace.SpanFromContext(ctx).SpanContext()
		if !sc.IsValid() {
			return nil
		}
		return []any{
			"trace_id", sc.TraceID().String(),
			"span_id", sc.SpanID().String(),
		}
	}
}

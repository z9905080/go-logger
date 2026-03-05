package logger

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

// OtelEnricher returns an EnricherFunc that extracts the OpenTelemetry
// trace ID and span ID from the context and appends them as log fields.
// Fields are only appended when a valid, sampled span is present.
func OtelEnricher() EnricherFunc {
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

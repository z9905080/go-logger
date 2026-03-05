package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// ZapLogger is a Logger implementation backed by go.uber.org/zap.
type ZapLogger struct {
	s         *zap.SugaredLogger
	enrichers []EnricherFunc
}

// NewZapLogger creates a production ZapLogger.
func NewZapLogger() (*ZapLogger, error) {
	z, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &ZapLogger{s: z.Sugar()}, nil
}

// NewZapLoggerFrom wraps an existing *zap.Logger.
func NewZapLoggerFrom(z *zap.Logger) *ZapLogger {
	return &ZapLogger{s: z.Sugar()}
}

// WithEnricher returns a new ZapLogger with the given EnricherFuncs appended.
func (l *ZapLogger) WithEnricher(fns ...EnricherFunc) *ZapLogger {
	newEnrichers := make([]EnricherFunc, len(l.enrichers)+len(fns))
	copy(newEnrichers, l.enrichers)
	copy(newEnrichers[len(l.enrichers):], fns)
	return &ZapLogger{s: l.s, enrichers: newEnrichers}
}

// enrich runs all registered EnricherFuncs and returns the combined KV fields.
func (l *ZapLogger) enrich(ctx context.Context) []any {
	var fields []any
	for _, fn := range l.enrichers {
		fields = append(fields, fn(ctx)...)
	}
	return fields
}

func (l *ZapLogger) Debug(ctx context.Context, msg string, keysAndValues ...any) {
	l.s.Debugw(msg, append(l.enrich(ctx), keysAndValues...)...)
}

func (l *ZapLogger) DebugF(ctx context.Context, msgFormat string, params ...any) {
	l.s.Debugw(fmt.Sprintf(msgFormat, params...), l.enrich(ctx)...)
}

func (l *ZapLogger) Info(ctx context.Context, msg string, keysAndValues ...any) {
	l.s.Infow(msg, append(l.enrich(ctx), keysAndValues...)...)
}

func (l *ZapLogger) InfoF(ctx context.Context, msgFormat string, params ...any) {
	l.s.Infow(fmt.Sprintf(msgFormat, params...), l.enrich(ctx)...)
}

func (l *ZapLogger) Warn(ctx context.Context, msg string, keysAndValues ...any) {
	l.s.Warnw(msg, append(l.enrich(ctx), keysAndValues...)...)
}

func (l *ZapLogger) WarnF(ctx context.Context, msgFormat string, params ...any) {
	l.s.Warnw(fmt.Sprintf(msgFormat, params...), l.enrich(ctx)...)
}

func (l *ZapLogger) Error(ctx context.Context, msg string, keysAndValues ...any) {
	l.s.Errorw(msg, append(l.enrich(ctx), keysAndValues...)...)
}

func (l *ZapLogger) ErrorF(ctx context.Context, msgFormat string, params ...any) {
	l.s.Errorw(fmt.Sprintf(msgFormat, params...), l.enrich(ctx)...)
}

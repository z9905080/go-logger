package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// ZapLogger is a Logger implementation backed by go.uber.org/zap.
type ZapLogger struct {
	z *zap.Logger
}

// NewZapLogger creates a production ZapLogger.
func NewZapLogger() (*ZapLogger, error) {
	z, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &ZapLogger{z: z}, nil
}

// NewZapLoggerFrom wraps an existing *zap.Logger.
func NewZapLoggerFrom(z *zap.Logger) *ZapLogger {
	return &ZapLogger{z: z}
}

// toZapFields converts keysAndValues pairs to []zap.Field.
// Pairs must be even; if odd an extra error field is appended.
func toZapFields(keysAndValues []any) []zap.Field {
	if len(keysAndValues)%2 != 0 {
		keysAndValues = append(keysAndValues, "(MISSING)")
	}
	fields := make([]zap.Field, 0, len(keysAndValues)/2)
	for i := 0; i < len(keysAndValues); i += 2 {
		key, ok := keysAndValues[i].(string)
		if !ok {
			key = fmt.Sprintf("%v", keysAndValues[i])
		}
		fields = append(fields, zap.Any(key, keysAndValues[i+1]))
	}
	return fields
}

func (l *ZapLogger) Debug(ctx context.Context, msg string, keysAndValues ...any) {
	l.z.Debug(msg, toZapFields(keysAndValues)...)
}

func (l *ZapLogger) DebugF(ctx context.Context, msgFormat string, params ...any) {
	l.z.Debug(fmt.Sprintf(msgFormat, params...))
}

func (l *ZapLogger) Info(ctx context.Context, msg string, keysAndValues ...any) {
	l.z.Info(msg, toZapFields(keysAndValues)...)
}

func (l *ZapLogger) InfoF(ctx context.Context, msgFormat string, params ...any) {
	l.z.Info(fmt.Sprintf(msgFormat, params...))
}

func (l *ZapLogger) Warn(ctx context.Context, msg string, keysAndValues ...any) {
	l.z.Warn(msg, toZapFields(keysAndValues)...)
}

func (l *ZapLogger) WarnF(ctx context.Context, msgFormat string, params ...any) {
	l.z.Warn(fmt.Sprintf(msgFormat, params...))
}

func (l *ZapLogger) Error(ctx context.Context, msg string, keysAndValues ...any) {
	l.z.Error(msg, toZapFields(keysAndValues)...)
}

func (l *ZapLogger) ErrorF(ctx context.Context, msgFormat string, params ...any) {
	l.z.Error(fmt.Sprintf(msgFormat, params...))
}

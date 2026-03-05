package logger

import (
	"context"

	"go.uber.org/zap"
)

// ZapLogger is a Logger implementation backed by go.uber.org/zap.
type ZapLogger struct {
	s *zap.SugaredLogger
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

func (l *ZapLogger) Debug(ctx context.Context, msg string, keysAndValues ...any) {
	l.s.Debugw(msg, keysAndValues...)
}

func (l *ZapLogger) DebugF(ctx context.Context, msgFormat string, params ...any) {
	l.s.Debugf(msgFormat, params...)
}

func (l *ZapLogger) Info(ctx context.Context, msg string, keysAndValues ...any) {
	l.s.Infow(msg, keysAndValues...)
}

func (l *ZapLogger) InfoF(ctx context.Context, msgFormat string, params ...any) {
	l.s.Infof(msgFormat, params...)
}

func (l *ZapLogger) Warn(ctx context.Context, msg string, keysAndValues ...any) {
	l.s.Warnw(msg, keysAndValues...)
}

func (l *ZapLogger) WarnF(ctx context.Context, msgFormat string, params ...any) {
	l.s.Warnf(msgFormat, params...)
}

func (l *ZapLogger) Error(ctx context.Context, msg string, keysAndValues ...any) {
	l.s.Errorw(msg, keysAndValues...)
}

func (l *ZapLogger) ErrorF(ctx context.Context, msgFormat string, params ...any) {
	l.s.Errorf(msgFormat, params...)
}

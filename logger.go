package logger

import "context"

// Logger defines the logging interface.
// A second PR will provide a zap-based implementation.
type Logger interface {
	Debug(ctx context.Context, message string, values ...string)
	DebugF(ctx context.Context, format string, params ...string)

	Info(ctx context.Context, message string, values ...string)
	InfoF(ctx context.Context, format string, params ...string)

	Warn(ctx context.Context, message string, values ...string)
	WarnF(ctx context.Context, format string, params ...string)

	Error(ctx context.Context, message string, values ...string)
	ErrorF(ctx context.Context, format string, params ...string)
}

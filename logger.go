package logger

import "context"

// EnricherFunc extracts additional log fields from a context.
// The returned slice must be even-length KV pairs, e.g. ("key1", val1, "key2", val2).
type EnricherFunc func(ctx context.Context) []any

// LogLevel represents the severity of a log message.
type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

// Logger defines the logging interface.
// A second PR will provide a zap-based implementation.
//
// For structured methods (Debug, Info, Warn, Error), keysAndValues must be
// provided in even-length KV pairs, e.g. ("key1", val1, "key2", val2).
type Logger interface {
	Debug(ctx context.Context, msg string, keysAndValues ...any)
	DebugF(ctx context.Context, msgFormat string, params ...any)

	Info(ctx context.Context, msg string, keysAndValues ...any)
	InfoF(ctx context.Context, msgFormat string, params ...any)

	Warn(ctx context.Context, msg string, keysAndValues ...any)
	WarnF(ctx context.Context, msgFormat string, params ...any)

	Error(ctx context.Context, msg string, keysAndValues ...any)
	ErrorF(ctx context.Context, msgFormat string, params ...any)
}

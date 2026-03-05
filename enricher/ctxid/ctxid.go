package ctxid

import (
	"context"

	logger "github.com/z9905080/go-logger"
)

// key is the unexported context key for the context ID.
type key struct{}

// NewContext stores a custom ID in the context for later extraction by Enricher.
func NewContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, key{}, id)
}

// Enricher returns a logger.EnricherFunc that extracts the custom context ID
// set by NewContext and appends it as a "ctx_id" log field.
// The field is omitted when no ID is present in the context.
func Enricher() logger.EnricherFunc {
	return func(ctx context.Context) []any {
		id, ok := ctx.Value(key{}).(string)
		if !ok || id == "" {
			return nil
		}
		return []any{"ctx_id", id}
	}
}

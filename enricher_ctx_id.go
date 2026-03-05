package logger

import "context"

// ctxIDKey is the unexported key used to store a context ID in a context.
type ctxIDKey struct{}

// ContextWithID stores a custom ID in the context for later extraction by CtxIDEnricher.
func ContextWithID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ctxIDKey{}, id)
}

// CtxIDEnricher returns an EnricherFunc that extracts the custom context ID
// set by ContextWithID and appends it as a "ctx_id" log field.
// The field is omitted when no ID is present in the context.
func CtxIDEnricher() EnricherFunc {
	return func(ctx context.Context) []any {
		id, ok := ctx.Value(ctxIDKey{}).(string)
		if !ok || id == "" {
			return nil
		}
		return []any{"ctx_id", id}
	}
}

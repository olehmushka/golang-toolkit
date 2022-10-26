package contexttools

import "context"

func SetValue(ctx context.Context, key, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

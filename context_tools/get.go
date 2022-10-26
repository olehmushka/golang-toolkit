package contexttools

import "context"

func GetValue(ctx context.Context, key string) string {
	if value, ok := ctx.Value(key).(string); ok {
		return value
	}

	return ""
}

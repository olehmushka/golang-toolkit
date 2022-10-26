package ptrtools

func PtrsToValues[T any](in []*T) []T {
	out := make([]T, len(in))
	for i := range out {
		out[i] = PtrToValue(in[i])
	}

	return out
}

func PtrToValue[T any](in *T) T {
	return *in
}

func ValuesToPtrs[T any](in []T) []*T {
	out := make([]*T, len(in))
	for i := range out {
		out[i] = ValueToPtr(in[i])
	}

	return out
}

func ValueToPtr[T any](in T) *T {
	return &in
}

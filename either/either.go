package either

type Either[T any] struct {
	Value T
	Err   error
}

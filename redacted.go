package redacted

const redactedValue = "****"

type Redacted[T any] struct {
	value T
}

func (r Redacted[T]) Value() T {
	return r.value
}

func (Redacted[T]) String() string {
	return redactedValue
}

func (Redacted[T]) GoString() string {
	return redactedValue
}

func New[T any](value T) Redacted[T] {
	return Redacted[T]{value: value}
}

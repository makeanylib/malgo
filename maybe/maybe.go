package maybe

import "errors"

var (
	ErrCannotUnwrapNone = errors.New("cannot unwrap none")
)

type T[V any] struct {
	value *V
}

func (t T[V]) IsSome() bool {
	return t.value != nil
}

func (t T[V]) IsNone() bool {
	return t.value == nil
}

func (t T[V]) UnsafeUnwrap() V {
	if t.value == nil {
		panic(ErrCannotUnwrapNone)
	}
	return *t.value
}

func (t T[V]) TryUnwrap() (V, bool) {
	if t.value == nil {
		var zero V
		return zero, false
	}
	return *t.value, true
}

func (t T[V]) ToPointer() *V {
	return t.value
}

func (t T[V]) UnwrapOr(v V) V {
	if t.value != nil {
		return *t.value
	}
	return v
}

func Some[V any](v V) T[V] {
	return T[V]{
		value: &v,
	}
}

func None[V any]() T[V] {
	return T[V]{
		value: nil,
	}
}

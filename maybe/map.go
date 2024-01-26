package maybe

func Map[U any, V any](t T[U], f func(U) V) T[V] {
	if t.value == nil {
		return None[V]()
	}
	return Some(f(*t.value))
}

func TryMap[U any, V any](t T[U], f func(U) (V, error)) (T[V], error) {
	if t.value == nil {
		return None[V](), nil
	}

	v, err := f(*t.value)
	if err != nil {
		return None[V](), err
	}

	return Some(v), nil
}

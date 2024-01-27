package slicex

func Map[U any, V any](s []U, f func(U) V) []V {
	if s == nil {
		return nil
	}

	r := make([]V, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func TryMap[U any, V any](s []U, f func(U) (V, error)) ([]V, error) {
	if s == nil {
		return nil, nil
	}

	r := make([]V, len(s))
	for i, v := range s {
		v, err := f(v)
		if err != nil {
			return nil, err
		}
		r[i] = v
	}
	return r, nil
}

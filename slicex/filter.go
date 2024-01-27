package slicex

func Filter[T any](s []T, f func(T) bool) []T {
	if s == nil {
		return nil
	}

	r := make([]T, 0)
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

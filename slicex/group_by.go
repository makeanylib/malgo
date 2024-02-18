package slicex

func GroupBy[T any, K comparable](s []T, f func(T) K) map[K][]T {
	groups := make(map[K][]T)
	for _, v := range s {
		k := f(v)
		groups[k] = append(groups[k], v)
	}
	return groups
}

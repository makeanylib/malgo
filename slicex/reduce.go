package slicex

func Reduce[T any](s []T, f func(acc T, cur T) T, initial T) T {
	if len(s) == 0 {
		return initial
	}

	acc := initial
	for _, cur := range s {
		acc = f(acc, cur)
	}

	return acc
}

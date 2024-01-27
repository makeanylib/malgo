package union2

type T[T1, T2 any] struct {
	tag int8
	obj any
}

func NewT1[T1, T2 any](t1 T1) T[T1, T2] {
	return T[T1, T2]{
		tag: 1,
		obj: t1,
	}
}

func NewT2[T1, T2 any](t2 T2) T[T1, T2] {
	return T[T1, T2]{
		tag: 2,
		obj: t2,
	}
}

func (u T[T1, T2]) IsT1() bool {
	return u.tag == 1
}

func (u T[T1, T2]) IsT2() bool {
	return u.tag == 2
}

func (u T[T1, T2]) TryGetT1() (T1, bool) {
	if u.tag == 1 {
		return u.obj.(T1), true
	}
	var zero T1
	return zero, false
}

func (u T[T1, T2]) TryGetT2() (T2, bool) {
	if u.tag == 2 {
		return u.obj.(T2), true
	}
	var zero T2
	return zero, false
}

func Match[T1, T2, TReturn any](
	u T[T1, T2],
	t1 func(T1) TReturn,
	t2 func(T2) TReturn,
) TReturn {
	switch u.tag {
	case 1:
		return t1(u.obj.(T1))
	case 2:
		return t2(u.obj.(T2))
	}

	panic("union is empty")
}

func TryMatch[T1, T2, TReturn any](
	u T[T1, T2],
	t1 func(T1) (TReturn, error),
	t2 func(T2) (TReturn, error),
) (TReturn, error) {
	switch u.tag {
	case 1:
		return t1(u.obj.(T1))
	case 2:
		return t2(u.obj.(T2))
	}

	panic("union is empty")
}

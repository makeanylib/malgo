package union4

type T[T1, T2, T3, T4 any] struct {
	tag int8
	obj any
}

func NewT1[T1, T2, T3, T4 any](t1 T1) T[T1, T2, T3, T4] {
	return T[T1, T2, T3, T4]{
		tag: 1,
		obj: t1,
	}
}

func NewT2[T1, T2, T3, T4 any](t2 T2) T[T1, T2, T3, T4] {
	return T[T1, T2, T3, T4]{
		tag: 2,
		obj: t2,
	}
}

func NewT3[T1, T2, T3, T4 any](t3 T3) T[T1, T2, T3, T4] {
	return T[T1, T2, T3, T4]{
		tag: 3,
		obj: t3,
	}
}

func NewT4[T1, T2, T3, T4 any](t4 T4) T[T1, T2, T3, T4] {
	return T[T1, T2, T3, T4]{
		tag: 4,
		obj: t4,
	}
}

func (u T[T1, T2, T3, T4]) IsT1() bool {
	return u.tag == 1
}

func (u T[T1, T2, T3, T4]) IsT2() bool {
	return u.tag == 2
}

func (u T[T1, T2, T3, T4]) IsT3() bool {
	return u.tag == 3
}

func (u T[T1, T2, T3, T4]) IsT4() bool {
	return u.tag == 4
}

func (u T[T1, T2, T3, T4]) TryGetT1() (T1, bool) {
	if u.tag == 1 {
		return u.obj.(T1), true
	}
	var zero T1
	return zero, false
}

func (u T[T1, T2, T3, T4]) TryGetT2() (T2, bool) {
	if u.tag == 2 {
		return u.obj.(T2), true
	}
	var zero T2
	return zero, false
}

func (u T[T1, T2, T3, T4]) TryGetT3() (T3, bool) {
	if u.tag == 3 {
		return u.obj.(T3), true
	}
	var zero T3
	return zero, false
}

func (u T[T1, T2, T3, T4]) TryGetT4() (T4, bool) {
	if u.tag == 4 {
		return u.obj.(T4), true
	}
	var zero T4
	return zero, false
}

func Match[T1, T2, T3, T4, TReturn any](
	u T[T1, T2, T3, T4],
	t1 func(T1) TReturn,
	t2 func(T2) TReturn,
	t3 func(T3) TReturn,
	t4 func(T4) TReturn,
) TReturn {
	switch u.tag {
	case 1:
		return t1(u.obj.(T1))
	case 2:
		return t2(u.obj.(T2))
	case 3:
		return t3(u.obj.(T3))
	case 4:
		return t4(u.obj.(T4))
	}

	panic("union is empty")
}

func TryMatch[T1, T2, T3, T4, TReturn any](
	u T[T1, T2, T3, T4],
	t1 func(T1) (TReturn, error),
	t2 func(T2) (TReturn, error),
	t3 func(T3) (TReturn, error),
	t4 func(T4) (TReturn, error),
) (TReturn, error) {
	switch u.tag {
	case 1:
		return t1(u.obj.(T1))
	case 2:
		return t2(u.obj.(T2))
	case 3:
		return t3(u.obj.(T3))
	case 4:
		return t4(u.obj.(T4))
	}

	panic("union is empty")
}

package maybe

func FromPointer[V any](v *V) T[V] {
	if v == nil {
		return None[V]()
	}
	return Some(*v)
}

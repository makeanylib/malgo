package channel

type Chan[T any] struct {
	closed bool
	ch     chan T
}

func NewChan[T any](capacity int) *Chan[T] {
	return &Chan[T]{
		ch: make(chan T, capacity),
	}
}

func (c *Chan[T]) Send(v T) bool {
	if c.closed {
		return false
	}
	c.ch <- v
	return true
}

func (c *Chan[T]) TrySend(v T) bool {
	if c.closed {
		return false
	}

	select {
	case c.ch <- v:
		return true
	default:
		return false
	}
}

func (c *Chan[T]) Receive() (T, bool) {
	v, ok := <-c.ch
	return v, ok
}

func (c *Chan[T]) TryReceive() (T, bool) {
	if c.closed {
		var zero T
		return zero, false
	}

	select {
	case v, ok := <-c.ch:
		if !ok {
			return v, false
		}
		return v, true
	default:
		var zero T
		return zero, false
	}
}

func (c *Chan[T]) Close() {
	if c.closed {
		return
	}
	c.closed = true
	close(c.ch)
}

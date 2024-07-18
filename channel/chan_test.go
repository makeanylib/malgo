package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChan_SendAndReceive(t *testing.T) {
	ch := NewChan[int](1)
	sent := ch.Send(42)
	assert.True(t, sent)

	v, ok := ch.Receive()
	assert.True(t, ok)
	assert.Equal(t, 42, v)

	ch.Close()

	v, ok = ch.Receive()
	assert.False(t, ok)
	assert.Zero(t, v)

	sent = ch.Send(42)
	assert.False(t, sent)
}

func TestChan_TrySendAndReceive(t *testing.T) {
	ch := NewChan[int](0)

	sent := ch.TrySend(42)
	assert.False(t, sent)

	v, ok := ch.TryReceive()
	assert.False(t, ok)
	assert.Zero(t, v)

	ch.Close()

	sent = ch.TrySend(42)
	assert.False(t, sent)

	v, ok = ch.TryReceive()
	assert.False(t, ok)
	assert.Zero(t, v)
}

func TestChan_TrySendAndReceive_WhenDataIsAvailable(t *testing.T) {
	ch := NewChan[int](1)

	sent := ch.TrySend(42)
	assert.True(t, sent)

	v, ok := ch.TryReceive()
	assert.True(t, ok)
	assert.Equal(t, 42, v)

	ch.Close()

	sent = ch.TrySend(42)
	assert.False(t, sent)

	v, ok = ch.TryReceive()
	assert.False(t, ok)
	assert.Zero(t, v)
}

func TestChan_DoubleClose(t *testing.T) {
	assert.NotPanics(t, func() {
		ch := NewChan[int](1)
		ch.Close()
		ch.Close()
	})
}

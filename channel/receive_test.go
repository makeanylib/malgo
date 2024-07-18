package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryReceive(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 42
	v, ok := TryReceive(ch)
	assert.True(t, ok)
	assert.Equal(t, 42, v)

	v, ok = TryReceive(ch)
	assert.False(t, ok)
	assert.Zero(t, v)
}

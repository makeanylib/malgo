package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		f    func(int) bool
		want []int
	}{
		{
			name: "should return nil slice, when filter on nil slice",
			s:    nil,
			f:    func(i int) bool { return true },
			want: nil,
		},
		{
			name: "should return empty slice, when filter on empty slice",
			s:    []int{},
			f:    func(i int) bool { return true },
			want: []int{},
		},
		{
			name: "should return filtered slice",
			s:    []int{1, 2, 3, 4, 5},
			f:    func(i int) bool { return i%2 == 0 },
			want: []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.s, tt.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

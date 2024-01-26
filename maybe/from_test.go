package maybe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaybe_FromPointer(t *testing.T) {
	tests := []struct {
		name    string
		pointer *int
		want    T[int]
	}{
		{
			name: "should map non-nil pointer correctly",
			pointer: func() *int {
				v := 123
				return &v
			}(),
			want: Some[int](123),
		},
		{
			name:    "should map nil pointer correctly",
			pointer: nil,
			want:    None[int](),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromPointer(tt.pointer)
			assert.Equal(t, tt.want, got)
		})
	}
}

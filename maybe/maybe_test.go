package maybe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaybe(t *testing.T) {
	tests := []struct {
		name   string
		maybe  T[int]
		assert func(t *testing.T, got T[int])
	}{
		{
			name:  "should handle some value correctly",
			maybe: Some[int](42),
			assert: func(t *testing.T, got T[int]) {
				assert.True(t, got.IsSome())
				assert.False(t, got.IsNone())

				v, ok := got.TryUnwrap()
				assert.True(t, ok)
				assert.Equal(t, 42, v)

				assert.NotPanics(t, func() {
					v := got.UnsafeUnwrap()
					assert.Equal(t, 42, v)
				})

				assert.Equal(t, 42, *got.ToPointer())
			},
		},
		{
			name:  "should handle none value correctly",
			maybe: None[int](),
			assert: func(t *testing.T, got T[int]) {
				assert.False(t, got.IsSome())
				assert.True(t, got.IsNone())

				v, ok := got.TryUnwrap()
				assert.False(t, ok)
				assert.Zero(t, v)

				assert.Panics(t, func() {
					got.UnsafeUnwrap()
				})

				assert.Nil(t, got.ToPointer())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assert(t, tt.maybe)
		})
	}
}

func TestMaybe_UnwrapOr(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		maybe  T[int]
		args   args
		assert func(t *testing.T, got int)
	}{
		{
			name:  "should unwrap some value correctly",
			maybe: Some[int](42),
			args: args{
				v: 10,
			},
			assert: func(t *testing.T, got int) {
				assert.Equal(t, 42, got)
			},
		},
		{
			name:  "should unwrap none value correctly",
			maybe: None[int](),
			args: args{
				v: 10,
			},
			assert: func(t *testing.T, got int) {
				assert.Equal(t, 10, got)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assert(t, tt.maybe.UnwrapOr(tt.args.v))
		})
	}
}

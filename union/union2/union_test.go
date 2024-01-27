package union2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion2(t *testing.T) {
	tests := []struct {
		name   string
		union  T[int, string]
		assert func(t *testing.T, v T[int, string])
	}{
		{
			name:  "T1",
			union: NewT1[int, string](42),
			assert: func(t *testing.T, v T[int, string]) {
				assert.True(t, v.IsT1())
				assert.False(t, v.IsT2())

				t1, ok := v.TryGetT1()
				assert.True(t, ok)
				assert.Equal(t, 42, t1)

				t2, ok := v.TryGetT2()
				assert.False(t, ok)
				assert.Zero(t, t2)

				r := Match(
					v,
					func(t1 int) int {
						return 1
					},
					func(t2 string) int {
						return 2
					},
				)
				assert.Equal(t, 1, r)

				r, err := TryMatch(
					v,
					func(t1 int) (int, error) {
						return 1, nil
					},
					func(t2 string) (int, error) {
						return 2, nil
					},
				)
				assert.NoError(t, err)
				assert.Equal(t, 1, r)

				r, err = TryMatch(
					v,
					func(t1 int) (int, error) {
						var zero int
						return zero, assert.AnError
					},
					func(t2 string) (int, error) {
						return 2, nil
					},
				)
				assert.Error(t, err)
				assert.Zero(t, r)
			},
		},
		{
			name:  "T2",
			union: NewT2[int, string]("HELLO"),
			assert: func(t *testing.T, v T[int, string]) {
				assert.False(t, v.IsT1())
				assert.True(t, v.IsT2())

				t1, ok := v.TryGetT1()
				assert.False(t, ok)
				assert.Zero(t, t1)

				t2, ok := v.TryGetT2()
				assert.True(t, ok)
				assert.Equal(t, "HELLO", t2)

				r := Match(
					v,
					func(t1 int) int {
						return 1
					},
					func(t2 string) int {
						return 2
					},
				)
				assert.Equal(t, 2, r)

				r, err := TryMatch(
					v,
					func(t1 int) (int, error) {
						return 1, nil
					},
					func(t2 string) (int, error) {
						return 2, nil
					},
				)
				assert.NoError(t, err)
				assert.Equal(t, 2, r)

				r, err = TryMatch(
					v,
					func(t1 int) (int, error) {
						return 1, nil
					},
					func(t2 string) (int, error) {
						var zero int
						return zero, assert.AnError
					},
				)
				assert.Error(t, err)
				assert.Zero(t, r)
			},
		},
		{
			name:  "empty",
			union: T[int, string]{},
			assert: func(t *testing.T, v T[int, string]) {
				assert.False(t, v.IsT1())
				assert.False(t, v.IsT2())

				t1, ok := v.TryGetT1()
				assert.False(t, ok)
				assert.Zero(t, t1)

				t2, ok := v.TryGetT2()
				assert.False(t, ok)
				assert.Zero(t, t2)

				assert.Panics(t, func() {
					Match(
						v,
						func(t1 int) int {
							return 1
						},
						func(t2 string) int {
							return 2
						},
					)
				})

				assert.Panics(t, func() {
					TryMatch(
						v,
						func(t1 int) (int, error) {
							return 1, nil
						},
						func(t2 string) (int, error) {
							return 2, nil
						},
					)
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assert(t, tt.union)
		})
	}
}

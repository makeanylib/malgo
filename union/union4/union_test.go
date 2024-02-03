package union4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion2(t *testing.T) {
	tests := []struct {
		name   string
		union  T[int, string, bool, float64]
		assert func(t *testing.T, v T[int, string, bool, float64])
	}{
		{
			name:  "T1",
			union: NewT1[int, string, bool, float64](42),
			assert: func(t *testing.T, v T[int, string, bool, float64]) {
				assert.True(t, v.IsT1())
				assert.False(t, v.IsT2())
				assert.False(t, v.IsT3())
				assert.False(t, v.IsT4())

				t1, ok := v.TryGetT1()
				assert.True(t, ok)
				assert.Equal(t, 42, t1)

				t2, ok := v.TryGetT2()
				assert.False(t, ok)
				assert.Zero(t, t2)

				t3, ok := v.TryGetT3()
				assert.False(t, ok)
				assert.Zero(t, t3)

				t4, ok := v.TryGetT4()
				assert.False(t, ok)
				assert.Zero(t, t4)

				r := Match(
					v,
					func(t1 int) int {
						return 1
					},
					func(t2 string) int {
						return 2
					},
					func(t3 bool) int {
						return 3
					},
					func(t4 float64) int {
						return 4
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
					func(t3 bool) (int, error) {
						return 3, nil
					},
					func(t4 float64) (int, error) {
						return 4, nil
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
					func(t3 bool) (int, error) {
						return 3, nil
					},
					func(t4 float64) (int, error) {
						return 4, nil
					},
				)
				assert.Error(t, err)
				assert.Zero(t, r)
			},
		},
		{
			name:  "T2",
			union: NewT2[int, string, bool, float64]("HELLO"),
			assert: func(t *testing.T, v T[int, string, bool, float64]) {
				assert.False(t, v.IsT1())
				assert.True(t, v.IsT2())
				assert.False(t, v.IsT3())
				assert.False(t, v.IsT4())

				t1, ok := v.TryGetT1()
				assert.False(t, ok)
				assert.Zero(t, t1)

				t2, ok := v.TryGetT2()
				assert.True(t, ok)
				assert.Equal(t, "HELLO", t2)

				t3, ok := v.TryGetT3()
				assert.False(t, ok)
				assert.Zero(t, t3)

				t4, ok := v.TryGetT4()
				assert.False(t, ok)
				assert.Zero(t, t4)

				r := Match(
					v,
					func(t1 int) int {
						return 1
					},
					func(t2 string) int {
						return 2
					},
					func(t3 bool) int {
						return 3
					},
					func(t4 float64) int {
						return 4
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
					func(t3 bool) (int, error) {
						return 3, nil
					},
					func(t4 float64) (int, error) {
						return 4, nil
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
					func(t3 bool) (int, error) {
						return 3, nil
					},
					func(t4 float64) (int, error) {
						return 4, nil
					},
				)
				assert.Error(t, err)
				assert.Zero(t, r)
			},
		},
		{
			name:  "T3",
			union: NewT3[int, string, bool, float64](true),
			assert: func(t *testing.T, v T[int, string, bool, float64]) {
				assert.False(t, v.IsT1())
				assert.False(t, v.IsT2())
				assert.True(t, v.IsT3())
				assert.False(t, v.IsT4())

				t1, ok := v.TryGetT1()
				assert.False(t, ok)
				assert.Zero(t, t1)

				t2, ok := v.TryGetT2()
				assert.False(t, ok)
				assert.Zero(t, t2)

				t3, ok := v.TryGetT3()
				assert.True(t, ok)
				assert.Equal(t, true, t3)

				t4, ok := v.TryGetT4()
				assert.False(t, ok)
				assert.Zero(t, t4)

				r := Match(
					v,
					func(t1 int) int {
						return 1
					},
					func(t2 string) int {
						return 2
					},
					func(t3 bool) int {
						return 3
					},
					func(t4 float64) int {
						return 4
					},
				)
				assert.Equal(t, 3, r)

				r, err := TryMatch(
					v,
					func(t1 int) (int, error) {
						return 1, nil
					},
					func(t2 string) (int, error) {
						return 2, nil
					},
					func(t3 bool) (int, error) {
						return 3, nil
					},
					func(t4 float64) (int, error) {
						return 4, nil
					},
				)
				assert.NoError(t, err)
				assert.Equal(t, 3, r)

				r, err = TryMatch(
					v,
					func(t1 int) (int, error) {
						return 1, nil
					},
					func(t2 string) (int, error) {
						return 2, nil
					},
					func(t3 bool) (int, error) {
						var zero int
						return zero, assert.AnError
					},
					func(f float64) (int, error) {
						return 4, nil
					},
				)
				assert.Error(t, err)
				assert.Zero(t, r)
			},
		},
		{
			name:  "T4",
			union: NewT4[int, string, bool, float64](12.34),
			assert: func(t *testing.T, v T[int, string, bool, float64]) {
				assert.False(t, v.IsT1())
				assert.False(t, v.IsT2())
				assert.False(t, v.IsT3())
				assert.True(t, v.IsT4())

				t1, ok := v.TryGetT1()
				assert.False(t, ok)
				assert.Zero(t, t1)

				t2, ok := v.TryGetT2()
				assert.False(t, ok)
				assert.Zero(t, t2)

				t3, ok := v.TryGetT3()
				assert.False(t, ok)
				assert.Zero(t, t3)

				t4, ok := v.TryGetT4()
				assert.True(t, ok)
				assert.Equal(t, 12.34, t4)

				r := Match(
					v,
					func(t1 int) int {
						return 1
					},
					func(t2 string) int {
						return 2
					},
					func(t3 bool) int {
						return 3
					},
					func(t4 float64) int {
						return 4
					},
				)
				assert.Equal(t, 4, r)

				r, err := TryMatch(
					v,
					func(t1 int) (int, error) {
						return 1, nil
					},
					func(t2 string) (int, error) {
						return 2, nil
					},
					func(t3 bool) (int, error) {
						return 3, nil
					},
					func(t4 float64) (int, error) {
						return 4, nil
					},
				)
				assert.NoError(t, err)
				assert.Equal(t, 4, r)

				r, err = TryMatch(
					v,
					func(t1 int) (int, error) {
						return 1, nil
					},
					func(t2 string) (int, error) {
						return 2, nil
					},
					func(t3 bool) (int, error) {
						return 3, assert.AnError
					},
					func(f float64) (int, error) {
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
			union: T[int, string, bool, float64]{},
			assert: func(t *testing.T, v T[int, string, bool, float64]) {
				assert.False(t, v.IsT1())
				assert.False(t, v.IsT2())
				assert.False(t, v.IsT3())
				assert.False(t, v.IsT4())

				t1, ok := v.TryGetT1()
				assert.False(t, ok)
				assert.Zero(t, t1)

				t2, ok := v.TryGetT2()
				assert.False(t, ok)
				assert.Zero(t, t2)

				t3, ok := v.TryGetT3()
				assert.False(t, ok)
				assert.Zero(t, t3)

				t4, ok := v.TryGetT4()
				assert.False(t, ok)
				assert.Zero(t, t4)

				assert.Panics(t, func() {
					Match(
						v,
						func(t1 int) int {
							return 1
						},
						func(t2 string) int {
							return 2
						},
						func(t3 bool) int {
							return 3
						},
						func(t4 float64) int {
							return 4
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
						func(t3 bool) (int, error) {
							return 3, nil
						},
						func(t4 float64) (int, error) {
							return 4, nil
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

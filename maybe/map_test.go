package maybe

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaybe_Map(t *testing.T) {
	tests := []struct {
		name  string
		maybe T[int]
		want  T[string]
	}{
		{
			name:  "should map some value correctly",
			maybe: Some[int](123),
			want:  Some[string]("123"),
		},
		{
			name:  "should map none value correctly",
			maybe: None[int](),
			want:  None[string](),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(
				tt.maybe,
				func(v int) string {
					return strconv.FormatInt(int64(v), 10)
				},
			)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMaybe_TryMap(t *testing.T) {
	tests := []struct {
		name   string
		maybe  T[string]
		assert func(t *testing.T, got T[int64], err error)
	}{
		{
			name:  "should map some value correctly with no error",
			maybe: Some("123"),
			assert: func(t *testing.T, got T[int64], err error) {
				assert.NoError(t, err)

				value, ok := got.TryUnwrap()
				assert.True(t, ok)
				assert.Equal(t, int64(123), value)
			},
		},
		{
			name:  "should map none value correctly with an error",
			maybe: Some("xyz"),
			assert: func(t *testing.T, got T[int64], err error) {
				assert.Error(t, err)
				assert.True(t, got.IsNone())
			},
		},
		{
			name:  "should map none value correctly with no error",
			maybe: None[string](),
			assert: func(t *testing.T, got T[int64], err error) {
				assert.NoError(t, err)
				assert.True(t, got.IsNone())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TryMap(
				tt.maybe,
				func(v string) (int64, error) {
					return strconv.ParseInt(v, 10, 64)
				},
			)

			tt.assert(t, got, err)
		})
	}
}

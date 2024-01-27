package slicex

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		f    func(int) string
		want []string
	}{
		{
			name: "should return nil slice, when map on nil slice",
			s:    nil,
			f:    func(i int) string { return strconv.FormatInt(int64(i), 10) },
			want: nil,
		},
		{
			name: "should return empty slice, when map on empty slice",
			s:    []int{},
			f:    func(i int) string { return strconv.FormatInt(int64(i), 10) },
			want: []string{},
		},
		{
			name: "should return mapped slice",
			s:    []int{1, 2, 3, 4, 5},
			f:    func(i int) string { return strconv.FormatInt(int64(i), 10) },
			want: []string{"1", "2", "3", "4", "5"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.s, tt.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTryMap(t *testing.T) {
	tests := []struct {
		name   string
		s      []string
		f      func(string) (int64, error)
		assert func(t *testing.T, got []int64, err error)
	}{
		{
			name: "should return nil slice, when map on nil slice",
			s:    nil,
			f:    func(s string) (int64, error) { return strconv.ParseInt(s, 10, 64) },
			assert: func(t *testing.T, got []int64, err error) {
				assert.Nil(t, got)
				assert.Nil(t, err)
			},
		},
		{
			name: "should return empty slice, when map on empty slice",
			s:    []string{},
			f:    func(s string) (int64, error) { return strconv.ParseInt(s, 10, 64) },
			assert: func(t *testing.T, got []int64, err error) {
				assert.Empty(t, got)
				assert.Nil(t, err)
			},
		},
		{
			name: "should return mapped slice",
			s:    []string{"1", "2", "3", "4", "5"},
			f:    func(s string) (int64, error) { return strconv.ParseInt(s, 10, 64) },
			assert: func(t *testing.T, got []int64, err error) {
				assert.Equal(t, []int64{1, 2, 3, 4, 5}, got)
				assert.Nil(t, err)
			},
		},
		{
			name: "should return error, when map on slice with invalid element",
			s:    []string{"1", "2", "3", "4", "5", "a"},
			f:    func(s string) (int64, error) { return strconv.ParseInt(s, 10, 64) },
			assert: func(t *testing.T, got []int64, err error) {
				assert.Nil(t, got)
				assert.Error(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TryMap(tt.s, tt.f)
			tt.assert(t, got, err)
		})
	}
}

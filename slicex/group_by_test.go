package slicex

import (
	"reflect"
	"testing"
)

func TestGroupBy(t *testing.T) {
	type args struct {
		s []int
		f func(int) string
	}

	tests := []struct {
		name string
		args args
		want map[string][]int
	}{
		{
			name: "empty slice",
			args: args{
				s: []int{},
				f: func(i int) string {
					if i%2 == 0 {
						return "even"
					}
					return "odd"
				},
			},
			want: map[string][]int{},
		},
		{
			name: "single element",
			args: args{
				s: []int{1},
				f: func(i int) string {
					if i%2 == 0 {
						return "even"
					}
					return "odd"
				},
			},
			want: map[string][]int{
				"odd": {1},
			},
		},
		{
			name: "multiple elements",
			args: args{
				s: []int{1, 2, 3, 4, 5},
				f: func(i int) string {
					if i%2 == 0 {
						return "even"
					}
					return "odd"
				},
			},
			want: map[string][]int{
				"odd":  {1, 3, 5},
				"even": {2, 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

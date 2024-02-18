package slicex

import "testing"

func TestReduce(t *testing.T) {
	type args struct {
		s       []int
		f       func(acc int, cur int) int
		initial int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty slice",
			args: args{
				s:       []int{},
				f:       func(acc int, cur int) int { return acc + cur },
				initial: 123,
			},
			want: 123,
		},
		{
			name: "single element",
			args: args{
				s:       []int{1},
				f:       func(acc int, cur int) int { return acc + cur },
				initial: 0,
			},
			want: 1,
		},
		{
			name: "multiple elements",
			args: args{
				s:       []int{1, 2, 3, 4, 5},
				f:       func(acc int, cur int) int { return acc + cur },
				initial: 0,
			},
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.s, tt.args.f, tt.args.initial); got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

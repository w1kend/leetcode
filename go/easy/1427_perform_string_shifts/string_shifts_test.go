package performstringshifts

import "testing"

func Test_stringShift(t *testing.T) {
	type args struct {
		s     string
		shift [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "e1",
			args: args{
				s: "abc",
				shift: [][]int{
					{0, 1},
					{1, 2},
				},
			},
			want: "cab",
		},
		{
			name: "e2",
			args: args{
				s: "abcdefg",
				shift: [][]int{
					{1, 1},
					{1, 1},
					{0, 2},
					{1, 3},
				},
			},
			want: "efgabcd",
		},
		{
			name: "e3",
			args: args{
				s: "wpdhhcj",
				shift: [][]int{
					{0, 7},
					{1, 7},
					{1, 0},
					{1, 3},
					{0, 3},
					{0, 6},
					{1, 2},
				},
			},
			want: "hcjwpdh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringShift(tt.args.s, tt.args.shift); got != tt.want {
				t.Errorf("stringShift() = %v, want %v", got, tt.want)
			}
		})
	}
}

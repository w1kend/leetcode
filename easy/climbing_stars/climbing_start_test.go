package climbingstars

import (
	"leetcode/testsupport"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2 => 2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "3 => 3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "1 => 1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "4 => 5",
			args: args{
				n: 4,
			},
			want: 5,
		},
		{
			name: "30 => 100000",
			args: args{
				n: 30,
			},
			want: 1346269,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClimbStairs(tt.args.n)
			testsupport.AssertEqual(t, got, tt.want, "unexpected count")
		})
	}
}

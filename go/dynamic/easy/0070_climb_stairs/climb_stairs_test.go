package climbstairs

import (
	"leetcode/go/test"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "e1",
			n:    1,
			want: 1,
		},
		{
			name: "e2",
			n:    2,
			want: 2,
		},
		{
			name: "e3",
			n:    3,
			want: 3,
		},
		{
			name: "e4",
			n:    4,
			want: 5,
		},
		{
			name: "e5",
			n:    5,
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

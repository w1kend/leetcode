package trappingrainwater

import (
	"leetcode/go/test"
	"testing"
)

func Test_trap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "e1",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "e2",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			name:   "e3",
			height: []int{4, 2, 3},
			want:   1,
		},
		{
			name:   "m4",
			height: []int{4, 3, 2, 0, 1, 2, 3, 0, 0, 2},
			want:   11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trap(tt.height)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

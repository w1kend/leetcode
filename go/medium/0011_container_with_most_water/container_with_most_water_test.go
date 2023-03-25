package containerwithmostwater

import (
	"leetcode/go/test"
	"testing"
)

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "e1",
			heights: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:    49,
		},
		{
			name:    "e2",
			heights: []int{1, 1},
			want:    1,
		},
		{
			name:    "m3",
			heights: []int{1, 0, 1, 1, 0, 1, 0, 1, 10},
			want:    8,
		},
		{
			name:    "e4",
			heights: []int{2, 3, 4, 5, 18, 17, 6},
			want:    17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.heights)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

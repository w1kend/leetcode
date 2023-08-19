package subsets

import (
	"leetcode/go/test"
	"testing"
)

func Test_subsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "e1",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			name: "e2",
			nums: []int{1, 2, 3},
			want: [][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.nums)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

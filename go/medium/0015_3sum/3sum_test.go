package threesum

import (
	"leetcode/go/test"
	"sort"
	"testing"
)

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want [][]int
	}{
		{
			name: "e1",
			in:   []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "e2",
			in:   []int{-2, 0, 1, 1, 2},
			want: [][]int{
				{-2, 0, 2},
				{-2, 1, 1},
			},
		},
		{
			name: "",
			in:   []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.in)
			for i := range tt.want {
				sort.Ints(tt.want[i])
				sort.Ints(got[i])
			}
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

package medianoftwosortedarrays

import (
	"leetcode/go/test"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "e1",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2,
		},
		{
			name:  "e2",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			name:  "m3",
			nums1: []int{1, 2, 2, 3, 4},
			nums2: []int{2, 3, 4, 5, 6},
			want:  3,
		},
		{
			name:  "e4",
			nums1: []int{},
			nums2: []int{1},
			want:  1,
		},
		{
			name:  "m5",
			nums1: []int{1, 2, 3},
			nums2: []int{},
			want:  2,
		},
		{
			name:  "m6",
			nums1: []int{1, 2, 3, 4},
			nums2: []int{},
			want:  2.5,
		},
		{
			name:  "e7",
			nums1: []int{1},
			nums2: []int{2, 3, 4},
			want:  2.5,
		},
		{
			name:  "e8",
			nums1: []int{4},
			nums2: []int{1, 2, 3},
			want:  2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

package twosum2

import (
	"leetcode/go/test"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "e1",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "e2",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "e3",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
		{
			name:    "e4",
			numbers: []int{1, 3, 4, 5, 6, 8, 9, 10, 11, 12, 13},
			target:  8,
			want:    []int{2, 4},
		},
		{
			name:    "e5",
			numbers: []int{4, 4},
			target:  8,
			want:    []int{1, 2},
		},
		{
			name:    "e6",
			numbers: []int{-100, -50, -10, 10, 20, 25, 200},
			target:  -25,
			want:    []int{2, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.numbers, tt.target)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

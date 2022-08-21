package binarysearch

import (
	"leetcode/testsupport"
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "-1,0,3,5,9,12  9",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			name: "[-1,0,3,5,9,12] 2",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			want: -1,
		},
		{
			name: "[-1,0,3,5,9,12] 13",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 13,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.args.nums, tt.args.target)
			testsupport.AssertEqual(t, got, tt.want, "check slices")
		})
	}
}

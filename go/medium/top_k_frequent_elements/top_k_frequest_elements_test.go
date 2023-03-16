package topkfrequentelements

import (
	"leetcode/go/test"
	"sort"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "e1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "e2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "e3",
			args: args{
				nums: []int{1, 1, 1, 1, 1, 1, 3, 3, 3, 3, 3, 2, 2, 2, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6},
				k:    3,
			},
			want: []int{1, 6, 3},
		},
		{
			name: "e4",
			args: args{
				nums: []int{1, 2},
				k:    2,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.args.nums, tt.args.k)
			sort.Ints(got)
			sort.Ints(tt.want)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

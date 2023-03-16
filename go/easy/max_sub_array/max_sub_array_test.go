package max_sub_array

import (
	"fmt"
	"leetcode/go/test"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "-2, 1, -3, 4, -1, 2, 1, -5, 4 => 6",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "-2, 1, -3 => 1",
			args: args{
				nums: []int{-2, 1, -3},
			},
			want: 1,
		},
		{
			name: "-3, -4, 5, -1, 2, -4, 6, -1 => 8",
			args: args{
				nums: []int{-3, -4, 5, -1, 2, -4, 6, -1},
			},
			want: 8,
		},
		{
			name: "-2, 3, -1, 2 => 4",
			args: args{
				nums: []int{-2, 3, -1, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("divide_and_conquer__%s", tt.name), func(t *testing.T) {
			got := maxSubArray(tt.args.nums)
			test.AssertEqual(t, got, tt.want, "wrong sum")
		})

		t.Run(fmt.Sprintf("letcoode_solution__%s", tt.name), func(t *testing.T) {
			got := maxSubArray2(tt.args.nums)
			test.AssertEqual(t, got, tt.want, "wrong sum")
		})

		t.Run(fmt.Sprintf("kadane's_algorithm__%s", tt.name), func(t *testing.T) {
			got := maxSubArray3(tt.args.nums)
			test.AssertEqual(t, got, tt.want, "wrong sum")
		})
	}
}

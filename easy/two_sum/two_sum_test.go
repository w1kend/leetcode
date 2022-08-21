package two_sum_test

import (
	"leetcode/easy/two_sum"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[2,7,11,15]",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "[3,2,4]",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "[0,4,3,0]",
			args: args{
				nums:   []int{0, 4, 3, 0},
				target: 0,
			},
			want: []int{0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := two_sum.TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

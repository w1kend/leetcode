package wigglesort

import (
	"errors"
	"fmt"
	"testing"
)

func Test_wiggleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "e1",
			args: args{
				nums: []int{3, 5, 2, 1, 6, 4},
			},
		},
		{
			name: "e2",
			args: args{
				nums: []int{6, 6, 5, 6, 3, 8},
			},
		},
		{
			name: "e3",
			args: args{
				nums: []int{8, 8, 4, 5, 1, 4, 3, 5, 6},
			},
		},
		{
			name: "m4",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
		},
		{
			name: "m5",
			args: args{
				nums: []int{},
			},
		},
		{
			name: "m6",
			args: args{
				nums: []int{1, 2},
			},
		},
		{
			name: "m7",
			args: args{
				nums: []int{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSort(tt.args.nums)
			if err := checkResult(tt.args.nums); err != nil {
				t.Fatal(err, tt.args.nums)
			}
		})
	}
}

// nums[0] <= nums[1] >= nums[2] <= nums[3]
func checkResult(arr []int) error {
	if len(arr) <= 1 {
		return nil
	}
	for i := 0; i < len(arr)-1; i++ {
		if i%2 == 0 && arr[i] > arr[i+1] {
			return errors.New(fmt.Sprintf(
				"nums[%d]=%d should be <= than nums[%d]=%d",
				i,
				arr[i],
				i+1,
				arr[i+1],
			))
		}

		if i%2 != 0 && arr[i] < arr[i+1] {
			return errors.New(fmt.Sprintf(
				"nums[%d]=%d should be <= than nums[%d]=%d",
				i,
				arr[i],
				i+1,
				arr[i+1],
			))
		}
	}

	return nil
}

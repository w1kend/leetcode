package majorityelement

import (
	"leetcode/go/testsupport"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "7, 7, 5, 7, 5, 1 | 5, 7 | 5, 5, 7, 7 | 7, 7, 7, 7 => 7",
			args: args{
				nums: []int{7, 7, 5, 7, 5, 1, 5, 7, 5, 5, 7, 7, 7, 7, 7, 7},
			},
			want: 7,
		},
		{
			name: "7, 7, 5, 7, 5, 1 | 5, 7 | 5, 5, 7, 7 | 5,5,5,5 => 5",
			args: args{
				nums: []int{7, 7, 5, 7, 5, 1, 5, 7, 5, 5, 7, 7, 5, 5, 5, 5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.args.nums)
			testsupport.AssertEqual(t, got, tt.want, "unexpected num")
		})
	}
}

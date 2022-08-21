package reverselinkedlist

import (
	"leetcode/pkg/list"
	"leetcode/testsupport"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *list.ListNode
	}
	tests := []struct {
		name string
		args args
		want *list.ListNode
	}{
		{
			name: "1,2,3,4,5 => 5,4,3,2,1",
			args: args{
				head: list.FromSlice([]int{1, 2, 3, 4, 5}),
			},
			want: list.FromSlice([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "1,2 => 2,1",
			args: args{
				head: list.FromSlice([]int{1, 2}),
			},
			want: list.FromSlice([]int{2, 1}),
		},
		{
			name: "1,2 => 2,1",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.args.head)
			testsupport.AssertEqual(t, got.ToSlice(), tt.want.ToSlice(), "unexpedted list")
		})
	}
}

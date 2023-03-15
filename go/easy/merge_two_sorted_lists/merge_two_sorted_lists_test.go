package mergetwosortedlists

import (
	"leetcode/go/testsupport"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,4]  [1,3,4]",
			args: args{
				list1: []int{1, 2, 4},
				list2: []int{1, 3, 4},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := sliceToList(tt.args.list1)
			list2 := sliceToList(tt.args.list2)

			got := MergeTwoLists(&list1, &list2)
			testsupport.AssertEqual(t, got == nil, false, "check result != nil")

			gotSize := 0
			allValues := make([]int, 0, len(tt.want))
			for {
				allValues = append(allValues, got.Val)
				testsupport.AssertEqual(t, got.Val, tt.want[gotSize], "check value")
				gotSize++

				if got.Next == nil {
					break
				}
				got = got.Next
			}
			t.Log(allValues)
			testsupport.AssertEqual(t, gotSize, len(tt.want), "check size")
		})
	}
}

func sliceToList(s []int) ListNode {
	root := ListNode{}

	var curNode *ListNode
	for i := range s {
		if curNode == nil {
			root.Val = s[i]
			curNode = &root
			continue
		}

		curNode.Next = &ListNode{
			Val: s[i],
		}
		curNode = curNode.Next
	}

	return root
}

package invertbinarytree

import (
	"fmt"
	"leetcode/go/pkg/tree"
	"leetcode/go/test"
	"testing"
)

func Test_recusrive(t *testing.T) {
	tests := []struct {
		name string
		root *tree.Node
		want *tree.Node
	}{
		{
			name: "e1",
			root: tree.FromSlice(tree.ToPointersSlice([]int{4, 2, 7, 1, 3, 6, 9})),
			want: tree.FromSlice(tree.ToPointersSlice([]int{4, 7, 2, 9, 6, 3, 1})),
		},
		{
			name: "e2",
			root: tree.FromSlice(tree.ToPointersSlice([]int{2, 1, 3})),
			want: tree.FromSlice(tree.ToPointersSlice([]int{2, 3, 1})),
		},
		{
			name: "e3",
			root: nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run("recursive_"+tt.name, func(t *testing.T) {
			got := recusrive(tt.root)
			if got == nil || tt.want == nil {
				test.AssertEqual(t, got, tt.want, "")
			} else {
				test.AssertEqual(t, got.ToSlice(), tt.want.ToSlice(), "")
			}
		})
		t.Run("iterative_"+tt.name, func(t *testing.T) {
			got := iterative(tt.root)
			if got == nil || tt.want == nil {
				test.AssertEqual(t, got, tt.want, "")
			} else {
				test.AssertEqual(t, got.ToSlice(), tt.want.ToSlice(), fmt.Sprintf("want-%s, have-%s", ptrSliceToStr(tt.want.ToSlice()), ptrSliceToStr(got.ToSlice())))
			}
		})
	}
}

func ptrSliceToStr(arr []*int) string {
	s := "["
	for _, v := range arr {
		if v == nil {
			s += "nil,"
		} else {
			s += fmt.Sprintf("%d,", *v)
		}
	}
	return s[:len(s)-1] + "]"
}

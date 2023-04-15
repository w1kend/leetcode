package subtreeofanothertree

import (
	"leetcode/go/pkg/tree"
	"leetcode/go/test"
	"testing"

	"github.com/samber/lo"
)

func Test_isSubtree(t *testing.T) {
	tests := []struct {
		name    string
		root    *tree.Node
		subRoot *tree.Node
		want    bool
	}{
		{
			name: "e1",
			root: tree.FromSlice(tree.ToPointersSlice([]int{
				3,
				4, 5,
				1, 2,
			})),
			subRoot: tree.FromSlice(tree.ToPointersSlice([]int{
				4,
				1, 2,
			})),
			want: true,
		},
		{
			name: "e2",
			root: tree.FromSlice([]*int{
				lo.ToPtr(3),
				lo.ToPtr(4), lo.ToPtr(5),
				lo.ToPtr(1), lo.ToPtr(2), nil, nil,
				nil, nil, lo.ToPtr(0), nil,
			}),
			subRoot: tree.FromSlice(tree.ToPointersSlice([]int{4, 1, 2})),
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run("recursion"+tt.name, func(t *testing.T) {
			got := isSubtreeRecursive(tt.root, tt.subRoot)
			test.AssertEqual(t, got, tt.want, "")
		})
		t.Run("iteration_"+tt.name, func(t *testing.T) {
			got := isSubtreeIterative(tt.root, tt.subRoot)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

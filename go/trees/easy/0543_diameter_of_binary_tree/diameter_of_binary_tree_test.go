package diameterofbinarytree

import (
	"leetcode/go/pkg/tree"
	"leetcode/go/test"
	"testing"

	"github.com/samber/lo"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		root *tree.Node
		want int
	}{
		{
			name: "e1",
			root: tree.FromSlice(tree.ToPointersSlice([]int{1, 2, 3, 4, 5})),
			want: 3,
		},
		{
			name: "e2",
			root: tree.FromSlice(tree.ToPointersSlice([]int{1, 2})),
			want: 1,
		},
		{
			name: "m3",
			root: tree.FromSlice([]*int{
				lo.ToPtr(1),
				lo.ToPtr(2), lo.ToPtr(3),
				lo.ToPtr(4), lo.ToPtr(5), nil, lo.ToPtr(6),
				lo.ToPtr(7), nil, lo.ToPtr(8), nil, nil, nil, nil, lo.ToPtr(9),
				lo.ToPtr(10), nil, nil, nil, lo.ToPtr(11), lo.ToPtr(12), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil, nil, nil, nil, lo.ToPtr(13),
			}),
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diameterOfBinaryTree(tt.root)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

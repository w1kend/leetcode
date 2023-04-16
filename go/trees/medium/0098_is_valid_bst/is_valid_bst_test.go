package isvalidbst

import (
	"leetcode/go/pkg/tree"
	"leetcode/go/test"
	"testing"

	"github.com/samber/lo"
)

func Test_isValidBST(t *testing.T) {
	tests := []struct {
		name string
		root *tree.Node
		want bool
	}{
		{
			name: "e1",
			root: tree.FromSlice(tree.ToPointersSlice([]int{2, 1, 3})),
			want: true,
		},
		{
			name: "e2",
			root: tree.FromSlice([]*int{
				lo.ToPtr(5),
				lo.ToPtr(1), lo.ToPtr(4),
				nil, nil, lo.ToPtr(3), lo.ToPtr(6),
			}),
			want: false,
		},
		{
			name: "e3",
			root: tree.FromSlice([]*int{
				lo.ToPtr(5),
				lo.ToPtr(4), lo.ToPtr(6),
				nil, nil, lo.ToPtr(3), lo.ToPtr(7),
			}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidBST(tt.root)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

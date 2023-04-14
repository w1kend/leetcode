package maxdepthofbinarytree

import (
	"leetcode/go/pkg/tree"
	"leetcode/go/test"
	"testing"

	"github.com/samber/lo"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *tree.Node
		want int
	}{
		{
			name: "e1",
			root: tree.FromSlice([]*int{lo.ToPtr(3), lo.ToPtr(9), lo.ToPtr(20), nil, nil, lo.ToPtr(15), lo.ToPtr(7)}),
			want: 3,
		},
		{
			name: "e2",
			root: tree.FromSlice([]*int{lo.ToPtr(1), nil, lo.ToPtr(2)}),
			want: 2,
		},
		{
			name: "m3",
			root: nil,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := recursion(tt.root)
			test.AssertEqual(t, got, tt.want, "")
		})
		t.Run(tt.name, func(t *testing.T) {
			got := iteration(tt.root)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

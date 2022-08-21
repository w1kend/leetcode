package lowestcommonancestor

import (
	"leetcode/pkg/sugar"
	"leetcode/pkg/tree"
	"leetcode/testsupport"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name       string
		tree       *tree.Node
		p, q, want func(root *tree.Node) *tree.Node
	}{
		{
			name: "root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8 ==> 6",
			tree: func() *tree.Node {
				t := tree.FromSlice([]*int{sugar.P(6), sugar.P(2), sugar.P(8), sugar.P(0), sugar.P(4), sugar.P(7), sugar.P(9), nil, nil, sugar.P(3), sugar.P(5)})
				return &t
			}(),
			p: func(root *tree.Node) *tree.Node {
				return root.Left
			},
			q: func(root *tree.Node) *tree.Node {
				return root.Right
			},
			want: func(root *tree.Node) *tree.Node {
				return root
			},
		},
		{
			name: "root = [6,2,8,0,4,7,9,null,null,3,5], p = 0, q = 8 ==> 6",
			tree: func() *tree.Node {
				t := tree.FromSlice([]*int{sugar.P(6), sugar.P(2), sugar.P(8), sugar.P(0), sugar.P(4), sugar.P(7), sugar.P(9), nil, nil, sugar.P(3), sugar.P(5)})
				return &t
			}(),
			p: func(root *tree.Node) *tree.Node {
				return root.Left.Left
			},
			q: func(root *tree.Node) *tree.Node {
				return root.Right
			},
			want: func(root *tree.Node) *tree.Node {
				return root
			},
		},
		{
			name: "[2,1], p = 2, q = 1; => 2",
			tree: func() *tree.Node {
				t := tree.FromSlice(tree.ToPointersSlice([]int{2, 1}))
				return &t
			}(),
			p: func(root *tree.Node) *tree.Node {
				return root
			},
			q: func(root *tree.Node) *tree.Node {
				return root.Left
			},
			want: func(root *tree.Node) *tree.Node {
				return root
			},
		},
		{
			name: "root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4 ==> 2",
			tree: func() *tree.Node {
				t := tree.FromSlice([]*int{sugar.P(6), sugar.P(2), sugar.P(8), sugar.P(0), sugar.P(4), sugar.P(7), sugar.P(9), nil, nil, sugar.P(3), sugar.P(5)})
				return &t
			}(),
			p: func(root *tree.Node) *tree.Node {
				return root.Left
			},
			q: func(root *tree.Node) *tree.Node {
				return root.Left.Right
			},
			want: func(root *tree.Node) *tree.Node {
				return root.Left
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, q, want := tt.p(tt.tree), tt.q(tt.tree), tt.want(tt.tree)
			got := lowestCommonAncestor2(tt.tree, p, q)
			testsupport.AssertEqual(t, got, want, "unexpected tree")
		})
	}
}

package balancedbinarytree

import (
	"leetcode/go/pkg/tree"
	"leetcode/go/test"
	"testing"

	"github.com/samber/lo"
)

func Test_isBalanced(t *testing.T) {
	tests := []struct {
		name string
		root *tree.Node
		want bool
	}{
		{
			name: "e1",
			root: tree.FromSlice([]*int{
				lo.ToPtr(3),
				lo.ToPtr(9), lo.ToPtr(20),
				nil, nil, lo.ToPtr(15), lo.ToPtr(7),
			}),
			want: true,
		},
		{
			name: "m2",
			root: tree.FromSlice([]*int{
				lo.ToPtr(3),
				lo.ToPtr(9), lo.ToPtr(20),
				nil, nil, lo.ToPtr(15), lo.ToPtr(7),
				nil, nil, nil, nil, nil, nil, lo.ToPtr(10), lo.ToPtr(20),
			}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBalanced(tt.root)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

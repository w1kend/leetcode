package levelordertraversal

import (
	"leetcode/go/pkg/tree"
	"leetcode/go/test"
	"testing"

	"github.com/samber/lo"
)

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *tree.Node
		want [][]int
	}{
		{
			name: "e1",
			root: tree.FromSlice([]*int{
				lo.ToPtr(3),
				lo.ToPtr(9), lo.ToPtr(20),
				nil, nil, lo.ToPtr(15), lo.ToPtr(7),
			}),
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(tt.root)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

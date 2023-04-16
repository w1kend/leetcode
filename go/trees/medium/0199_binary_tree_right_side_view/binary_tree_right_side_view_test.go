package binarytreerightsideview

import (
	"leetcode/go/pkg/tree"
	"leetcode/go/test"
	"testing"

	"github.com/samber/lo"
)

func Test_rightSideView(t *testing.T) {
	tests := []struct {
		name string
		root *tree.Node
		want []int
	}{
		{
			name: "e1",
			root: tree.FromSlice([]*int{
				lo.ToPtr(1),
				lo.ToPtr(2), lo.ToPtr(3),
				nil, lo.ToPtr(5), nil, lo.ToPtr(4),
				nil, nil, nil, lo.ToPtr(8),
			}),
			want: []int{1, 3, 4, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rightSideView(tt.root)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

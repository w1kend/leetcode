package balancedbinarytree

import (
	"leetcode/go/pkg/sugar"
	"leetcode/go/pkg/tree"
	"leetcode/go/test"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *tree.Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "bananced",
			args: args{
				root: tree.FromSlice([]*int{sugar.P(3), sugar.P(9), sugar.P(20), nil, nil, sugar.P(15), sugar.P(7)}),
			},
			want: true,
		},
		{
			name: "not bananced",
			args: args{
				root: tree.FromSlice([]*int{sugar.P(1), sugar.P(2), sugar.P(2), sugar.P(3), sugar.P(3), nil, nil, sugar.P(4), sugar.P(4)}),
			},
			want: false,
		},
		{
			name: "bananced",
			args: args{
				root: &tree.Node{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBalanced(tt.args.root)
			test.AssertEqual(t, got, tt.want, "unexpected result")
		})
	}
}

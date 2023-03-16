package tree

import (
	"fmt"
	"leetcode/go/pkg/sugar"
	"leetcode/go/test"
	"testing"
)

func TestTreeNode_Value(t *testing.T) {
	type fields struct {
		Val   int
		Left  *Node
		Right *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := n.Value(); got != tt.want {
				t.Errorf("TreeNode.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		node *Node
		want []*int
	}{
		{
			name: "toslice",
			node: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
				},
				Right: &Node{
					Val:   3,
					Right: &Node{Val: 7},
				},
			},
			want: []*int{sugar.P(1), sugar.P(2), sugar.P(3), nil, nil, nil, sugar.P(7)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.node.ToSlice()

			for i := range got {
				test.AssertEqual(t, got[i], tt.want[i], "unexpected value")
			}
		})
	}
}

func TestFromSlice(t *testing.T) {
	type args struct {
		s []*int
	}
	tests := []struct {
		name string
		args args
		want []*int
	}{
		{
			name: "[1, 2, 3, 4, 5, 6]",
			args: args{
				s: ToPointersSlice([]int{1, 2, 3, 4, 5, 6}),
			},
			want: ToPointersSlice([]int{1, 2, 3, 4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromSlice(tt.args.s)

			for _, v := range got.ToSlice() {
				if v == nil {
					fmt.Println("nil")
				} else {
					fmt.Println(*v)
				}
			}
			test.AssertEqual(t, got.ToSlice(), tt.want, "unexpected value")
		})
	}
}

func TestTreeNode_Length(t *testing.T) {
	tests := []struct {
		name string
		node *Node
		want int
	}{
		{
			name: "3",
			node: &Node{
				Val:   1,
				Left:  &Node{Val: 3},
				Right: &Node{Val: 2},
			},
			want: 3,
		},
		{
			name: "11",
			node: &Node{
				Val: 1,
				Left: &Node{
					Val:  3,
					Left: &Node{Val: 4},
					Right: &Node{
						Val:   1,
						Right: &Node{Val: 1},
					},
				},
				Right: &Node{
					Val:  2,
					Left: &Node{Val: 2},
				},
			},
			want: 11,
		},
		{
			name: "7",
			node: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
				},
				Right: &Node{
					Val:   3,
					Right: &Node{Val: 7},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test.AssertEqual(t, tt.node.SliceLen(), tt.want, "unexpected tree length")
		})
	}
}

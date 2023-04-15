package diameterofbinarytree

import (
	"leetcode/go/pkg/sugar/cmpr"
	"leetcode/go/pkg/tree"
)

func diameterOfBinaryTree(root *tree.Node) int {
	if root == nil || root.Left == nil && root.Right == nil {
		return 0
	}

	max := 0
	depth(root, func(a, b int) {
		if a+b > max {
			max = a + b
		}
	})

	return max
}

func depth(root *tree.Node, fn func(a, b int)) int {
	if root == nil || root.Left == nil && root.Right == nil {
		return 0
	}
	ld, rd := 0, 0
	if root.Left != nil {
		ld = 1 + depth(root.Left, fn)
	}
	if root.Right != nil {
		rd = 1 + depth(root.Right, fn)
	}
	fn(ld, rd)

	return cmpr.Max(ld, rd)
}

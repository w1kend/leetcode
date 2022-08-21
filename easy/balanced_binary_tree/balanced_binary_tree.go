package balancedbinarytree

import (
	"leetcode/pkg/sugar"
	"leetcode/pkg/tree"
)

// https://leetcode.com/problems/balanced-binary-tree
// https://leetcode.com/submissions/detail/754351525
func isBalanced(root *tree.Node) bool {
	if root == nil {
		return true
	}
	_, ok := isTreeBalanced(root)
	return ok
}

func isTreeBalanced(node *tree.Node) (int, bool) {
	if node == nil {
		return 0, true
	}

	ld, lf := isTreeBalanced(node.Left)
	rd, rf := isTreeBalanced(node.Right)

	if !lf || !rf || ld-rd > 1 || rd-ld > 1 {
		return 0, false
	}

	return sugar.Max(ld, rd) + 1, true
}

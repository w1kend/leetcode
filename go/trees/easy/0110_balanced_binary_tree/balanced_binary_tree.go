package balancedbinarytree

import (
	"leetcode/go/pkg/sugar/cmpr"
	"leetcode/go/pkg/tree"
)

func isBalanced(root *tree.Node) bool {
	balanced := true
	depth(root, func(a, b int) bool {
		if a-b > 1 || b-a > 1 {
			balanced = false
			return false
		}
		return true
	})

	return balanced
}

func depth(root *tree.Node, fn func(a, b int) bool) (int, bool) {
	if root == nil || root.Left == nil && root.Right == nil {
		return 0, true
	}
	ld, rd := 0, 0
	var ok bool
	if root.Left != nil {
		ld, ok = depth(root.Left, fn)
		if !ok {
			return 0, false
		}
		ld++
	}
	if root.Right != nil {
		rd, ok = depth(root.Right, fn)
		if !ok {
			return 0, false
		}
		rd++
	}
	if !fn(ld, rd) {
		return 0, false
	}

	return cmpr.Max(ld, rd), true
}

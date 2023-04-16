package isvalidbst

import (
	"leetcode/go/pkg/tree"
)

func isValidBST(root *tree.Node) bool {
	return isValid(root, nil, nil)
}

func isValid(node, min, max *tree.Node) bool {
	if node == nil {
		return true
	}

	if min != nil && node.Val <= min.Val {
		return false
	}

	if max != nil && node.Val >= max.Val {
		return false
	}

	return isValid(node.Left, min, node) && isValid(node.Right, node, max)
}

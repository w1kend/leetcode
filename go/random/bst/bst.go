package main

import "leetcode/go/pkg/xmath"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) IsNil() bool {
	return n == nil
}

func main() {

}

func insert(node *TreeNode, val int) {
	if node == nil {
		return
	}

	if node.Val > val {
		if node.Left.IsNil() {
			node.Left = &TreeNode{Val: val}
			return
		}
		insert(node.Left, val)
		return
	}

	if node.Val < val {
		if node.Right.IsNil() {
			node.Right = &TreeNode{Val: val}
			return
		}
		insert(node.Right, val)
		return
	}
}

func balance(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	return node
}

func IsBalanced(node *TreeNode) bool {
	b, _ := isBalanced(node)
	return b
}

func isBalanced(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}

	lb, lh := isBalanced(node.Left)
	if lb == false {
		// early exit if left sub-tree is not balanced
		return false, 0
	}

	rb, rh := isBalanced(node.Right)

	return lb && rb && xmath.Abs(lh-rh) <= 1, max(lh, rh) + 1
}

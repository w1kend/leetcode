package invertbinarytree

/**
* Definition for a binary tree node.
* type TreeNode struct {
	*     Val int
	*     Left *TreeNode
	*     Right *TreeNode
	* }
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/invert-binary-tree/
// https://leetcode.com/submissions/detail/736899381/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	}

	reverse(root)

	return root
}

func reverse(node *TreeNode) {
	if node == nil {
		return
	}
	node.Left, node.Right = node.Right, node.Left
	reverse(node.Left)
	reverse(node.Right)
}

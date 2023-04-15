package invertbinarytree

import (
	"leetcode/go/pkg/stack"
	"leetcode/go/pkg/tree"
)

func invertTree(root *tree.Node) *tree.Node {
	return recusrive(root)
}

func recusrive(root *tree.Node) *tree.Node {
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	}

	l, r := recusrive(root.Left), recusrive(root.Right)
	root.Left = r
	root.Right = l

	return root
}

func iterative(root *tree.Node) *tree.Node {
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	}

	stack := stack.Stack[*tree.Node]{}

	for stack.Size() > 0 {
		node := stack.Pop()
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			stack.Push(node.Left)
		}
		if node.Right != nil {
			stack.Push(node.Right)
		}
	}

	return root
}

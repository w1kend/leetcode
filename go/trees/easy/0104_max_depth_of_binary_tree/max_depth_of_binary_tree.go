package maxdepthofbinarytree

import (
	"container/list"
	"leetcode/go/pkg/tree"

	"github.com/samber/lo"
)

func maxDepth(root *tree.Node) int {
	return recursion(root)
}

func recursion(root *tree.Node) int {
	if root == nil {
		return 0
	}

	lLen, rLen := recursion(root.Left), recursion(root.Right)

	return 1 + lo.Max([]int{lLen, rLen})
}

func iteration(root *tree.Node) int {
	if root == nil {
		return 0
	}

	level := 0

	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		for i := 0; i < q.Len(); i++ {
			val := q.Front()
			node := val.Value.(*tree.Node)
			q.Remove(val)

			if node.Left != nil {
				q.PushBack(node.Left)
			}

			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}

		level++
	}

	return level
}

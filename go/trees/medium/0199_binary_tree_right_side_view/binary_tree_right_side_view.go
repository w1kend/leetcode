package binarytreerightsideview

import (
	"container/list"
	"leetcode/go/pkg/tree"
)

func rightSideView(root *tree.Node) []int {
	if root == nil {
		return []int{}
	}

	st := list.New()
	st.PushBack(root)
	res := make([]int, 0, 2)

	for st.Len() > 0 {
		res = append(res, st.Back().Value.(*tree.Node).Val)

		for n := st.Len(); n > 0; n-- {
			node := st.Remove(st.Front()).(*tree.Node)

			if node.Left != nil {
				st.PushBack(node.Left)
			}

			if node.Right != nil {
				st.PushBack(node.Right)
			}
		}
	}

	return res
}

package levelordertraversal

import (
	"container/list"
	"leetcode/go/pkg/tree"
)

func levelOrder(root *tree.Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	st := list.New()
	st.PushBack(root)

	res := make([][]int, 0, 3)

	for st.Len() > 0 {
		levelValues := make([]int, 0, st.Len())
		for n := st.Len(); n > 0; n-- {
			node := st.Remove(st.Front()).(*tree.Node)
			levelValues = append(levelValues, node.Val)

			if node.Left != nil {
				st.PushBack(node.Left)
			}

			if node.Right != nil {
				st.PushBack(node.Right)
			}
		}
		res = append(res, levelValues)
	}

	return res
}

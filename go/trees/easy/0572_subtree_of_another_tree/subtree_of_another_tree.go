package subtreeofanothertree

import (
	"container/list"
	"leetcode/go/pkg/tree"
)

func isSubtreeRecursive(root *tree.Node, subRoot *tree.Node) bool {
	if root == nil {
		return false
	}
	// compare trees
	if root.Val == subRoot.Val && compareRecusrive(root, subRoot) {
		return true
	}

	return isSubtreeRecursive(root.Left, subRoot) || isSubtreeRecursive(root.Right, subRoot)
}

func isSubtreeIterative(root *tree.Node, subRoot *tree.Node) bool {
	if root == nil {
		return false
	}
	// compare trees
	if root.Val == subRoot.Val && compareIterate(root, subRoot) {
		return true
	}

	return isSubtreeIterative(root.Left, subRoot) || isSubtreeIterative(root.Right, subRoot)
}

type pair struct {
	main *tree.Node
	sub  *tree.Node
}

func compareIterate(root, sub *tree.Node) bool {
	st := list.New()
	st.PushBack(pair{main: root, sub: sub})

	for st.Len() > 0 {
		n := st.Len()
		for i := 0; i < n; i++ {
			p := st.Remove(st.Front()).(pair)
			// if both are nil
			if p.main == nil && p.sub == nil {
				continue
			}
			// if some of them is nil or values are not equal
			if p.main == nil || p.sub == nil || p.main.Val != p.sub.Val {
				return false
			}
			st.PushBack(pair{main: p.main.Left, sub: p.sub.Left})
			st.PushBack(pair{main: p.main.Right, sub: p.sub.Right})
		}
	}

	return true
}

func compareRecusrive(root, sub *tree.Node) bool {
	if root == nil && sub == nil {
		return true
	}
	if root == nil || sub == nil || root.Val != sub.Val {
		return false
	}

	return compareRecusrive(root.Left, sub.Left) && compareRecusrive(root.Right, sub.Right)
}

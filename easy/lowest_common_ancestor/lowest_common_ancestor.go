package lowestcommonancestor

import (
	"fmt"
	"leetcode/pkg/tree"
)

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// https://leetcode.com/submissions/detail/749702329/
func lowestCommonAncestor(root, p, q *tree.Node) *tree.Node {
	pAnc, qAnc := &tree.Node{Left: &tree.Node{}}, &tree.Node{Left: &tree.Node{}}
	lowestAncestor(root, p, q, pAnc, qAnc)

	return pAnc
}

func lowestAncestor(n, p, q, pA, qA *tree.Node) {
	if n == nil {
		return
	}

	if n == p {
		*pA = *n
	}
	if n == q {
		*qA = *n
	}

	if *pA != *p || *qA != *q {
		fmt.Printf("call LEFT from %+v,\npA - %v, qA - %v\np - %v, q - %v\n *qA == *q %v, *pA == *p %v\n", n, pA, qA, p, q, *qA == *q, *pA == *p)
		lowestAncestor(n.Left, p, q, pA, qA)
	}

	if *pA != *p || *qA != *q {
		fmt.Printf("call RIGHT from %+v,\npA - %v, qA - %v\np - %v, q - %v\n *qA == *q %v, *pA == *p %v\n", n, pA, qA, p, q, *qA == *q, *pA == *p)
		lowestAncestor(n.Right, p, q, pA, qA)
	}

	fmt.Printf("AFTER CALLING left/right from - %d, pA - %v, qA = %v\n", n.Val, pA, qA)

	if n.Left != nil {
		if *pA != *qA && *n.Left == *pA {
			*pA = *n
		}

		if *pA != *qA && *n.Left == *qA {
			*qA = *n
		}
	}

	if n.Right != nil {
		if *pA != *qA && *n.Right == *pA {
			*pA = *n
		}

		if *pA != *qA && *n.Right == *qA {
			*qA = *n
		}
	}

	return
}

// https://leetcode.com/submissions/detail/751742490
// considering the property of binary search tree ==> x <= x.right, x >= x.left
func lowestCommonAncestor2(root, p, q *tree.Node) *tree.Node {
	cur := root

	for cur != nil {
		if p.Val <= cur.Val && q.Val >= cur.Val || p.Val >= cur.Val && q.Val <= cur.Val {
			return cur
		}

		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	return cur
}

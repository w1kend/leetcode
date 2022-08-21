package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/valid-parentheses
// https://leetcode.com/submissions/detail/736709181/
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	switch {
	case list1 == nil && list2 == nil:
		return nil
	case list1 == nil:
		return list2
	case list2 == nil:
		return list1
	}

	root := &ListNode{}

	currNode := root
	l1Node := list1
	l2Node := list2
	for l1Node != nil || l2Node != nil {
		switch {
		case l1Node == nil || (l2Node != nil && l2Node.Val <= l1Node.Val):
			currNode.Next = &ListNode{
				Val: l2Node.Val,
			}
			l2Node = l2Node.Next
		case l2Node == nil || (l1Node != nil && l1Node.Val <= l2Node.Val):
			currNode.Next = &ListNode{
				Val: l1Node.Val,
			}
			l1Node = l1Node.Next
		}

		currNode = currNode.Next
	}

	return root.Next
}

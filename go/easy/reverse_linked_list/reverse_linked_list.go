package reverselinkedlist

import "leetcode/go/pkg/list"

// https://leetcode.com/problems/reverse-linked-list
// https://leetcode.com/submissions/detail/755387800
func reverseList(head *list.ListNode) *list.ListNode {
	if head == nil {
		return head
	}

	cur := head

	prev := (*list.ListNode)(nil)
	for cur.Next != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	cur.Next = prev

	return cur
}

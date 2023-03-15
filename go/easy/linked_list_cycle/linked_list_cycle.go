package linkedlistcycle

import "leetcode/go/pkg/list"

// https://leetcode.com/problems/linked-list-cycle
// https://leetcode.com/submissions/detail/754610519
func hasCycle(head *list.ListNode) bool {
	if head == nil {
		return false
	}

	cur := head
	slowed := head

	s := false
	for cur.Next != slowed {
		if cur.Next == nil {
			return false
		}

		cur = cur.Next
		switch s {
		case true:
			slowed = slowed.Next
			s = false
		case false:
			s = true
		}
	}

	return true
}

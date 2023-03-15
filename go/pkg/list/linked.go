package list

// ListNode - a node in a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

// FromSlice - returns a list from a slice.
func FromSlice(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}

	head := &ListNode{Val: s[0]}
	cur := head
	for i := 1; i < len(s); i++ {
		cur.Next = &ListNode{
			Val:  s[i],
			Prev: cur,
		}
		cur = cur.Next
	}

	return head
}

// ToSlice - returns a slice of the list.
func (l *ListNode) ToSlice() []int {
	if l == nil {
		return []int{}
	}

	s := []int{l.Val}
	for l.Next != nil {
		l = l.Next
		s = append(s, l.Val)
	}

	return s
}

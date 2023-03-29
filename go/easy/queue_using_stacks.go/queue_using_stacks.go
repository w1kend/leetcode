package queueusingstacksgo

import "leetcode/go/pkg/stack"

// https://leetcode.com/problems/implement-queue-using-stacks
// https://leetcode.com/submissions/detail/754714290
type StackQueue struct {
	in  stack.Stack[int]
	out stack.Stack[int]
}

func (sq *StackQueue) Push(item int) {
	sq.in.Push(item)
}

func (sq *StackQueue) Pop() int {
	item := sq.Peek()
	sq.out.Pop()

	return item
}

// in 1,2,3,4
// pop
// in []
// out - 4,3,2
// push
// in [5]
// out [4,3,2]

//pop
// in [5]
// out[4,3]
//pop
// in [5]
//out[4]
//pop
// in [5]
//out[]
// in [], out [5] => out []

func (sq *StackQueue) Peek() int {
	if sq.Empty() {
		return -1
	}
	if sq.out.IsEmpty() {
		sq.fillOut()
	}

	return sq.out.Peek()
}

func (sq *StackQueue) Empty() bool {
	return sq.in.Size() == 0 && sq.out.Size() == 0
}

func (sq *StackQueue) fillOut() {
	for !sq.in.IsEmpty() {
		sq.out.Push(sq.in.Pop())
	}
}

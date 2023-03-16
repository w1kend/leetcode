package queueusingstacksgo

import (
	"leetcode/go/test"
	"testing"
)

func TestStackQueue(t *testing.T) {
	tests := []struct {
		name  string
		check func(t *testing.T)
	}{
		{
			name: `["MyQueue", "push", "push", "peek", "pop", "empty"] => [null, null, null, 1, 1, false]`,
			check: func(t *testing.T) {
				q := StackQueue{}

				q.Push(1)
				q.Push(2)
				test.AssertEqual(t, q.Peek(), 1, "unexpected value")
				test.AssertEqual(t, q.Pop(), 1, "unexpected value")
				test.AssertEqual(t, q.Empty(), false, "unexpected value")
				test.AssertEqual(t, q.Peek(), 2, "unexpected value")
			},
		},
		{
			// in = [[],[1],[2],[3],[4],[],[5],[],[],[],[]]
			// out = [null,null,null,null,null,1,null,2,3,4,5]
			name: `["MyQueue","push","push","push","push","pop","push","pop","pop","pop","pop"]`,
			check: func(t *testing.T) {
				q := StackQueue{}

				q.Push(1)
				q.Push(2)
				q.Push(3)
				q.Push(4)
				test.AssertEqual(t, q.Pop(), 1, "unexpected value")
				q.Push(5)
				test.AssertEqual(t, q.Pop(), 2, "unexpected value")
				test.AssertEqual(t, q.Pop(), 3, "unexpected value")
				test.AssertEqual(t, q.Pop(), 4, "unexpected value")
				test.AssertEqual(t, q.Pop(), 5, "unexpected value")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.check(t)
		})
	}
}

package minstack

import (
	"leetcode/go/test"
	"testing"
)

func TestMinStack(t *testing.T) {
	t.Run("e1", func(t *testing.T) {
		s := Constructor()
		s.Push(-2)
		s.Push(0)
		s.Push(-3)
		test.AssertEqual(t, s.GetMin(), -3, "")
		s.Pop()
		test.AssertEqual(t, s.Top(), 0, "")
		test.AssertEqual(t, s.GetMin(), -2, "")
	})

	t.Run("m2", func(t *testing.T) {
		s := Constructor()
		s.Push(-1)
		s.Push(-2)
		s.Push(-3)
		s.Push(-4)
		test.AssertEqual(t, s.GetMin(), -4, "")
		s.Pop()
		test.AssertEqual(t, s.GetMin(), -3, "")
		s.Pop()
		test.AssertEqual(t, s.GetMin(), -2, "")
		s.Pop()
		test.AssertEqual(t, s.GetMin(), -1, "")
		s.Pop()
	})

	t.Run("m3", func(t *testing.T) {
		s := Constructor()
		s.Push(-2)
		s.Push(-2)
		s.Push(-2)
		s.Push(0)
		s.Push(1)
		s.Push(-20)
		test.AssertEqual(t, s.GetMin(), -20, "")
		s.Pop()
		test.AssertEqual(t, s.GetMin(), -2, "")
		s.Pop()
		test.AssertEqual(t, s.GetMin(), -2, "")
		s.Pop()
		test.AssertEqual(t, s.GetMin(), -2, "")
		s.Pop()
		test.AssertEqual(t, s.GetMin(), -2, "")
		s.Pop()
		test.AssertEqual(t, s.GetMin(), -2, "")
		s.Pop()
	})
}

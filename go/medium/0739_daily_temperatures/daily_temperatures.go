package dailytemperatures

import "leetcode/go/pkg/stack"

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	st := stack.Stack[int]{}

	for i, t := range temperatures {
		for !st.IsEmpty() && temperatures[st.Peek()] < t {
			idx := st.Pop()
			res[idx] = i - idx
		}
		st.Push(i)
	}

	return res
}

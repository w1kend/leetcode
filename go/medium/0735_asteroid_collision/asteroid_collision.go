package asteroidcollision

import (
	"leetcode/go/pkg/stack"
)

func asteroidCollision(asteroids []int) []int {
	st := stack.Stack[int]{}

	for _, size := range asteroids {
		if size > 0 || st.IsEmpty() {
			st.Push(size)
			continue
		}

		for !st.IsEmpty() {
			last := st.Peek()
			if last < 0 {
				st.Push(size)
				break
			}
			// destroy previous asteroid
			if last < -size {
				st.Pop()
				if st.IsEmpty() {
					st.Push(size)
					break
				}
				continue
			}
			// prev. asteroid is bigger in size. destroy current
			if last > -size {
				break
			}

			// size is equal. destroy both
			if last == -size {
				st.Pop()
				break
			}
		}
	}

	return st.Source()
}

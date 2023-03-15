package climbingstars

// https://leetcode.com/problems/climbing-stairs
// https://leetcode.com/submissions/detail/755260035
func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}

	f := 1
	s := 1
	for i := 1; i < n; i++ {
		s, f = f+s, s
	}

	return s
}

// the solution is right, but it takes to much time
// basically the right solution is fibonacci sequence
func climb(step, n, cur int) int {
	if step+cur > n {
		return 0
	}
	if step+cur == n {
		return 1
	}

	return climb(1, n, cur+step) + climb(2, n, cur+step)
}

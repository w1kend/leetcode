package climbstairs

func climbStairs(n int) int {
	c1 := 0
	c2 := 1

	for i := 1; i <= n; i++ {
		c1, c2 = c2, c1+c2
	}

	return c2
}

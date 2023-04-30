package coinchange

func coinChange(coins []int, amount int) int {
	c := bruteForceChanger{availableCoins: coins, counts: make(map[int]int, amount/2)}
	return c.change(amount)
}

type bruteForceChanger struct {
	availableCoins []int
	counts         map[int]int
}

func newBrutForceChanger(coins []int) *bruteForceChanger {
	return &bruteForceChanger{
		availableCoins: coins,
		counts:         make(map[int]int, len(coins)/2),
	}
}

func (c *bruteForceChanger) change(amount int) int {
	if amount == 0 {
		return 0
	}
	//if we've already calculated it
	if cnt, ok := c.counts[amount]; ok {
		return cnt
	}

	min := amount
	for _, coin := range c.availableCoins {
		newAmount := amount - coin

		if newAmount == 0 {
			c.counts[amount] = 1
			return 1
		}
		if newAmount < 0 {
			continue
		}

		// calculate how many coins we should pick for (amount - coin)
		cnt := 1 + c.change(newAmount)

		// on each step we can have len(availableCoins) different changes
		// so we just calculate each of the paths and choose the minimum
		if cnt != 0 && cnt <= min {
			c.counts[amount] = cnt
			min = cnt
		}
	}

	// if we didn't find min value
	if c.counts[amount] == 0 {
		c.counts[amount] = -1
		return -1
	}

	return min
}

type iterativeChanger struct {
	availableCoins []int
	dp             []int
}

func newIterativeChanger(coins []int) *iterativeChanger {
	return &iterativeChanger{
		availableCoins: coins,
	}
}

func (c *iterativeChanger) change(amount int) int {
	if amount == 0 {
		return 0
	}

	c.dp = make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		for _, coin := range c.availableCoins {
			diff := i - coin
			if diff < 0 {
				continue
			}

			diffCnt := c.dp[diff]
			if diffCnt == 0 && diff != 0 {
				continue
			}
			currCnt := diffCnt + 1

			if c.dp[i] == 0 || currCnt < c.dp[i] {
				c.dp[i] = currCnt
			}
		}
	}

	if c.dp[amount] == 0 {
		return -1
	}

	return c.dp[amount]
}

package besttimetobuysellstock

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// https://leetcode.com/submissions/detail/736739834/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0
	lowestPrice := prices[0]

	//7,1,5,3,6,4
	for i := range prices {
		// possible profit
		pp := prices[i] - lowestPrice
		if pp > profit {
			profit = pp
		}

		if prices[i] < lowestPrice {
			lowestPrice = prices[i]
		}
	}

	return profit
}

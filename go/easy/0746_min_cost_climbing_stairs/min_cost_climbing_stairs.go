package mincostclimbingstairs

func minCostClimbingStairs(cost []int) int {
	prevStair, currStair := 0, cost[0]

	for i := 2; i <= len(cost); i++ {
		nextStair := min(currStair, prevStair) + cost[i-1]
		prevStair, currStair = currStair, nextStair
	}

	return min(currStair, prevStair)
}

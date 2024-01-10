package deleteandearn

import "sort"

func deleteAndEarn(nums []int) int {
	sort.Ints(nums)
	sums := make([]int, 0, len(nums))

	sum, lastNum := 0, nums[0]
	for _, num := range nums {
		if num == lastNum {
			sum += num
			continue
		}

		sums = append(sums, sum)
		if num != lastNum+1 {
			// add an empty space to be able to delete both nums
			sums = append(sums, 0)
		}
		sum = num
		lastNum = num
	}
	sums = append(sums, sum)

	dp := make([]int, len(sums)+1)
	dp[0] = 0
	dp[1] = sums[0]

	for i := 2; i <= len(sums); i++ {
		dp[i] = max(sums[i-1]+dp[i-2], dp[i-1])
	}

	return dp[len(sums)]
}

package houserobber

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	robAmounts := make([]int, len(nums))
	//base case
	robAmounts[0] = nums[0]
	robAmounts[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		robAmounts[i] = max(nums[i]+robAmounts[i-2], robAmounts[i-1])
	}

	return robAmounts[len(nums)-1]
}

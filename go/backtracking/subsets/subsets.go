package subsets

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	subset := subsets(nums[:len(nums)-1])

	n := len(subset)
	val := nums[len(nums)-1]
	for i := 1; i < n; i++ {
		nset := append(append([]int{}, subset[i]...), val)
		subset = append(subset, nset)
	}

	subset = append(subset, []int{val})

	return subset
}

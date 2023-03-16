package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	for i := range left {
		left[i], right[i] = 1, 1
	}

	for i, n := range nums {
		res := n
		if i > 0 {
			res *= left[i-1]
		}
		left[i] = res

		rIdx := len(nums) - i - 1
		res = nums[rIdx]
		if rIdx < len(nums)-1 {
			res *= right[rIdx+1]
		}
		right[rIdx] = res
	}

	res := make([]int, len(nums))
	for i := range nums {
		total := 1
		if i > 0 {
			total *= left[i-1]
		}
		if i < len(nums)-1 {
			total *= right[i+1]
		}
		res[i] = total
	}

	return res
}

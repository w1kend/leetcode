package majorityelement

// 2,2,1,1,1,2,2
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	el := nums[0]
	counter := 0
	for i := range nums {
		if counter == 0 {
			el = nums[i]
		}

		if nums[i] == el {
			counter++
		} else {
			counter--
		}
	}

	return el
}

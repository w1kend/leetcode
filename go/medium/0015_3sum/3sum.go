package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0, 2)
	var iVal, jVal int

	for k, kVal := range nums {
		if kVal > 0 {
			break
		}
		if k > 0 && nums[k-1] == kVal {
			continue
		}

		i := k + 1
		j := len(nums) - 1
		for i < j {
			jVal = nums[j]
			iVal = nums[i]
			sum := kVal + iVal + jVal

			if sum > 0 {
				j--
				continue
			}
			if sum < 0 {
				i++
				continue
			}

			res = append(res, []int{iVal, jVal, kVal})
			for i < j && nums[i] == nums[i+1] {
				i++
			}
			i++
		}
	}

	return res
}

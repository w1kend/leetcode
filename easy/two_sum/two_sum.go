package two_sum

// https://leetcode.com/problems/two-sum
// https://leetcode.com/submissions/detail/736602720/
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i := range nums {
		j, exists := m[nums[i]]
		if exists {
			return []int{j, i}
		}
		m[target-nums[i]] = i
	}

	return []int{}
}

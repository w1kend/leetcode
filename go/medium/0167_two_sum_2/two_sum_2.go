package twosum2

func twoSum(numbers []int, target int) []int {
	r := len(numbers) - 1
	l := 0

	for l < r {
		sum := numbers[l] + numbers[r]

		switch {
		case sum > target:
			r--
		case sum == target:
			return []int{l + 1, r + 1}
		case sum < target:
			l++
		}
	}

	return nil
}

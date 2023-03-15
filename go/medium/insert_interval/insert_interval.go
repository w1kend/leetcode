package insertinterval

// https://leetcode.com/problems/insert-interval
// https://leetcode.com/submissions/detail/755738937
func insert(intervals [][]int, newInterval []int) [][]int {
	i := len(intervals)

	for n := range intervals {
		if intervals[n][0] >= newInterval[0] {
			i = n
			break
		}
	}

	intervals = append(intervals[:i], append([][]int{newInterval}, intervals[i:]...)...)

	prevI := i
	for j := max(i-1, 0); j < len(intervals); j++ {
		if j == i || j < i && intervals[j][1] < intervals[i][0] || i < j && intervals[i][1] < intervals[j][0] {
			continue
		}

		l, r := min(intervals[i][0], intervals[j][0]), max(intervals[i][1], intervals[j][1])

		if i < j {
			intervals[j] = []int{l, r}
			i = j
		} else {
			prevI = max(prevI-1, 0)
			intervals[i] = []int{l, r}
			j++
		}
	}

	if prevI != i {
		intervals = append(intervals[:prevI], intervals[i:]...)
	}

	return intervals
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

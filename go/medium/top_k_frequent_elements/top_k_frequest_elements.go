package topkfrequentelements

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int, len(nums))
	topn := make([]int, 0, k)

	for _, n := range nums {
		counts[n]++
	}

	countsArr := make([][]int, len(nums)+1)
	for n, cnt := range counts {
		countsArr[cnt] = append(countsArr[cnt], n)
	}

	for i := len(countsArr) - 1; i >= 0; i-- {
		if len(topn) == k {
			break
		}
		values := countsArr[i]
		if len(values) > 0 {
			topn = append(topn, values...)
		}
	}

	return topn
}

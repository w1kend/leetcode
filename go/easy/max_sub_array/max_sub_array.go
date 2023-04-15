package max_sub_array

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/maximum-subarray
// https://leetcode.com/submissions/detail/746662273

// leetcode discussion
func maxSubArray2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	l := 0

	res := math.MinInt
	sum := 0
	for i := range nums {
		sum = sum + nums[i]
		if sum > res {
			res = sum
		}

		for sum < 0 {
			sum -= nums[l]
			l++
		}
	}

	return res
}

// divide and conquer alg
func maxSubArray(nums []int) int {
	i, j := subArr(nums)
	return sum(nums[i : j+1]...)
}

// i - biggest sub array start
// j - biggest sub array end
func subArr(v []int) (i int, j int) {
	if len(v) == 1 {
		return 0, 0
	}

	m := int(math.Floor(float64(len(v) / 2)))
	left := v[:m]
	right := v[m:]

	li, lj := subArr(left)
	ri, rj := subArr(right)

	lsum := sum(left[li : lj+1]...)
	rsum := sum(right[ri : rj+1]...)

	nsum := lsum
	maxNsum := nsum
	for i := lj + 1; i <= len(left)+rj; i++ {
		nsum += v[i]
		if nsum > maxNsum {
			maxNsum = nsum
			lj = i
		}
	}

	if maxNsum > lsum {
		lsum = maxNsum
	}

	if lsum > rsum {
		return li, lj
	}

	return len(left) + ri, len(left) + rj
}

func sum(vals ...int) int {
	s := 0
	for i := range vals {
		s += vals[i]
	}

	return s
}

func bprint(i, j, sum int, arr []int, str string) {
	fmt.Printf("%s-%+v = %d. soruce - %v\n", str, arr[i:j+1], sum, arr)
}

// Kadane's algorithm
func maxSubArray3(nums []int) int {
	localSum := 0
	maxSum := math.MinInt

	for i := range nums {
		localSum = max(nums[i], nums[i]+localSum)
		maxSum = max(localSum, maxSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

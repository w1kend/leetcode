package longestincreasingsubsequence

import "slices"

func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	s := NewSolver(nums)
	return s.solveTopDown(0, -1)
}

func lengthOfLIS2(nums []int) int {
	s := NewSolver(nums)
	return s.solveBottomUp()
}

type solver struct {
	nums []int
	memo [][]int
}

func NewSolver(nums []int) solver {
	s := solver{
		nums: nums,
		memo: make([][]int, 0, len(nums)+1),
	}

	for i := 0; i <= len(nums); i++ {
		ar := make([]int, len(nums)+1)
		for i := range ar {
			ar[i] = -1
		}
		s.memo = append(s.memo, ar)
	}

	return s
}

func (s *solver) solveTopDown(i, prev int) int {
	if i >= len(s.nums) {
		return 0
	}

	if s.memo[i][prev+1] != -1 {
		return s.memo[i][prev+1]
	}

	longest := 0

	if prev == -1 || s.nums[i] > s.nums[prev] {
		longest = 1 + s.solveTopDown(i+1, i)
	}

	longest = max(longest, s.solveTopDown(i+1, prev))

	s.memo[i][prev+1] = longest

	return longest
}

func (s *solver) solveBottomUp() int {
	seq := make([]int, len(s.nums))
	for i := range seq {
		seq[i] = 1
	}

	for i := range s.nums {
		for j := 0; j < i; j++ {
			if s.nums[i] > s.nums[j] {
				seq[i] = max(seq[i], 1+seq[j])
			}
		}
	}

	return slices.Max(seq)
}

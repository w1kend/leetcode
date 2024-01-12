package maxscoreperformingmultiplicationoperations

func maximumScore(nums []int, multipliers []int) int {
	return NewSolver(nums, multipliers).topDown(0, 0)
}

func maximumScore2(nums []int, multipliers []int) int {
	return NewSolver(nums, multipliers).bottopUp()
}

type solver struct {
	nums        []int
	multipliers []int
	memo        [][]int
}

func NewSolver(nums, multipliers []int) *solver {
	s := solver{
		nums:        nums,
		multipliers: multipliers,
		memo:        make([][]int, 0, len(multipliers)),
	}

	for i := 0; i < len(multipliers); i++ {
		s.memo = append(s.memo, make([]int, len(nums)))
	}

	return &s
}

func (s *solver) rightIndex(i, left int) int {
	return len(s.nums) - 1 - (i - left)
}

func (s *solver) topDown(i, left int) int {
	if i == len(s.memo) {
		return 0
	}

	if s.memo[i][left] != 0 {
		return s.memo[i][left]
	}

	maxScore := max(
		s.topDown(i+1, left+1)+s.nums[left]*s.multipliers[i],                // use a num from the left
		s.topDown(i+1, left)+s.nums[s.rightIndex(i, left)]*s.multipliers[i], // use a num from the right
	)

	s.memo[i][left] = maxScore

	return maxScore
}

func (s *solver) bottopUp() int {
	m := len(s.multipliers)
	dp := make([][]int, 0, m+1)
	for i := 0; i < m+1; i++ {
		dp = append(dp, make([]int, m+1))
	}

	for i := m - 1; i >= 0; i-- {
		for left := i; left >= 0; left-- {
			dp[i][left] = max(
				dp[i+1][left+1]+s.nums[left]*s.multipliers[i],                // use a num from the left
				dp[i+1][left]+s.nums[s.rightIndex(i, left)]*s.multipliers[i], // use a num from the right
			)
		}
	}

	return dp[0][0]
}

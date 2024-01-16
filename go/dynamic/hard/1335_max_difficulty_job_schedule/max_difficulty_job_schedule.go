package maxdifficultyjobschedule

import (
	"math"
	"slices"
)

func minDifficulty(jobDifficulty []int, d int) int {
	if len(jobDifficulty) < d {
		return -1
	}

	s := NewSolver(jobDifficulty, d)
	return s.solve(0, 1)
}

type solver struct {
	difficulties []int
	days         int

	memo [][]int
}

func NewSolver(difficulties []int, days int) solver {
	memo := make([][]int, 0, len(difficulties))
	for i := 0; i < len(difficulties); i++ {
		s := make([]int, days)
		for i := range s {
			s[i] = math.MinInt64
		}
		memo = append(memo, s)
	}

	return solver{
		difficulties: difficulties,
		days:         days,
		memo:         memo,
	}
}

func (s *solver) solve(firstJob, day int) int {
	if day == s.days {
		// the last day
		return slices.Max(s.difficulties[firstJob:])
	}

	if s.memo[firstJob][day] != math.MinInt64 {
		return s.memo[firstJob][day]
	}

	best := math.MaxInt64
	hardest := 0

	for j := firstJob; j < s.maxJobsAtDay(day); j++ {
		hardest = max(hardest, s.difficulties[j])
		best = min(best, hardest+s.solve(j+1, day+1))
	}

	s.memo[firstJob][day] = best

	return best
}

func (s *solver) maxJobsAtDay(day int) int {
	// total_jobs - min_jobs_needed_for_future_days
	return len(s.difficulties) - (s.days - day)
}

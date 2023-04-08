package kokoeatingbananas

import (
	"leetcode/go/pkg/sugar/cmpr"
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, v := range piles {
		max = cmpr.Max(max, v)
	}

	l := 1
	r := max

	var mid int
	for l <= r {
		mid = (l + r) / 2

		hours := calcHours(piles, mid)

		if l == r {
			break
		}

		if hours <= h {
			r = mid
		}

		if hours > h {
			l = mid + 1
		}
	}

	return mid
}

func calcHours(piles []int, perHour int) int {
	sum := 0
	for _, p := range piles {
		sum += int(math.Ceil(float64(p) / float64(perHour)))
	}

	return sum
}

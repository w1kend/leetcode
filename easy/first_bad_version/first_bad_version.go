package firstbadversion

import (
	"fmt"
	"math"
)

type Srv struct {
	FirstBadVersion int
}

// https://leetcode.com/problems/first-bad-version
// https://leetcode.com/submissions/detail/754867554
func (s Srv) IsBadVersion(version int) bool {
	fmt.Println(version)
	return s.FirstBadVersion <= version
}

func firstBadVersion(n int, s Srv) int {
	l := 0
	r := n
	for {
		mid := int(math.Floor(float64(l+r) / 2))

		isBad := s.IsBadVersion(mid)
		if isBad {
			b := s.IsBadVersion(mid - 1)
			if !b {
				return mid
			}
			r = mid - 1
		} else {
			b := s.IsBadVersion(mid + 1)
			if b {
				return mid + 1
			}
			l = mid + 1
		}
	}
}

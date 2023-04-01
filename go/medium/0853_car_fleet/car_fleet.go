package carfleet

import (
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	pts := make(map[int]int, len(position)) // position to speed

	for i, p := range position {
		pts[p] = speed[i]
	}

	sort.Ints(position)

	times := make([]float32, len(position))
	n := 0
	for i := len(position) - 1; i >= 0; i-- {
		pos := position[i]
		timeToTarget := float32(target-pos) / float32(pts[pos])
		times[n] = timeToTarget
		n++

		if n > 1 && times[n-1] <= times[n-2] {
			n--
		}
	}

	return n
}

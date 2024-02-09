package performstringshifts

func stringShift(s string, shift [][]int) string {
	if len(s) <= 1 {
		return s
	}

	finalShift := 0

	for _, s := range shift {
		direction, amount := s[0], s[1]
		if direction == 0 {
			finalShift -= amount
		} else {
			finalShift += amount
		}
	}

	finalShift %= len(s)

	if finalShift < 0 {
		return s[-finalShift:] + s[:-finalShift]
	}

	return s[len(s)-finalShift:] + s[:len(s)-finalShift]
}

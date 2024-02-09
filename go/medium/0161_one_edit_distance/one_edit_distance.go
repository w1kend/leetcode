package oneeditdistance

func isOneEditDistance(s string, t string) bool {
	if abs(len(s)-len(t)) > 1 {
		return false
	}

	i := 0

	minLen := min(len(s), len(t))

	for i < minLen && s[i] == t[i] {
		i++
	}

	switch {
	case len(s) == len(t):
		// replace. compare the rest parts
		return i != minLen && s[i+1:] == t[i+1:]
	case len(s) > len(t):
		// delete.
		return i == minLen || s[i+1:] == t[i:]
	case len(s) < len(t):
		// insert.
		return i == minLen || s[i:] == t[i+1:]
	default:
		return false
	}
}

func abs(v int) int {
	if v > 0 {
		return v
	}

	return -v
}

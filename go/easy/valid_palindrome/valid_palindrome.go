package validpalindrome

import (
	"math"
)

// https://leetcode.com/problems/valid-palindrome/
// https://leetcode.com/submissions/detail/736851778/
func IsPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	rawLetters := make([]rune, 0, len(s))

	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			rawLetters = append(rawLetters, r)
		}
		if r >= 'A' && r <= 'Z' {
			rawLetters = append(rawLetters, r-('A'-'a'))
		}
	}

	for i := 0; i < int(math.Floor(float64(len(rawLetters)/2))); i++ {
		if rawLetters[i] != rawLetters[len(rawLetters)-1-i] {
			return false
		}
	}

	return true
}

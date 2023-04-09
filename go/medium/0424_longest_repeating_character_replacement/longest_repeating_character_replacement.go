package longestrepeatingcharacterreplacement

import "leetcode/go/pkg/sugar/cmpr"

func characterReplacement(s string, k int) int {
	if len(s) == 1 {
		return 1
	}

	counts := make(map[rune]int16, 26)
	max := 0
	maxCount := int16(0)
	l := 0

	// build a window, check if it's a valid substring(we can replace only k letters)
	// if it's not then move a window from the left
	runes := []rune(s)
	for r := 0; r < len(runes); r++ {
		char := runes[r]
		// update a counter for the letter
		charCount := counts[char]
		charCount++
		counts[char] = charCount

		// calculate max frequent letter in out window.
		// because the other letters we are going to replace k times
		maxCount = cmpr.Max(maxCount, charCount)
		// curr window length
		currLength := r - l + 1

		// if out substring is not valid. (there are more different letters than we can replace)
		if currLength-int(maxCount) > k {
			counts[runes[l]]--
			l++
			continue
		}
		// save max length as a result
		max = cmpr.Max(currLength, max)
	}

	return max
}

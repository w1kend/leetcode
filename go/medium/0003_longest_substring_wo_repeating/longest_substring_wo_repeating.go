package longestsubstringworepeating

import "leetcode/go/pkg/sugar/cmpr"

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int, len(s))
	startIdx := 0 // an index from which we start calculating a length of a substring
	maxLen := 0   // max len of subscring without repeating characters
	for i, c := range s {
		idx, ok := m[c]
		if ok && idx >= startIdx {
			// if the char c presents between 'startIdx' and 'i'
			// then we start looking for a new interval from 'c's last appearence
			m[c] = i
			startIdx = i
		} else {
			// save index of c and calculate maxLen
			m[c] = i
			maxLen = cmpr.Max(maxLen, i-startIdx+1)
		}
	}

	return maxLen
}

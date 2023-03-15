package longestpalindrome

// https://leetcode.com/problems/longest-palindrome
// https://leetcode.com/submissions/detail/755288443
func longestPalindrome(s string) int {
	letters := make(map[byte]int, 56)
	for i := range s {
		letters[s[i]]++
	}

	maxLen := 0
	maxOdd := 0
	for l := range letters {
		switch {
		case letters[l]%2 == 0:
			maxLen += letters[l]
		case letters[l]-1 > 0:
			maxLen += letters[l] - 1
			maxOdd = 1
		case letters[l] == 1:
			maxOdd = 1
		}
	}

	return maxLen + maxOdd
}

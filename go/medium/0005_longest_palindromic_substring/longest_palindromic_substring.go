package longestpalindromicsubstring

func longestPalindrome(s string) string {
	longest := string(s[0])

	for i := 0; i < len(s)-len(longest); i++ {
		for j := i + len(longest) + 1; j <= len(s); j++ {
			substr := s[i:j]
			if isPalindrome(substr) && len(substr) > len(longest) {
				longest = substr
			}
		}
	}
	return longest
}

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}

func dpSolution(s string) string {
	palindromes := make([][]bool, 0, len(s))
	for range s {
		palindromes = append(palindromes, make([]bool, len(s)))
	}

	for i := range palindromes {
		// one letter palindromes
		palindromes[i][i] = true
	}
	longest := string(s[0])

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			// two letters palindromes
			palindromes[i][i+1] = true
			longest = s[i : i+2]
		}
	}

	for strLen := 3; strLen <= len(s); strLen++ {
		for i := 0; i <= len(s)-strLen; i++ {
			j := i + strLen - 1
			if palindromes[i+1][j-1] && s[j] == s[i] {
				palindromes[i][j] = true
				longest = s[i : j+1]
			}
		}
	}

	return longest
}

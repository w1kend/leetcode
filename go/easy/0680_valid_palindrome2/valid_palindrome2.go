package validpalindrome2

// aaavvfaaa
func validPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return isPalindrom(s, i, j-1) || isPalindrom(s, i+1, j)
		}
		i++
		j--
	}

	return true
}

func isPalindrom(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

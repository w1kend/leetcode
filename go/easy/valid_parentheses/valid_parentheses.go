package validparentheses

// https://leetcode.com/submissions/detail/736636274/
func IsValid(s string) bool {
	openFor := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]rune, len(s))
	j := 0
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack[j] = r
			j++
		case ')', ']', '}':
			if j < 1 || stack[j-1] != openFor[r] {
				return false
			}
			j--
		}
	}

	return j == 0
}

package validanagram

// todo: try without hash, using slice where 'a' index = 'a' - 'a', i letter index = l[i] - 'a'
// https://leetcode.com/problems/valid-anagram/
// https://leetcode.com/submissions/detail/739759746
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	words := make(map[byte]int32, 28)

	for i := range s {
		words[s[i]-'a'] += 1
		words[t[i]-'a'] -= 1
	}

	for i := range words {
		if words[i] < 0 {
			return false
		}
	}

	return true
}

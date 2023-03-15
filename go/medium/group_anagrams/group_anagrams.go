package groupanagrams

// https://leetcode.com/problems/group-anagrams/submissions/915923240/

func groupAnagrams(strs []string) [][]string {
	// hold an index to a result slice
	groups := make(map[[26]byte]int, len(strs)/2)

	res := make([][]string, 0, len(strs)/4)
	// to modify 'a' => 0, 'b' => 1, etc.
	// and fill an !array to get a key for a hashmap
	abyte := byte('a')
	emptyKey := [26]byte{}

	key := [26]byte{}
	for _, s := range strs {
		key = emptyKey
		for _, b := range []byte(s) {
			key[b-abyte]++
		}
		idx, ok := groups[key]
		if !ok {
			res = append(res, make([]string, 0, 2))
			idx = len(res) - 1
			groups[key] = idx
		}
		res[idx] = append(res[idx], s)
	}

	return res
}

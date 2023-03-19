package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	s := make(map[int]struct{}, len(nums))

	for _, n := range nums {
		s[n] = struct{}{}
	}

	longest := 0
	curr := 0
	for _, n := range nums {
		if _, ok := s[n-1]; !ok {
			//check a sequence
			idx := n + 1
			curr = 1
			t := true
			for t {
				if _, ok := s[idx]; ok {
					curr++
					idx++
				} else {
					t = false
					if curr > longest {
						longest = curr
					}
				}
			}
		}
	}
	if curr > longest {
		longest = curr
	}

	return longest
}

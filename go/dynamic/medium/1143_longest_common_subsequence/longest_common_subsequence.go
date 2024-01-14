package longestcommonsubsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	s := solver{text1: text1, text2: text2}
	mem := make([][]int, 0, len(text1))
	for i := 0; i < len(text1); i++ {
		mem = append(mem, make([]int, len(text2)))
	}

	return s.calc(0, 0, mem)
}

type solver struct {
	text1 string
	text2 string
}

func (s *solver) calc(i, j int, mem [][]int) int {
	if i >= len(s.text1) || j >= len(s.text2) {
		return 0
	}

	if mem[i][j] != 0 {
		return mem[i][j]
	}

	maxCnt := 0
	if s.text1[i] == s.text2[j] {
		maxCnt = max(maxCnt, s.calc(i+1, j+1, mem)+1)
	} else {
		maxCnt = max(
			s.calc(i+1, j, mem),
			s.calc(i, j+1, mem),
		)
	}

	mem[i][j] = maxCnt

	return maxCnt
}

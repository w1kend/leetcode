package wordbreak

func wordBreak(s string, wordDict []string) bool {
	sl := NewSolver(s, wordDict)

	return sl.solve()
}

type solver struct {
	dict map[string]bool
	word string
	memo []bool
}

func NewSolver(word string, dict []string) solver {
	s := solver{
		dict: map[string]bool{},
		word: word,
		memo: make([]bool, len(word)),
	}

	for _, w := range dict {
		s.dict[w] = true
	}

	return s
}

func (s *solver) solve() bool {
	for i := range s.word {
		for j := 0; j <= i; j++ {
			half := s.word[j+1 : i+1]
			wholeWord := s.word[:i+1]
			if s.memo[j] && s.dict[half] || s.dict[wholeWord] {
				s.memo[i] = true
			}
		}
	}

	return s.memo[len(s.memo)-1]
}

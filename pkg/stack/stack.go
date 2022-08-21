package stack

// TODO: add tests

type Stack struct {
	i     int
	items []int
}

func (s *Stack) Push(item int) {
	if s.i >= len(s.items) {
		n := make([]int, len(s.items)*2+1)
		for i := range s.items {
			n[i] = s.items[i]
		}
		s.items = n
	}

	s.items[s.i] = item
	s.i++
}

func (s *Stack) Pop() int {
	s.i--
	return s.items[s.i]
}

func (s *Stack) IsEmpty() bool {
	return s.i == 0
}

func (s *Stack) Peek() int {
	return s.items[s.i-1]
}

func (s *Stack) Size() int {
	return s.i
}

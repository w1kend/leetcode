package stack

// TODO: add tests

type Stack struct {
	elems []int
}

func (s *Stack) Push(item int) {
	s.elems = append(s.elems, item)
}

func (s *Stack) Pop() int {
	el := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return el
}

func (s *Stack) IsEmpty() bool {
	return len(s.elems) == 0
}

func (s *Stack) Peek() int {
	return s.elems[len(s.elems)-1]
}

func (s *Stack) Size() int {
	return len(s.elems)
}

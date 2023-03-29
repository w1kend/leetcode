package stack

type Stack[T any] struct {
	elems []T
}

func (s *Stack[T]) Push(item T) {
	s.elems = append(s.elems, item)
}

func (s *Stack[T]) Pop() T {
	el := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return el
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elems) == 0
}

func (s *Stack[T]) Peek() T {
	return s.elems[len(s.elems)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.elems)
}

func (s *Stack[T]) Source() []T {
	return s.elems
}

package minstack

type MinStack struct {
	elems []int
	mins  []int
}

func Constructor() MinStack {
	return MinStack{
		elems: make([]int, 0, 30),
		mins:  make([]int, 0, 30),
	}
}

func (s *MinStack) Push(val int) {
	s.elems = append(s.elems, val)

	if len(s.mins) == 0 || val <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, val)
	}
}

func (s *MinStack) Pop() {
	el := s.elems[len(s.elems)-1]
	if el == s.mins[len(s.mins)-1] {
		s.mins = s.mins[:len(s.mins)-1]
	}
	s.elems = s.elems[:len(s.elems)-1]
}

func (s *MinStack) Top() int {
	return s.elems[len(s.elems)-1]
}

func (s *MinStack) GetMin() int {
	return s.mins[len(s.mins)-1]
}

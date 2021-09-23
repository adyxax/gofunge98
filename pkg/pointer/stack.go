package pointer

type Stack struct {
	size   int
	height int
	data   []int
	next   *Stack // Pointer to the next element in the stack stack
}

func NewStack() *Stack {
	return &Stack{
		size:   32,
		height: 0,
		data:   make([]int, 32),
		next:   nil,
	}
}

func (s *Stack) Clear() {
	s.height = 0
}

func (s *Stack) Duplicate() {
	if s.height > 0 {
		s.Push(s.data[s.height-1])
	}
}

func (s *Stack) Pop() int {
	if s.height > 0 {
		s.height--
		return s.data[s.height]
	}
	return 0
}

func (s *Stack) Push(value int) {
	if s.height >= s.size {
		s.size += 32
		s.data = append(s.data, make([]int, 32)...)
	}
	s.data[s.height] = value
	s.height++
}

func (s *Stack) Swap() {
	a := s.Pop()
	b := s.Pop()
	s.Push(a)
	s.Push(b)
}

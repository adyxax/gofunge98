package stack

type Stack struct {
	size   int
	height int
	data   []int
	next   *Stack // Pointer to the next element in the stack stack
}

func NewStack(size int, next *Stack) *Stack {
	return &Stack{
		size:   size,
		height: 0,
		data:   make([]int, size),
		next:   next,
	}
}

func (s *Stack) Clear() {
	s.height = 0
}

func (s *Stack) Duplicate() {
	if s.height > 0 {
		s.Push(s.data[s.height-1])
	} else {
		s.Push(0)
		s.Push(0)
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

func (s Stack) GetHeights() []int {
	if s.next != nil {
		return append(s.next.GetHeights(), s.height)
	} else {
		return []int{s.height}
	}
}

func (toss *Stack) Transfert(soss *Stack, n int) {
	// Implements a value transfert between two stacks, intended for use with the '{'
	// (aka begin) and '}' (aka end) stackstack commands
	toss.height += n
	if toss.height > toss.size {
		toss.data = append(toss.data, make([]int, toss.height-toss.size)...)
		toss.size = toss.height
	}
	for i := 1; i <= min(soss.height, n); i++ {
		toss.data[toss.height-i] = soss.data[soss.height-i]
		for i := min(soss.height, n) + 1; i <= n; i++ {
			toss.data[toss.height-i] = 0
		}
	}
	soss.height -= n
	if soss.height < 0 {
		soss.height = 0
	}
}

func (s Stack) Next() *Stack {
	return s.next
}

func (s *Stack) Discard(n int) {
	// Implements a discard mechanism intended for use with the '}'(aka end) stackstack command
	s.height -= n
	if s.height < 0 {
		s.height = 0
	}
}

func (s *Stack) YCommandPick(n int, h int) {
	if n > s.height {
		s.height = 1
		s.data[0] = 0
	} else {
		v := s.data[s.height-n]
		s.height = h
		s.Push(v)
	}
}

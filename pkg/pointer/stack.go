package pointer

func (ss *StackStack) Pop() int {
	return ss.head.Pop()
}

func (ss *StackStack) Push(v int) {
	ss.head.Push(v)
}

func (s *StackStack) YCommandPick(n int, h int) {
	s.head.YCommandPick(n, h)
}

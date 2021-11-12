package pointer

import "git.adyxax.org/adyxax/gofunge98/pkg/stack"

type StackStack struct {
	head   *stack.Stack
	height int
}

func NewStackStack() *StackStack {
	return &StackStack{
		head:   stack.NewStack(32, nil),
		height: 1,
	}
}

func (ss *StackStack) Begin(p *Pointer) {
	ss.height++
	soss := ss.head
	n := soss.Pop()
	np := n
	if np < 0 {
		np = -np
	}
	toss := stack.NewStack(np, soss)
	ss.head = toss
	if n > 0 {
		toss.Transfert(soss, n)
	} else if n < 0 {
		for i := 0; i < np; i++ {
			soss.Push(0)
		}
	}
	x, y := p.GetStorageOffset()
	soss.Push(x)
	soss.Push(y)
	p.CalculateNewStorageOffset()
}

func (ss *StackStack) End(p *Pointer) (reflect bool) {
	if ss.height == 1 {
		return true
	}
	toss := ss.head
	soss := ss.head.Next()
	n := ss.head.Pop()
	y := soss.Pop()
	x := soss.Pop()
	p.SetStorageOffset(x, y)
	if n > 0 {
		soss.Transfert(toss, n)
	} else if n < 0 {
		soss.Discard(-n)
	}
	ss.height--
	ss.head = soss
	return false
}

func (ss *StackStack) Under() (reflect bool) {
	if ss.height == 1 {
		return true
	}
	n := ss.head.Pop()
	if n > 0 {
		for i := 0; i < n; i++ {
			ss.head.Push(ss.head.Next().Pop())
		}
	} else {
		for i := 0; i < -n; i++ {
			ss.head.Next().Push(ss.head.Pop())
		}
	}
	return false
}

func (s StackStack) GetHeights() []int {
	return s.head.GetHeights()
}

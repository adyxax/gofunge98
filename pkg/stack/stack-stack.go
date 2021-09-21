package stack

import (
	"git.adyxax.org/adyxax/gofunge/pkg/pointer"
)

type StackStack struct {
	head   *Stack
	height int
}

func NewStackStack() *StackStack {
	return &StackStack{
		head:   NewStack(),
		height: 1,
	}
}

func (ss *StackStack) Begin(p *pointer.Pointer) {
	ss.height++
	soss := ss.head
	n := soss.Pop()
	np := n
	if np < 0 {
		np = -np
	}
	toss := &Stack{
		size:   np,
		height: np,
		data:   make([]int, np),
		next:   soss,
	}
	ss.head = toss
	max := n - soss.height
	if max < 0 {
		max = 0
	}
	for i := n - 1; i >= max; i-- {
		toss.data[i] = soss.data[soss.height-n+i]
	}
	x, y := p.GetStorageOffset()
	soss.Push(x)
	soss.Push(y)
	p.CalculateNewStorageOffset()
}

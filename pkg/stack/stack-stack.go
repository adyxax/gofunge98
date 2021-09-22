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

func (ss *StackStack) End(p *pointer.Pointer) (reflect bool) {
	if ss.height == 1 {
		return true
	}
	soss := ss.head.next
	n := ss.head.Pop()
	y := soss.Pop()
	x := soss.Pop()
	p.SetStorageOffset(x, y)
	if n > 0 {
		soss.height += n
		if soss.size < soss.height {
			soss.data = append(soss.data, make([]int, soss.height-soss.size)...)
			soss.size = soss.height
		}
		for i := n; i > 0; i-- {
			soss.data[soss.height-i] = ss.head.data[ss.head.height-i]
		}
	} else {
		soss.height += n
		if soss.height < 0 {
			soss.height = 0
		}
	}
	ss.height--
	ss.head = ss.head.next
	return false
}

func (ss *StackStack) Under() (reflect bool) {
	if ss.height == 1 {
		return true
	}
	n := ss.head.Pop()
	if n > 0 {
		for i := 0; i < n; i++ {
			ss.head.Push(ss.head.next.Pop())
		}
	} else {
		for i := 0; i < -n; i++ {
			ss.head.next.Push(ss.head.Pop())
		}
	}
	return false
}

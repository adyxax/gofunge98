package pointer

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

func (ss *StackStack) Begin(p *Pointer) {
	ss.height++
	soss := ss.head
	n := soss.Pop()
	np := n
	if np < 0 {
		np = -np
	}
	toss := &Stack{
		size: np,
		data: make([]int, np),
		next: soss,
	}
	ss.head = toss
	if n > 0 {
		toss.height = n
		elts := soss.height - n
		if elts < 0 {
			elts = soss.height
		} else {
			elts = n
		}
		for i := 1; i <= elts; i++ {
			toss.data[toss.height-i] = soss.data[soss.height-i]
		}
		soss.height -= elts
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
	soss := ss.head.next
	n := ss.head.Pop()
	y := soss.Pop()
	x := soss.Pop()
	p.SetStorageOffset(x, y)
	if n > 0 {
		if n > ss.head.height {
			for i := n; i > ss.head.height; i-- {
				soss.Push(0)
			}
			n = ss.head.height
		}
		soss.height += n
		if soss.size < soss.height {
			soss.data = append(soss.data, make([]int, soss.height-soss.size)...)
			soss.size = soss.height
		}
		for i := n; i > 0; i-- {
			soss.data[soss.height-i] = ss.head.data[ss.head.height-i]
		}
	} else if n < 0 {
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

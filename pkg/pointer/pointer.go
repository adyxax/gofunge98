package pointer

import (
	"math/rand"

	"git.adyxax.org/adyxax/gofunge/pkg/field"
)

type Pointer struct {
	// the position
	x int
	y int
	// The delta
	dx int
	dy int
	// The Storage offset
	sox int
	soy int
	// The stack
	Ss *StackStack
	// The next element for the multi-"threaded" b98 interpreter
	Next *Pointer
}

func NewPointer() *Pointer {
	return &Pointer{dx: 1, Ss: NewStackStack()}
}

func (p Pointer) Split() *Pointer {
	return &p // p is already a copy
}

func (p *Pointer) Step(f field.Field) {
	p.x, p.y = f.Step(p.x, p.y, p.dx, p.dy)
}

func (p Pointer) Get(f field.Field) int {
	return f.Get(p.x, p.y)
}

func (p *Pointer) StepAndGet(f field.Field) int {
	p.Step(f)
	return p.Get(f)
}

func (p *Pointer) Set(x, y int) {
	p.x, p.y = x, y
}

func (p *Pointer) RedirectTo(dx, dy int) {
	p.dx, p.dy = dx, dy
}

func (p *Pointer) Reverse() {
	p.dx, p.dy = -p.dx, -p.dy
}

func (p *Pointer) Redirect(c int) bool {
	switch c {
	case '^':
		p.dx, p.dy = 0, -1
	case '>':
		p.dx, p.dy = 1, 0
	case 'v':
		p.dx, p.dy = 0, 1
	case '<':
		p.dx, p.dy = -1, 0
	case '?':
		directions := []int{0, -1, 1, 0, 0, 1, -1, 0}
		r := 2 * rand.Intn(4)
		p.dx, p.dy = directions[r], directions[r+1]
	case '[':
		p.dx, p.dy = p.dy, -p.dx
	case ']':
		p.dx, p.dy = -p.dy, p.dx
	case 'r':
		p.Reverse()
	case 'x':
		dy := p.Ss.Pop()
		dx := p.Ss.Pop()
		p.RedirectTo(dx, dy)
	default:
		return false
	}
	return true
}

package pointer

import "git.adyxax.org/adyxax/gofunge/pkg/field"

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
	// The next element for the multi-"threaded" b98 interpreter
	Next *Pointer
}

func NewPointer() *Pointer {
	return &Pointer{dx: 1}
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

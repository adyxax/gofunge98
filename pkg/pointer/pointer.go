package pointer

import "git.adyxax.org/adyxax/gofunge/pkg/field"

type Pointer struct {
	x     int
	y     int
	delta *Delta
}

func NewPointer() *Pointer {
	return &Pointer{delta: NewDelta(1, 0)}
}

func (p Pointer) ForkPointer() *Pointer {
	return &p // p is already a copy
}

func (p *Pointer) Step(f field.Field) {
	p.x, p.y = f.Step(p.x, p.y, p.delta.x, p.delta.y)
}

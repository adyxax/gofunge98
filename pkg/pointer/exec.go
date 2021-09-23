package pointer

import (
	"log"

	"git.adyxax.org/adyxax/gofunge/pkg/field"
)

func (p *Pointer) Exec(f *field.Field) (done bool, returnValue *int) {
	c := p.Get(*f)
	for jumpingMode := false; jumpingMode || c == ' ' || c == ';'; c = p.StepAndGet(*f) {
		if jumpingMode {
			if c == ';' {
				jumpingMode = false
			}
			continue
		}
	}
	switch c {
	case '@':
		return true, nil
	case '#':
		p.Step(*f)
	case 'j':
		n := p.Ss.Pop()
		for j := 0; j < n; j++ {
			p.Step(*f)
		}
	case 'q':
		v := p.Ss.Pop()
		return true, &v
	default:
		if !p.Redirect(c) {
			log.Fatalf("Non implemented instruction code %d : %c", c, c)
		}
	}
	p.Step(*f)
	return
}

package interpreter

import (
	"log"

	"git.adyxax.org/adyxax/gofunge/pkg/field"
	"git.adyxax.org/adyxax/gofunge/pkg/pointer"
)

type Interpreter struct {
	f *field.Field
	p *pointer.Pointer
}

func NewInterpreter(f *field.Field, p *pointer.Pointer) *Interpreter {
	return &Interpreter{f: f, p: p}
}

func (i *Interpreter) Run() {
	for i.p != nil {
		i.Step()
	}
}

func (i *Interpreter) Step() {
	var prev *pointer.Pointer = nil
	for p := i.p; p != nil; p = p.Next {
		switch p.Get(*i.f) {
		case '@':
			if prev == nil {
				i.p = p.Next
			} else {
				prev.Next = p.Next
			}
			break
		default:
			log.Fatalf("Non implemented instruction code %d : %c", p.Get(*i.f), p.Get(*i.f))
		}
	}
}

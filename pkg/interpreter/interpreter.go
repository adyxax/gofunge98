package interpreter

import (
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

func (i *Interpreter) Run() int {
	for i.p != nil {
		if v := i.step(); v != nil {
			return *v
		}
	}
	return 0
}

func (i *Interpreter) step() *int {
	var prev *pointer.Pointer = nil
	for p := i.p; p != nil; p = p.Next {
		done, v := p.Exec(i.f)
		if v != nil {
			return v
		}
		if done {
			if prev == nil {
				i.p = p.Next
			} else {
				prev.Next = p.Next
			}
		}
	}
	return nil
}

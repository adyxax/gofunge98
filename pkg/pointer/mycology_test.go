package pointer

import (
	"fmt"
	"os"
	"testing"

	"git.adyxax.org/adyxax/gofunge98/pkg/field"
	"github.com/stretchr/testify/require"
)

func TestMycology(t *testing.T) {
	file, err := os.Open("../../mycology/mycology.b98")
	if err != nil {
		t.Skip("mycology test suite not found, skipping")
	}

	f, err := field.Load(file)
	require.NoError(t, err)
	p := NewPointer()
	p.Argv = []string{"../../mycology/mycology.b98"}
	output := ""
	// TODO test expected output
	p.CharacterOutput = func(c int) {
		output += fmt.Sprintf("%c", c)
	}
	i := NewInterpreter(f, p)

	v := i.Run()
	require.Equal(t, 15, v)
}

type Interpreter struct {
	f *field.Field
	p *Pointer
}

func NewInterpreter(f *field.Field, p *Pointer) *Interpreter {
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
	var prev *Pointer = nil
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

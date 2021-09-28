package interpreter

import (
	"fmt"
	"os"
	"testing"

	"git.adyxax.org/adyxax/gofunge98/pkg/field"
	"git.adyxax.org/adyxax/gofunge98/pkg/pointer"
	"github.com/stretchr/testify/require"
)

func TestMycology(t *testing.T) {
	file, err := os.Open("../../mycology/mycology.b98")
	if err != nil {
		t.Skip("mycology test suite not found, skipping")
	}

	f, err := field.Load(file)
	require.NoError(t, err)
	p := pointer.NewPointer()
	p.Argv = []string{"../../mycology/mycology.b98"}
	// TODO test expected output
	output := ""
	p.CharacterOutput = func(c int) {
		output += fmt.Sprintf("%c", c)
	}
	i := NewInterpreter(f, p)

	v := i.Run()
	require.Equal(t, 15, v)
}

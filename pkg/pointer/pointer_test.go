package pointer

import (
	"os"
	"testing"

	"git.adyxax.org/adyxax/gofunge/pkg/field"
	"github.com/stretchr/testify/require"
)

func TestNewPointer(t *testing.T) {
	require.Equal(t, NewPointer(), &Pointer{delta: &Delta{1, 0}})
}

func TestSplit(t *testing.T) {
	file, err := os.Open("../field/test_data/hello.b98")
	require.NoError(t, err, "Failed to open file")
	defer file.Close()
	f, err := field.Load(file)
	require.NoError(t, err)
	p := NewPointer()
	p2 := p.Split()
	// We check that p2 is a real copy
	p.Step(*f)
	p2.Step(*f)
	require.Equal(t, p, &Pointer{x: 1, y: 0, delta: &Delta{1, 0}})
	require.Equal(t, p2, &Pointer{x: 1, y: 0, delta: &Delta{1, 0}})
}

func TestStep(t *testing.T) { // Step is thoroughly tested in the field package
	defaultPointer := NewPointer()
	// File of one char
	file, err := os.Open("../field/test_data/minimal.b98")
	require.NoError(t, err, "Failed to open file")
	defer file.Close()
	f, err := field.Load(file)
	require.NoError(t, err)
	p := NewPointer()
	p.Step(*f)
	require.Equal(t, defaultPointer, p)
}

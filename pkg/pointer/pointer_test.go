package pointer

import (
	"os"
	"testing"

	"git.adyxax.org/adyxax/gofunge/pkg/field"
	"github.com/stretchr/testify/require"
)

func TestNewPointer(t *testing.T) {
	require.Equal(t, NewPointer(), &Pointer{dx: 1, ss: NewStackStack()})
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
	require.Equal(t, &Pointer{x: 1, y: 0, dx: 1, ss: NewStackStack()}, p)
	require.Equal(t, &Pointer{x: 1, y: 0, dx: 1, ss: NewStackStack()}, p2)
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

func TestGet(t *testing.T) {
	// File of one char
	file, err := os.Open("../field/test_data/minimal.b98")
	require.NoError(t, err, "Failed to open file")
	defer file.Close()
	f, err := field.Load(file)
	p := NewPointer()
	v := p.Get(*f)
	require.Equal(t, int('@'), v)
}

func TestSet(t *testing.T) {
	p := NewPointer()
	p.Set(3, 14)
	require.Equal(t, 3, p.x)
	require.Equal(t, 14, p.y)
}

func TestRedirectTo(t *testing.T) {
	p := NewPointer()
	p.RedirectTo(3, 14)
	require.Equal(t, 3, p.dx)
	require.Equal(t, 14, p.dy)
}

func TestReverse(t *testing.T) {
	p := NewPointer()
	p.RedirectTo(3, 14)
	p.Reverse()
	require.Equal(t, -3, p.dx)
	require.Equal(t, -14, p.dy)
}

func TestRedirect(t *testing.T) {
	testCases := []struct {
		name       string
		input      byte
		expectedDx int
		expectedDy int
	}{
		{"up", '^', 0, -1},
		{"right", '>', 1, 0},
		{"down", 'v', 0, 1},
		{"left", '<', -1, 0},
		{"turn left", '[', 14, -3},
		{"turn right", ']', -14, 3},
		{"reverse", 'r', -3, -14},
		{"redirectTo", 'x', 2, 7},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewPointer()
			p.RedirectTo(3, 14)
			p.ss.Push(2)
			p.ss.Push(7)
			v := p.Redirect(int(tc.input))
			require.Equal(t, true, v)
			require.Equal(t, tc.expectedDx, p.dx, "Invalid dx value")
			require.Equal(t, tc.expectedDy, p.dy, "Invalid dy value")
		})
	}
	// We cannot really test random, can we? This just gives coverage
	p := NewPointer()
	v := p.Redirect(int('?'))
	require.Equal(t, true, v)
	// A character that does not redirect should return false
	require.Equal(t, false, p.Redirect(int('@')))
}

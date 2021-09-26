package field

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlankInside(t *testing.T) {
	input := Field{
		x:  0,
		y:  0,
		lx: 3,
		ly: 1,
		lines: []Line{
			Line{x: 0, l: 3, columns: []int{'@', 'a', 'b'}},
		},
	}
	expected := Field{
		x:  0,
		y:  0,
		lx: 3,
		ly: 1,
		lines: []Line{
			Line{x: 0, l: 3, columns: []int{'@', ' ', 'b'}},
		},
	}
	// Test cases
	testCases := []struct {
		name     string
		input    *Field
		inputX   int
		inputY   int
		expected *Field
	}{
		{"inside", &input, 1, 0, &expected},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.input.Blank(tc.inputX, tc.inputY)
			require.Equal(t, tc.expected, tc.input)
		})
	}
}

func TestBlankInsideLine(t *testing.T) {
	input := Field{
		x:  0,
		y:  0,
		lx: 3,
		ly: 3,
		lines: []Line{
			Line{x: 0, l: 3, columns: []int{'@', 'a', 'b'}},
			Line{x: 0, l: 1, columns: []int{'d'}},
			Line{x: 0, l: 1, columns: []int{'c'}},
		},
	}
	expected := Field{
		x:  0,
		y:  0,
		lx: 3,
		ly: 3,
		lines: []Line{
			Line{x: 0, l: 3, columns: []int{'@', 'a', 'b'}},
			Line{columns: []int{}},
			Line{x: 0, l: 1, columns: []int{'c'}},
		},
	}
	// Test cases
	testCases := []struct {
		name     string
		input    *Field
		inputX   int
		inputY   int
		expected *Field
	}{
		{"inside", &input, 0, 1, &expected},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.input.Blank(tc.inputX, tc.inputY)
			require.Equal(t, tc.expected, tc.input)
		})
	}
}

func TestBlankOutside(t *testing.T) {
	input := Field{
		x:  0,
		y:  0,
		lx: 1,
		ly: 1,
		lines: []Line{
			Line{x: 0, l: 1, columns: []int{'@'}},
		},
	}
	expected := Field{
		x:  0,
		y:  0,
		lx: 1,
		ly: 1,
		lines: []Line{
			Line{x: 0, l: 1, columns: []int{'@'}},
		},
	}
	// Test cases
	testCases := []struct {
		name     string
		input    *Field
		inputX   int
		inputY   int
		expected *Field
	}{
		{"xappend", &input, 1, 0, &expected},
		{"xprepend", &input, -1, 0, &expected},
		{"yappend", &input, 0, 1, &expected},
		{"yprepend", &input, 0, -1, &expected},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.input.Blank(tc.inputX, tc.inputY)
			require.Equal(t, tc.expected, tc.input)
		})
	}
}

func TestBlankOutsideLine(t *testing.T) {
	input := Field{
		x:  -1,
		y:  0,
		lx: 3,
		ly: 2,
		lines: []Line{
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
			Line{x: 0, l: 1, columns: []int{'@'}},
		},
	}
	expected := Field{
		x:  -1,
		y:  0,
		lx: 3,
		ly: 2,
		lines: []Line{
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
			Line{x: -0, l: 1, columns: []int{'@'}},
		},
	}
	// Test cases
	testCases := []struct {
		name     string
		input    *Field
		inputX   int
		inputY   int
		expected *Field
	}{
		{"xappend", &input, 1, 1, &expected},
		{"xprepend", &input, -1, 1, &expected},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.input.Blank(tc.inputX, tc.inputY)
			require.Equal(t, tc.expected, tc.input)
		})
	}
}

func TestBlankLineTrim(t *testing.T) {
	expected := Field{
		x:  -1,
		y:  0,
		lx: 3,
		ly: 1,
		lines: []Line{
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
		},
	}
	bottomCenter := Field{
		x:  -1,
		y:  0,
		lx: 3,
		ly: 3,
		lines: []Line{
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
			Line{},
			Line{x: 0, l: 1, columns: []int{'@'}},
		},
	}
	bottomLeft := Field{
		x:  -4,
		y:  0,
		lx: 6,
		ly: 3,
		lines: []Line{
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
			Line{},
			Line{x: -4, l: 1, columns: []int{'@'}},
		},
	}
	bottomRight := Field{
		x:  -1,
		y:  0,
		lx: 6,
		ly: 3,
		lines: []Line{
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
			Line{},
			Line{x: 4, l: 1, columns: []int{'@'}},
		},
	}
	topCenter := Field{
		x:  -1,
		y:  -2,
		lx: 3,
		ly: 3,
		lines: []Line{
			Line{x: 0, l: 1, columns: []int{'@'}},
			Line{},
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
		},
	}
	topLeft := Field{
		x:  -4,
		y:  -2,
		lx: 6,
		ly: 3,
		lines: []Line{
			Line{x: -4, l: 1, columns: []int{'@'}},
			Line{},
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
		},
	}
	topRight := Field{
		x:  -1,
		y:  -2,
		lx: 6,
		ly: 3,
		lines: []Line{
			Line{x: 4, l: 1, columns: []int{'@'}},
			Line{},
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
		},
	}
	// Test cases
	testCases := []struct {
		name     string
		input    *Field
		inputX   int
		inputY   int
		expected *Field
	}{
		{"bottomCenter", &bottomCenter, 0, 2, &expected},
		{"bottomLeft", &bottomLeft, -4, 2, &expected},
		{"bottomRight", &bottomRight, 4, 2, &expected},
		{"topCenter", &topCenter, 0, -2, &expected},
		{"topLeft", &topLeft, -4, -2, &expected},
		{"topRight", &topRight, 4, -2, &expected},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.input.Blank(tc.inputX, tc.inputY)
			require.Equal(t, tc.expected, tc.input)
		})
	}
}

func TestBlankcolumnsTrim(t *testing.T) {
	expectedBottom := Field{
		x:  -1,
		y:  0,
		lx: 3,
		ly: 2,
		lines: []Line{
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
			Line{x: 0, l: 1, columns: []int{'@'}},
		},
	}
	expectedTop := Field{
		x:  -1,
		y:  -1,
		lx: 3,
		ly: 2,
		lines: []Line{
			Line{x: 0, l: 1, columns: []int{'@'}},
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
		},
	}
	bottomLeft := Field{
		x:  -4,
		y:  0,
		lx: 6,
		ly: 2,
		lines: []Line{
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
			Line{x: -4, l: 5, columns: []int{'@', ' ', ' ', ' ', '@'}},
		},
	}
	bottomRight := Field{
		x:  -1,
		y:  0,
		lx: 6,
		ly: 2,
		lines: []Line{
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
			Line{x: 0, l: 5, columns: []int{'@', ' ', ' ', ' ', '@'}},
		},
	}
	topLeft := Field{
		x:  -4,
		y:  -1,
		lx: 6,
		ly: 2,
		lines: []Line{
			Line{x: -4, l: 5, columns: []int{'@', ' ', ' ', ' ', '@'}},
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
		},
	}
	topRight := Field{
		x:  -1,
		y:  -1,
		lx: 6,
		ly: 2,
		lines: []Line{
			Line{x: 0, l: 5, columns: []int{'@', ' ', ' ', ' ', '@'}},
			Line{x: -1, l: 3, columns: []int{'@', '@', '@'}},
		},
	}
	// Test cases
	testCases := []struct {
		name     string
		input    *Field
		inputX   int
		inputY   int
		expected *Field
	}{
		{"bottomLeft", &bottomLeft, -4, 1, &expectedBottom},
		{"bottomRight", &bottomRight, 4, 1, &expectedBottom},
		{"topLeft", &topLeft, -4, -1, &expectedTop},
		{"topRight", &topRight, 4, -1, &expectedTop},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.input.Blank(tc.inputX, tc.inputY)
			require.Equal(t, tc.expected, tc.input)
		})
	}
}

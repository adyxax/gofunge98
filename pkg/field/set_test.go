package field

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetMinimalAppendOne(t *testing.T) {
	hashField := Field{
		x:  0,
		y:  0,
		lx: 1,
		ly: 1,
		lines: []Line{
			Line{x: 0, l: 1, columns: []int{'#'}},
		},
	}
	xAppendField := Field{
		x:  0,
		y:  0,
		lx: 2,
		ly: 1,
		lines: []Line{
			Line{x: 0, l: 2, columns: []int{'@', '#'}},
		},
	}
	xPrependField := Field{
		x:  -1,
		y:  0,
		lx: 2,
		ly: 1,
		lines: []Line{
			Line{x: -1, l: 2, columns: []int{'#', '@'}},
		},
	}
	yAppendField := Field{
		x:  0,
		y:  0,
		lx: 1,
		ly: 2,
		lines: []Line{
			Line{x: 0, l: 1, columns: []int{'@'}},
			Line{x: 0, l: 1, columns: []int{'#'}},
		},
	}
	yPrependField := Field{
		x:  0,
		y:  -1,
		lx: 1,
		ly: 2,
		lines: []Line{
			Line{x: 0, l: 1, columns: []int{'#'}},
			Line{x: 0, l: 1, columns: []int{'@'}},
		},
	}
	// Test cases
	testCases := []struct {
		name     string
		input    string
		inputX   int
		inputY   int
		inputV   int
		expected *Field
	}{
		{"simple", "test_data/minimal.b98", 0, 0, '#', &hashField},
		{"xappend", "test_data/minimal.b98", 1, 0, '#', &xAppendField},
		{"xprepend", "test_data/minimal.b98", -1, 0, '#', &xPrependField},
		{"yappend", "test_data/minimal.b98", 0, 1, '#', &yAppendField},
		{"yprepend", "test_data/minimal.b98", 0, -1, '#', &yPrependField},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.input)
			require.NoError(t, err, "Failed to open file")
			defer file.Close()
			field, err := Load(file)
			field.Set(tc.inputX, tc.inputY, tc.inputV)
			require.NoError(t, err)
			require.Equal(t, tc.expected, field, "Invalid value")
		})
	}
}

func TestSetMinimalAppendTwo(t *testing.T) {
	bottomRight := Field{
		x:  0,
		y:  0,
		lx: 6,
		ly: 3,
		lines: []Line{
			Line{x: 0, l: 1, columns: []int{'@'}},
			Line{},
			Line{x: 5, l: 1, columns: []int{'#'}},
		},
	}
	bottomLeft := Field{
		x:  -5,
		y:  0,
		lx: 6,
		ly: 4,
		lines: []Line{
			Line{x: 0, l: 1, columns: []int{'@'}},
			Line{},
			Line{},
			Line{x: -5, l: 1, columns: []int{'#'}},
		},
	}
	topRight := Field{
		x:  0,
		y:  -3,
		lx: 9,
		ly: 4,
		lines: []Line{
			Line{x: 8, l: 1, columns: []int{'#'}},
			Line{},
			Line{},
			Line{x: 0, l: 1, columns: []int{'@'}},
		},
	}
	topLeft := Field{
		x:  -12,
		y:  -4,
		lx: 13,
		ly: 5,
		lines: []Line{
			Line{x: -12, l: 1, columns: []int{'#'}},
			Line{},
			Line{},
			Line{},
			Line{x: 0, l: 1, columns: []int{'@'}},
		},
	}
	// Test cases
	testCases := []struct {
		name     string
		input    string
		inputX   int
		inputY   int
		inputV   int
		expected *Field
	}{
		{"bottomRight", "test_data/minimal.b98", 5, 2, '#', &bottomRight},
		{"bottomLeft", "test_data/minimal.b98", -5, 3, '#', &bottomLeft},
		{"topRight", "test_data/minimal.b98", 8, -3, '#', &topRight},
		{"topLeft", "test_data/minimal.b98", -12, -4, '#', &topLeft},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.input)
			require.NoError(t, err, "Failed to open file")
			defer file.Close()
			field, err := Load(file)
			field.Set(tc.inputX, tc.inputY, tc.inputV)
			require.NoError(t, err)
			require.Equal(t, tc.expected, field, "Invalid value")
		})
	}
}

func TestSetMinimalAppendThree(t *testing.T) {
	xAppendField := Field{
		x:  0,
		y:  0,
		lx: 4,
		ly: 1,
		lines: []Line{
			Line{x: 0, l: 4, columns: []int{'@', ' ', ' ', '#'}},
		},
	}
	xPrependField := Field{
		x:  -3,
		y:  0,
		lx: 4,
		ly: 1,
		lines: []Line{
			Line{x: -3, l: 4, columns: []int{'#', ' ', ' ', '@'}},
		},
	}
	yAppendField := Field{
		x:  0,
		y:  0,
		lx: 1,
		ly: 4,
		lines: []Line{
			Line{x: 0, l: 1, columns: []int{'@'}},
			Line{},
			Line{},
			Line{x: 0, l: 1, columns: []int{'#'}},
		},
	}
	yPrependField := Field{
		x:  0,
		y:  -3,
		lx: 1,
		ly: 4,
		lines: []Line{
			Line{x: 0, l: 1, columns: []int{'#'}},
			Line{},
			Line{},
			Line{x: 0, l: 1, columns: []int{'@'}},
		},
	}
	// Test cases
	testCases := []struct {
		name     string
		input    string
		inputX   int
		inputY   int
		inputV   int
		expected *Field
	}{
		{"xappend", "test_data/minimal.b98", 3, 0, '#', &xAppendField},
		{"xprepend", "test_data/minimal.b98", -3, 0, '#', &xPrependField},
		{"yappend", "test_data/minimal.b98", 0, 3, '#', &yAppendField},
		{"yprepend", "test_data/minimal.b98", 0, -3, '#', &yPrependField},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.input)
			require.NoError(t, err, "Failed to open file")
			defer file.Close()
			field, err := Load(file)
			field.Set(tc.inputX, tc.inputY, tc.inputV)
			require.NoError(t, err)
			require.Equal(t, tc.expected, field, "Invalid value")
		})
	}
}

func TestSetAppendResize(t *testing.T) {
	base := Field{
		x: -1, y: -1, lx: 3, ly: 3, lines: []Line{
			Line{x: 0, l: 1, columns: []int{'u'}},
			Line{x: -1, l: 3, columns: []int{'l', '0', 'r'}},
			Line{x: 0, l: 1, columns: []int{'d'}},
		},
	}
	xappend := Field{
		x: -1, y: -1, lx: 5, ly: 3, lines: []Line{
			Line{x: 0, l: 1, columns: []int{'u'}},
			Line{x: -1, l: 3, columns: []int{'l', '0', 'r'}},
			Line{x: 0, l: 4, columns: []int{'d', ' ', ' ', 'n'}},
		},
	}
	xprepend := Field{
		x: -3, y: -1, lx: 5, ly: 3, lines: []Line{
			Line{x: 0, l: 1, columns: []int{'u'}},
			Line{x: -1, l: 3, columns: []int{'l', '0', 'r'}},
			Line{x: -3, l: 4, columns: []int{'n', ' ', ' ', 'd'}},
		},
	}
	xprependyprepend := Field{
		x: -3, y: -3, lx: 5, ly: 5, lines: []Line{
			Line{x: -3, l: 1, columns: []int{'n'}},
			Line{},
			Line{x: 0, l: 1, columns: []int{'u'}},
			Line{x: -1, l: 3, columns: []int{'l', '0', 'r'}},
			Line{x: 0, l: 1, columns: []int{'d'}},
		},
	}
	xappendyprepend := Field{
		x: -1, y: -3, lx: 5, ly: 5, lines: []Line{
			Line{x: 3, l: 1, columns: []int{'n'}},
			Line{},
			Line{x: 0, l: 1, columns: []int{'u'}},
			Line{x: -1, l: 3, columns: []int{'l', '0', 'r'}},
			Line{x: 0, l: 1, columns: []int{'d'}},
		},
	}
	xprependyappend := Field{
		x: -3, y: -1, lx: 5, ly: 5, lines: []Line{
			Line{x: 0, l: 1, columns: []int{'u'}},
			Line{x: -1, l: 3, columns: []int{'l', '0', 'r'}},
			Line{x: 0, l: 1, columns: []int{'d'}},
			Line{},
			Line{x: -3, l: 1, columns: []int{'n'}},
		},
	}
	xappendyappend := Field{
		x: -1, y: -1, lx: 5, ly: 5, lines: []Line{
			Line{x: 0, l: 1, columns: []int{'u'}},
			Line{x: -1, l: 3, columns: []int{'l', '0', 'r'}},
			Line{x: 0, l: 1, columns: []int{'d'}},
			Line{},
			Line{x: 3, l: 1, columns: []int{'n'}},
		},
	}
	// Test cases
	testCases := []struct {
		name     string
		input    Field
		inputX   int
		inputY   int
		inputV   int
		expected Field
	}{
		{"xappend", base, 3, 1, 'n', xappend},
		{"xprepend", base, -3, 1, 'n', xprepend},
		{"xprependyprepend", base, -3, -3, 'n', xprependyprepend},
		{"xappendyprepend", base, 3, -3, 'n', xappendyprepend},
		{"xprependyappend", base, -3, 3, 'n', xprependyappend},
		{"xappendyappend", base, 3, 3, 'n', xappendyappend},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			field := tc.input
			field.lines = make([]Line, field.ly)
			copy(field.lines, tc.input.lines)
			field.Set(tc.inputX, tc.inputY, tc.inputV)
			require.Equal(t, tc.expected, field, "Invalid value")
		})
	}
}

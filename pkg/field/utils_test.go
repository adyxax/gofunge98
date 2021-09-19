package field

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		input    string
		inputX   int
		inputY   int
		expected int
	}{
		{"minimal0,0", "test_data/minimal.b98", 0, 0, '@'},
		{"minimal-1,0", "test_data/minimal.b98", -1, 0, ' '},
		{"minimal1,0", "test_data/minimal.b98", 1, 0, ' '},
		{"minimal0,-1", "test_data/minimal.b98", 0, -1, ' '},
		{"minimal0,1", "test_data/minimal.b98", 0, 1, ' '},
		{"hello3,0", "test_data/hello.b98", 3, 0, '"'},
		{"hello3,1", "test_data/hello.b98", 3, 1, ' '},
		{"factorial0,1", "test_data/factorial.b98", 0, 1, ' '},
		{"factorial14,1", "test_data/factorial.b98", 14, 1, ' '},
		{"factorial15,1", "test_data/factorial.b98", 15, 1, ' '},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.input)
			require.NoError(t, err, "Failed to open file")
			defer file.Close()
			field, err := LoadFile(file)
			valid := field.Get(tc.inputX, tc.inputY)
			require.NoError(t, err)
			require.Equal(t, tc.expected, valid, "Invalid value")
		})
	}
}

func TestIsIn(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		input    string
		inputX   int
		inputY   int
		expected bool
	}{
		{"minimal0,0", "test_data/minimal.b98", 0, 0, true},
		{"minimal-1,0", "test_data/minimal.b98", -1, 0, false},
		{"minimal1,0", "test_data/minimal.b98", 1, 0, false},
		{"minimal0,-1", "test_data/minimal.b98", 0, -1, false},
		{"minimal0,1", "test_data/minimal.b98", 0, 1, false},
		{"hello3,0", "test_data/hello.b98", 3, 0, true},
		{"hello3,1", "test_data/hello.b98", 3, 1, false},
		{"factorial0,1", "test_data/factorial.b98", 0, 1, true},
		{"factorial14,1", "test_data/factorial.b98", 14, 1, true},
		{"factorial15,1", "test_data/factorial.b98", 15, 1, false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.input)
			require.NoError(t, err, "Failed to open file")
			defer file.Close()
			field, err := LoadFile(file)
			valid := field.isIn(tc.inputX, tc.inputY)
			require.NoError(t, err)
			require.Equal(t, tc.expected, valid, "Invalid value")
		})
	}
}

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
			field, err := LoadFile(file)
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
		lx: 5,
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
			field, err := LoadFile(file)
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
			field, err := LoadFile(file)
			field.Set(tc.inputX, tc.inputY, tc.inputV)
			require.NoError(t, err)
			require.Equal(t, tc.expected, field, "Invalid value")
		})
	}
}

// get and put on an empty line
func TestGetAndSetOnEmptyLines(t *testing.T) {
	f := Field{
		x:  0,
		y:  -4,
		lx: 1,
		ly: 8,
		lines: []Line{
			Line{x: 0, l: 1, columns: []int{'b'}},
			Line{},
			Line{},
			Line{x: 0, l: 1, columns: []int{'c'}},
			Line{x: 0, l: 1, columns: []int{'@'}},
			Line{},
			Line{x: 0, l: 1, columns: []int{'a'}},
			Line{x: 0, l: 1, columns: []int{'#'}},
		},
	}
	file, err := os.Open("test_data/minimal.b98")
	require.NoError(t, err, "Failed to open file")
	defer file.Close()
	field, err := LoadFile(file)
	field.Set(0, 3, '#')
	v := field.Get(0, 3)
	require.Equal(t, v, int('#'))
	v = field.Get(0, 2)
	require.Equal(t, v, int(' '))
	field.Set(0, 2, 'a')
	field.Set(0, -4, 'b')
	v = field.Get(0, -1)
	require.Equal(t, v, int(' '))
	field.Set(0, -1, 'c')
	require.Equal(t, field, &f)
}

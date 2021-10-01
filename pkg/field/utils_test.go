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
			field, err := Load(file)
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
			field, err := Load(file)
			valid := field.isIn(tc.inputX, tc.inputY)
			require.NoError(t, err)
			require.Equal(t, tc.expected, valid, "Invalid value")
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
	field, err := Load(file)
	field.Set(0, 3, '#')
	v := field.Get(0, 3)
	require.Equal(t, int('#'), v)
	v = field.Get(0, 2)
	require.Equal(t, int(' '), v)
	field.Set(0, 2, 'a')
	field.Set(0, -4, 'b')
	v = field.Get(0, -1)
	require.Equal(t, int(' '), v)
	field.Set(0, -1, 'c')
	require.Equal(t, &f, field)
}

func TestDump(t *testing.T) {
	// nothing to test really
	f := Field{}
	f.Dump()
}

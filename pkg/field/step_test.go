package field

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStep(t *testing.T) {
	// Test cases
	testCases := []struct {
		name  string
		input string
		x     int
		y     int
		dx    int
		dy    int
		ex    int // expectedX
		ey    int
	}{
		{"minimal0", "test_data/minimal.b98", 0, 0, 0, 0, 0, 0},
		{"minimal1", "test_data/minimal.b98", 0, 0, 1, 0, 0, 0},
		{"hello0", "test_data/hello.b98", 3, 0, 0, 0, 3, 0},
		{"hello1", "test_data/hello.b98", 3, 0, 1, 0, 4, 0},
		{"dna0", "test_data/dna.b98", 3, 3, 2, 1, 5, 4},
		{"dna1", "test_data/dna.b98", 6, 1, 1, 1, 5, 0},
		{"dna2", "test_data/dna.b98", 1, 4, -2, 2, 5, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.input)
			require.NoError(t, err, "Failed to open file")
			defer file.Close()
			field, err := LoadFile(file)
			x, y := field.Step(tc.x, tc.y, tc.dx, tc.dy)
			require.NoError(t, err)
			require.Equal(t, tc.ex, x, "Invalid x value")
			require.Equal(t, tc.ey, y, "Invalid y value")
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
			valid := field.IsIn(tc.inputX, tc.inputY)
			require.NoError(t, err)
			require.Equal(t, tc.expected, valid, "Invalid value")
		})
	}
}

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
			field, err := Load(file)
			x, y := field.Step(tc.x, tc.y, tc.dx, tc.dy)
			require.NoError(t, err)
			require.Equal(t, tc.ex, x, "Invalid x value")
			require.Equal(t, tc.ey, y, "Invalid y value")
		})
	}
}

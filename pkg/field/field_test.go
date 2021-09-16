package field

import (
	"io"
	"os"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/require"
)

func TestLoadFile(t *testing.T) {
	// minimal b98 file
	minimalField := Field{
		x:  0,
		y:  0,
		lx: 1,
		ly: 1,
		lines: []Line{
			Line{
				x:       0,
				l:       1,
				columns: []int{'@'},
			},
		},
	}
	// hello b98 file
	helloField := Field{
		x:  0,
		y:  0,
		lx: 24,
		ly: 1,
		lines: []Line{
			Line{
				x:       0,
				l:       24,
				columns: []int{'6', '4', '+', '"', '!', 'd', 'l', 'r', 'o', 'W', ' ', ',', 'o', 'l', 'l', 'e', 'H', '"', '>', ':', '#', ',', '_', '@'},
			},
		},
	}
	// factorial b98 file
	factorialField := Field{
		x:  0,
		y:  0,
		lx: 15,
		ly: 2,
		lines: []Line{
			Line{
				x:       0,
				l:       15,
				columns: []int{'&', '>', ':', '1', '-', ':', 'v', ' ', 'v', ' ', '*', '_', '$', '.', '@'},
			},
			Line{
				x:       1,
				l:       11,
				columns: []int{'^', ' ', ' ', ' ', ' ', '_', '$', '>', '\\', ':', '^'},
			},
		},
	}
	// dna b98 file
	dnaField := Field{
		x:  0,
		y:  0,
		lx: 7,
		ly: 8,
		lines: []Line{
			Line{
				x:       0,
				l:       7,
				columns: []int{'7', '^', 'D', 'N', '>', 'v', 'A'},
			},
			Line{
				x:       0,
				l:       7,
				columns: []int{'v', '_', '#', 'v', '?', ' ', 'v'},
			},
			Line{
				x:       0,
				l:       7,
				columns: []int{'7', '^', '<', '"', '"', '"', '"'},
			},
			Line{
				x:       0,
				l:       7,
				columns: []int{'3', ' ', ' ', 'A', 'C', 'G', 'T'},
			},
			Line{
				x:       0,
				l:       7,
				columns: []int{'9', '0', '!', '"', '"', '"', '"'},
			},
			Line{
				x:       0,
				l:       7,
				columns: []int{'4', '*', ':', '>', '>', '>', 'v'},
			},
			Line{
				x:       0,
				l:       7,
				columns: []int{'+', '8', '^', '-', '1', ',', '<'},
			},
			Line{
				x:       0,
				l:       7,
				columns: []int{'>', ' ', ',', '+', ',', '@', ')'},
			},
		},
	}
	// \r\n file b98 file
	rnField := Field{
		x:  0,
		y:  0,
		lx: 24,
		ly: 1,
		lines: []Line{
			Line{
				x:       0,
				l:       24,
				columns: []int{'6', '4', '+', '"', '!', 'd', 'l', 'r', 'o', 'W', ' ', ',', 'o', 'l', 'l', 'e', 'H', '"', '>', ':', '#', ',', '_', '@'},
			},
		},
	}
	// Test cases
	type addError func(r io.Reader) io.Reader
	testCases := []struct {
		name          string
		input         string
		inputAddError addError
		expected      *Field
		expectedError error
	}{
		{"Empty file", "test_data/empty.b98", nil, nil, &DecodeError{}},
		{"Invalid file content", "test_data/invalid.b98", nil, nil, &DecodeError{}},
		{"io error", "test_data/minimal.b98", iotest.TimeoutReader, nil, &ReadError{}},
		{"minimal", "test_data/minimal.b98", nil, &minimalField, nil},
		{"hello", "test_data/hello.b98", nil, &helloField, nil},
		{"factorial", "test_data/factorial.b98", nil, &factorialField, nil},
		{"dna", "test_data/dna.b98", nil, &dnaField, nil},
		{"\\r\\n file", "test_data/rn.b98", nil, &rnField, nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var fd io.Reader
			file, err := os.Open(tc.input)
			require.NoError(t, err, "Failed to open file")
			defer file.Close()
			if tc.inputAddError != nil {
				fd = tc.inputAddError(file)
			} else {
				fd = file
			}
			valid, err := LoadFile(fd)
			if tc.expectedError != nil {
				require.Error(t, err)
				requireErrorTypeMatch(t, err, tc.expectedError)
				require.Nil(t, valid)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expected, valid, "Invalid value")
		})
	}
}

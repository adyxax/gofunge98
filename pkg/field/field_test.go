package field

import (
	"io"
	"os"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadFile(t *testing.T) {
	// minimal b98 file
	minimalField := Field{
		firstLineIndex: 0,
		length:         1,
		lines: []Line{
			Line{
				firstColumnIndex: 0,
				length:           1,
				columns:          []byte{'@'},
			},
		},
	}
	// hello b98 file
	helloField := Field{
		firstLineIndex: 0,
		length:         1,
		lines: []Line{
			Line{
				firstColumnIndex: 0,
				length:           24,
				columns:          []byte{'6', '4', '+', '"', '!', 'd', 'l', 'r', 'o', 'W', ' ', ',', 'o', 'l', 'l', 'e', 'H', '"', '>', ':', '#', ',', '_', '@'},
			},
		},
	}
	// factorial b98 file
	factorialField := Field{
		firstLineIndex: 0,
		length:         2,
		lines: []Line{
			Line{
				firstColumnIndex: 0,
				length:           15,
				columns:          []byte{'&', '>', ':', '1', '-', ':', 'v', ' ', 'v', ' ', '*', '_', '$', '.', '@'},
			},
			Line{
				firstColumnIndex: 1,
				length:           11,
				columns:          []byte{'^', ' ', ' ', ' ', ' ', '_', '$', '>', '\\', ':', '^'},
			},
		},
	}
	// dna b98 file
	dnaField := Field{
		firstLineIndex: 0,
		length:         8,
		lines: []Line{
			Line{
				firstColumnIndex: 0,
				length:           7,
				columns:          []byte{'7', '^', 'D', 'N', '>', 'v', 'A'},
			},
			Line{
				firstColumnIndex: 0,
				length:           7,
				columns:          []byte{'v', '_', '#', 'v', '?', ' ', 'v'},
			},
			Line{
				firstColumnIndex: 0,
				length:           7,
				columns:          []byte{'7', '^', '<', '"', '"', '"', '"'},
			},
			Line{
				firstColumnIndex: 0,
				length:           7,
				columns:          []byte{'3', ' ', ' ', 'A', 'C', 'G', 'T'},
			},
			Line{
				firstColumnIndex: 0,
				length:           7,
				columns:          []byte{'9', '0', '!', '"', '"', '"', '"'},
			},
			Line{
				firstColumnIndex: 0,
				length:           7,
				columns:          []byte{'4', '*', ':', '>', '>', '>', 'v'},
			},
			Line{
				firstColumnIndex: 0,
				length:           7,
				columns:          []byte{'+', '8', '^', '-', '1', ',', '<'},
			},
			Line{
				firstColumnIndex: 0,
				length:           7,
				columns:          []byte{'>', ' ', ',', '+', ',', '@', ')'},
			},
		},
	}
	// \r\n file b98 file
	rnField := Field{
		firstLineIndex: 0,
		length:         1,
		lines: []Line{
			Line{
				firstColumnIndex: 0,
				length:           24,
				columns:          []byte{'6', '4', '+', '"', '!', 'd', 'l', 'r', 'o', 'W', ' ', ',', 'o', 'l', 'l', 'e', 'H', '"', '>', ':', '#', ',', '_', '@'},
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
			assert.Equal(t, tc.expected, valid, "Invalid value")
		})
	}
}

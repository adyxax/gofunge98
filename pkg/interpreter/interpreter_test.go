package interpreter

import (
	"os"
	"testing"

	"git.adyxax.org/adyxax/gofunge98/pkg/field"
	"git.adyxax.org/adyxax/gofunge98/pkg/pointer"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	file, err := os.Open("../field/test_data/minimal.b98")
	require.NoError(t, err)
	defer file.Close()
	f, err := field.Load(file)
	require.NoError(t, err)
	NewInterpreter(f, pointer.NewPointer()).Run()
}

func TestStep(t *testing.T) {
	testCases := []struct {
		name            string
		filename        string
		pointer         *pointer.Pointer
		expectedField   *field.Field
		expectedPointer *pointer.Pointer
	}{
		{"minimal", "../field/test_data/minimal.b98", nil, nil, nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.filename)
			require.NoError(t, err)
			defer file.Close()
			f, err := field.Load(file)
			require.NoError(t, err)
			if tc.pointer == nil {
				tc.pointer = pointer.NewPointer()
			}
			NewInterpreter(f, tc.pointer).step()
			if tc.expectedField != nil {
				require.Equal(t, tc.expectedField, f)
			}
			if tc.expectedPointer != nil {
				require.Equal(t, tc.expectedPointer, tc.pointer)
			}
		})
	}
}

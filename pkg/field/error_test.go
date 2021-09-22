package field

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func requireErrorTypeMatch(t *testing.T, err error, expected error) {
	require.Equalf(t, reflect.TypeOf(expected), reflect.TypeOf(err), "Invalid error type. Got %s but expected %s", reflect.TypeOf(err), reflect.TypeOf(expected))
}

func TestErrorsCoverage(t *testing.T) {
	readErr := ReadError{}
	_ = readErr.Error()
	_ = readErr.Unwrap()
	decodeError := DecodeError{}
	_ = decodeError.Error()
}

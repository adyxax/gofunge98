package field

import "fmt"

// Read error
type ReadError struct {
	err error
}

func (e ReadError) Error() string {
	return fmt.Sprintf("Failed to decode file")
}
func (e ReadError) Unwrap() error { return e.err }

func newReadError(err error) error {
	return &ReadError{
		err: err,
	}
}

// Funge decoding error
type DecodeError struct {
	msg string
}

func (e DecodeError) Error() string {
	return fmt.Sprintf("Failed to decode file with message : %s", e.msg)
}

func newDecodeError(msg string) error {
	return &DecodeError{
		msg: msg,
	}
}

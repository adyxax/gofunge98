package pointer

import (
	"fmt"
	"os"
)

// We keep the last input char to handle a pain point in the funge-98 spec :
// when reading from the decimal input you need to read until you encounter a
// non numeric char, but not drop it
var defaultInputLastChar *int = nil

func DefaultCharacterInput() (int, error) {
	if defaultInputLastChar != nil {
		c := *defaultInputLastChar
		defaultInputLastChar = nil
		return c, nil
	}
	b := make([]byte, 1)
	i, err := os.Stdin.Read(b)
	if err != nil {
		return 0, fmt.Errorf("Error in DefaultCharacterInput { b: %c, i: %d, err: %w }", b[0], i, err)
	}
	return int(b[0]), nil
}

func DefaultDecimalInput() (int, error) {
	var v int
	for {
		c, err := DefaultCharacterInput()
		if err != nil {
			return 0, err
		}
		if c >= '0' && c <= '9' {
			v = c - '0'
			break
		}
	}
	for {
		c, err := DefaultCharacterInput()
		if err != nil {
			break
		}
		if c >= '0' && c <= '9' {
			v = v*10 + c - '0'
		} else {
			defaultInputLastChar = &c
			break
		}
	}
	return v, nil
}

func DefaultCharacterOutput(c int) {
	fmt.Printf("%c", c)
}

func DefaultDecimalOutput(c int) {
	fmt.Printf("%d ", c)
}

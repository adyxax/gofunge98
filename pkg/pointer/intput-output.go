package pointer

import (
	"fmt"
	"log"
	"os"
)

var defaultInputLastChar *int = nil

func DefaultCharacterInput() int {
	if defaultInputLastChar != nil {
		c := *defaultInputLastChar
		defaultInputLastChar = nil
		return c
	}
	b := make([]byte, 1)
	i, err := os.Stdin.Read(b)
	if err != nil {
		log.Fatalf("Error in DefaultCharacterInput { b: %c, i: %d, err: %+v }", b[0], i, err)
	}
	return int(b[0])
}

func DefaultDecimalInput() int {
	var v int
	for {
		c := DefaultCharacterInput()
		if c >= '0' && c <= '9' {
			v = c - '0'
			break
		}
	}
	for {
		c := DefaultCharacterInput()
		if c >= '0' && c <= '9' {
			v = v*10 + c - '0'
		} else {
			defaultInputLastChar = &c
			break
		}
	}
	return v
}

func DefaultCharacterOutput(c int) {
	fmt.Printf("%c", c)
}

func DefaultDecimalOutput(c int) {
	fmt.Printf("%d ", c)
}

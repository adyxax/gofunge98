package field

import (
	"bytes"
	"io"
)

//            32-bit Befunge-98
//           =================
//                   |-2,147,483,648
//                   |
//                   |    x
//              -----+-----
// -2,147,483,648    |    2,147,483,647
//                   |
//                  y|2,147,483,647

type Field struct {
	firstLineIndex int
	length         int
	lines          []Line
}

type Line struct {
	firstColumnIndex int
	length           int
	columns          []byte
}

func LoadFile(fd io.Reader) (*Field, error) {
	f := new(Field)
	l := new(Line)
	trailingSpaces := 0
	for {
		data := make([]byte, 4096)
		if n, errRead := fd.Read(data); errRead != nil {
			if errRead == io.EOF {
				if f.length == 0 && l.length == 0 {
					return nil, newDecodeError("No instruction on the first line of the file produces an unusable program in Befunge98")
				}
				if l.length > 0 {
					f.length++
					f.lines = append(f.lines, *l)
				}
				break
			} else {
				return nil, newReadError(errRead)
			}
		} else {
			for i := 0; i < n; i++ {
				if data[i] == '\n' || data[i] == '\r' {
					if f.length == 0 && l.length == 0 {
						return nil, newDecodeError("No instruction on the first line of the file produces an unusable program in Befunge98")
					}
					f.length++
					f.lines = append(f.lines, *l)
					l = new(Line)
					trailingSpaces = 0
					if i+1 < n && data[i] == '\r' && data[i+1] == '\n' {
						i++
					}
				} else {
					if l.length == 0 && data[i] == ' ' {
						l.firstColumnIndex++ // trim leading spaces
					} else {
						if data[i] == ' ' {
							trailingSpaces++
						} else {
							l.columns = append(l.columns, bytes.Repeat([]byte{' '}, trailingSpaces)...)
							l.length += trailingSpaces
							trailingSpaces = 0
							l.length++
							l.columns = append(l.columns, data[i])
						}
					}
				}
			}
		}
	}
	return f, nil
}

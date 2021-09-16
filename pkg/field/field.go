package field

import (
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
	x     int
	y     int
	lx    int
	ly    int
	lines []Line
}

type Line struct {
	x       int
	l       int
	columns []int
}

func LoadFile(fd io.Reader) (*Field, error) {
	f := new(Field)
	l := new(Line)
	trailingSpaces := 0
	for {
		data := make([]byte, 4096)
		if n, errRead := fd.Read(data); errRead != nil {
			if errRead == io.EOF {
				if f.ly == 0 && l.l == 0 {
					return nil, newDecodeError("No instruction on the first line of the file produces an unusable program in Befunge98")
				}
				if l.l > 0 {
					f.ly++
					if f.lx < l.l {
						f.lx = l.l
					}
					f.lines = append(f.lines, *l)
				}
				break
			} else {
				return nil, newReadError(errRead)
			}
		} else {
			for i := 0; i < n; i++ {
				if data[i] == '\n' || data[i] == '\r' {
					if f.ly == 0 && l.l == 0 {
						return nil, newDecodeError("No instruction on the first line of the file produces an unusable program in Befunge98")
					}
					f.ly++
					if f.lx < l.l {
						f.lx = l.l
					}
					f.lines = append(f.lines, *l)
					l = new(Line)
					trailingSpaces = 0
					if i+1 < n && data[i] == '\r' && data[i+1] == '\n' {
						i++
					}
				} else {
					if l.l == 0 && data[i] == ' ' {
						l.x++ // trim leading spaces
					} else {
						if data[i] == ' ' {
							trailingSpaces++
						} else {
							for j := 0; j < trailingSpaces; j++ {
								l.columns = append(l.columns, ' ')
							}
							l.l += trailingSpaces + 1
							trailingSpaces = 0
							l.columns = append(l.columns, int(data[i]))
						}
					}
				}
			}
		}
	}
	return f, nil
}

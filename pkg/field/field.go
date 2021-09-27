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

func Load(fd io.Reader) (*Field, error) {
	f := Field{
		lx:    1,
		ly:    1,
		lines: []Line{Line{l: 1, columns: []int{'>'}}},
	}
	var x, y int
	for {
		data := make([]byte, 4096)
		if n, errRead := fd.Read(data); errRead != nil {
			if errRead == io.EOF {
				break
			} else {
				return nil, newReadError(errRead)
			}
		} else {
			for i := 0; i < n; i++ {
				if data[i] == '' {
					continue
				}
				if data[i] == '\n' || data[i] == '\r' {
					x = 0
					y++
					if i+1 < n && data[i] == '\r' && data[i+1] == '\n' {
						i++
					}
				} else {
					f.Set(x, y, int(data[i]))
					x++
				}
			}
		}
	}
	if f.x == 0 && f.lx == 1 && f.lines[0].columns[0] == '>' {
		return nil, newDecodeError("No instruction on the first line of the file produces an unusable program in Befunge98")
	}
	return &f, nil
}

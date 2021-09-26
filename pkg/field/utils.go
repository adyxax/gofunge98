package field

func (f Field) Get(x, y int) int {
	if y >= f.y && y < f.y+f.ly {
		l := f.lines[y-f.y]
		if x >= l.x && x < l.x+l.l {
			return l.columns[x-l.x]
		}
	}
	return ' '
}

func (f Field) isIn(x, y int) bool {
	return x >= f.x && x < f.x+f.lx && y >= f.y && y < f.y+f.ly
}

func (f *Field) Set(x, y, v int) {
	if v == ' ' {
		f.Blank(x, y)
		return
	}
	if y >= f.y {
		if y < f.y+f.ly {
			l := &f.lines[y-f.y]
			if l.l == 0 {
				// an empty line is a special case
				l.x = x
				l.l = 1
				l.columns = append(l.columns, v)
			} else if x >= l.x {
				if x < l.x+l.l {
					// just set the value
					l.columns[x-l.x] = v
				} else {
					// append columns
					newL := x - l.x + 1
					for i := l.l; i < newL-1; i++ {
						l.columns = append(l.columns, ' ')
					}
					l.columns = append(l.columns, v)
					l.l = newL
					if f.lx-f.x < l.l-l.x {
						f.lx = l.l - l.x
					}
				}
			} else {
				// prepend columns
				newL := l.l + l.x - x
				c := make([]int, newL)
				c[0] = v
				for i := 1; i < l.x-x; i++ {
					c[i] = ' '
				}
				for i := l.x - x; i < newL; i++ {
					c[i] = l.columns[i-l.x+x]
				}
				l.columns = c
				l.x = x
				l.l = newL
				if f.x > x {
					f.lx = f.lx + f.x - x
					f.x = x
				}
			}
		} else {
			// append lines
			newLy := y - f.y + 1
			for i := f.ly; i < newLy-1; i++ {
				f.lines = append(f.lines, Line{})
			}
			f.lines = append(f.lines, Line{x: x, l: 1, columns: []int{v}})
			f.ly = newLy
			if f.x > x {
				f.lx = f.lx + f.x - x
				f.x = x
			}
			if f.lx-f.x < x {
				f.lx = x - f.x
			}
		}
	} else {
		// prepend lines
		newLy := f.ly + f.y - y
		lines := make([]Line, newLy)
		lines[0] = Line{x: x, l: 1, columns: []int{v}}
		for j := f.y - y; j < newLy; j++ {
			lines[j] = f.lines[j-f.y+y]
		}
		f.lines = lines
		f.y = y
		f.ly = newLy
		if f.x > x {
			f.lx = f.lx + f.x - x
			f.x = x
		}
		if f.lx-f.x < x {
			f.lx = x - f.x + 1
		}
	}
}

func (f Field) Dump() (int, int, int, int) {
	return f.x, f.y, f.lx, f.ly
}

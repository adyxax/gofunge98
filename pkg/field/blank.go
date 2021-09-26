package field

func (f *Field) Blank(x, y int) {
	if !f.isIn(x, y) {
		// outside the field, nothing to do
		return
	}
	l := &f.lines[y-f.y]
	if x < l.x || x >= l.x+l.l {
		// outside the current line's columns
		return
	}
	if x > l.x && x < l.x+l.l-1 {
		// just set the value
		l.columns[x-l.x] = ' '
		return
	}
	// do we need to trim this line?
	if l.l == 1 {
		// this was the last character on the line
		if y == f.y {
			// We need to trim the leading lines
			leadingLines := 1
			for i := 1; i < f.ly; i++ {
				if f.lines[i].l == 0 {
					leadingLines++
				} else {
					break
				}
			}
			f.y += leadingLines
			f.ly -= leadingLines
			lines := make([]Line, f.ly)
			for i := 0; i < f.ly; i++ {
				lines[i] = f.lines[leadingLines+i]
			}
			f.lines = lines
		} else if y == f.y+f.ly-1 {
			// We need to trim the trailing lines
			trailingLines := 1
			for i := f.ly + f.y - 2; i >= 0; i-- {
				if f.lines[i].l == 0 {
					trailingLines++
				} else {
					break
				}
			}
			f.ly -= trailingLines
			f.lines = f.lines[:f.ly]
		} else {
			// it was a line in the middle
			l.l = 0
			l.columns = make([]int, 0)
		}
	} else if x == l.x {
		// We need to remove leading spaces
		leadingSpaces := 1
		for i := 1; i < l.l; i++ {
			if l.columns[i] == ' ' {
				leadingSpaces++
			} else {
				break
			}
		}
		l.x += leadingSpaces
		l.l -= leadingSpaces
		columns := make([]int, l.l)
		for i := 0; i < l.l; i++ {
			columns[i] = l.columns[leadingSpaces+i]
		}
		l.columns = columns
	} else if x == l.l+l.x-1 {
		// we need to remove trailing spaces
		trailingSpaces := 1
		for i := l.l - 2; i >= 0; i-- {
			if l.columns[i] == ' ' {
				trailingSpaces++
			} else {
				break
			}
		}
		l.l -= trailingSpaces
		l.columns = l.columns[:l.l]
	}
	// we now need to find the new field limits
	f.x = f.lines[0].x
	x2 := f.lines[0].l + f.lines[0].x
	for i := 1; i < f.ly; i++ {
		if f.x > f.lines[i].x {
			f.x = f.lines[i].x
		}
		if x2 < f.lines[i].x+f.lines[i].l {
			x2 = f.lines[i].x + f.lines[i].l
		}
	}
	f.lx = x2 - f.x
}

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

func (f Field) Dump() (int, int, int, int) {
	return f.x, f.y, f.lx, f.ly
}

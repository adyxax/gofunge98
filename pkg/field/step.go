package field

func (f Field) Step(x, y, dx, dy int) (int, int) {
	x2, y2 := x+dx, y+dy
	if f.IsIn(x2, y2) {
		return x2, y2
	}
	// We are stepping outside, we need to wrap the Lahey-space
	for {
		x2, y2 := x-dx, y-dy
		if !f.IsIn(x2, y2) {
			return x, y
		}
		x, y = x2, y2
	}
}

func (f Field) IsIn(x, y int) bool {
	return x >= f.x && x < f.x+f.lx && y >= f.y && y < f.y+f.ly
}

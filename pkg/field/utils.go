package field

func (f Field) isIn(x, y int) bool {
	return x >= f.x && x < f.x+f.lx && y >= f.y && y < f.y+f.ly
}

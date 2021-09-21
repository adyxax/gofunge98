package pointer

func (p Pointer) GetStorageOffset() (x, y int) {
	return p.sox, p.soy
}

func (p *Pointer) CalculateNewStorageOffset() {
	p.sox, p.soy = p.x+p.dx, p.y+p.dy
}

func (p *Pointer) SetStorageOffset(x, y int) {
	p.sox, p.soy = x, y
}

package pointer

type Delta struct {
	x int
	y int
}

func NewDelta(x, y int) *Delta {
	return &Delta{x: x, y: y}
}

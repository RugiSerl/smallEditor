package math

type Vec2i struct {
	X, Y int
}

func NewVec2i(x int, y int) Vec2i {
	return Vec2i{X: x, Y: y}
}

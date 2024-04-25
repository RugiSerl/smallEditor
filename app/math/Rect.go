package math

import rl "github.com/gen2brain/raylib-go/raylib"

type Rect struct {
	Position, Size Vec2f
}

func NewRect(position, size Vec2f) Rect {
	return Rect{Position: position, Size: size}
}

func (r Rect) PointCollision(v Vec2f) bool {
	return v.X > r.Position.X && v.X <= r.Position.X+r.Size.X && v.Y > r.Position.Y && v.Y <= r.Position.Y+r.Size.Y
}

func (r Rect) ToRL() rl.Rectangle {
	return rl.NewRectangle(float32(r.Position.X), float32(r.Position.Y), float32(r.Size.X), float32(r.Size.Y))
}

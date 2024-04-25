package math

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Vec2f struct {
	X, Y float64
}

func (v Vec2f) ToRL() rl.Vector2 {
	return rl.NewVector2(float32(v.X), float32(v.Y))
}

func FromRL(vec rl.Vector2) Vec2f {
	return NewVec2(float64(vec.X), float64(vec.Y))
}

func NewVec2(x float64, y float64) Vec2f {
	return Vec2f{X: x, Y: y}
}

// ---------------
// linear function
func (v Vec2f) Add(v2 Vec2f) Vec2f {
	return Vec2f{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func (v Vec2f) Substract(v2 Vec2f) Vec2f {
	return Vec2f{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v Vec2f) Scale(scalar float64) Vec2f {
	return Vec2f{X: v.X * scalar, Y: v.Y * scalar}
}

func (v Vec2f) GetNorm() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// null vector unhandled
func (v Vec2f) Normalize() Vec2f {
	return v.Scale(1 / v.GetNorm())
}

// null vector unhandled
func (v Vec2f) ScaleToNorm(norm float64) Vec2f {
	return v.Normalize().Scale(norm)
}

// ------------------
// rotation functions
// trigonometric circle
func (v Vec2f) GetAngle() float64 {
	a := math.Acos(v.Normalize().X)
	if v.Y > 0 {
		return a
	} else {
		return -a
	}
}

func (v Vec2f) Rotate(angle float64) Vec2f {
	x := math.Cos(angle)*v.X - math.Sin(angle)*v.Y
	y := math.Cos(angle)*v.Y + math.Sin(angle)*v.X
	return Vec2f{X: x, Y: y}
}

package input

import (
	"github.com/RugiSerl/smallEditor/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MouseButton int

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonRight
	MouseButtonMiddle
	MouseButtonSide
	MouseButtonExtra
	MouseButtonForward
	MouseButtonBack
)

func GetMousePosition() math.Vec2f {
	return math.FromRL(rl.GetMousePosition())
}

func GetMouseDelta() math.Vec2f {
	return math.FromRL(rl.GetMouseDelta())
}

func IsMouseClicked(b MouseButton) bool {
	return rl.IsMouseButtonPressed(int32(b))
}
func IsMouseDown(b MouseButton) bool {
	return rl.IsMouseButtonDown(int32(b))
}

func MouseWheelDelta() float64 {
	return float64(rl.GetMouseWheelMove())

}

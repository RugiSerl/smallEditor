package graphic

import (
	"image/color"

	"github.com/RugiSerl/smallEditor/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ROUNDED_SEGMENTS = 10
)

func DrawRect(rectangle math.Rect, color color.RGBA) {
	rl.DrawRectangleV(rectangle.Position.ToRL(), rectangle.Size.ToRL(), color)
}

func DrawRectRounded(rectangle math.Rect, roundness float64, color color.RGBA) {
	rl.DrawRectangleRounded(rectangle.ToRL(), float32(roundness), ROUNDED_SEGMENTS, color)
}

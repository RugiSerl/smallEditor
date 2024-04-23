package graphic

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GetDeltaTime() float64 {
	return float64(rl.GetFrameTime())

}
func ClearBackground(color color.RGBA) {
	rl.ClearBackground(color)
}

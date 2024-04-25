package graphic

import (
	"github.com/RugiSerl/smallEditor/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Texture struct {
	rl.Texture2D
}

func NewTexture(path string) Texture {
	return Texture{rl.LoadTexture(path)}
}

func DrawTexture(texture Texture, position math.Vec2f) {
	rl.DrawTextureV(texture.Texture2D, position.ToRL(), rl.White)

}

func DrawTextureRect(texture Texture, sourceRect math.Rect, destRect math.Rect, origin math.Vec2f, rotation float32) {
	rl.DrawTexturePro(texture.Texture2D, sourceRect.ToRL(), destRect.ToRL(), origin.ToRL(), rotation, rl.White)
}

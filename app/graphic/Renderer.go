package graphic

import (
	"github.com/RugiSerl/smallEditor/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Texture used to render things on it
type Renderer struct {
	rl.RenderTexture2D
	Size          math.Vec2
	QualityFactor float64
}

func NewRenderTexture(size math.Vec2, qualityFactor float64) *Renderer {

	return &Renderer{rl.LoadRenderTexture(int32(size.X*qualityFactor), int32(size.Y*qualityFactor)), size, qualityFactor}
}

func (r *Renderer) Resize(newSize math.Vec2) *Renderer {
	rl.UnloadRenderTexture(r.RenderTexture2D)
	return NewRenderTexture(newSize, r.QualityFactor)
}

func (r *Renderer) GetTexture() Texture {
	return Texture{r.Texture}
}

func (r *Renderer) Draw(position math.Vec2) {
	DrawTextureRect(Texture{r.Texture}, math.NewRect(math.NewVec2(0, 0), math.NewVec2(r.Size.X, -r.Size.Y).Scale(r.QualityFactor)), math.NewRect(position, r.Size), math.NewVec2(0, 0), 0)
}

func (r *Renderer) Begin() {

	rl.BeginTextureMode(r.RenderTexture2D)
	rl.ClearBackground(rl.NewColor(0, 0, 0, 0))
}

func (r *Renderer) End() {
	rl.EndTextureMode()
}

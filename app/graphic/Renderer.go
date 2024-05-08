package graphic

import (
	"github.com/RugiSerl/smallEditor/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Texture used to render things on it
type Renderer struct {
	rl.RenderTexture2D
	Size          math.Vec2f
	QualityFactor float64
}

// Returns a new instance of Renderer
func NewRenderTexture(size math.Vec2f, qualityFactor float64) *Renderer {
	r := &Renderer{rl.LoadRenderTexture(int32(size.X*qualityFactor), int32(size.Y*qualityFactor)), size, qualityFactor}
	rl.SetTextureFilter(r.Texture, rl.TextureFilterLinear)
	return r

}

// Resize the renderer by creating a new one and disposing of the last one to avoid memory leaks
func (r *Renderer) Resize(newSize math.Vec2f) *Renderer {
	rl.UnloadRenderTexture(r.RenderTexture2D)
	return NewRenderTexture(newSize, r.QualityFactor)
}

// Get the texture of the renderer
func (r *Renderer) GetTexture() Texture {
	return Texture{r.Texture}
}

// Draw the renderer
// NOTE: The texture has been flipped upside down because for some reason it is upside down when displaying normally
func (r *Renderer) Draw(position math.Vec2f) {
	DrawTextureRect(r.GetTexture(), math.NewRect(math.NewVec2(0, 0), math.NewVec2(r.Size.X, -r.Size.Y).Scale(r.QualityFactor)), math.NewRect(position, r.Size), math.NewVec2(0, 0), 0)
}

// Begin renderer mode
func (r *Renderer) Begin() {
	rl.BeginTextureMode(r.RenderTexture2D)
	rl.ClearBackground(rl.NewColor(0, 0, 0, 0)) // Get rid of what was before on the renderer
}

// End renderer mode
func (r *Renderer) End() {
	rl.EndTextureMode()
}

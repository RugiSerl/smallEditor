package graphic

import (
	"math"

	m "github.com/RugiSerl/smallEditor/app/math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CAMERA_SPEED  = 1000
	ZOOM_AMOUNT   = 0.3
	CAMERA_SMOOTH = .05
	ZOOM_SMOOTH   = .05
	INITIAL_ZOOM  = 1.5
)

type Camera2D struct {
	rl.Camera2D
	targetPosition        m.Vec2
	logarithmicTargetZoom float32 // Here we take a logarithmic zoom, which is modified linearly by the user. This value will then be plugged in exp() to get the true zoom
	logarithmicSmoothZoom float32 // logarithmicSmoothZoom "follows" logarithmicTargetZoom smoothly
}

func NewCamera() *Camera2D {
	c := new(Camera2D)
	c.Camera2D = rl.NewCamera2D(rl.NewVector2(0, 0), rl.NewVector2(0, 0), 0, 1)
	c.targetPosition = m.NewVec2(0, 0)

	c.logarithmicTargetZoom = INITIAL_ZOOM
	c.logarithmicSmoothZoom = 0.5 // to create an animation at the beginning

	return c
}

func (c *Camera2D) UpdateCamera(boundingBox m.Rect) {
	c.updatePosition()
	c.updateZoom()

}

func (c *Camera2D) updatePosition() {

	c.Target = m.FromRL(c.Target).Add(c.targetPosition.Substract(m.FromRL(c.Target)).Scale(float64(rl.GetFrameTime()) / CAMERA_SMOOTH)).ToRL()
	// décalage de la caméra, pour que la cible, c'est-à-dire les coordonnées de la caméra, se trouve au milieu de l'écran
	c.Offset = rl.NewVector2(0, 0)

}
func (c *Camera2D) updateZoom() {

	c.logarithmicSmoothZoom += (c.logarithmicTargetZoom - c.logarithmicSmoothZoom) / ZOOM_SMOOTH * float32(GetDeltaTime())

	c.Zoom = float32(math.Exp(float64(c.logarithmicSmoothZoom)))
}

func (c *Camera2D) UpdateZoomInput() {
	c.logarithmicTargetZoom += rl.GetMouseWheelMove() * ZOOM_AMOUNT

}
func (c *Camera2D) UpdateMoveInput() {
	var speed float64 = CAMERA_SPEED * float64(rl.GetFrameTime()) / float64(c.Zoom)

	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		c.targetPosition.X -= speed
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		c.targetPosition.X += speed
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		c.targetPosition.Y -= speed
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		c.targetPosition.Y += speed
	}
}

func (c *Camera2D) ConvertToWorldCoordinates(coordinates m.Vec2) m.Vec2 {
	return coordinates.Substract(m.FromRL(c.Offset)).Scale(1 / float64(c.Zoom)).Add(m.FromRL(c.Target))
}

func (c *Camera2D) Begin() {
	rl.BeginMode2D(c.Camera2D)
}

func (c *Camera2D) End() {
	rl.EndMode2D()
}

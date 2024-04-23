package component

import (
	"image/color"

	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/input"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/settings"
	"github.com/RugiSerl/smallEditor/app/ui/component/container"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

const SCALE = 1.0

const (
	WINDOW_BAR_LENGTH = 35 // px
)

type windowState int

const (
	ANCHORED windowState = iota
	FREE
)

type Window struct {
	Container   container.Resizable
	State       windowState
	closeButton *ImageButton
	Content     *graphic.Renderer
	Padding     float64
	Closed      bool
}

func NewWindow(pos utils.RelativePosition, size math.Vec2, state windowState, qualityFactor float64) *Window {
	w := new(Window)
	w.Container = container.NewResizable(utils.RelativeRect{Position: pos, Size: size})
	w.State = state
	w.Padding = 5 * SCALE
	w.Closed = false
	w.Content = graphic.NewRenderTexture(w.GetRendererSize(size), qualityFactor)
	w.closeButton = NewImageButton("assets/exit.png", utils.RelativeRect{
		Position: utils.RelativePosition{
			HorizontalAnchor: utils.ANCHOR_RIGHT,
			VerticalAnchor:   utils.ANCHOR_TOP,
			Vec2:             math.NewVec2(0, 0),
		},
		Size: math.NewVec2(SCALE*WINDOW_BAR_LENGTH, SCALE*WINDOW_BAR_LENGTH),
	})
	return w
}

func (w *Window) Update(boundingBox math.Rect) {
	windowAbsoluteRect := w.GetWindowRect(boundingBox)
	// Resize the renderer if necessary
	if w.GetRendererSize(windowAbsoluteRect.Size) != w.Content.Size {
		w.Resize(w.GetRendererSize(windowAbsoluteRect.Size))
	}

	barRect := math.NewRect(windowAbsoluteRect.Position, math.NewVec2(windowAbsoluteRect.Size.X, WINDOW_BAR_LENGTH))
	w.renderWindow(windowAbsoluteRect, barRect)
	w.handleWindowMovement(boundingBox, barRect)

	if w.closeButton.Clicked {
		w.Closed = true
	}

}

func (w *Window) renderWindow(windowRect math.Rect, barRect math.Rect) {
	graphic.DrawRectRounded(windowRect, 0.03, settings.SettingInstance.Theme.WindowTheme.BackgroundColor)
	w.Content.Draw(w.GetRendererPosition(windowRect))
	w.closeButton.Update(windowRect)

	graphic.DrawRect(barRect, color.RGBA{255, 255, 255, 5})

}

func (w *Window) handleWindowMovement(boundingBox math.Rect, barRect math.Rect) {
	if w.State == FREE {
		w.Container.UpdateResize(boundingBox)
		if barRect.PointCollision(input.GetMousePosition()) && !w.closeButton.Hovered && !w.Container.Resizing {
			w.Container.UpdateDrag(boundingBox)
		}

	}
	if barRect.PointCollision(input.GetMousePosition()) || w.Container.Hovering {
		if w.Container.HandleDoubleClick(boundingBox) {
			if w.State == FREE { // junky code (no real choice)
				w.State = ANCHORED
			} else {
				w.State = FREE
			}

		}
	}
}

func (w *Window) GetRendererSize(windowSize math.Vec2) math.Vec2 {
	return windowSize.Substract(math.NewVec2(w.Padding*2, WINDOW_BAR_LENGTH+w.Padding*2))
}

func (w *Window) GetRendererPosition(windowAbsoluteRect math.Rect) math.Vec2 {
	return windowAbsoluteRect.Position.Add(math.NewVec2(w.Padding, WINDOW_BAR_LENGTH+w.Padding))
}

func (w *Window) GetWindowRect(boundingBox math.Rect) math.Rect {
	var windowAbsoluteRect math.Rect // rect with absolute coordinates
	if w.State == FREE {
		windowAbsoluteRect = w.Container.GetAbsoluteRect(boundingBox)
	} else {
		windowAbsoluteRect = boundingBox
	}
	return windowAbsoluteRect
}

func (w *Window) Resize(newSize math.Vec2) {
	w.Content = w.Content.Resize(newSize)
}

func (w *Window) ConvertPositionToRenderer(v math.Vec2, boundingBox math.Rect) math.Vec2 {
	return v.
		Substract(w.GetRendererPosition(w.GetWindowRect(boundingBox))). // substract the position of the renderer
		Scale(w.Content.QualityFactor)                                  // unstretch what was stretched by the QualityFactor
}

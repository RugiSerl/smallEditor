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

const (
	// Height of the bar at the top of the window
	WINDOW_BAR_LENGTH = 35 // px
)

// Enums of the type of the window
type windowState int

const (
	ANCHORED windowState = iota // The window cannot move and is set to its bounding box
	FREE                        // The window can be moved and resized
)

// Extends from ScrollableContainer.
type Window struct {
	container.ScrollableContainer
	State       windowState
	closeButton *ImageButton
	Padding     float64 // Inner padding
	Closed      bool    // Keep track if the window has been closed, waiting to be removed
}

func NewWindow(rect utils.RelativeRect, state windowState, qualityFactor float64) *Window {
	w := new(Window)
	w.ScrollableContainer = container.NewScrollableContainer(rect, w.GetRendererSize(rect.Size), qualityFactor)
	w.State = state
	w.Padding = 5 * settings.SettingInstance.Scale
	w.Closed = false

	w.closeButton = NewImageButton("assets/exit.png", utils.RelativeRect{
		Position: utils.RelativePosition{
			HorizontalAnchor: utils.ANCHOR_RIGHT,
			VerticalAnchor:   utils.ANCHOR_TOP,
			Vec2:             math.NewVec2(0, 0),
		},
		Size: math.NewVec2(WINDOW_BAR_LENGTH, WINDOW_BAR_LENGTH).Scale(settings.SettingInstance.Scale),
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

	// Set the window as closed
	if w.closeButton.Clicked {
		w.Closed = true
	}

}

func (w *Window) renderWindow(windowRect math.Rect, barRect math.Rect) {
	// TODO: the 0.03 is probably the roundness of the edges of the rectangle compared to its size. -> implement as setting
	// Draw the window background
	graphic.DrawRectRounded(windowRect, 0.03, settings.SettingInstance.Theme.WindowTheme.BackgroundColor)
	// Draw the content inside the window
	w.Content.Draw(w.GetRendererPosition(windowRect))
	// Update the button to close the window
	w.closeButton.Update(windowRect)
	// Draw the bar at the top of the window
	graphic.DrawRect(barRect, color.RGBA{255, 255, 255, 5})

}

func (w *Window) handleWindowMovement(boundingBox math.Rect, barRect math.Rect) {
	// Handle resizing/moving the window
	if w.State == FREE {
		w.UpdateResize(boundingBox)
		if !w.closeButton.Hovered && !w.Resizing {
			w.UpdateDrag(boundingBox)
		}
	}
	// Handle double click to switch state of the window
	if barRect.PointCollision(input.GetMousePosition()) || w.Hovering {
		if w.HandleDoubleClick(boundingBox) {
			if w.State == FREE { // junky code (no real choice). I honestly don't see any elegant way to switch state ((+1)%n maybe..)
				w.State = ANCHORED
			} else {
				w.State = FREE
			}
		}
	}
}

// Get the size of the renderer
func (w *Window) GetRendererSize(windowSize math.Vec2) math.Vec2 {
	return windowSize.Substract(math.NewVec2(w.Padding*2, WINDOW_BAR_LENGTH+w.Padding*2))
}

// Get the position of the renderer
func (w *Window) GetRendererPosition(windowAbsoluteRect math.Rect) math.Vec2 {
	return windowAbsoluteRect.Position.Add(math.NewVec2(w.Padding, WINDOW_BAR_LENGTH+w.Padding))
}

// Get the window Rect in absolute coordinates
func (w *Window) GetWindowRect(boundingBox math.Rect) math.Rect {
	var windowAbsoluteRect math.Rect // Rect with absolute coordinates
	if w.State == FREE {
		windowAbsoluteRect = w.GetAbsoluteRect(boundingBox)
	} else { // The rect is fixed to the size of the bounding box
		windowAbsoluteRect = boundingBox
	}
	return windowAbsoluteRect
}

// Resize (reload) the renderer
func (w *Window) Resize(newSize math.Vec2) {
	w.Content = w.Content.Resize(newSize)
}

// convert position from the position in pixel from the top left of the window (Raylib), to the position in the window
func (w *Window) ConvertPositionToRenderer(v math.Vec2, boundingBox math.Rect) math.Vec2 {
	return v.
		Substract(w.GetRendererPosition(w.GetWindowRect(boundingBox))). // substract the position of the renderer
		Scale(w.Content.QualityFactor)                                  // unstretch what was stretched by the QualityFactor
}

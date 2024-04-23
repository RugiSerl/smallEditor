package container

import (
	"math"

	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/input"
	m "github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const RESIZE_LENIENCY = 5        // px
const DOUBLE_CLICK_INTERVAL = .3 // second - the time between two clicks to be considered as a double click

type ResizeSide int

const (
	None ResizeSide = iota
	Top
	Bottom
	Left
	Right
)

// a resizable object extends from draggable
type Resizable struct {
	Draggable
	Hovering, Resizing bool
	side               ResizeSide
	lastClick          float64 // The last time the user clicked to handle double click
}

func NewResizable(rect utils.RelativeRect) Resizable {
	return Resizable{NewDraggable(rect), false, false, None, 5}
}

func (r *Resizable) UpdateResize(boundingBox m.Rect) {

	r.handleCursor()
	r.startResizing()
	r.resize()
	r.EndResizing()

}

// Check if the user has started resizing
func (r *Resizable) startResizing() {
	if input.IsMouseClicked(input.MouseButtonLeft) {
		r.side = r.getSide()
		r.Resizing = r.side != None

	}

}

// The User is resizing the container
func (r *Resizable) resize() {
	if r.Resizing {

		switch r.side {
		case Top:
			r.Position.Y += input.GetMouseDelta().Y
			r.Size.Y -= input.GetMouseDelta().Y // the bottom side must remain unchanged
		case Bottom:
			r.Size.Y += input.GetMouseDelta().Y
		case Left:
			r.Position.X += input.GetMouseDelta().X
			r.Size.X -= input.GetMouseDelta().X // the bottom side must remain unchanged
		case Right:
			r.Size.X += input.GetMouseDelta().X
		}
	}
}

// Check if the user has stopped resizing
func (r *Resizable) EndResizing() {
	if !input.IsMouseDown(input.MouseButtonLeft) {
		r.Resizing = false
		r.side = None
	}
}

// Get which side of the rectangle, if not none, the mouse is hovering
func (r *Resizable) getSide() ResizeSide {
	mousePos := input.GetMousePosition()
	switch {
	case math.Abs(mousePos.X-r.Position.X) < RESIZE_LENIENCY:
		return Left
	case math.Abs(mousePos.X-(r.Position.X+r.Size.X)) < RESIZE_LENIENCY:
		return Right
	case math.Abs(mousePos.Y-r.Position.Y) < RESIZE_LENIENCY:
		return Top
	case math.Abs(mousePos.Y-(r.Position.Y+r.Size.Y)) < RESIZE_LENIENCY:
		return Bottom
	default:
		return None
	}

}

// For better visual appearence, displays a different cursor if the user can resize
// TODO: implement the diagonal arrows
func (r *Resizable) handleCursor() {
	mousePos := input.GetMousePosition()
	switch {
	case math.Abs(mousePos.X-r.Position.X) < RESIZE_LENIENCY:
		rl.SetMouseCursor(rl.MouseCursorResizeEW)
	case math.Abs(mousePos.X-(r.Position.X+r.Size.X)) < RESIZE_LENIENCY:
		rl.SetMouseCursor(rl.MouseCursorResizeEW)
	case math.Abs(mousePos.Y-r.Position.Y) < RESIZE_LENIENCY:
		rl.SetMouseCursor(rl.MouseCursorResizeNS)
	case math.Abs(mousePos.Y-(r.Position.Y+r.Size.Y)) < RESIZE_LENIENCY:
		rl.SetMouseCursor(rl.MouseCursorResizeNS)
	default:
		rl.SetMouseCursor(rl.MouseCursorDefault)
	}
}

// Returns bool - whether the user want to change the state of the container (like ANCHORED or FREE)
func (r *Resizable) HandleDoubleClick(boundingBox m.Rect) bool {
	r.lastClick += graphic.GetDeltaTime()
	if input.IsMouseClicked(input.MouseButtonLeft) {
		if r.lastClick < DOUBLE_CLICK_INTERVAL {
			r.lastClick = DOUBLE_CLICK_INTERVAL + 1
			s := r.getSide()
			switch s {
			case Left, Right:
				r.Position.X = boundingBox.Position.X
				r.Size.X = boundingBox.Size.X
			case Top, Bottom:
				r.Position.Y = boundingBox.Position.Y
				r.Size.Y = boundingBox.Size.Y
			default:
				return true

			}
		} else {
			r.lastClick = 0
		}

	}
	return false
}

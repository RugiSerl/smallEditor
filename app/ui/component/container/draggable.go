package container

import (
	"github.com/RugiSerl/smallEditor/app/input"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

type Draggable struct {
	utils.RelativeRect
	Dragging bool
}

func NewDraggable(rect utils.RelativeRect) Draggable {
	return Draggable{rect, false}
}

func (d *Draggable) UpdateDrag(boundingBox math.Rect) {
	if input.IsMouseClicked(input.MouseButtonLeft) && d.GetAbsoluteRect(boundingBox).PointCollision(input.GetMousePosition()) {
		d.Dragging = true
	}

	if input.IsMouseDown(input.MouseButtonLeft) && d.Dragging {
		d.Position.Vec2 = d.Position.Add(input.GetMouseDelta())
	}

	if !input.IsMouseDown(input.MouseButtonLeft) && d.Dragging {
		d.Dragging = false
	}
}

package container

import (
	"github.com/RugiSerl/smallEditor/app/input"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

type DraggableContainer struct {
	utils.RelativeRect
	Dragging                 bool
	DragOrigin, DragPosition math.Vec2f
}

func NewDraggableContainer(rect utils.RelativeRect) DraggableContainer {
	return DraggableContainer{rect, false, math.NewVec2(0, 0), math.NewVec2(0, 0)}
}

func (d *DraggableContainer) UpdateDrag(boundingBox math.Rect) {

	d.StartDrag(boundingBox)
	d.Drag()
	d.EndDrag()

}

func (d *DraggableContainer) StartDrag(boundingBox math.Rect) {
	// Has the user started dragging ?
	if input.IsMouseClicked(input.MouseButtonLeft) && d.GetAbsoluteRect(boundingBox).PointCollision(input.GetMousePosition()) {
		d.Dragging = true
		d.DragOrigin = input.GetMousePosition().Substract(d.Position.Vec2f)
	}

}

func (d *DraggableContainer) Drag() {
	// Is the user still dragging ?
	if input.IsMouseDown(input.MouseButtonLeft) && d.Dragging {
		d.Position.Vec2f = input.GetMousePosition().Substract(d.DragOrigin)

	}
}

func (d *DraggableContainer) EndDrag() {
	// Has the user finished dragging ?
	if !input.IsMouseDown(input.MouseButtonLeft) && d.Dragging {
		d.Dragging = false
	}
}

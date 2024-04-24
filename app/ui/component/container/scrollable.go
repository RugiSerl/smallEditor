package container

import (
	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

// Extends from ResizableContainer, since in most cases, a scrollable item is resizable
type ScrollableContainer struct {
	ResizableContainer                   // Contains size and relative position of the container
	RealSize           math.Vec2         // Real size of the elements in the textBox. This value must be given by the element using the container
	Camera             *graphic.Camera2D // Camera used to handle the movement
	Content            *graphic.Renderer // Renderer in which all the content is renderer

}

func NewScrollableContainer(rect utils.RelativeRect) ScrollableContainer {
	return ScrollableContainer{
		ResizableContainer: NewResizableContainer(rect),
		RealSize:           math.NewVec2(0, 0), // default value
		Camera:             graphic.NewCamera(),
	}
}

func (s *ScrollableContainer) UpdateScroll() {

}

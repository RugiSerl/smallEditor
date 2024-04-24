package container

import (
	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/input"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/settings"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

// Extends from ResizableContainer, since in most cases, a scrollable item is resizable
type ScrollableContainer struct {
	ResizableContainer                   // Contains size and relative position of the container
	RealSize           math.Vec2         // Real size of the elements in the textBox. This value must be given by the element using the container
	Camera             *graphic.Camera2D // Camera used to handle the movement
	Content            *graphic.Renderer // Renderer in which all the content is renderer
	scrollOffset       math.Vec2
}

func NewScrollableContainer(rect utils.RelativeRect, size math.Vec2, qualityFactor float64) ScrollableContainer {
	return ScrollableContainer{
		ResizableContainer: NewResizableContainer(rect),
		RealSize:           math.NewVec2(0, 0), // default value
		Camera:             graphic.NewCamera(),
		Content:            graphic.NewRenderTexture(size, qualityFactor),
	}
}

func (s *ScrollableContainer) UpdateScroll() {
	// Handling vertical scroll
	// BUG: The scroll sometimes stutters
	s.scrollOffset.Y -= input.MouseWheelDelta() * settings.SettingInstance.ScrollSpeed * float64(settings.SettingInstance.FontSize)
	if s.scrollOffset.Y < 0 {
		s.scrollOffset.Y = 0
	}
	s.Camera.SetTargetPosition(s.scrollOffset)

}

func (s *ScrollableContainer) BeginRendering(boundingBox math.Rect) {
	s.Camera.UpdateCamera(boundingBox)
	s.Content.Begin()
	s.Camera.Begin()

}

func (s *ScrollableContainer) EndRendering() {
	s.Camera.End()
	s.Content.End()
}

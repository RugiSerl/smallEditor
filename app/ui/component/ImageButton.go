package component

import (
	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/input"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/settings"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

// ImageLabel with more stats
type ImageButton struct {
	image            *ImageLabel
	Clicked, Hovered bool
}

func NewImageButton(imagePath string, rect utils.RelativeRect) *ImageButton {
	i := new(ImageButton)
	i.image = NewImageLabel(rect, imagePath)

	return i
}

func (i *ImageButton) Update(boundingBox math.Rect) {
	i.handleInput(boundingBox)
	i.render(boundingBox)

}

func (i *ImageButton) render(boundingBox math.Rect) {
	if i.Hovered {
		graphic.DrawRect(i.image.GetRect(boundingBox), settings.SettingInstance.Theme.ButtonTheme.HoverColor)
	}
	i.image.Update(boundingBox)

}

func (i *ImageButton) handleInput(boundingBox math.Rect) {
	i.Clicked = false
	i.Hovered = false
	if i.image.GetRect(boundingBox).PointCollision(input.GetMousePosition()) {
		i.Hovered = true
		if input.IsMouseClicked(input.MouseButtonLeft) {
			i.Clicked = true
		}
	}
}

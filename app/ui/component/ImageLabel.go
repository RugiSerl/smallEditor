package component

import (
	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

type ImageLabel struct {
	rect        utils.RelativeRect
	textureSize math.Vec2f
	texture     graphic.Texture
}

func NewImageLabel(rect utils.RelativeRect, imagePath string) *ImageLabel {
	i := new(ImageLabel)
	i.rect = rect
	i.texture = graphic.NewTexture(imagePath)
	i.textureSize = math.NewVec2(float64(i.texture.Width), float64(i.texture.Height))
	return i
}

func (i *ImageLabel) Update(boundingBox math.Rect) {
	graphic.DrawTextureRect(i.texture, math.NewRect(math.NewVec2(0, 0), i.textureSize), i.rect.GetAbsoluteRect(boundingBox), math.NewVec2(0, 0), 0)
}

func (i *ImageLabel) GetRect(boundingBox math.Rect) math.Rect {
	return i.rect.GetAbsoluteRect(boundingBox)
}

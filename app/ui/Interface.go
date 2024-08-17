package ui

import (
	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/settings"
	"github.com/RugiSerl/smallEditor/app/ui/component"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

type Interface struct {
	windowManager *component.WindowManager
	logo          *component.ImageLabel
}

func NewInterface() *Interface {
	i := new(Interface)
	i.logo = component.NewImageLabel(utils.RelativeRect{Position: utils.RelativePosition{HorizontalAnchor: utils.ANCHOR_MIDDLE, VerticalAnchor: utils.ANCHOR_MIDDLE, Vec2f: math.NewVec2(0, 0)}, Size: math.NewVec2(128, 128)}, "assets/logo.png")
	i.windowManager = component.NewWindowManager()

	return i
}

func (i *Interface) Update(boundingBox math.Rect) {
	graphic.ClearBackground(settings.SettingInstance.Theme.InterfaceTheme.BackgroundColor)
	i.logo.Update(boundingBox)
	i.windowManager.Update(boundingBox)

}

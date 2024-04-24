package ui

import (
	"github.com/RugiSerl/smallEditor/app/IO"
	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/settings"
	"github.com/RugiSerl/smallEditor/app/ui/component"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

type Interface struct {
	text *component.TextEditor
	logo *component.ImageLabel
}

func NewInterface() *Interface {
	i := new(Interface)
	f, _ := IO.ParseFile("app/app.go")
	i.text = component.NewTextEditor(utils.RelativeRect{Position: utils.RelativePosition{HorizontalAnchor: utils.ANCHOR_LEFT, VerticalAnchor: utils.ANCHOR_TOP, Vec2: math.NewVec2(0, 0)}, Size: math.NewVec2(600, 400)}, component.ANCHORED, f.GetText())
	i.logo = component.NewImageLabel(utils.RelativeRect{Position: utils.RelativePosition{HorizontalAnchor: utils.ANCHOR_MIDDLE, VerticalAnchor: utils.ANCHOR_MIDDLE, Vec2: math.NewVec2(0, 0)}, Size: math.NewVec2(128, 128)}, "assets/logo.png")

	return i
}

func (i *Interface) Update(boundingBox math.Rect) {
	graphic.ClearBackground(settings.SettingInstance.Theme.InterfaceTheme.BackgroundColor)
	i.logo.Update(boundingBox)
	i.text.Update(boundingBox)

}

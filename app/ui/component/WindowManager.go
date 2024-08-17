package component

import (
	"github.com/RugiSerl/smallEditor/app/IO"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

type WindowManager struct {
	data []IWindow
}

func NewWindowManager() *WindowManager {
	f, _ := IO.ParseFile("assets/shader/blur.fs")
	return &WindowManager{
		// By default there is only one window open
		data: []IWindow{NewTextEditor(utils.RelativeRect{Position: utils.RelativePosition{HorizontalAnchor: utils.ANCHOR_LEFT, VerticalAnchor: utils.ANCHOR_TOP, Vec2f: math.NewVec2(0, 0)}, Size: math.NewVec2(600, 400)}, ANCHORED, f.GetText())},
	}

}

func (w *WindowManager) Update(boundingBox math.Rect) {
	for i := len(w.data) - 1; i >= 0; i-- {
		w.data[i].Update(boundingBox)
	}
}

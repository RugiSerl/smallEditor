package component

import (
	"github.com/RugiSerl/smallEditor/app/IO"
	"github.com/RugiSerl/smallEditor/app/math"
	u "github.com/RugiSerl/smallEditor/app/ui/utils"
	"github.com/RugiSerl/smallEditor/app/utils"
)

type WindowManager struct {
	data []IWindow
}

func NewWindowManager() *WindowManager {
	f, _ := IO.ParseFile("assets/shader/blur.fs")
	return &WindowManager{
		// By default there is only one window open
		data: []IWindow{NewTextEditor(u.RelativeRect{Position: u.RelativePosition{HorizontalAnchor: u.ANCHOR_LEFT, VerticalAnchor: u.ANCHOR_TOP, Vec2f: math.NewVec2(0, 0)}, Size: math.NewVec2(600, 400)}, ANCHORED, f.GetText())},
	}

}

func (w *WindowManager) Update(boundingBox math.Rect) {
	var toDelete = []int{}

	for i := len(w.data) - 1; i >= 0; i-- {
		w.data[i].Update(boundingBox)
		if w.data[i].IsClosed() {
			toDelete = append(toDelete, i)
		}
	}

	for _, i := range toDelete {
		w.data = utils.Remove[IWindow](w.data, i)
	}
}

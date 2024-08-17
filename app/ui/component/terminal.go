package component

import (
	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/settings"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

// Extends from windows
type Terminal struct {
	*Window
	textBox *TextBox
	History []string
}

func NewTerminal() *Terminal {
	textBoxFont := graphic.NewFont(settings.SettingInstance.FontPath, int32(float64(settings.SettingInstance.FontSize)*settings.SettingInstance.TextSettings.QualityFactor))
	WindowRect := utils.RelativeRect{
		Position: utils.RelativePosition{
			Vec2f:            math.NewVec2(0, 0),
			HorizontalAnchor: utils.ANCHOR_LEFT,
			VerticalAnchor:   utils.ANCHOR_TOP,
		},
		Size: math.NewVec2(100, 20),
	}
	commandRect := WindowRect
	commandRect.Size.Y *= .25

	return &Terminal{
		textBox: NewTextBox(commandRect, textBoxFont),
		History: []string{},
		Window:  NewWindow(WindowRect, FREE, settings.SettingInstance.TextSettings.QualityFactor),
	}
}

func (t *Terminal) Update(boundingBox math.Rect) {
	// Start to draw inside the window renderer
	t.Window.BeginRendering(boundingBox)

	// Actual rendering
	t.textBox.Update(boundingBox)

	// End of renderer mode
	t.Window.EndRendering()

	// Update the containing window, (which will also draw the renderer in which we draw)
	t.Window.Update(boundingBox)

}

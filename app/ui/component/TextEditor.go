package component

import (
	"strings"

	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/input"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/settings"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

const (
	SCROLL_SPEED  = 2
	SMOOTH_AMOUNT = 20
)

type TextEditor struct {
	window             *Window
	text               string
	font               *graphic.Font
	scrollOffset       float64
	smoothScrollOffset float64
	camera             *graphic.Camera2D
}

func NewTextEditor(pos utils.RelativePosition, size math.Vec2, state windowState) *TextEditor {
	t := new(TextEditor)
	t.text = ""
	t.window = NewWindow(pos, size, state, settings.SettingInstance.TextSettings.QualityFactor)
	t.font = graphic.NewFont(settings.SettingInstance.FontPath, int32(float64(settings.SettingInstance.FontSize)*settings.SettingInstance.TextSettings.QualityFactor))
	t.camera = graphic.NewCamera()
	return t
}

func (t *TextEditor) AppendText(text string) {
	t.text += text
}

func (t *TextEditor) Update(boundingBox math.Rect) {
	t.handleInput()
	t.camera.UpdateCamera(boundingBox)

	t.window.Content.Begin()

	t.camera.Begin()

	t.Render(boundingBox)
	t.window.Content.End()
	t.camera.End()

	t.window.Update(boundingBox)

}

func (t *TextEditor) Render(boundingBox math.Rect) {
	t.text = strings.ReplaceAll(t.text, string(13), "") // Remove vertical tab (ascii 13) (wtf is even that xd)
	lines := strings.Split(t.text, "\n")
	lineHeight := t.font.GetSize("A!").X
	for i, line := range lines {
		t.font.Draw(line, math.NewVec2(0, t.smoothScrollOffset+float64(i)*lineHeight), settings.SettingInstance.TextEditorTheme.TextColor)
	}

}

func (t *TextEditor) handleInput() {
	// Handling zoom
	if input.IsKeyDown(input.KeyLeftControl) {
		t.camera.UpdateZoomInput()

		// Handling scroll
	} else {
		t.scrollOffset += input.MouseWheelDelta() * SCROLL_SPEED * float64(settings.SettingInstance.FontSize)
		if t.scrollOffset > 0 {
			t.scrollOffset = 0
		}
		t.smoothScrollOffset += (t.scrollOffset - t.smoothScrollOffset) / SMOOTH_AMOUNT
	}
	t.AppendText(input.GetKeysPressed())
	if strings.Contains(t.text, string(8)) {
		t.text = handleBackspace(t.text)

	}

}

func (t *TextEditor) convertPosition(v math.Vec2, boundingBox math.Rect) math.Vec2 {
	return t.camera.ConvertToWorldCoordinates(t.window.ConvertPositionToRenderer(v, boundingBox)) // get the conversion from the camera
}

func handleBackspace(text string) string {
	reversedString := ""
	backSpaceAmount := 0
	for i := len(text) - 1; i >= 0; i-- {
		if text[i] == 8 {
			backSpaceAmount++
		} else {
			if backSpaceAmount > 0 {
				backSpaceAmount--
			} else {
				reversedString += string(text[i])
			}
		}
	}
	newString := ""
	for i := len(reversedString) - 1; i >= 0; i-- {
		newString += string(reversedString[i])
	}
	return newString
}

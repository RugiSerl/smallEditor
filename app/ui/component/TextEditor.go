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
	SCROLL_SPEED         = 2
	SCROLL_SMOOTH_AMOUNT = 20 // Amount of smooth given for TextEditor.smoothScrollOffset
)

type TextEditor struct {
	*Window                   // Window containing of the TextEditor
	camera  *graphic.Camera2D // Camera used to zoom/de-zoom
	textbox *TextBox
}

func NewTextEditor(rect utils.RelativeRect, state windowState, text string) *TextEditor {
	t := new(TextEditor)
	t.Window = NewWindow(rect, state, settings.SettingInstance.TextSettings.QualityFactor)

	textBoxText := strings.ReplaceAll(text, string(13), "") // Remove vertical tab (ascii 13) (wtf is even that xd)
	textBoxFont := graphic.NewFont(settings.SettingInstance.FontPath, int32(float64(settings.SettingInstance.FontSize)*settings.SettingInstance.TextSettings.QualityFactor))
	t.textbox = NewTextBox(rect, textBoxFont)
	t.textbox.InsertText(textBoxText)

	t.camera = graphic.NewCamera()
	return t
}

// Main function called each frame
func (t *TextEditor) Update(boundingBox math.Rect) {
	t.handleInput()
	// Update zoom
	t.camera.UpdateCamera(boundingBox)

	// Start to draw inside the window renderer
	t.Content.Begin()
	// Start camera mode to allow zoom
	t.camera.Begin()
	// Actual rendering
	t.textbox.Update(boundingBox, settings.SettingInstance.Theme.TextEditorTheme.TextColor)

	// End of camera mode
	t.camera.End()
	// End of renderer mode
	t.Content.End()

	//update the containing window, (which will also draw the renderer in which we draw)
	t.Window.Update(boundingBox)

}

// Handle user input
func (t *TextEditor) handleInput() {
	// Handling zoom
	if input.IsKeyDown(input.KeyLeftControl) {
		t.camera.UpdateZoomInput()

	}

}

// convert position from the position in pixel from the top left of the window (Raylib), to the position in the textBox
func (t *TextEditor) convertPosition(v math.Vec2, boundingBox math.Rect) math.Vec2 {
	return t.camera.ConvertToWorldCoordinates(t.ConvertPositionToRenderer(v, boundingBox)) // get the conversion from the camera
}

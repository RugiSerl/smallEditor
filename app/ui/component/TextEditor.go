package component

import (
	"strings"

	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/input"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/settings"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

// Extends from Window.
// Object used to edit a file of code.
type TextEditor struct {
	*Window          // Window containing of the TextEditor
	textbox *TextBox // Object containing the text
}

func NewTextEditor(rect utils.RelativeRect, state windowState, text string) *TextEditor {
	t := new(TextEditor)
	t.Window = NewWindow(rect, state, settings.SettingInstance.TextSettings.QualityFactor)

	textBoxText := strings.ReplaceAll(text, string(13), "") // Remove vertical tab (ascii 13) (wtf is even that xd)
	textBoxFont := graphic.NewFont(settings.SettingInstance.FontPath, int32(float64(settings.SettingInstance.FontSize)*settings.SettingInstance.TextSettings.QualityFactor))
	t.textbox = NewTextBox(rect, textBoxFont)
	t.textbox.InsertText(textBoxText)

	t.Camera = graphic.NewCamera()
	return t
}

// Main function called each frame
func (t *TextEditor) Update(boundingBox math.Rect) {

	// Start to draw inside the window renderer
	t.Window.BeginRendering(boundingBox)
	t.handleInput(boundingBox)

	// Actual rendering
	t.textbox.Update(boundingBox)

	// End of renderer mode
	t.Window.EndRendering()

	// Update the containing window, (which will also draw the renderer in which we draw)
	t.Window.Update(boundingBox)

}

// Handle user input
func (t *TextEditor) handleInput(boundingBox math.Rect) {
	// Handling zoom
	if input.IsKeyDown(input.KeyLeftControl) {
		t.Camera.UpdateZoomInput()

	} else {
		t.UpdateScroll()
	}

	if input.IsMouseClicked(input.MouseButtonLeft) {
		t.textbox.SetCursorPosition(t.convertPosition(input.GetMousePosition(), boundingBox))
	}
}

// convert position from the position in pixel from the top left of the window (Raylib), to the position in the textBox
func (t *TextEditor) convertPosition(v math.Vec2f, boundingBox math.Rect) math.Vec2f {
	return t.Camera.ConvertToWorldCoordinates(t.ConvertPositionToRenderer(v, boundingBox)) // get the conversion from the camera
}

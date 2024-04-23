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
	window             *Window           // Window containing of the TextEditor
	text               string            // Text displayed (content of the file)
	font               *graphic.Font     // Font used to draw text
	scrollOffset       float64           // Vertical offset triggered by scrolling
	smoothScrollOffset float64           // Value "following" scrollOffset but smoothly
	camera             *graphic.Camera2D // Camera used to zoom/de-zoom
	cursor             int               // Index where the cursor is located to insert text
}

func NewTextEditor(pos utils.RelativePosition, size math.Vec2, state windowState) *TextEditor {
	t := new(TextEditor)
	t.text = ""
	t.window = NewWindow(pos, size, state, settings.SettingInstance.TextSettings.QualityFactor)
	t.font = graphic.NewFont(settings.SettingInstance.FontPath, int32(float64(settings.SettingInstance.FontSize)*settings.SettingInstance.TextSettings.QualityFactor))
	t.camera = graphic.NewCamera()
	return t
}

// Add text at the end of the text
func (t *TextEditor) AppendText(text string) {
	t.text += text
}

// Main function called each frame
func (t *TextEditor) Update(boundingBox math.Rect) {
	t.handleInput()
	// Update zoom
	t.camera.UpdateCamera(boundingBox)

	// Start to draw inside the window renderer
	t.window.Content.Begin()
	// Start camera mode to allow zoom
	t.camera.Begin()
	// Actual rendering
	t.Render(boundingBox)

	// End of camera mode
	t.camera.End()
	// End of renderer mode
	t.window.Content.End()

	//update the containing window, (which will also draw the renderer in which we draw)
	t.window.Update(boundingBox)

}

// Drawing the TextBox
func (t *TextEditor) Render(boundingBox math.Rect) {
	t.text = strings.ReplaceAll(t.text, string(13), "") // Remove vertical tab (ascii 13) (wtf is even that xd)
	lines := strings.Split(t.text, "\n")                // split into line
	lineHeight := t.font.GetSize("A!").X                // Get approximatively the height of a line
	for i, line := range lines {                        //display line by line to get more control over text spacing
		t.font.Draw(line, math.NewVec2(0, t.smoothScrollOffset+float64(i)*lineHeight), settings.SettingInstance.TextEditorTheme.TextColor)
	}

}

// Handle user input
func (t *TextEditor) handleInput() {
	// Handling zoom
	if input.IsKeyDown(input.KeyLeftControl) {
		t.camera.UpdateZoomInput()

	} else {
		// Handling scroll
		t.scrollOffset += input.MouseWheelDelta() * SCROLL_SPEED * float64(settings.SettingInstance.FontSize)
		if t.scrollOffset > 0 {
			t.scrollOffset = 0
		}
		t.smoothScrollOffset += (t.scrollOffset - t.smoothScrollOffset) / SCROLL_SMOOTH_AMOUNT
	}

	// Get Text entered by user
	t.AppendText(input.GetKeysPressed())

	// Remove characteres from string
	if input.IsKeyDown(input.KeyBackspace) && len(t.text) > 0 {
		t.text = t.text[:len(t.text)-1]
	}

}

// convert position from the position in pixel from the top left of the window (Raylib), to the position in the textBox
func (t *TextEditor) convertPosition(v math.Vec2, boundingBox math.Rect) math.Vec2 {
	return t.camera.ConvertToWorldCoordinates(t.window.ConvertPositionToRenderer(v, boundingBox)) // get the conversion from the camera
}

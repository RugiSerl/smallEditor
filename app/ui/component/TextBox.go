package component

import (
	"strings"

	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/input"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/settings"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cursor struct {
	Line, Column    int
	Index           int // Index where the cursor is located to insert text
	SelectionOrigin int // Index of the origin of the selection. -1 for no selection
}

type TextBox struct {
	rect   utils.RelativeRect
	cursor Cursor
	text   string        // Text displayed
	lines  []string      // Text displayed but stored as a slice of the lines
	font   *graphic.Font // Font used to draw text
}

func NewTextBox(rect utils.RelativeRect, font *graphic.Font) *TextBox {
	t := new(TextBox)
	t.rect = rect
	t.text = ""
	t.lines = []string{}
	t.font = font
	t.cursor.Index = 0            // Beginning of the document
	t.cursor.SelectionOrigin = -1 // -1 for no selection
	return t
}

//---------------------------------------------------------------------
// Update functions (Called each frame)--------------------------------

func (t *TextBox) Update(boundingBox math.Rect) {
	t.lines = strings.Split(t.text, "\n")
	t.handleInput()
	t.render(boundingBox)
	t.UpdateCursorFromIndex()

}

func (t *TextBox) handleInput() {
	// Get Text entered by user
	t.InsertText(input.GetKeysPressed())
	t.handleSpecialKeys()

}

// Drawing the TextBox
func (t *TextBox) render(boundingBox math.Rect) {
	lines := strings.Split(t.text, "\n") // Split into line
	lineHeight := t.GetCharSize().Y
	for i, line := range lines { // Display line by line to get more control over text spacing
		t.font.Draw(line, math.NewVec2(0, float64(i)*lineHeight), settings.SettingInstance.TextColor)
	}
	graphic.DrawRect(math.NewRect(t.GetCursorRealPosition(), math.NewVec2(2, lineHeight)), rl.White)
}

//---------------------------------------------------------------------
// Size functions------------------------------------------------------

// Get approximatively the dimension of a charactere, in px
func (t *TextBox) GetCharSize() math.Vec2f {
	return t.font.GetSize("A")
}

// Get the size in row and column of the text
func (t *TextBox) GetTextSize() math.Vec2i {
	size := math.Vec2i{}
	size.Y = strings.Count(t.text, "\n")
	size.X = 0
	// Search for the maximum
	for _, line := range strings.Split(t.text, "\n") {
		if len(line) > size.X {
			size.X = len(line)
		}
	}
	return size
}

// Get the size in px of the text
func (t *TextBox) GetRealSize() math.Vec2f {
	charSize := t.GetCharSize()
	textSize := t.GetTextSize()
	return math.NewVec2(charSize.X*float64(textSize.X), charSize.Y*float64(textSize.Y))
}

//---------------------------------------------------------------------
// Cursor functions----------------------------------------------------

// Add text at the position of the cursor
func (t *TextBox) InsertText(text string) {
	t.text = t.text[:t.cursor.Index] + text + t.text[t.cursor.Index:]
	t.cursor.Index += len(text)
}

// Update cusor position depending on its index
func (t *TextBox) UpdateCursorFromIndex() {
	t.cursor.Line = strings.Count(t.text[:t.cursor.Index], "\n")
	t.cursor.Column = t.cursor.Index - strings.LastIndex(t.text[:t.cursor.Index], "\n") - 1
}

// Update cursor position depending on its line and column
func (t *TextBox) UpdateCursorFromPosition() {
	// Make sure the line and column exist
	t.cursor.Line = math.ClampInt(0, len(t.lines), t.cursor.Line)
	t.cursor.Column = math.ClampInt(0, len(t.lines[t.cursor.Line]), t.cursor.Column)

	t.cursor.Index = t.cursor.Column
	for i := 0; i < t.cursor.Line; i++ {
		t.cursor.Index += len(t.lines[i]) + 1
	}
}

// Get the cursor position onscreen
func (t *TextBox) GetCursorRealPosition() math.Vec2f {
	charSize := t.GetCharSize()
	return math.NewVec2(float64(t.cursor.Column)*charSize.X, float64(t.cursor.Line)*charSize.Y)
}

// Set the cursor index from the position onscreen
func (t *TextBox) SetCursorPosition(position math.Vec2f) {
	charSize := t.GetCharSize()
	textSize := t.GetRealSize()
	t.cursor.Line = int(math.Clamp(0, textSize.Y, position.Y/charSize.Y))
	t.cursor.Column = math.RoundWithThreshold(math.Clamp(0, float64(len(t.lines[t.cursor.Line])), position.X/charSize.X), 0.67)
	t.UpdateCursorFromPosition()
	t.cursor.SelectionOrigin = t.cursor.Index
}

//---------------------------------------------------------------------
// Editor functions----------------------------------------------------

func (t *TextBox) handleSelection() {
	if input.IsMouseDown(input.MouseButtonLeft) {

	}

}

func (t *TextBox) handleSpecialKeys() {
	// Remove characteres from string
	if input.IsKeyDownUsingCoolDown(input.KeyBackspace) {
		t.ctrlAction(t.deleteAction, false) // False for left
	}

	// handle horizontal key cursor movement
	if input.IsKeyDownUsingCoolDown(input.KeyLeft) {
		t.ctrlAction(t.moveAction, false) // False for left
	}
	if input.IsKeyDownUsingCoolDown(input.KeyRight) {
		t.ctrlAction(t.moveAction, true) // True for right
	}

	// handle vertical key cursor movement
	if input.IsKeyDownUsingCoolDown(input.KeyUp) && t.cursor.Line > 0 {
		t.cursor.Line--
		t.UpdateCursorFromPosition()
	}
	if input.IsKeyDownUsingCoolDown(input.KeyDown) && t.cursor.Line < len(t.lines)-1 {
		t.cursor.Line++
		t.UpdateCursorFromPosition()
	}

	// NOTE: Enter is sadly not registered as string in GetCharPressed(), so we have to manually, which will not respect the order of the keys for a low framerate
	if input.IsKeyDownUsingCoolDown(input.KeyEnter) && !input.IsKeyDown(input.KeyLeftControl) { // TODO: ctrl+enter shortcut is reserved for terminal
		t.InsertText("\n")
	}
	if input.IsKeyDownUsingCoolDown(input.KeyTab) {
		t.InsertText("\t")
	}
}

type action func(int)

// Execute an action but with the specification of handling the ctrl+key shortcut (like ctrl+left/right).
// Direction is whether the action is going in the right
func (t *TextBox) ctrlAction(inputAction action, direction bool) {
	if input.IsKeyDown(input.KeyLeftControl) { // ctrl+key action
		if !direction { // To the left
			inputAction(strings.LastIndexAny(t.text[:t.cursor.Index-1], " \n\t.,:)}+-*\"'") + 1 - t.cursor.Index)
		} else { // To the right
			inputAction(strings.IndexAny(t.text[t.cursor.Index+1:], " \n\t.,:({+-*\"'") + 1)
		}

	} else { // Normal action
		if !direction { // To the left
			inputAction(-1)
		} else { // To the right
			inputAction(1)
		}
	}
}

// Straight forward
func (t *TextBox) deleteAction(offset int) {
	if offset >= -t.cursor.Index {
		t.text = t.text[:t.cursor.Index+offset] + t.text[t.cursor.Index:]
		t.cursor.Index += offset
	}
}

// Adds "offset" to the cursor index position
func (t *TextBox) moveAction(offset int) {
	if t.cursor.Index+offset >= 0 && t.cursor.Index+offset <= len(t.text)-1 {
		t.cursor.Index += offset
	}
}

package component

import (
	"image/color"
	"strings"

	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/input"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
)

type TextBox struct {
	rect   utils.RelativeRect
	cursor int           // Index where the cursor is located to insert text
	text   string        // Text displayed
	font   *graphic.Font // Font used to draw text
}

func NewTextBox(rect utils.RelativeRect, font *graphic.Font) *TextBox {
	t := new(TextBox)
	t.rect = rect
	t.text = ""
	t.font = font
	t.cursor = 0

	return t
}

// Add text at the position of the cursor
func (t *TextBox) InsertText(text string) {
	t.text = t.text[:t.cursor] + text + t.text[t.cursor:]
	t.cursor += len(text)
}

func (t *TextBox) Update(boundingBox math.Rect, color color.RGBA) {
	t.handleInput()
	t.Render(boundingBox, color)
}

func (t *TextBox) handleInput() {
	// Get Text entered by user
	t.InsertText(input.GetKeysPressed())

	// TODO: regulate the deletion of characteres
	// Remove characteres from string
	if input.IsKeyDown(input.KeyBackspace) && t.cursor > 0 {
		t.text = t.text[:t.cursor-1] + t.text[t.cursor:]
		t.cursor -= 1
	}

	if input.IsKeyPressed(input.KeyLeft) && t.cursor > 0 {
		t.cursor -= 1
	}
	if input.IsKeyPressed(input.KeyRight) && t.cursor < len(t.text)-1 {
		t.cursor += 1
	}
}

// Drawing the TextBox
func (t *TextBox) Render(boundingBox math.Rect, color color.RGBA) {
	lines := strings.Split(t.text, "\n") // split into line
	lineHeight := t.font.GetSize("A!").X // Get approximatively the height of a line
	for i, line := range lines {         //display line by line to get more control over text spacing
		t.font.Draw(line, math.NewVec2(0, float64(i)*lineHeight), color)
	}

}

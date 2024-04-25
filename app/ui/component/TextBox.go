package component

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/input"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cursor struct {
	Line, Column int
	Index        int // Index where the cursor is located to insert text
}

type TextBox struct {
	rect   utils.RelativeRect
	cursor Cursor
	text   string // Text displayed
	lines  []string
	font   *graphic.Font // Font used to draw text
}

func NewTextBox(rect utils.RelativeRect, font *graphic.Font) *TextBox {
	t := new(TextBox)
	t.rect = rect
	t.text = ""
	t.font = font
	t.cursor.Index = 0
	t.lines = []string{}
	return t
}

// Add text at the position of the cursor
func (t *TextBox) InsertText(text string) {
	t.text = t.text[:t.cursor.Index] + text + t.text[t.cursor.Index:]
	t.cursor.Index += len(text)
}

func (t *TextBox) Update(boundingBox math.Rect, color color.RGBA) {
	t.lines = strings.Split(t.text, "\n")
	t.handleInput()
	t.Render(boundingBox, color)
	t.UpdateCursor()
	fmt.Println(t.GetRealSize())

}

func (t *TextBox) handleInput() {
	// Get Text entered by user
	t.InsertText(input.GetKeysPressed())

	// Remove characteres from string
	if input.IsKeyDownUsingCoolDown(input.KeyBackspace) && t.cursor.Index > 0 {
		t.text = t.text[:t.cursor.Index-1] + t.text[t.cursor.Index:]
		t.cursor.Index--
	}

	if input.IsKeyDownUsingCoolDown(input.KeyLeft) && t.cursor.Index > 0 {
		if input.IsKeyDown(input.KeyLeftControl) {
			t.cursor.Index = strings.LastIndexAny(t.text[:t.cursor.Index-1], " \n\t.,:)}+-*/\"'") + 1
		} else {
			t.cursor.Index--
		}

	}
	if input.IsKeyDownUsingCoolDown(input.KeyRight) && t.cursor.Index < len(t.text)-1 {
		if input.IsKeyDown(input.KeyLeftControl) {
			t.cursor.Index = t.cursor.Index + strings.IndexAny(t.text[t.cursor.Index+1:], " \n\t.,:({+-*/\"'") + 1
			fmt.Println(strings.IndexAny(t.text[t.cursor.Index:], " \n.:({+"))
		} else {
			t.cursor.Index++
		}
	}

}

// Drawing the TextBox
func (t *TextBox) Render(boundingBox math.Rect, color color.RGBA) {
	lines := strings.Split(t.text, "\n") // Split into line
	lineHeight := t.GetCharSize().Y
	for i, line := range lines { // Display line by line to get more control over text spacing
		t.font.Draw(line, math.NewVec2(0, float64(i)*lineHeight), color)
	}
	graphic.DrawRect(math.NewRect(t.GetCursorPosition(), math.NewVec2(2, lineHeight)), rl.White)
}

// Get approximatively the dimension of a charactere, in px
func (t *TextBox) GetCharSize() math.Vec2f {
	return t.font.GetSize("A")
}

// Get the size in row and column of the text
func (t *TextBox) GetTextSize() math.Vec2i {
	size := math.Vec2i{}
	size.Y = strings.Count(t.text, "\n")
	size.X = 0
	// Search for the minimum
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

// Update cusor position depending on its index
func (t *TextBox) UpdateCursor() {
	t.cursor.Line = strings.Count(t.text[:t.cursor.Index], "\n")
	t.cursor.Column = t.cursor.Index - strings.LastIndex(t.text[:t.cursor.Index], "\n") - 1
}

// Get the cursor position onscreen
func (t *TextBox) GetCursorPosition() math.Vec2f {
	charSize := t.GetCharSize()
	return math.NewVec2(float64(t.cursor.Column)*charSize.X, float64(t.cursor.Line)*charSize.Y)
}

// TODO: Make this work
// Set the cursor index from the position onscreen
func (t *TextBox) SetCursorPosition(position math.Vec2f) {
	charSize := t.GetCharSize()
	textSize := t.GetRealSize()
	line := int(math.Clamp(0, textSize.Y, position.Y/charSize.X))
	column := int(math.Clamp(0, float64(len(t.lines[line])), position.Y/charSize.X))
	t.cursor.Index = column
	for i := 0; i < line-1; i++ {
		t.cursor.Index += len(t.lines[i])
	}

}

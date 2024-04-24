package component

import (
	"image/color"
	"strings"

	"github.com/RugiSerl/smallEditor/app/graphic"
	"github.com/RugiSerl/smallEditor/app/input"
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	KEY_COOLDOWN = 0.03 // Amount of time to wait when a key is down between each action it triggers
)

type TextBox struct {
	rect        utils.RelativeRect
	cursor      int                   // Index where the cursor is located to insert text
	text        string                // Text displayed
	font        *graphic.Font         // Font used to draw text
	keyCoolDown map[input.Key]float64 // Used to regulate the amount of key per second
}

func NewTextBox(rect utils.RelativeRect, font *graphic.Font) *TextBox {
	t := new(TextBox)
	t.rect = rect
	t.text = ""
	t.font = font
	t.cursor = 0
	t.keyCoolDown = make(map[input.Key]float64)
	t.keyCoolDown[input.KeyBackspace] = 0

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

	// TODO: create an object to generalize action of key down (first pontual action, then repeat it)
	// Remove characteres from string
	if input.IsKeyDown(input.KeyBackspace) && t.cursor > 0 && t.keyCoolDown[input.KeyBackspace] > KEY_COOLDOWN {
		t.text = t.text[:t.cursor-1] + t.text[t.cursor:]
		t.cursor--
		t.keyCoolDown[input.KeyBackspace] = 0
	}

	if input.IsKeyPressed(input.KeyLeft) && t.cursor > 0 {
		t.cursor--
	}
	if input.IsKeyPressed(input.KeyRight) && t.cursor < len(t.text)-1 {
		t.cursor++
	}
	t.keyCoolDown[input.KeyBackspace] += graphic.GetDeltaTime()
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

func (t *TextBox) GetCharSize() math.Vec2 {
	return t.font.GetSize("A") // Get approximatively the dimension of a charactere, in px
}

func (t *TextBox) GetCursorPosition() math.Vec2 {
	line := strings.Count(t.text[:t.cursor], "\n")
	column := t.cursor - strings.LastIndex(t.text[:t.cursor], "\n") - 1
	charSize := t.GetCharSize()
	return math.NewVec2(float64(column)*charSize.X, float64(line)*charSize.Y)
}

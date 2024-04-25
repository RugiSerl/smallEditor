package graphic

import (
	"image/color"

	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Font struct {
	rl.Font
	path  string
	chars []rune
}

func NewFont(path string, size int32) *Font {
	chars := []rune("azertyuiopqsdfghjklmwxcvbnAZERTYUIOPQSDFGHJKLMWXCVBN0123456789²&é\"'(è_çà)=~#{[|`\\^@]}^$ù*+/-,;:!¨£%µ?.§\t<>")
	f := rl.LoadFontEx(path, size, chars, int32(len(chars)))
	rl.SetTextureFilter(f.Texture, rl.FilterBilinear)
	return &Font{f, path, chars}
}

func (f *Font) Reload() *Font {
	s := f.BaseSize
	rl.UnloadFont(f.Font)
	return NewFont(f.path, s)
}

func (f *Font) Draw(text string, position math.Vec2f, color color.RGBA) {
	rl.DrawTextEx(f.Font, text, position.ToRL(), float32(settings.SettingInstance.FontSize), 0, color)
}

func (f *Font) GetSize(text string) math.Vec2f {
	return math.FromRL(rl.MeasureTextEx(f.Font, text, float32(settings.SettingInstance.FontSize), 0))

}

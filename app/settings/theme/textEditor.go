package theme

import "image/color"

type TextEditorTheme struct {
	TextColor, BackgroundColor color.RGBA
}

func GetDefaultTextEditorTheme() TextEditorTheme {
	return TextEditorTheme{
		TextColor: color.RGBA{255, 255, 255, 255},
	}
}

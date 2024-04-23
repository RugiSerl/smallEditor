package theme

import "image/color"

type WindowTheme struct {
	BackgroundColor color.RGBA
}

func GetDefaultWindowTheme() WindowTheme {
	return WindowTheme{
		BackgroundColor: color.RGBA{48, 56, 65, 255},
	}
}

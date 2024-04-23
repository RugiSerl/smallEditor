package theme

import "image/color"

type ButtonTheme struct {
	HoverColor color.RGBA
}

func GetDefaultButtonTheme() ButtonTheme {
	return ButtonTheme{
		HoverColor: color.RGBA{128, 128, 128, 128},
	}
}

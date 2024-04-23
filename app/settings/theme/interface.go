package theme

import "image/color"

type InterfaceTheme struct {
	BackgroundColor color.RGBA
}

func GetDefaultInterfaceTheme() InterfaceTheme {
	return InterfaceTheme{
		BackgroundColor: color.RGBA{30, 35, 40, 255},
	}
}

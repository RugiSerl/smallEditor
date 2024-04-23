package theme

type Theme struct {
	TextEditorTheme
	WindowTheme
	ButtonTheme
	InterfaceTheme
}

func GetDefaultTheme() Theme {
	return Theme{
		TextEditorTheme: GetDefaultTextEditorTheme(),
		WindowTheme:     GetDefaultWindowTheme(),
		ButtonTheme:     GetDefaultButtonTheme(),
		InterfaceTheme:  GetDefaultInterfaceTheme(),
	}
}

package settings

type InterfaceSettings struct {
	Scale       float64 // Used to change the size of the elements in the interface.
	ScrollSpeed float64 // Scroll speed when the user uses the mouse wheel.
}

func GetDefaultInterfaceSettings() InterfaceSettings {
	return InterfaceSettings{
		Scale:       1.0,
		ScrollSpeed: 2.0,
	}

}

package settings

type TextSettings struct {
	FontSize      int32   // Size of global font.
	FontPath      string  // Path of the global font.
	QualityFactor float64 // Ratio of the definition of a renderer on its size.
}

func GetDefaultTextSettings() TextSettings {
	return TextSettings{
		FontSize:      20,
		FontPath:      "assets/Consolas.ttf",
		QualityFactor: 4,
	}

}

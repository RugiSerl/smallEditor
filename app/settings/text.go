package settings

type TextSettings struct {
	FontSize      int32   // Size of global font.
	FontPath      string  // Path of the global font. Make sure to use a monospace font ! (same size for all characteres), or it will break the editor
	QualityFactor float64 // Ratio of the definition of a renderer on its size.
}

func GetDefaultTextSettings() TextSettings {
	return TextSettings{
		FontSize:      20,
		FontPath:      "assets/font/Go-Mono.ttf",
		QualityFactor: 4,
	}

}

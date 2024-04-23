package settings

type TextSettings struct {
	FontSize      int32
	FontPath      string
	QualityFactor float64
}

func GetdefaultTextSettings() TextSettings {
	return TextSettings{
		FontSize:      20,
		FontPath:      "assets/Consolas.ttf",
		QualityFactor: 4,
	}

}

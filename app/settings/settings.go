package settings

import "github.com/RugiSerl/smallEditor/app/settings/theme"

// Main settings object
type Settings struct {
	TextSettings
	InterfaceSettings
	GraphicSettings
	theme.Theme
}

var (
	SettingInstance Settings
)

func LoadSettings() {
	SettingInstance = GetDefaultSettings()

}
func GetDefaultSettings() Settings {
	return Settings{
		InterfaceSettings: GetDefaultInterfaceSettings(),
		GraphicSettings:   GetDefaultGraphicSettings(),
		TextSettings:      GetDefaultTextSettings(),
		Theme:             theme.GetDefaultTheme(),
	}
}

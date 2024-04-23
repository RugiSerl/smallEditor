package settings

import "github.com/RugiSerl/smallEditor/app/settings/theme"

type Settings struct {
	TextSettings
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
		TextSettings: GetdefaultTextSettings(),
		Theme:        theme.GetDefaultTheme(),
	}
}

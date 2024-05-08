package settings

type GraphicSettings struct {
	FpsCap int32 // Maximum fps. -1 for unlimited
}

func GetDefaultGraphicSettings() GraphicSettings {
	return GraphicSettings{
		FpsCap: -1,
	}

}

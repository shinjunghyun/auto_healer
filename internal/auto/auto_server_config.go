package auto

type HpMpControl struct {
	ClientMinHpPercent float32 `json:"clientMinHpPercent"`
	ClientMaxHpPercent float32 `json:"clientMaxHpPercent"`
	ClientMinMpPercent float32 `json:"clientMinMpPercent"`
	ServerMinHpPercent float32 `json:"serverMinHpPercent"`
}

type MapInfo struct {
	PrevMapHash        string `json:"prevMapHash"`
	CurrMapHash        string `json:"currMapHash"`
	HeldKeyOnMapChange uint8  `json:"heldKeyOnMapChange"`
}

type CastingConfig struct {
	BaekHoCooldownMilliseconds     uint64 `json:"baekHoCooldownMilliseconds"`
	BaekHoChumCooldownMilliseconds uint64 `json:"baekHoChumCooldownMilliseconds"`
}

type CastingHotkeys struct {
	HonMa      string `json:"honMa"`
	GuiYum     string `json:"guiYum"`
	KiWon      string `json:"kiWon"`
	GongRyuk   string `json:"gongRyuk"`
	BaekHo     string `json:"baekHo"`
	BaekHoChum string `json:"baekHoChum"`
	PaRyuk     string `json:"paRyuk"`
	BooHwal    string `json:"booHwal"`
	SiHoi      string `json:"siHoi"`
	PaHon      string `json:"paHon"`
	Bun        string `json:"bun"`
}

type ConfigExternal struct {
	HpMpControl    HpMpControl    `json:"hpMpControl"`
	MapData        MapInfo        `json:"mapData"`
	CastingConfig  CastingConfig  `json:"castingConfig"`
	CastingHotkeys CastingHotkeys `json:"castingHotkeys"`
}

package speedtrackertypes

type Swing struct {
	Colour   string `json:"colour"`
	Position string `json:"position"`
	Side     string `json:"side"`
	Speed    int    `json:"speed"`
}

func SwingsHaveSameCharacteristics(a Swing, b Swing) (haveSameSwingType bool) {
	return a.Colour == b.Colour && a.Position == b.Position && a.Side == b.Side
}

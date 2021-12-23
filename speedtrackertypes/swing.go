package speedtrackertypes

type Swing struct {
	Colour   string `json:"colour"`
	Position string `json:"position"`
	Side     string `json:"side"`
	Speed    int    `json:"speed"`
}
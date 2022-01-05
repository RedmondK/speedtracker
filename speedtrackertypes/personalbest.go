package speedtrackertypes

import "time"

type PersonalBest struct {
	Date  time.Time `json:"Date"`
	Swing Swing     `json:"Swing"`
}

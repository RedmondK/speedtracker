package speedtrackertypes

import "time"

type PersonalBest struct {
	Date  time.Time `json:"date"`
	Swing Swing     `json:"swing"`
}

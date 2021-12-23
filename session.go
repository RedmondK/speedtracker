package speedtrackertypes

import "time"

type Session struct {
	Protocol string `json:"protocol"`
	SequenceNumber int `json:"sequenceNumber"`
	SessionDate time.Time `json:"date"`
	Swings      []Swing   `json:"swings"`
}
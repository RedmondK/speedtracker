package speedtrackertypes

import "time"

type Session struct {
	Protocol string `json:"protocol"`
	SessionDate time.Time `json:"sessionDate"`
	Swings      []Swing   `json:"swings"`
}
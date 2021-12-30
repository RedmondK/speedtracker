package speedtrackertypes

import "time"

type Session struct {
	Protocol string `json:"protocol"`
	Date time.Time `json:"date"`
	Swings      []Swing   `json:"swings"`
	TotalSpeed int `json:"totalSpeed"`
	AverageSpeed int `json:"averageSpeed"`
}
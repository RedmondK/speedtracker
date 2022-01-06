package speedtrackertypes

type PersonalBestHistoryRecord struct {
	Speed        int          `json:"speed"`
	PersonalBest PersonalBest `json:"personalBest"`
}

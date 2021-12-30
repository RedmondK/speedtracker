package speedtrackertypes

import (
	"time"
)

func CalculateSessionSpeedStatistics(session *Session) {
	session.MinimumSpeed = session.Swings[0].Speed
	session.MaximumSpeed = session.Swings[0].Speed

	for i := 0; i < len(session.Swings); i++ {
		session.TotalSpeed = session.TotalSpeed + session.Swings[i].Speed

		if(session.Swings[i].Speed > session.MaximumSpeed) {
			session.MaximumSpeed = session.Swings[i].Speed
		}

		if(session.Swings[i].Speed < session.MinimumSpeed) {
			session.MinimumSpeed = session.Swings[i].Speed
		}
	}

	session.AverageSpeed = session.TotalSpeed / len(session.Swings)
}

func UpdatePersonalBests(sessionDate time.Time, swings []Swing, personalBests []PersonalBest) (newPBs []PersonalBest) {
	for i := 0; i < len(swings); i++ {
		swing := swings[i];
		pbFound := false
		pb := PersonalBest {}

		for j := 0; j < len(personalBests); j++ {
			pb = personalBests[j];
			if(swing.Colour == pb.Swing.Colour && swing.Position == pb.Swing.Position && swing.Side == pb.Swing.Side) {
				pbFound = true

				if(swing.Speed > pb.Swing.Speed) {
					personalBests[j].Date = sessionDate
					personalBests[j].Swing = swing
				}
			}
		}

		if(!pbFound) {
			newPB := PersonalBest { Date: sessionDate, Swing: swing }
			personalBests = append(personalBests, newPB)
		}
	}

	return personalBests
}
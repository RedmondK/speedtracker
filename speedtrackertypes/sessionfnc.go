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

func UpdatePersonalBests(sessionDate time.Time, swings []Swing, existingPersonalBests []PersonalBest) (newPBs []PersonalBest, replacedPBs []PersonalBest) {
	replacedPBs = []PersonalBest{};
	
	for i := 0; i < len(swings); i++ {
		swing := swings[i];
		pbFound := false
		pb := PersonalBest {}

		for j := 0; j < len(existingPersonalBests); j++ {
			pb = existingPersonalBests[j];
			if(swing.Colour == pb.Swing.Colour && swing.Position == pb.Swing.Position && swing.Side == pb.Swing.Side) {
				pbFound = true

				if(swing.Speed > pb.Swing.Speed) {
					replacedPBs = append(replacedPBs, existingPersonalBests[j])

					existingPersonalBests[j].Date = sessionDate
					existingPersonalBests[j].Swing = swing
				}
			}
		}

		if(!pbFound) {
			newPB := PersonalBest { Date: sessionDate, Swing: swing }
			existingPersonalBests = append(existingPersonalBests, newPB)
		}
	}

	return existingPersonalBests, replacedPBs
}
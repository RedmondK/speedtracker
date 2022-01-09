package speedtrackertypes

import "time"

func UpdatePersonalBests(sessionDate time.Time, swings []Swing, existingPersonalBests []PersonalBest) (newPBs []PersonalBest, pbHistory []PersonalBest) {
	pbHistory = []PersonalBest{}

	for i := 0; i < len(swings); i++ {
		swing := swings[i]
		pbFound := false
		pb := PersonalBest{}

		for j := 0; j < len(existingPersonalBests); j++ {
			pb = existingPersonalBests[j]
			if swing.Colour == pb.Swing.Colour && swing.Position == pb.Swing.Position && swing.Side == pb.Swing.Side {
				pbFound = true

				if swing.Speed > pb.Swing.Speed {
					existingPersonalBests[j].Date = sessionDate
					existingPersonalBests[j].Swing = swing

					newPB := PersonalBest{Date: sessionDate, Swing: swing}
					pbHistory = append(pbHistory, newPB)
				}
			}
		}

		if !pbFound {
			newPB := PersonalBest{Date: sessionDate, Swing: swing}
			existingPersonalBests = append(existingPersonalBests, newPB)
			pbHistory = append(pbHistory, newPB)
		}
	}

	return existingPersonalBests, pbHistory
}

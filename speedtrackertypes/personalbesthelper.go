package speedtrackertypes

import "time"

func GetUpdatedPersonalBestsUsingNewSwings(sessionDate time.Time, personalBests []PersonalBest, newSwings []Swing) (newPersonalBests []PersonalBest, createdPBs []PersonalBest) {
	for i := 0; i < len(newSwings); i++ {
		swing := newSwings[i]
		pbFound := false
		pb := PersonalBest{}

		for j := 0; j < len(personalBests); j++ {
			pb = personalBests[j]
			if swing.Colour == pb.Swing.Colour && swing.Position == pb.Swing.Position && swing.Side == pb.Swing.Side {
				pbFound = true

				if swing.Speed > pb.Swing.Speed {
					personalBests[j].Date = sessionDate
					personalBests[j].Swing = swing

					newPB := PersonalBest{Date: sessionDate, Swing: swing}
					createdPBs = append(createdPBs, newPB)
				}
			}
		}

		if !pbFound {
			newPB := PersonalBest{Date: sessionDate, Swing: swing}
			createdPBs = append(createdPBs, newPB)
			personalBests = append(personalBests, newPB)
		}
	}

	return personalBests, createdPBs
}

func GetObsoletePBHistoryWithNewPBs(personalBestNotInHistory PersonalBest, existingPBHistory []PersonalBestHistoryRecord) (obsoleteHistoryRecords []PersonalBestHistoryRecord) {
	for i := 0; i < len(existingPBHistory); i++ {
		if existingPBHistory[i].PersonalBest.Date.Before(personalBestNotInHistory.Date) { //older pbs cannot be made obsolete by a faster one
			continue
		}

		if existingPBHistory[i].Speed < personalBestNotInHistory.Swing.Speed { //a newer pb which is actually slower than this one is obsolete
			obsoleteHistoryRecords = append(obsoleteHistoryRecords, existingPBHistory[i])
		}
	}

	return obsoleteHistoryRecords
}

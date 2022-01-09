package speedtrackertypes

import "time"

func GetUpdatedPersonalBestsUsingNewSwings(sessionDate time.Time, personalBests []PersonalBest, newSwings []Swing) (newPersonalBests []PersonalBest, createdPBs []PersonalBest) {
	for i := 0; i < len(newSwings); i++ {
		swing := newSwings[i]
		pbFound := false
		pbCreatedForThisSwing := false
		pb := PersonalBest{}

		for j := 0; j < len(personalBests); j++ {
			pb = personalBests[j]
			if swing.Colour == pb.Swing.Colour && swing.Position == pb.Swing.Position && swing.Side == pb.Swing.Side {
				pbFound = true

				if swing.Speed > pb.Swing.Speed && !pbCreatedForThisSwing {
					personalBests[j].Date = sessionDate
					personalBests[j].Swing = swing

					newPB := PersonalBest{Date: sessionDate, Swing: swing}
					createdPBs = append(createdPBs, newPB)
					pbCreatedForThisSwing = true
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

func GetObsoletePBHistoryWithNewPBs(pbNotInHistory PersonalBest, existingPBHistory []PersonalBestHistoryRecord) (obsoleteHistoryRecords []PersonalBestHistoryRecord) {
	for i := 0; i < len(existingPBHistory); i++ {
		historyPBDate := existingPBHistory[i].PersonalBest.Date
		historyPBSwing := existingPBHistory[i].PersonalBest.Swing

		if historyPBDate.Before(pbNotInHistory.Date) { //older pbs cannot be made obsolete by a faster one
			continue
		}

		if historyPBSwing.Colour != pbNotInHistory.Swing.Colour || historyPBSwing.Position != pbNotInHistory.Swing.Position || historyPBSwing.Side != pbNotInHistory.Swing.Side {
			continue
		}

		if existingPBHistory[i].Speed < pbNotInHistory.Swing.Speed { //a newer pb which is actually slower than this one is obsolete
			obsoleteHistoryRecords = append(obsoleteHistoryRecords, existingPBHistory[i])
		}
	}

	return obsoleteHistoryRecords
}

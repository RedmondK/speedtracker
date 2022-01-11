package speedtrackertypes

import "time"

func GetUpdatedPersonalBestData(sessionDate time.Time, personalBests []PersonalBest, newSwings []Swing, existingPBHistory []PersonalBestHistoryRecord) (newPersonalBests []PersonalBest, createdPBs []PersonalBest, obsoleteHistoryRecords []PersonalBestHistoryRecord) {
	for i := 0; i < len(newSwings); i++ {
		swing := newSwings[i]
		pbFound := false
		pbCreated := false

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
					pbCreated = true
					break
				}
			}
		}

		if !pbFound {
			newPB := PersonalBest{Date: sessionDate, Swing: swing}
			createdPBs = append(createdPBs, newPB)
			personalBests = append(personalBests, newPB)
		}

		if pbFound {
			if !pbCreated { //this PB has neither updated a PB or been created as a PB but may have been a PB at time of occuring
				for i := 0; i < len(existingPBHistory); i++ {
					if existingPBHistory[i].PersonalBest.Date.Before(sessionDate) { //this swing cannot supercede a PB in the past
						continue
					}

					if existingPBHistory[i].PersonalBest.Swing.Speed <= swing.Speed && existingPBHistory[i].PersonalBest.Swing.Colour == swing.Colour && existingPBHistory[i].PersonalBest.Swing.Position == swing.Position && existingPBHistory[i].PersonalBest.Swing.Side == swing.Side { //this swing is slower than this new swing and is invalid in the PB History
						if !pbCreated { //check to avoid creating duplicate pbs for this swing when it supercedes multiple history items
							newPB := PersonalBest{Date: sessionDate, Swing: swing}
							createdPBs = append(createdPBs, newPB) //this PB has been created now but will not become part of the user personal bests
							pbCreated = true
						}
						obsoleteHistoryRecords = append(obsoleteHistoryRecords, existingPBHistory[i]) //mark the obsolete pb history for deletion
					}
				}
			} else { // this PB has updated a PB and may have caused obsolete PB records which need to be marked obsolete
				for i := 0; i < len(existingPBHistory); i++ {
					if existingPBHistory[i].PersonalBest.Date.Before(sessionDate) { //this swing cannot supercede a PB in the past
						continue
					}

					if existingPBHistory[i].PersonalBest.Swing.Speed <= swing.Speed && existingPBHistory[i].PersonalBest.Swing.Colour == swing.Colour && existingPBHistory[i].PersonalBest.Swing.Position == swing.Position && existingPBHistory[i].PersonalBest.Swing.Side == swing.Side { //this swing is slower than this new swing and is invalid in the PB History
						obsoleteHistoryRecords = append(obsoleteHistoryRecords, existingPBHistory[i]) //mark the obsolete pb history for deletion
					}
				}
			}
		}
	}

	return personalBests, createdPBs, obsoleteHistoryRecords
}

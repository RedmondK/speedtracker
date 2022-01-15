package speedtrackertypes

import (
	"sort"
	"time"
)

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

		if !pbCreated { //this PB has neither updated a PB or been created as a PB but may have been a PB at time of occuring
			for i := 0; i < len(existingPBHistory); i++ {
				if existingPBHistory[i].PersonalBest.Date.Before(sessionDate) { //this swing cannot supercede a PB in the past
					continue
				}

				var existingPBSwing = existingPBHistory[i].PersonalBest.Swing
				if existingPBSwing.Speed <= swing.Speed && existingPBSwing.Colour == swing.Colour && existingPBSwing.Position == swing.Position && existingPBSwing.Side == swing.Side { //this swing is slower than this new swing and is invalid in the PB History
					if !pbCreated { //check to avoid creating duplicate pbs for this swing when it supercedes multiple history items
						newPB := PersonalBest{Date: sessionDate, Swing: swing}
						createdPBs = append(createdPBs, newPB) //this PB has been created now but will not become part of the user personal bests
						pbCreated = true
					}

					if !existingPBHistory[i].PersonalBest.Date.Equal(sessionDate) {
						obsoleteHistoryRecords = append(obsoleteHistoryRecords, existingPBHistory[i]) //mark the obsolete pb history for deletion
					}
				}
			}
		} else { // this PB has updated a PB and may have caused obsolete PB records which need to be marked obsolete or may have caused future swings to become PBs
			for i := 0; i < len(existingPBHistory); i++ {
				if existingPBHistory[i].PersonalBest.Date.Before(sessionDate) { //this swing cannot supercede a PB in the past
					continue
				}

				var existingPBSwing = existingPBHistory[i].PersonalBest.Swing

				if existingPBSwing.Speed <= swing.Speed && existingPBSwing.Colour == swing.Colour && existingPBSwing.Position == swing.Position && existingPBSwing.Side == swing.Side { //this swing is slower than this new swing and is invalid in the PB History
					if !existingPBHistory[i].PersonalBest.Date.Equal(sessionDate) {
						obsoleteHistoryRecords = append(obsoleteHistoryRecords, existingPBHistory[i]) //mark the obsolete pb history for deletion
					}
				}
			}
		}
	}

	return personalBests, createdPBs, obsoleteHistoryRecords
}

func GetPersonalBestData(sessions []Session) (personalBests []PersonalBest, personalBestHistory []PersonalBestHistoryRecord) {
	sort.Slice(sessions, func(i, j int) bool {
		return sessions[i].Date.Before(sessions[j].Date)
	})

	for _, session := range sessions {
		for _, swing := range session.Swings {
			swingPB := GetPBForSwing(swing, personalBests)

			if swingPB == nil {
				newPB := PersonalBest{Date: session.Date, Swing: swing}
				newPBHistoryRecord := PersonalBestHistoryRecord{Speed: newPB.Swing.Speed, PersonalBest: newPB}

				personalBests = append(personalBests, newPB)
				personalBestHistory = append(personalBestHistory, newPBHistoryRecord)
				continue
			}

			if swingPB.Swing.Speed < swing.Speed {
				newPB := PersonalBest{Date: session.Date, Swing: swing}

				UpdatePersonalBestsWithNewPB(newPB, personalBests)
				personalBestHistory = UpdatePersonalBestHistoryWithNewPB(newPB, personalBestHistory)
			}
		}
	}

	return personalBests, personalBestHistory
}

func GetPBForSwing(swing Swing, personalBests []PersonalBest) (personalBest *PersonalBest) {
	for _, pb := range personalBests {
		if SwingsHaveSameCharacteristics(pb.Swing, swing) {
			return &pb
		}
	}

	return nil
}

func UpdatePersonalBestsWithNewPB(newPersonalBest PersonalBest, personalBests []PersonalBest) {
	for index, pb := range personalBests {
		if SwingsHaveSameCharacteristics(pb.Swing, newPersonalBest.Swing) {
			personalBests[index] = newPersonalBest
		}
	}
}

func UpdatePersonalBestHistoryWithNewPB(newPersonalBest PersonalBest, personalBestHistory []PersonalBestHistoryRecord) (updatedPersonalBestHistory []PersonalBestHistoryRecord) {
	updatedPersonalBestHistory = personalBestHistory

	for index, pbHistoryRecord := range updatedPersonalBestHistory {
		if !SwingsHaveSameCharacteristics(pbHistoryRecord.PersonalBest.Swing, newPersonalBest.Swing) {
			continue
		}

		if pbHistoryRecord.PersonalBest.Date.Equal(newPersonalBest.Date) {
			updatedPersonalBestHistory[index].Speed = newPersonalBest.Swing.Speed
			updatedPersonalBestHistory[index].PersonalBest = newPersonalBest
			continue
		}

		newPersonalBestHistoryRecord := PersonalBestHistoryRecord{Speed: newPersonalBest.Swing.Speed, PersonalBest: newPersonalBest}
		if !HistoryRecordExists(newPersonalBestHistoryRecord, updatedPersonalBestHistory) {
			updatedPersonalBestHistory = append(updatedPersonalBestHistory, newPersonalBestHistoryRecord)
		}
	}

	return updatedPersonalBestHistory
}

func HistoryRecordExists(historyRecord PersonalBestHistoryRecord, currentHistory []PersonalBestHistoryRecord) (exists bool) {
	for _, currentRecord := range currentHistory {
		if historyRecord.PersonalBest.Date.Equal(currentRecord.PersonalBest.Date) {
			return true
		}
	}

	return false
}

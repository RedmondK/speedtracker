package speedtrackertypes

import (
	"sort"
)

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

func IdentifyObsoletePBHistoryRecords(previousHistory []PersonalBestHistoryRecord, newHistory []PersonalBestHistoryRecord) (obsoleteRecords []PersonalBestHistoryRecord) {
	sort.Slice(previousHistory, func(i, j int) bool {
		return previousHistory[i].PersonalBest.Date.Before(previousHistory[j].PersonalBest.Date)
	})

	sort.Slice(newHistory, func(i, j int) bool {
		return newHistory[i].PersonalBest.Date.Before(newHistory[j].PersonalBest.Date)
	})

	for _, previousRecord := range previousHistory {
		foundInNewHistory := false

		for _, newRecord := range newHistory {
			if newRecord.PersonalBest.Date.Equal(previousRecord.PersonalBest.Date) {
				foundInNewHistory = true
				break
			}
		}

		if !foundInNewHistory {
			obsoleteRecords = append(obsoleteRecords, previousRecord)
		}
	}

	return obsoleteRecords
}

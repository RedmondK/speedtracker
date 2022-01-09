package main

import (
	"testing"
	"time"

	"github.com/RedmondK/speedtracker/speedtrackertypes"
)

func TestUpdatePBsFromEmpty(t *testing.T) {
	testSwings := []speedtrackertypes.Swing{}

	testTime := time.Now()
	testPBs := []speedtrackertypes.PersonalBest{}
	testSwings = append(testSwings, speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 144})
	testPersonalBestHistory := []speedtrackertypes.PersonalBestHistoryRecord{}

	newCurrentPBs, newPBsForHistory := speedtrackertypes.GetUpdatedPersonalBestsUsingNewSwings(testTime, testPBs, testSwings)
	obsoletePBs := speedtrackertypes.GetObsoletePBHistoryWithNewPBs(newCurrentPBs[0], testPersonalBestHistory)

	if len(newCurrentPBs) == 0 {
		t.Error("PBs not set")
	}

	if len(newCurrentPBs) > 1 {
		t.Error("Too many PBs")
	}

	if len(newPBsForHistory) == 0 {
		t.Error("PBs not set for History")
	}

	if len(newPBsForHistory) > 1 {
		t.Error("Too many PBs for History")
	}

	if len(obsoletePBs) > 0 {
		t.Error("Obsolete PBs should not be set")
	}

	if newCurrentPBs[0].Date != testTime {
		t.Error("Incorrect date on pb")
	}

	if newCurrentPBs[0].Swing.Colour != "green" {
		t.Error("Incorrect colour on pb")
	}

	if newCurrentPBs[0].Swing.Position != "normal" {
		t.Error("Incorrect position on pb")
	}

	if newCurrentPBs[0].Swing.Side != "dominant" {
		t.Error("Incorrect side on pb")
	}
	if newCurrentPBs[0].Swing.Speed != 144 {
		t.Error("Incorrect speed on pb")
	}
}

func TestUpdatePBFromExisting(t *testing.T) {
	testSwings := []speedtrackertypes.Swing{}

	testTime := time.Now()
	testPBs := []speedtrackertypes.PersonalBest{}

	timeOfOldPB := time.Now()

	testPBs = append(testPBs, speedtrackertypes.PersonalBest{Date: timeOfOldPB, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 140}})
	testSwings = append(testSwings, speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 160})
	testPersonalBestHistory := []speedtrackertypes.PersonalBestHistoryRecord{}

	newCurrentPBs, newPBsForHistory := speedtrackertypes.GetUpdatedPersonalBestsUsingNewSwings(testTime, testPBs, testSwings)
	obsoletePBs := speedtrackertypes.GetObsoletePBHistoryWithNewPBs(newCurrentPBs[0], testPersonalBestHistory)

	if len(newCurrentPBs) == 0 {
		t.Error("New Current PBs not set")
	}

	if len(newCurrentPBs) > 1 {
		t.Error("Too many PBs")
	}

	if len(obsoletePBs) > 0 {
		t.Error("Obsolete PBs should not be set")
	}

	if len(newPBsForHistory) == 0 {
		t.Error("PBs for history should be set")
	}

	if len(newPBsForHistory) > 1 {
		t.Error("Too many PBs for History")
	}

	if newCurrentPBs[0].Date != testTime {
		t.Error("Incorrect date on pb")
	}

	if newCurrentPBs[0].Swing.Colour != "green" {
		t.Error("Incorrect colour on pb")
	}

	if newCurrentPBs[0].Swing.Position != "normal" {
		t.Error("Incorrect position on pb")
	}

	if newCurrentPBs[0].Swing.Side != "dominant" {
		t.Error("Incorrect side on pb")
	}
	if newCurrentPBs[0].Swing.Speed != 160 {
		t.Error("Incorrect speed on pb")
	}
	if newPBsForHistory[0].Swing.Speed != 160 {
		t.Error("Incorrect speed on new pb for history")
	}
	if newPBsForHistory[0].Date != testTime {
		t.Error("Incorrect date on new pb for history")
	}
}

func TestUpdateMultiplePBSomeExisting(t *testing.T) {
	newSwings := []speedtrackertypes.Swing{}

	oldPBTime := time.Now().Add(time.Duration(-600))
	testSessionTime := time.Now()
	currentPBs := []speedtrackertypes.PersonalBest{}

	oldPB := speedtrackertypes.PersonalBest{Date: oldPBTime, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 140}}

	currentPBs = append(currentPBs, oldPB)
	newSwings = append(newSwings, speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 160})
	newSwings = append(newSwings, speedtrackertypes.Swing{Side: "non-dominant", Colour: "red", Position: "max out", Speed: 125})

	testPersonalBestHistory := []speedtrackertypes.PersonalBestHistoryRecord{}
	testPersonalBestHistory = append(testPersonalBestHistory, speedtrackertypes.PersonalBestHistoryRecord{Speed: oldPB.Swing.Speed, PersonalBest: oldPB})

	newCurrentPBs, newPBsForHistory := speedtrackertypes.GetUpdatedPersonalBestsUsingNewSwings(testSessionTime, currentPBs, newSwings)
	obsoleteGreenPBs := speedtrackertypes.GetObsoletePBHistoryWithNewPBs(newCurrentPBs[0], testPersonalBestHistory)
	obsoleteRedPBs := speedtrackertypes.GetObsoletePBHistoryWithNewPBs(newCurrentPBs[1], testPersonalBestHistory)

	if len(newCurrentPBs) == 0 {
		t.Error("PBs not set")
	}

	if len(newCurrentPBs) > 2 {
		t.Error("Too many PBs")
	}

	if len(newCurrentPBs) < 2 {
		t.Error("Too few PBs")
	}

	if len(newPBsForHistory) == 0 {
		t.Error("PBs For History not set")
	}

	if len(newPBsForHistory) > 2 {
		t.Error("Too many PBs for History")
	}

	if newCurrentPBs[0].Date != testSessionTime {
		t.Error("Incorrect date on pb")
	}

	if newCurrentPBs[0].Swing.Colour != "green" {
		t.Error("Incorrect colour on pb")
	}

	if newCurrentPBs[0].Swing.Position != "normal" {
		t.Error("Incorrect position on pb")
	}

	if newCurrentPBs[0].Swing.Side != "dominant" {
		t.Error("Incorrect side on pb")
	}
	if newCurrentPBs[0].Swing.Speed != 160 {
		t.Error("Incorrect speed on pb")
	}

	if newCurrentPBs[1].Date != testSessionTime {
		t.Error("Incorrect date on pb")
	}

	if newCurrentPBs[1].Swing.Colour != "red" {
		t.Error("Incorrect colour on pb")
	}

	if newCurrentPBs[1].Swing.Position != "max out" {
		t.Error("Incorrect position on pb")
	}

	if newCurrentPBs[1].Swing.Side != "non-dominant" {
		t.Error("Incorrect side on pb")
	}
	if newCurrentPBs[1].Swing.Speed != 125 {
		t.Error("Incorrect speed on pb")
	}

	if len(obsoleteGreenPBs) > 0 {
		t.Error("Obsolete PBs should not be set")
	}

	if len(obsoleteRedPBs) > 0 {
		t.Error("Obsolete PBs should not be set")
	}
}

func TestUpdateMultiplePB(t *testing.T) {
	newSwings := []speedtrackertypes.Swing{}

	testSessionTime := time.Now()
	currentPBs := []speedtrackertypes.PersonalBest{}

	oldDate, _ := time.Parse(time.RFC3339, "2021-01-01T22:45:02Z")

	oldPB := speedtrackertypes.PersonalBest{Date: oldDate, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 140}}

	currentPBs = append(currentPBs, speedtrackertypes.PersonalBest{Date: time.Now(), Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 140}})
	currentPBs = append(currentPBs, speedtrackertypes.PersonalBest{Date: time.Now(), Swing: speedtrackertypes.Swing{Side: "non-dominant", Colour: "blue", Position: "normal", Speed: 115}})
	currentPBs = append(currentPBs, speedtrackertypes.PersonalBest{Date: time.Now(), Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "red", Position: "happy gilmore", Speed: 123}})
	currentPBs = append(currentPBs, speedtrackertypes.PersonalBest{Date: time.Now(), Swing: speedtrackertypes.Swing{Side: "non-dominant", Colour: "red", Position: "max out", Speed: 122}})
	currentPBs = append(currentPBs, speedtrackertypes.PersonalBest{Date: time.Now(), Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "step-change", Speed: 140}})
	currentPBs = append(currentPBs, speedtrackertypes.PersonalBest{Date: oldDate, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "olddate", Speed: 140}})

	newSwings = append(newSwings, speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 160})    // update
	newSwings = append(newSwings, speedtrackertypes.Swing{Side: "non-dominant", Colour: "blue", Position: "normal", Speed: 125}) // update
	newSwings = append(newSwings, speedtrackertypes.Swing{Side: "non-dominant", Colour: "red", Position: "max out", Speed: 128}) // update
	newSwings = append(newSwings, speedtrackertypes.Swing{Side: "dominant", Colour: "red", Position: "max out", Speed: 125})     // new
	newSwings = append(newSwings, speedtrackertypes.Swing{Side: "non-dominant", Colour: "blue", Position: "sprint", Speed: 156}) //new

	testPersonalBestHistory := []speedtrackertypes.PersonalBestHistoryRecord{}
	testPersonalBestHistory = append(testPersonalBestHistory, speedtrackertypes.PersonalBestHistoryRecord{Speed: oldPB.Swing.Speed, PersonalBest: oldPB})

	newCurrentPBs, newPBsForHistory := speedtrackertypes.GetUpdatedPersonalBestsUsingNewSwings(testSessionTime, currentPBs, newSwings)
	obsoletePBs := speedtrackertypes.GetObsoletePBHistoryWithNewPBs(newCurrentPBs[0], testPersonalBestHistory)

	if len(newPBsForHistory) == 0 {
		t.Error("history records not set")
	}

	if len(obsoletePBs) > 0 {
		t.Error("obsolete PBs should not be set")
	}

	if len(newPBsForHistory) != 5 {
		t.Error("incorrect history records")
	}

	if len(newCurrentPBs) == 0 {
		t.Error("PBs not set")
	}

	if len(newCurrentPBs) > 8 {
		t.Error("Too many PBs")
	}

	if len(newCurrentPBs) < 8 {
		t.Error("Too few PBs")
	}

	if newCurrentPBs[0].Date != testSessionTime {
		t.Error("Incorrect date on pb")
	}

	if newCurrentPBs[0].Swing.Colour != "green" {
		t.Error("Incorrect colour on pb")
	}

	if newCurrentPBs[0].Swing.Position != "normal" {
		t.Error("Incorrect position on pb")
	}

	if newCurrentPBs[0].Swing.Side != "dominant" {
		t.Error("Incorrect side on pb")
	}
	if newCurrentPBs[0].Swing.Speed != 160 {
		t.Error("Incorrect speed on pb")
	}

	if newCurrentPBs[1].Date != testSessionTime {
		t.Error("Incorrect date on pb")
	}

	if newCurrentPBs[1].Swing.Colour != "blue" {
		t.Error("Incorrect colour on pb")
	}

	if newCurrentPBs[1].Swing.Position != "normal" {
		t.Error("Incorrect position on pb")
	}

	if newCurrentPBs[1].Swing.Side != "non-dominant" {
		t.Error("Incorrect side on pb")
	}
	if newCurrentPBs[1].Swing.Speed != 125 {
		t.Error("Incorrect speed on pb")
	}

	if newCurrentPBs[2].Date != testSessionTime {
		t.Error("Incorrect date on pb")
	}

	if newCurrentPBs[2].Swing.Colour != "red" {
		t.Error("Incorrect colour on pb")
	}

	if newCurrentPBs[2].Swing.Position != "happy gilmore" {
		t.Error("Incorrect position on pb")
	}

	if newCurrentPBs[2].Swing.Side != "dominant" {
		t.Error("Incorrect side on pb")
	}
	if newCurrentPBs[2].Swing.Speed != 123 {
		t.Error("Incorrect speed on pb")
	}

	if newCurrentPBs[3].Swing.Speed != 128 {
		t.Error("Incorrect speed on pb")
	}

	if newCurrentPBs[6].Swing.Speed != 125 {
		t.Error("Incorrect speed on pb")
	}

	if newCurrentPBs[7].Swing.Speed != 156 {
		t.Error("Incorrect speed on pb")
	}

	if newCurrentPBs[5].Date != oldDate {
		t.Error("Incorrect date on oldDate PB which should not change")
	}
}

func TestUpdateWithSessionInThePast(t *testing.T) {
	newSwings := []speedtrackertypes.Swing{}

	testSessionTime, _ := time.Parse(time.RFC3339, "2021-01-07T22:45:02Z")

	oldDate1, _ := time.Parse(time.RFC3339, "2021-01-01T22:45:02Z")
	oldDate2, _ := time.Parse(time.RFC3339, "2021-01-03T22:45:02Z")
	oldDate3, _ := time.Parse(time.RFC3339, "2021-01-05T22:45:02Z")
	oldDate4, _ := time.Parse(time.RFC3339, "2021-01-11T22:45:02Z")
	oldDate5, _ := time.Parse(time.RFC3339, "2021-01-13T22:45:02Z")

	oldPB1 := speedtrackertypes.PersonalBest{Date: oldDate1, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 115}}
	oldPB2 := speedtrackertypes.PersonalBest{Date: oldDate2, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 120}}
	oldPB3 := speedtrackertypes.PersonalBest{Date: oldDate3, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 125}}
	oldPB4 := speedtrackertypes.PersonalBest{Date: oldDate4, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 130}}
	oldPB5 := speedtrackertypes.PersonalBest{Date: oldDate5, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 135}}

	currentPBs := []speedtrackertypes.PersonalBest{}

	//our current PB is the newest PB for this swing type
	currentPBs = append(currentPBs, oldPB5)

	//add five historical pbs for this swing type
	testPersonalBestHistory := []speedtrackertypes.PersonalBestHistoryRecord{}
	testPersonalBestHistory = append(testPersonalBestHistory, speedtrackertypes.PersonalBestHistoryRecord{Speed: oldPB1.Swing.Speed, PersonalBest: oldPB1})
	testPersonalBestHistory = append(testPersonalBestHistory, speedtrackertypes.PersonalBestHistoryRecord{Speed: oldPB2.Swing.Speed, PersonalBest: oldPB2})
	testPersonalBestHistory = append(testPersonalBestHistory, speedtrackertypes.PersonalBestHistoryRecord{Speed: oldPB3.Swing.Speed, PersonalBest: oldPB3})
	testPersonalBestHistory = append(testPersonalBestHistory, speedtrackertypes.PersonalBestHistoryRecord{Speed: oldPB4.Swing.Speed, PersonalBest: oldPB4})
	testPersonalBestHistory = append(testPersonalBestHistory, speedtrackertypes.PersonalBestHistoryRecord{Speed: oldPB5.Swing.Speed, PersonalBest: oldPB5})

	//now the user submits a swing that is newer than two existing PBs, therefore goes into the middle of the history, but also means those shouldn't be PBs since it's swing
	//is faster than them
	newSwings = append(newSwings, speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 160}) // update
	newCurrentPBs, newPBsForHistory := speedtrackertypes.GetUpdatedPersonalBestsUsingNewSwings(testSessionTime, currentPBs, newSwings)
	obsoletePBs := speedtrackertypes.GetObsoletePBHistoryWithNewPBs(newCurrentPBs[0], testPersonalBestHistory)

	//only one swing type involved, so new current pbs should contain one pb
	if len(newCurrentPBs) > 1 {
		t.Error("Too many current pbs")
	}

	if newCurrentPBs[0].Swing.Speed != 160 {
		t.Error("Incorrect current PB")
	}

	if len(newPBsForHistory) > 1 {
		t.Error("Too many new pbs")
	}

	if len(newPBsForHistory) > 1 {
		t.Error("Too many new pbs")
	}

	if len(obsoletePBs) < 2 {
		t.Error("Should have two obsolete PBs")
	}

	if obsoletePBs[0].PersonalBest.Date != oldPB4.Date {
		t.Error("Incorrect first obsolete pb")
	}

	if obsoletePBs[1].PersonalBest.Date != oldPB5.Date {
		t.Error("Incorrect first obsolete pb")
	}
}

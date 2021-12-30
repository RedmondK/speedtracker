package speedtrackertypes

import (
	"testing"
	"time"
)

func TestUpdatePBsFromEmpty(t *testing.T) {
	testSwings := []Swing {
	}

	testTime := time.Now()
	testPBs := []PersonalBest {}
	testSwings = append(testSwings, Swing { Side: "dominant", Colour: "green", Position: "normal", Speed: 144 })
	
	testPBs = UpdatePersonalBests(testTime, testSwings, testPBs)

	if(len(testPBs) == 0) {
		t.Error("PBs not set")
	}

	if(len(testPBs) > 1) {
		t.Error("Too many PBs")
	}

	if(testPBs[0].Date != testTime) {
		t.Error("Incorrect date on pb")
	}

	if(testPBs[0].Swing.Colour != "green") {
		t.Error("Incorrect colour on pb")
	}

	if(testPBs[0].Swing.Position != "normal") {
		t.Error("Incorrect position on pb")
	}

	if(testPBs[0].Swing.Side != "dominant") {
		t.Error("Incorrect side on pb")
	}
	if(testPBs[0].Swing.Speed != 144) {
		t.Error("Incorrect speed on pb")
	}
}

func TestUpdatePBFromExisting(t *testing.T) {
	testSwings := []Swing {
	}

	testTime := time.Now()
	testPBs := []PersonalBest {}
	
	testPBs = append(testPBs, PersonalBest{ Date: time.Now(), Swing: Swing { Side: "dominant", Colour: "green", Position: "normal", Speed: 140 } })
	testSwings = append(testSwings, Swing { Side: "dominant", Colour: "green", Position: "normal", Speed: 160 })

	UpdatePersonalBests(testTime, testSwings, testPBs)

	if(len(testPBs) == 0) {
		t.Error("PBs not set")
	}

	if(len(testPBs) > 1) {
		t.Error("Too many PBs")
	}

	if(testPBs[0].Date != testTime) {
		t.Error("Incorrect date on pb")
	}

	if(testPBs[0].Swing.Colour != "green") {
		t.Error("Incorrect colour on pb")
	}

	if(testPBs[0].Swing.Position != "normal") {
		t.Error("Incorrect position on pb")
	}

	if(testPBs[0].Swing.Side != "dominant") {
		t.Error("Incorrect side on pb")
	}
	if(testPBs[0].Swing.Speed != 160) {
		t.Error("Incorrect speed on pb")
	}
}

func TestUpdateMultiplePBSomeExisting(t *testing.T) {
	testSwings := []Swing {
	}

	testTime := time.Now()
	testPBs := []PersonalBest {}
	
	testPBs = append(testPBs, PersonalBest{ Date: time.Now(), Swing: Swing { Side: "dominant", Colour: "green", Position: "normal", Speed: 140 } })
	testSwings = append(testSwings, Swing { Side: "dominant", Colour: "green", Position: "normal", Speed: 160 })
	testSwings = append(testSwings, Swing { Side: "non-dominant", Colour: "red", Position: "max out", Speed: 125 })

	testPBs = UpdatePersonalBests(testTime, testSwings, testPBs)

	if(len(testPBs) == 0) {
		t.Error("PBs not set")
	}

	if(len(testPBs) > 2) {
		t.Error("Too many PBs")
	}

	if(len(testPBs) < 2) {
		t.Error("Too few PBs")
	}

	if(testPBs[0].Date != testTime) {
		t.Error("Incorrect date on pb")
	}

	if(testPBs[0].Swing.Colour != "green") {
		t.Error("Incorrect colour on pb")
	}

	if(testPBs[0].Swing.Position != "normal") {
		t.Error("Incorrect position on pb")
	}

	if(testPBs[0].Swing.Side != "dominant") {
		t.Error("Incorrect side on pb")
	}
	if(testPBs[0].Swing.Speed != 160) {
		t.Error("Incorrect speed on pb")
	}

	if(testPBs[1].Date != testTime) {
		t.Error("Incorrect date on pb")
	}

	if(testPBs[1].Swing.Colour != "red") {
		t.Error("Incorrect colour on pb")
	}

	if(testPBs[1].Swing.Position != "max out") {
		t.Error("Incorrect position on pb")
	}

	if(testPBs[1].Swing.Side != "non-dominant") {
		t.Error("Incorrect side on pb")
	}
	if(testPBs[1].Swing.Speed != 125) {
		t.Error("Incorrect speed on pb")
	}
}

func TestUpdateMultiplePB(t *testing.T) {
	testSwings := []Swing {
	}

	testTime := time.Now()
	testPBs := []PersonalBest {}
	oldDate, _ := time.Parse(time.RFC3339, "2021-01-01T22:45:02Z")
	
	testPBs = append(testPBs, PersonalBest{ Date: time.Now(), Swing: Swing { Side: "dominant", Colour: "green", Position: "normal", Speed: 140 } })
	testPBs = append(testPBs, PersonalBest{ Date: time.Now(), Swing: Swing { Side: "non-dominant", Colour: "blue", Position: "normal", Speed: 115 } })
	testPBs = append(testPBs, PersonalBest{ Date: time.Now(), Swing: Swing { Side: "dominant", Colour: "red", Position: "happy gilmore", Speed: 123 } })
	testPBs = append(testPBs, PersonalBest{ Date: time.Now(), Swing: Swing { Side: "non-dominant", Colour: "red", Position: "max out", Speed: 122 } })
	testPBs = append(testPBs, PersonalBest{ Date: time.Now(), Swing: Swing { Side: "dominant", Colour: "green", Position: "step-change", Speed: 140 } })
	testPBs = append(testPBs, PersonalBest{ Date: oldDate, Swing: Swing { Side: "dominant", Colour: "green", Position: "olddate", Speed: 140 } })
	testSwings = append(testSwings, Swing { Side: "dominant", Colour: "green", Position: "normal", Speed: 160 }) // update
	testSwings = append(testSwings, Swing { Side: "non-dominant", Colour: "blue", Position: "normal", Speed: 125 }) // update
	testSwings = append(testSwings, Swing { Side: "non-dominant", Colour: "red", Position: "max out", Speed: 128 }) // update
	testSwings = append(testSwings, Swing { Side: "dominant", Colour: "red", Position: "max out", Speed: 125 }) // new
	testSwings = append(testSwings, Swing { Side: "non-dominant", Colour: "blue", Position: "sprint", Speed: 156 }) //new

	testPBs = UpdatePersonalBests(testTime, testSwings, testPBs)

	if(len(testPBs) == 0) {
		t.Error("PBs not set")
	}

	if(len(testPBs) > 8) {
		t.Error("Too many PBs")
	}

	if(len(testPBs) < 8) {
		t.Error("Too few PBs")
	}

	if(testPBs[0].Date != testTime) {
		t.Error("Incorrect date on pb")
	}

	if(testPBs[0].Swing.Colour != "green") {
		t.Error("Incorrect colour on pb")
	}

	if(testPBs[0].Swing.Position != "normal") {
		t.Error("Incorrect position on pb")
	}

	if(testPBs[0].Swing.Side != "dominant") {
		t.Error("Incorrect side on pb")
	}
	if(testPBs[0].Swing.Speed != 160) {
		t.Error("Incorrect speed on pb")
	}

	if(testPBs[1].Date != testTime) {
		t.Error("Incorrect date on pb")
	}

	if(testPBs[1].Swing.Colour != "blue") {
		t.Error("Incorrect colour on pb")
	}

	if(testPBs[1].Swing.Position != "normal") {
		t.Error("Incorrect position on pb")
	}

	if(testPBs[1].Swing.Side != "non-dominant") {
		t.Error("Incorrect side on pb")
	}
	if(testPBs[1].Swing.Speed != 125) {
		t.Error("Incorrect speed on pb")
	}

	if(testPBs[2].Date != testTime) {
		t.Error("Incorrect date on pb")
	}

	if(testPBs[2].Swing.Colour != "red") {
		t.Error("Incorrect colour on pb")
	}

	if(testPBs[2].Swing.Position != "happy gilmore") {
		t.Error("Incorrect position on pb")
	}

	if(testPBs[2].Swing.Side != "dominant") {
		t.Error("Incorrect side on pb")
	}
	if(testPBs[2].Swing.Speed != 123) {
		t.Error("Incorrect speed on pb")
	}

	if(testPBs[3].Swing.Speed != 128) {
		t.Error("Incorrect speed on pb")
	}

	if(testPBs[6].Swing.Speed != 125) {
		t.Error("Incorrect speed on pb")
	}

	if(testPBs[7].Swing.Speed != 156) {
		t.Error("Incorrect speed on pb")
	}

	if(testPBs[5].Date != oldDate) {
		t.Error("Incorrect date on oldDate PB which should not change")
	}
}
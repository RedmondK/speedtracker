package main

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/RedmondK/speedtracker/speedtrackertypes"
)

func TestFromEmpty(t *testing.T) {
	testTime := time.Now()

	testSessions := []speedtrackertypes.Session{
		{
			Date: testTime,
			Swings: []speedtrackertypes.Swing{
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 144},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 148},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 152},
			},
		},
	}

	newCurrentPBs, newPBsForHistory := speedtrackertypes.GetPersonalBestData(testSessions)

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
	if newCurrentPBs[0].Swing.Speed != 152 {
		t.Error("Incorrect speed on pb")
	}
}

func TestFromEmptyMultiSession(t *testing.T) {
	testTime, _ := time.Parse(time.RFC3339, "2022-01-08T22:30:02Z")
	testTime2, _ := time.Parse(time.RFC3339, "2022-01-08T22:35:02Z")

	testSessions := []speedtrackertypes.Session{
		{
			Date: testTime,
			Swings: []speedtrackertypes.Swing{
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 144},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 148},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 152},
			},
		},
		{
			Date: testTime2,
			Swings: []speedtrackertypes.Swing{
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 144},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 148},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 153},
			},
		},
	}

	newCurrentPBs, newPBsForHistory := speedtrackertypes.GetPersonalBestData(testSessions)

	if len(newCurrentPBs) == 0 {
		t.Error("PBs not set")
	}

	if len(newCurrentPBs) > 1 {
		t.Error("Too many PBs")
	}

	if newCurrentPBs[0].Date != testTime2 {
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

	if newCurrentPBs[0].Swing.Speed != 153 {
		t.Error("Incorrect speed on pb")
	}

	if len(newPBsForHistory) != 2 {
		t.Error("PB History incorrect length")
		log.Print(newPBsForHistory)
	}

	if newPBsForHistory[0].Speed != 152 {
		t.Errorf("PB History incorrect speed, not 152: %s", strconv.Itoa(newPBsForHistory[0].Speed))
		log.Print(newPBsForHistory)
	}

	if newPBsForHistory[1].Speed != 153 {
		t.Errorf("PB History incorrect speed, not 153: %s", strconv.Itoa(newPBsForHistory[1].Speed))
		log.Print(newPBsForHistory)
	}
}

func TestFromEmptyMultiSessionMultiType(t *testing.T) {
	testTime, _ := time.Parse(time.RFC3339, "2022-01-08T22:30:02Z")
	testTime2, _ := time.Parse(time.RFC3339, "2022-01-08T22:35:02Z")
	testTime3, _ := time.Parse(time.RFC3339, "2022-01-08T22:45:02Z")

	testSessions := []speedtrackertypes.Session{
		{
			Date: testTime,
			Swings: []speedtrackertypes.Swing{
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 144},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 148},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 152},
				{Side: "non-dominant", Colour: "green", Position: "normal", Speed: 85},
				{Side: "non-dominant", Colour: "green", Position: "normal", Speed: 90},
				{Side: "non-dominant", Colour: "green", Position: "normal", Speed: 95},
			},
		},
		{
			Date: testTime2,
			Swings: []speedtrackertypes.Swing{
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 144},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 148},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 153},
				{Side: "dominant", Colour: "red", Position: "step-change", Speed: 170},
				{Side: "dominant", Colour: "red", Position: "step-change", Speed: 172},
				{Side: "dominant", Colour: "red", Position: "step-change", Speed: 190},
			},
		},
		{
			Date: testTime3,
			Swings: []speedtrackertypes.Swing{
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 144},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 148},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 165},
				{Side: "non-dominant", Colour: "blue", Position: "step-change", Speed: 107},
				{Side: "non-dominant", Colour: "blue", Position: "step-change", Speed: 109},
				{Side: "non-dominant", Colour: "blue", Position: "step-change", Speed: 110},
			},
		},
	}

	newCurrentPBs, newPBsForHistory := speedtrackertypes.GetPersonalBestData(testSessions)

	if len(newCurrentPBs) != 4 {
		t.Error("PBs not set correctly")
		log.Print(newCurrentPBs)
	}

	if newCurrentPBs[0].Date != testTime3 {
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

	if newCurrentPBs[0].Swing.Speed != 165 {
		t.Error("Incorrect speed on pb")
	}

	if newCurrentPBs[1].Date != testTime {
		t.Error("Incorrect date on pb")
	}

	if newCurrentPBs[1].Swing.Colour != "green" {
		t.Error("Incorrect colour on pb")
	}

	if newCurrentPBs[1].Swing.Position != "normal" {
		t.Error("Incorrect position on pb")
	}

	if newCurrentPBs[1].Swing.Side != "non-dominant" {
		t.Error("Incorrect side on pb")
	}

	if newCurrentPBs[1].Swing.Speed != 95 {
		t.Error("Incorrect speed on pb")
	}

	if newCurrentPBs[2].Date != testTime2 {
		t.Error("Incorrect date on pb")
	}

	if newCurrentPBs[2].Swing.Colour != "red" {
		t.Error("Incorrect colour on pb")
	}

	if newCurrentPBs[2].Swing.Position != "step-change" {
		t.Error("Incorrect position on pb")
	}

	if newCurrentPBs[2].Swing.Side != "dominant" {
		t.Error("Incorrect side on pb")
	}

	if newCurrentPBs[2].Swing.Speed != 190 {
		t.Error("Incorrect speed on pb")
	}

	if newCurrentPBs[3].Date != testTime3 {
		t.Error("Incorrect date on pb")
	}

	if newCurrentPBs[3].Swing.Colour != "blue" {
		t.Error("Incorrect colour on pb")
	}

	if newCurrentPBs[3].Swing.Position != "step-change" {
		t.Error("Incorrect position on pb")
	}

	if newCurrentPBs[3].Swing.Side != "non-dominant" {
		t.Error("Incorrect side on pb")
	}

	if newCurrentPBs[3].Swing.Speed != 110 {
		t.Error("Incorrect speed on pb")
	}

	if len(newPBsForHistory) != 6 {
		t.Error("PB History incorrect length")
		log.Print(newPBsForHistory)
	}

	if newPBsForHistory[0].Speed != 152 {
		t.Errorf("PB History incorrect speed, not 152: %s", strconv.Itoa(newPBsForHistory[0].Speed))
		log.Print(newPBsForHistory)
	}

	if newPBsForHistory[1].Speed != 95 {
		t.Errorf("PB History incorrect speed, not 95: %s", strconv.Itoa(newPBsForHistory[1].Speed))
		log.Print(newPBsForHistory)
	}

	if newPBsForHistory[2].Speed != 153 {
		t.Errorf("PB History incorrect speed, not 153: %s", strconv.Itoa(newPBsForHistory[2].Speed))
		log.Print(newPBsForHistory)
	}
}

func TestFromEmptyMultiSessionMultiTypeNonSequential(t *testing.T) {
	testTime, _ := time.Parse(time.RFC3339, "2022-01-08T22:30:02Z")
	testTime2, _ := time.Parse(time.RFC3339, "2022-01-08T22:35:02Z")
	testTime3, _ := time.Parse(time.RFC3339, "2022-01-08T22:45:02Z")

	testSessions := []speedtrackertypes.Session{
		{
			Date: testTime,
			Swings: []speedtrackertypes.Swing{
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 144},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 148},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 152},
				{Side: "non-dominant", Colour: "green", Position: "normal", Speed: 85},
				{Side: "non-dominant", Colour: "green", Position: "normal", Speed: 90},
				{Side: "non-dominant", Colour: "green", Position: "normal", Speed: 95},
			},
		},
		{
			Date: testTime3,
			Swings: []speedtrackertypes.Swing{
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 144},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 148},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 165},
				{Side: "non-dominant", Colour: "blue", Position: "step-change", Speed: 107},
				{Side: "non-dominant", Colour: "blue", Position: "step-change", Speed: 109},
				{Side: "non-dominant", Colour: "blue", Position: "step-change", Speed: 110},
			},
		},
		{
			Date: testTime2,
			Swings: []speedtrackertypes.Swing{
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 144},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 148},
				{Side: "dominant", Colour: "green", Position: "normal", Speed: 153},
				{Side: "dominant", Colour: "red", Position: "step-change", Speed: 170},
				{Side: "dominant", Colour: "red", Position: "step-change", Speed: 172},
				{Side: "dominant", Colour: "red", Position: "step-change", Speed: 190},
			},
		},
	}

	newCurrentPBs, newPBsForHistory := speedtrackertypes.GetPersonalBestData(testSessions)

	if len(newCurrentPBs) != 4 {
		t.Error("PBs not set correctly")
		log.Print(newCurrentPBs)
	}

	if newCurrentPBs[0].Date != testTime3 {
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

	if newCurrentPBs[0].Swing.Speed != 165 {
		t.Error("Incorrect speed on pb")
	}

	if newCurrentPBs[1].Date != testTime {
		t.Error("Incorrect date on pb")
	}

	if newCurrentPBs[1].Swing.Colour != "green" {
		t.Error("Incorrect colour on pb")
	}

	if newCurrentPBs[1].Swing.Position != "normal" {
		t.Error("Incorrect position on pb")
	}

	if newCurrentPBs[1].Swing.Side != "non-dominant" {
		t.Error("Incorrect side on pb")
	}

	if newCurrentPBs[1].Swing.Speed != 95 {
		t.Error("Incorrect speed on pb")
	}

	if newCurrentPBs[2].Date != testTime2 {
		t.Error("Incorrect date on pb")
	}

	if newCurrentPBs[2].Swing.Colour != "red" {
		t.Error("Incorrect colour on pb")
	}

	if newCurrentPBs[2].Swing.Position != "step-change" {
		t.Error("Incorrect position on pb")
	}

	if newCurrentPBs[2].Swing.Side != "dominant" {
		t.Error("Incorrect side on pb")
	}

	if newCurrentPBs[2].Swing.Speed != 190 {
		t.Error("Incorrect speed on pb")
	}

	if newCurrentPBs[3].Date != testTime3 {
		t.Error("Incorrect date on pb")
	}

	if newCurrentPBs[3].Swing.Colour != "blue" {
		t.Error("Incorrect colour on pb")
	}

	if newCurrentPBs[3].Swing.Position != "step-change" {
		t.Error("Incorrect position on pb")
	}

	if newCurrentPBs[3].Swing.Side != "non-dominant" {
		t.Error("Incorrect side on pb")
	}

	if newCurrentPBs[3].Swing.Speed != 110 {
		t.Error("Incorrect speed on pb")
	}

	if len(newPBsForHistory) != 6 {
		t.Error("PB History incorrect length")
		log.Print(newPBsForHistory)
	}

	if newPBsForHistory[0].Speed != 152 {
		t.Errorf("PB History incorrect speed, not 152: %s", strconv.Itoa(newPBsForHistory[0].Speed))
		log.Print(newPBsForHistory)
	}

	if newPBsForHistory[1].Speed != 95 {
		t.Errorf("PB History incorrect speed, not 95: %s", strconv.Itoa(newPBsForHistory[1].Speed))
		log.Print(newPBsForHistory)
	}

	if newPBsForHistory[2].Speed != 153 {
		t.Errorf("PB History incorrect speed, not 153: %s", strconv.Itoa(newPBsForHistory[2].Speed))
		log.Print(newPBsForHistory)
	}
}

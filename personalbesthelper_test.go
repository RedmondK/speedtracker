package main

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/RedmondK/speedtracker/speedtrackertypes"
)

func TestIdentifyObsoleteRecords(t *testing.T) {
	existingDate1, _ := time.Parse(time.RFC3339, "2022-01-07T22:30:02Z")
	existingPB1 := speedtrackertypes.PersonalBest{Date: existingDate1, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 201}}

	existingDate2, _ := time.Parse(time.RFC3339, "2022-01-08T22:30:02Z")
	existingPB2 := speedtrackertypes.PersonalBest{Date: existingDate2, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 16500}}

	existingDate3, _ := time.Parse(time.RFC3339, "2022-01-11T22:30:02Z")
	existingPB3 := speedtrackertypes.PersonalBest{Date: existingDate3, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 18002}}

	existingDate4, _ := time.Parse(time.RFC3339, "2022-01-20T22:30:02Z")
	existingPB4 := speedtrackertypes.PersonalBest{Date: existingDate4, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 18005}}

	existing := []speedtrackertypes.PersonalBestHistoryRecord{
		{
			Speed:        200,
			PersonalBest: existingPB1,
		},
		{
			Speed:        16500,
			PersonalBest: existingPB2,
		},
		{
			Speed:        18002,
			PersonalBest: existingPB3,
		},
		{
			Speed:        18005,
			PersonalBest: existingPB4,
		},
	}

	newDate1, _ := time.Parse(time.RFC3339, "2022-01-07T22:30:02Z")
	newPB1 := speedtrackertypes.PersonalBest{Date: newDate1, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 201}}

	newDate2, _ := time.Parse(time.RFC3339, "2022-01-08T22:30:02Z")
	newPB2 := speedtrackertypes.PersonalBest{Date: newDate2, Swing: speedtrackertypes.Swing{Side: "dominant", Colour: "green", Position: "normal", Speed: 16500}}

	new := []speedtrackertypes.PersonalBestHistoryRecord{{
		Speed:        201,
		PersonalBest: newPB1,
	}, {
		Speed:        16500,
		PersonalBest: newPB2,
	}}

	obsolete := speedtrackertypes.IdentifyObsoletePBHistoryRecords(existing, new)

	if len(obsolete) != 2 {
		t.Errorf("Incorrect obsolete pbs, expect 2, got %s", strconv.Itoa(len(obsolete)))
		log.Print(obsolete)
	}
}

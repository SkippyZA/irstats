package irstats_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/skippyza/irstats"
	"github.com/stretchr/testify/assert"
)

func TestLastRaceStats(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := "/memberstats/member/GetLastRacesStats"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, loadFixture("testdata/last_race_stats.json"))
	})

	result, err := client.LastRaceStats(irstats.String("2000"))
	assert.NoError(t, err)

	expected := []irstats.RaceStats{
		{
			Date:                 "2018-05-04",
			WinnerName:           "Steven+Inskip",
			QualifyTime:          0,
			TrackID:              178,
			LicenseLevel:         17,
			Laps:                 32,
			TrackName:            "Indianapolis+Motor+Speedway",
			SOF:                  1576,
			CarID:                4,
			CarColor1:            "FFE200",
			CarColor2:            "132347",
			CarColor3:            "125B30",
			WinnerID:             188953,
			QualifyTimeFormatted: "00.000",
			Incidents:            8,
			ClubPoints:           10,
			SubSessionID:         23059878,
			ChampPoints:          92,
			WinnerHC1:            "FFE200",
			WinnerHC3:            "132347",
			WinnerHC2:            "125B30",
			CarClassID:           4,
			SeriesID:             44,
			WinnerHPattern:       64,
			StartPos:             2,
			CarPattern:           11,
			WinnerLL:             17,
			SeasonID:             2054,
			LapsLed:              9,
			Time:                 1525441500000,
			FinishPosition:       1,
		},
	}
	assert.Equal(t, expected, result)
}

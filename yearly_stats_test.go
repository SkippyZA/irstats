package irstats_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/skippyza/irstats"
	"github.com/stretchr/testify/assert"
)

func TestYearlyStats(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(irstats.UrlPathYearlyStats, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, loadFixture("testdata/yearly_stats.json"))
	})

	result, err := client.YearlyStats(irstats.String("2000"))
	assert.NoError(t, err)

	expected := &irstats.YearlyStats{
		{
			AvgFinish:     10,
			AvgIncPerRace: 5.264200210571289,
			AvgPtsPerRace: 63,
			AvgStart:      9,
			Category:      "Road",
			ClubPoints:    1076,
			LapsLed:       253,
			LapsLedPerc:   7.150000095367432,
			Poles:         17,
			Starts:        247,
			Top5:          69,
			Top5Perc:      27.940000534057617,
			TotalLaps:     3536,
			WinPerc:       6.070000171661377,
			Wins:          15,
			Year:          "2016",
		},
	}
	assert.Equal(t, expected, result)
}

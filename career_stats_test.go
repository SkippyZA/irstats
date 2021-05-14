package irstats_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/skippyza/irstats"
	"github.com/stretchr/testify/assert"
)

func TestCareerStats(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(irstats.UrlPathCareerStats, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, loadFixture("testdata/career_stats.json"))
	})

	result, _, err := client.CareerStats(irstats.String("2000"))
	assert.NoError(t, err)

	expected := &irstats.CareerStats{
		{
			AvgFinish:       10,
			AvgIncPerRace:   5.185299873352051,
			AvgPtsPerRace:   61,
			AvgStart:        9,
			Category:        "Road",
			LapsLed:         335,
			LapsLedPerC:     4.449999809265137,
			Poles:           19,
			Starts:          505,
			Top5:            134,
			Top5PerC:        26.530000686645508,
			TotalLaps:       7524,
			TotalClubPoints: 2401,
			WinPerc:         3.9600000381469727,
			Wins:            20,
		},
	}
	assert.Equal(t, expected, result)
}

package irstats

import (
	"net/http"

	"github.com/skippyza/irstats/internal/form"
)

type EventResults struct{}

type EventResultsRequest struct {
	CustID          *string    `form:"custid"`
	ShowRaces       *bool      `form:"showraces"`
	ShowQuals       *bool      `form:"showquals"`
	ShowTimeTrials  *bool      `form:"showtts"`
	ShowPractice    *bool      `form:"showops"`
	ShowOfficial    *bool      `form:"showofficial"`
	ShowUnofficial  *bool      `form:"showunofficial"`
	ShowRookie      *bool      `form:"showrookie"`
	ShowClassD      *bool      `form:"showclassd"`
	ShowClassC      *bool      `form:"showclassc"`
	ShowClassB      *bool      `form:"showclassb"`
	ShowClassA      *bool      `form:"showclassa"`
	ShowPro         *bool      `form:"showpro"`
	ShowProWC       *bool      `form:"showprowc"`
	LowerBound      *int       `form:"lowerbound"`
	UpperBound      *int       `form:"upperbound"`
	Sort            sortType   `form:"sort"`
	Order           order      `form:"order"`
	Format          *int       `form:"format"`
	Category        []category `form:"category[]"`
	SeasonYear      *int       `form:"seasonyear"`
	SeasonQuarter   *int       `form:"seasonquarter"`
	RaceWeek        *int       `form:"raceweek"`
	TrackID         *int       `form:"trackid"`
	CarClassID      *int       `form:"carclassid"`
	CarID           *int       `form:"carid"`
	StartLow        *int       `form:"start_low"`
	StartHigh       *int       `form:"start_high"`
	FinishLow       *int       `form:"finish_low"`
	FinishHigh      *int       `form:"finish_high"`
	IncidentsLow    *int       `form:"incidents_low"`
	IncidentsHigh   *int       `form:"incidents_high"`
	ChamppointsLow  *int       `form:"champpoints_low"`
	ChamppointsHigh *int       `form:"champpoints_high"`
}

// Results returns a list with an EventResults object for each of a driver's
// past events that meet the selected criteria.
func (c *Client) EventResults(opts *EventResultsRequest) (*map[string]interface{}, *http.Response, error) {
	v := form.ParseFormTags(*opts)
	results := &map[string]interface{}{}
	resp, err := c.do(URLPathResults, &v, results)
	return results, resp, err
}

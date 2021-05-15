package irstats

import (
	"net/http"

	"github.com/skippyza/irstats/internal/form"
)

type Results struct{}

type ResultsRequest struct {
	CustID          *string    `form:"custid"`
	ShowRaces       *int       `form:"showraces"`
	ShowQuals       *int       `form:"showquals"`
	ShowTimeTrials  *int       `form:"showtts"`
	ShowPractice    *int       `form:"showops"`
	ShowOfficial    *int       `form:"showofficial"`
	ShowUnofficial  *int       `form:"showunofficial"`
	ShowRookie      *int       `form:"showrookie"`
	ShowClassD      *int       `form:"showclassd"`
	ShowClassC      *int       `form:"showclassc"`
	ShowClassB      *int       `form:"showclassb"`
	ShowClassA      *int       `form:"showclassa"`
	ShowPro         *int       `form:"showpro"`
	ShowProWC       *int       `form:"showprowc"`
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

func (c *Client) Results(opts *ResultsRequest) (*map[string]interface{}, *http.Response, error) {
	params := form.ParseFormTags(*opts)
	results := &map[string]interface{}{}
	resp, err := c.do(URLPathResults, &params, results)

	return results, resp, err
}

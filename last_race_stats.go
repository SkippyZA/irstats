package irstats

import "net/http"

type RaceStatsItem struct {
	Date                 string `json:"date"`
	WinnerName           string `json:"winnerName"`
	QualifyTime          int    `json:"qualifyTime"`
	TrackID              int    `json:"trackID"`
	LicenseLevel         int    `json:"licenseLevel"`
	Laps                 int    `json:"laps"`
	TrackName            string `json:"trackName"`
	SOF                  int    `json:"sof"`
	CarID                int    `json:"carID"`
	CarColor1            string `json:"carColor1"`
	CarColor2            string `json:"carColor2"`
	CarColor3            string `json:"carColor3"`
	WinnerID             int    `json:"winnerID"`
	QualifyTimeFormatted string `json:"qualifyTimeFormatted"`
	Incidents            int    `json:"incidents"`
	ClubPoints           int    `json:"clubPoints"`
	SubSessionID         int    `json:"subsessionID"`
	ChampPoints          int    `json:"champPoints"`
	WinnerHC1            string `json:"winnerHC1"`
	WinnerHC3            string `json:"winnerHC3"`
	WinnerHC2            string `json:"winnerHC2"`
	CarClassID           int    `json:"carClassID"`
	SeriesID             int    `json:"seriesID"`
	WinnerHPattern       int    `json:"winnerHPattern"`
	StartPos             int    `json:"startPos"`
	CarPattern           int    `json:"carPattern"`
	WinnerLL             int    `json:"winnerLL"`
	SeasonID             int    `json:"seasonID"`
	LapsLed              int    `json:"lapsLed"`
	Time                 int64  `json:"time"`
	FinishPosition       int    `json:"finishPos"`
}

type RaceStats = []RaceStatsItem

// LastRaceStats returns stat summary for the driver's last 10 races as seen on the /CareerStats page.
func (c *Client) LastRaceStats(custID *string) (*RaceStats, *http.Response, error) {
	raceStats := &RaceStats{}
	resp, err := c.do(URLPathLastRaceStats, &map[string][]string{"custid": {*custID}}, raceStats)

	return raceStats, resp, err
}

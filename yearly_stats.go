package irstats

import "net/http"

type YearlyStatsItem struct {
	AvgFinish     int     `json:"avgFinish"`
	AvgIncPerRace float64 `json:"avgIncPerRace"`
	AvgPtsPerRace int     `json:"avgPtsPerRace"`
	AvgStart      int     `json:"avgStart"`
	Category      string  `json:"category"`
	ClubPoints    int     `json:"clubpoints"`
	LapsLed       int     `json:"lapsLed"`
	LapsLedPerc   float64 `json:"lapsLedPerc"`
	Poles         int     `json:"poles"`
	Starts        int     `json:"starts"`
	Top5          int     `json:"top5"`
	Top5Perc      float64 `json:"top5Perc"`
	TotalLaps     int     `json:"totalLaps"`
	WinPerc       float64 `json:"winPerc"`
	Wins          int     `json:"wins"`
	Year          string  `json:"year"`
}

type YearlyStats = []YearlyStatsItem

// YearlyStats returns the breakdown of career stats by year, as seen on the /CareerStats driver profile.
func (c *Client) YearlyStats(custID *string) (*YearlyStats, *http.Response, error) {
	p := map[string]string{
		"custid": *custID,
	}

	yearlyStats := &YearlyStats{}
	resp, err := c.do(URLPathYearlyStats, &p, yearlyStats)
	return yearlyStats, resp, err
}

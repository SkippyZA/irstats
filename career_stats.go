package irstats

type CareerStatItem struct {
	AvgFinish       int     `json:"avgFinish"`
	AvgIncPerRace   float64 `json:"avgIncPerRace"`
	AvgPtsPerRace   int     `json:"avgPtsPerRace"`
	AvgStart        int     `json:"avgStart"`
	Category        string  `json:"category"`
	LapsLed         int     `json:"lapsLed"`
	LapsLedPerC     float64 `json:"lapsLedPerc"`
	Poles           int     `json:"poles"`
	Starts          int     `json:"starts"`
	Top5            int     `json:"top5"`
	Top5PerC        float64 `json:"top5Perc"`
	TotalLaps       int     `json:"totalLaps"`
	TotalClubPoints int     `json:"totalclubpoints"`
	WinPerc         float64 `json:"winPerc"`
	Wins            int     `json:"wins"`
}

type CareerStats = []CareerStatItem

// CareerStats Returns driver career stats as seen on the driver profile page.
// E.g. Starts, Avg Inc., Win %, etc.
func (c *Client) CareerStats(custID *string) (*CareerStats, error) {
	p := map[string]string{
		"custid": *custID,
	}

	careerStats := &CareerStats{}
	err := c.do(UrlPathCareerStats, &p, careerStats)
	return careerStats, err
}

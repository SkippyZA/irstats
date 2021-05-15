package irstats

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/skippyza/irstats/internal/form"
)

type DriverStats struct {
	AveFieldSize         int     `json:"avefieldsize"`
	AveFinishingPosition int     `json:"avefinishingposition"`
	AveStartingPosition  int     `json:"avestartingposition"`
	AvgIncidents         float64 `json:"avgincidents"`
	AvgPoints            int     `json:"avgpoints"`
	ClubID               int     `json:"clubid"`
	ClubName             string  `json:"clubname"`
	ClubPoints           int     `json:"clubpoints"`
	CountryCode          string  `json:"countrycode"`
	CustID               int     `json:"custid"`
	DisplayName          string  `json:"displayname"`
	GroupLetter          string  `json:"groupletter"`
	GroupName            string  `json:"groupname"`
	HelmColor1           string  `json:"helmcolor1"`
	HelmColor2           string  `json:"helmcolor2"`
	HelmColor3           string  `json:"helmcolor3"`
	HelmFaceType         int     `json:"helmfacetype"`
	HelmHelmetType       int     `json:"helmhelmettype"`
	HelmPattern          int     `json:"helmpattern"`
	IRating              int     `json:"irating"`
	IRatingRank          int     `json:"irating_rank"`
	Laps                 int     `json:"laps"`
	LapsLead             int     `json:"lapslead"`
	LicenseClass         string  `json:"licenseclass"`
	LicenseClassRank     int     `json:"licenseclass_rank"`
	LicenseGroup         int     `json:"licensegroup"`
	LicenseLevel         int     `json:"licenselevel"`
	Points               int     `json:"points"`
	RN                   int     `json:"rn"`
	Rank                 int     `json:"rank"`
	Region               string  `json:"region"`
	Starts               int     `json:"starts"`
	SubLevel             int     `json:"sublevel"`
	TTRating             int     `json:"ttrating"`
	TTRatingRank         int     `json:"ttrating_rank"`
	Top25Pcnt            int     `json:"top25pcnt"`
	Wins                 int     `json:"wins"`
}

func (ds *DriverStats) UnmarshalJSON(b []byte) error {
	r := map[string]interface{}{}
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	rn := 0
	if v, ok := r["37"].(float64); ok {
		rn = int(v)
	}

	ds.Top25Pcnt = int(r["1"].(float64))
	ds.IRatingRank = int(r["2"].(float64))
	ds.IRating = int(r["3"].(float64))
	ds.CountryCode = r["4"].(string)
	ds.ClubID = int(r["5"].(float64))
	ds.TTRatingRank = int(r["6"].(float64))
	ds.Laps = int(r["7"].(float64))
	ds.AveStartingPosition = int(r["8"].(float64))
	ds.AveFinishingPosition = int(r["9"].(float64))
	ds.HelmHelmetType = int(r["10"].(float64))
	ds.Points = int(r["11"].(float64))
	ds.LicenseGroup = int(r["12"].(float64))
	ds.AveFieldSize = int(r["13"].(float64))
	ds.Rank = int(r["14"].(float64))
	ds.Starts = int(r["15"].(float64))
	ds.LicenseClass = r["16"].(string)
	ds.Wins = int(r["18"].(float64))
	ds.LicenseClassRank = int(r["19"].(float64))
	ds.ClubPoints = int(r["21"].(float64))
	ds.HelmPattern = int(r["22"].(float64))
	ds.LicenseLevel = int(r["23"].(float64))
	ds.GroupName = r["24"].(string)
	ds.TTRating = int(r["25"].(float64))
	ds.AvgIncidents = r["26"].(float64)
	ds.AvgPoints = int(r["27"].(float64))
	ds.HelmColor3 = r["28"].(string)
	ds.ClubName = r["29"].(string)
	ds.LapsLead = int(r["30"].(float64))
	ds.HelmColor1 = r["31"].(string)
	ds.DisplayName = r["32"].(string)
	ds.HelmColor2 = r["33"].(string)
	ds.CustID = int(r["34"].(float64))
	ds.SubLevel = int(r["35"].(float64))
	ds.HelmFaceType = int(r["36"].(float64))
	ds.RN = rn
	ds.Region = r["38"].(string)
	ds.GroupLetter = r["39"].(string)

	return nil
}

type DriverStatsResult struct {
	CustRow  int
	RowCount int
	Drivers  []DriverStats
}

func (ds *DriverStatsResult) UnmarshalJSON(b []byte) error {
	dsRaw := struct {
		D json.RawMessage `json:"d"`
	}{}

	err := json.Unmarshal(b, &dsRaw)
	if err != nil {
		return err
	}

	if len(dsRaw.D) == 2 {
		return nil
	}

	objRaw := struct {
		CustRow  interface{}   `json:"17"`
		RowCount interface{}   `json:"20"`
		Rows     []DriverStats `json:"r"`
	}{}
	err = json.Unmarshal(dsRaw.D, &objRaw)
	if err != nil {
		return err
	}

	custRow := 0
	if v, ok := objRaw.CustRow.(string); ok {
		custRow, err = strconv.Atoi(v)
		if err != nil {
			return err
		}
	} else {
		custRow = int(objRaw.CustRow.(float64))
	}

	rowCount := 0
	if v, ok := objRaw.RowCount.(string); ok {
		rowCount, err = strconv.Atoi(v)
		if err != nil {
			return err
		}
	} else {
		rowCount = int(objRaw.RowCount.(float64))
	}

	// Sort the slice by RN. RN is the "ranking" based you the sort + order
	sort.SliceStable(objRaw.Rows, func(i, j int) bool {
		return objRaw.Rows[i].RN < objRaw.Rows[j].RN
	})

	// When RN is 0 it is showing yourself compared to others.
	// TODO: perhaps this should be optional behaviour
	if objRaw.Rows[0].RN == 0 {
		objRaw.Rows = append(objRaw.Rows[:0], objRaw.Rows[1:]...)
	}

	ds.CustRow = custRow
	ds.RowCount = rowCount
	ds.Drivers = objRaw.Rows

	return nil
}

type DriverStatsRequest struct {
	Search           *string  `form:"search"`
	Friend           *string  `form:"friend"`
	Watched          *int     `form:"watched"`
	Recent           *int     `form:"recent"`
	Country          *string  `form:"country"`
	Category         category `form:"category"`
	ClassLow         *int     `form:"classlow"`
	ClassHigh        *int     `form:"classhigh"`
	IRatingLow       *int     `form:"iratinglow"`
	IRatingHigh      *int     `form:"iratinghigh"`
	TTRatingLow      *int     `form:"ttratinglow"`
	TTRatingHigh     *int     `form:"ttratinghigh"`
	AvgStartLow      *int     `form:"avgstartlow"`
	AvgStartHigh     *int     `form:"avgstarthigh"`
	AvgFinishLow     *int     `form:"avgfinishlow"`
	AvgFinishHigh    *int     `form:"avgfinishhigh"`
	AvgPointsLow     *int     `form:"avgpointslow"`
	AvgPointsHigh    *int     `form:"avgpointshigh"`
	AvgIncidentsLow  *int     `form:"avgincidentslow"`
	AvgIncidentsHigh *int     `form:"avgincidentshigh"`
	CustID           *string  `form:"custid"`
	LowerBound       *int     `form:"lowerbound"`
	UpperBound       *int     `form:"upperbound"`
	Sort             sortType `form:"sort"`
	Order            order    `form:"order"`
	Active           *int     `form:"active"`
}

// DriverStats returns a list of drivers that match the given parameters.
//
// This is the backend source for /DriverLookup.Do AKA 'Driver Stats.'
func (c *Client) DriverStats(opts *DriverStatsRequest) (*DriverStatsResult, *http.Response, error) {
	params := form.ParseFormTags(*opts)

	carsDriven := &DriverStatsResult{}
	resp, err := c.do(URLPathDriverStats, &params, carsDriven)
	return carsDriven, resp, err
}

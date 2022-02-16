package irstats

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/skippyza/irstats/internal/form"
)

type Result struct {
	HelmColor1            string `json:"helm_color1"`            // 1
	WinnerHelmColor2      string `json:"winnerhelmcolor2"`       // 2
	FinishingPosition     int    `json:"finishing_position"`     // 3
	WinnerHelmColor3      string `json:"winnerhelmcolor3"`       // 4
	WinnerHelmColor1      string `json:"winnerhelmcolor1"`       // 5
	BestQualLapTime       string `json:"bestquallaptime"`        // 6
	SubsessionBestLapTime string `json:"subsession_bestlaptime"` // 7
	RaceWeekNum           int    `json:"race_week_num"`          // 8
	SessionID             int    `json:"sessionid"`              // 9
	FinishedAt            int    `json:"finishedat"`             // 10
	RawStartTime          int64  `json:"raw_start_time"`         // 11
	StartingPosition      int    `json:"starting_position"`      // 12
	HelmColor3            string `json:"helm_color3"`            // 13
	HelmColor2            string `json:"helm_color2"`            // 14
	Clubpoints            int    `json:"clubpoints"`             // 16
	DropracePoints        int    `json:"dropracepoints"`         // 17
	OfficialSession       int    `json:"officialsession"`        // 18
	GroupName             string `json:"groupname"`              // 19
	SeriesID              int    `json:"seriesid"`               // 20
	StartTime             string `json:"start_time"`             // 21
	SeasonID              int    `json:"seasonid"`               // 22
	CustID                int    `json:"custid"`                 // 23
	HelmLicenseLevel      int    `json:"helm_licenselevel"`      // 24
	WinnerLicenseLevel    int    `json:"winnerlicenselevel"`     // 25
	RN                    int    `json:"rn"`                     // 26
	WinnersGroupID        int    `json:"winnersgroupid"`         // 27
	SesRank               int    `json:"sesrank"`                // 28
	CarClassID            int    `json:"carclassid"`             // 29
	TrackID               int    `json:"trackid"`                // 30
	WinnerDisplayName     string `json:"winnerdisplayname"`      // 31
	CarID                 int    `json:"carid"`                  // 32
	CatID                 int    `json:"catid"`                  // 33
	SeasonQuarter         int    `json:"season_quarter"`         // 34
	LicenseGroup          int    `json:"licensegroup"`           // 35
	WinnerHelmPattern     int    `json:"winnerhelmpattern"`      // 36
	EvtType               int    `json:"evttype"`                // 37
	BestLaptime           string `json:"bestlaptime"`            // 38
	Incidents             int    `json:"incidents"`              // 39
	ChampPoints           int    `json:"champpoints"`            // 40
	SubsessionID          int    `json:"subsessionid"`           // 41
	SeasonYear            int    `json:"season_year"`            // 42
	ChampPointsSort       int    `json:"champpointssort"`        // 43
	StartDate             string `json:"start_date"`             // 44
	StrengthOfField       int    `json:"strengthoffield"`        // 45
	HelmPattern           int    `json:"helm_pattern"`           // 46
	ClubPointsSort        int    `json:"clubpointssort"`         // 47
	DisplayName           string `json:"displayname"`            // 48
}

func (r *Result) UnmarshalJSON(b []byte) error {
	row := map[string]interface{}{}
	err := json.Unmarshal(b, &row)
	if err != nil {
		return err
	}

	rn := 0
	if v, ok := row["37"].(float64); ok {
		rn = int(v)
	}

	r.HelmColor1 = row["1"].(string)
	r.WinnerHelmColor2 = row["2"].(string)
	r.FinishingPosition = int(row["3"].(float64))
	r.WinnerHelmColor3 = row["4"].(string)
	r.WinnerHelmColor1 = row["5"].(string)
	r.BestQualLapTime = row["6"].(string)
	r.SubsessionBestLapTime = row["7"].(string)
	r.RaceWeekNum = int(row["8"].(float64))
	r.SessionID = int(row["9"].(float64))
	r.FinishedAt = int(row["10"].(float64))
	r.RawStartTime = int64(row["11"].(float64))
	r.StartingPosition = int(row["12"].(float64))
	r.HelmColor3 = row["13"].(string)
	r.HelmColor2 = row["14"].(string)
	r.Clubpoints = int(row["16"].(float64))
	r.DropracePoints = int(row["17"].(float64))
	r.OfficialSession = int(row["18"].(float64))
	r.GroupName = row["19"].(string)
	r.SeriesID = int(row["20"].(float64))
	r.StartTime = row["21"].(string)
	r.SeasonID = int(row["22"].(float64))
	r.CustID = int(row["23"].(float64))
	r.HelmLicenseLevel = int(row["24"].(float64))
	r.WinnerLicenseLevel = int(row["25"].(float64))
	r.RN = rn
	r.WinnersGroupID = int(row["27"].(float64))
	r.SesRank = int(row["28"].(float64))
	r.CarClassID = int(row["29"].(float64))
	r.TrackID = int(row["30"].(float64))
	r.WinnerDisplayName = row["31"].(string)
	r.CarID = int(row["32"].(float64))
	r.CatID = int(row["33"].(float64))
	r.SeasonQuarter = int(row["34"].(float64))
	r.LicenseGroup = int(row["35"].(float64))
	r.WinnerHelmPattern = int(row["36"].(float64))
	r.EvtType = int(row["37"].(float64))
	r.BestLaptime = row["38"].(string)
	r.Incidents = int(row["39"].(float64))
	r.ChampPoints = int(row["40"].(float64))
	r.SubsessionID = int(row["41"].(float64))
	r.SeasonYear = int(row["42"].(float64))
	r.ChampPointsSort = int(row["43"].(float64))
	r.StartDate = row["44"].(string)
	r.StrengthOfField = int(row["45"].(float64))
	r.HelmPattern = int(row["46"].(float64))
	r.ClubPointsSort = int(row["47"].(float64))
	r.DisplayName = row["48"].(string)

	return nil
}

type EventResults struct {
	RowCount int
	Results  []Result
}

func (r *EventResults) UnmarshalJSON(b []byte) error {
	dsRaw := struct {
		D json.RawMessage `json:"d"`
		M json.RawMessage `json:"m"`
	}{}

	err := json.Unmarshal(b, &dsRaw)
	if err != nil {
		return err
	}

	if string(dsRaw.M) == "{}" {
		return nil
	}

	objRaw := struct {
		RowCount interface{} `json:"15"`
		Rows     []Result    `json:"r"`
	}{}
	err = json.Unmarshal(dsRaw.D, &objRaw)
	if err != nil {
		return err
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
	r.RowCount = rowCount

	// Sort the slice by RN. RN is the "ranking" based you the sort + order
	sort.SliceStable(objRaw.Rows, func(i, j int) bool {
		return objRaw.Rows[i].RN < objRaw.Rows[j].RN
	})

	// When RN is 0 it is showing yourself compared to others.
	// TODO: perhaps this should be optional behaviour
	if objRaw.Rows[0].RN == 0 {
		objRaw.Rows = append(objRaw.Rows[:0], objRaw.Rows[1:]...)
	}
	r.Results = objRaw.Rows

	return nil
}

type EventResultsRequest struct {
	CustID          *string    `form:"custid"`
	ShowRaces       boolean    `form:"showraces"`
	ShowQuals       boolean    `form:"showquals"`
	ShowTimeTrials  boolean    `form:"showtts"`
	ShowPractice    boolean    `form:"showops"`
	ShowOfficial    boolean    `form:"showofficial"`
	ShowUnofficial  boolean    `form:"showunofficial"`
	ShowRookie      boolean    `form:"showrookie"`
	ShowClassD      boolean    `form:"showclassd"`
	ShowClassC      boolean    `form:"showclassc"`
	ShowClassB      boolean    `form:"showclassb"`
	ShowClassA      boolean    `form:"showclassa"`
	ShowPro         boolean    `form:"showpro"`
	ShowProWC       boolean    `form:"showprowc"`
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
	StartTimeLow    *int       `form:"starttime_low"`
	StartTimeHigh   *int       `form:"starttime_high"`
	FinishLow       *int       `form:"finish_low"`
	FinishHigh      *int       `form:"finish_high"`
	IncidentsLow    *int       `form:"incidents_low"`
	IncidentsHigh   *int       `form:"incidents_high"`
	ChampPointsLow  *int       `form:"champpoints_low"`
	ChampPointsHigh *int       `form:"champpoints_high"`
}

// Results returns a list with an EventResults object for each of a driver's
// past events that meet the selected criteria.
func (c *Client) EventResults(opts *EventResultsRequest) (*EventResults, *http.Response, error) {
	v := form.ParseFormTags(*opts)
	results := &EventResults{}
	resp, err := c.do(URLPathResults, &v, results)
	return results, resp, err
}

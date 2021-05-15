package irstats

import "errors"

const (
	sitePath  = "/membersite/member"
	statsPath = "/memberstats/member"
)

type urlPath = string

// iRacing API paths
var (
	URLPathLogin             = urlPath("/membersite/Login")
	URLPathSubSessionResults = urlPath(sitePath + "/GetSubsessionResults")
	URLPathLastRaceStats     = urlPath(statsPath + "/GetLastRacesStats")
	URLPathCarsDriven        = urlPath(statsPath + "/GetCarsDriven")
	URLPathCareerStats       = urlPath(statsPath + "/GetCareerStats")
	URLPathYearlyStats       = urlPath(statsPath + "/GetYearlyStats")
	URLPathDriverStats       = urlPath(statsPath + "/GetDriverStats")
	URLPathResults           = urlPath(statsPath + "/GetResults")
)

type sortType = string

// Request sorting values
const (
	SortAvgFinishingPosition = sortType("avgfinishingposition")
	SortAvgIncidents         = sortType("avgincidents")
	SortAvgPoints            = sortType("avgpoints")
	SortAvgStartingPosition  = sortType("avgstartingposition")
	SortClass                = sortType("class")
	SortClubName             = sortType("clubname")
	SortClubpoints           = sortType("clubpoints")
	SortCountryCode          = sortType("countrycode")
	SortDisplayName          = sortType("displayname")
	SortIRating              = sortType("irating")
	SortLaps                 = sortType("laps")
	SortLapsLead             = sortType("lapslead")
	SortPoints               = sortType("points")
	SortSessionName          = sortType("sessionname")
	SortStartTime            = sortType("start_time")
	SortStarts               = sortType("starts")
	SortTTRating             = sortType("ttrating")
	SortTop25pcnt            = sortType("top25pcnt")
	SortWins                 = sortType("wins")
)

type order = string

// Request ordering values
const (
	OrderDescending = order("desc")
	OrderAscending  = order("asc")
)

type boolean bool

func (lt boolean) IsValid() error {
	switch lt {
	case true, false:
		return nil
	}
	return errors.New("Invalid boolean type")
}

func (b boolean) String() string {
	if b {
		return "1"
	}
	return "0"
}

const (
	True  = boolean(true)
	False = boolean(false)
)

// // Holds the index for each license type
// const (
// 	LicenseRookie = iota + 1
// 	LicenseD
// 	LicenseC
// 	LicenseB
// 	LicenseA
// 	LicenseP
// 	LicensePWC
// )

type category = int

// Holds the index for each type of racing discipline
const (
	CategoryOval category = iota + 1
	CategoryRoad
	CategoryDirtOval
	CategoryDirtRoad
)

// // Holds the index for the charts available from stats_chart()
// const (
// 	ChartTypeIRating = iota + 1
// 	ChartTypeTTRating
// 	ChartTypeLicenseClass
// )

// // Holds the index for the session event type
// const (
// 	EventTypeTest = iota + 1
// 	EventTypePractice
// 	EventTypeQualify
// 	EventTypeTimeTrial
// 	EventTypeRace
// 	EventTypeOficial
// 	EventTypeUnofficial
// )

// const (
// 	SimSessionTypeRace     = 0
// 	SimSessionTypeQualify  = -1
// 	SimSessionTypePractice = -2
// )

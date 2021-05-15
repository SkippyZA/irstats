package irstats

const (
	sitePath  = "/membersite/member"
	statsPath = "/memberstats/member"
)

type urlPath = string

// iRacing API paths
var (
	URLPathLogin             urlPath = "/membersite/Login"
	URLPathSubSessionResults         = sitePath + "/GetSubsessionResults"
	URLPathLastRaceStats             = statsPath + "/GetLastRacesStats"
	URLPathCarsDriven                = statsPath + "/GetCarsDriven"
	URLPathCareerStats               = statsPath + "/GetCareerStats"
	URLPathYearlyStats               = statsPath + "/GetYearlyStats"
	URLPathDriverStats               = statsPath + "/GetDriverStats"
)

type sortType = string

// Request sorting values
const (
	SortAvgFinishingPosition sortType = "avgfinishingposition"
	SortAvgIncidents                  = "avgincidents"
	SortAvgPoints                     = "avgpoints"
	SortAvgStartingPosition           = "avgstartingposition"
	SortClass                         = "class"
	SortClubName                      = "clubname"
	SortClubpoints                    = "clubpoints"
	SortCountryCode                   = "countrycode"
	SortDisplayName                   = "displayname"
	SortIRating                       = "irating"
	SortLaps                          = "laps"
	SortLapsLead                      = "lapslead"
	SortPoints                        = "points"
	SortSessionName                   = "sessionname"
	SortStartTime                     = "start_time"
	SortStarts                        = "starts"
	SortTTRating                      = "ttrating"
	SortTop25pcnt                     = "top25pcnt"
	SortWins                          = "wins"
)

type order = string

// Request ordering values
const (
	OrderDescending order = "desc"
	OrderAscending        = "asc"
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

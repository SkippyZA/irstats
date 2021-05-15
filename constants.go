package irstats

const (
	sitePath  = "/membersite/member"
	statsPath = "/memberstats/member"
)

var (
	URLPathLogin = "/membersite/Login"
)

// Member site paths
var (
	URLPathSubSessionResults = sitePath + "/GetSubsessionResults"
)

// Member stats paths
var (
	URLPathLastRaceStats = statsPath + "/GetLastRacesStats"
	URLPathCarsDriven    = statsPath + "/GetCarsDriven"
	URLPathCareerStats   = statsPath + "/GetCareerStats"
	URLPathYearlyStats   = statsPath + "/GetYearlyStats"
	URLPathDriverStats   = statsPath + "/GetDriverStats"
)

// Holds the index for each license type
const (
	LicenseRookie = iota + 1
	LicenseD
	LicenseC
	LicenseB
	LicenseA
	LicenseP
	LicensePWC
)

// Holds the index for each type of racing discipline
const (
	CategoryOval = iota + 1
	CategoryRoad
	CategoryDirtOval
	CategoryDirtRoad
)

// Holds the index for the charts available from stats_chart()
const (
	ChartTypeIRating = iota + 1
	ChartTypeTTRating
	ChartTypeLicenseClass
)

// Holds the index for the session event type
const (
	EventTypeTest = iota + 1
	EventTypePractice
	EventTypeQualify
	EventTypeTimeTrial
	EventTypeRace
	EventTypeOficial
	EventTypeUnofficial
)

const (
	SimSessionTypeRace     = 0
	SimSessionTypeQualify  = -1
	SimSessionTypePractice = -2
)

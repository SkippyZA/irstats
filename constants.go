package irstats

const (
	sitePath  = "/membersite/member"
	statsPath = "/memberstats/member"
)

var (
	UrlPathLogin = "/membersite/Login"

	// Member site paths
	UrlPathSubSessionResults = sitePath + "/GetSubsessionResults"

	// Member stats paths
	UrlPathLastRaceStats = statsPath + "/GetLastRacesStats"
	UrlPathCarsDriven    = statsPath + "/GetCarsDriven"
	UrlPathCareerStats   = statsPath + "/GetCareerStats"
	UrlPathYearlyStats   = statsPath + "/GetYearlyStats"
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

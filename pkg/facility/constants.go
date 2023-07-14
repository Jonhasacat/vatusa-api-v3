package facility

type Facility = string

const (
	Academy      Facility = "ZAE"
	Headquarters Facility = "ZHQ"
	NonMember    Facility = "ZZN"
	Inactive     Facility = "ZZI"
	Albuquerque  Facility = "ZAB"
	Anchorage    Facility = "ZAN"
	Atlanta      Facility = "ZTL"
	Boston       Facility = "ZBW"
	Chicago      Facility = "ZAU"
	Cleveland    Facility = "ZOB"
	Denver       Facility = "ZDV"
	FortWorth    Facility = "ZFW"
	Honolulu     Facility = "HCF"
	Houston      Facility = "ZHU"
	Indianapolis Facility = "ZID"
	Jacksonville Facility = "ZJX"
	KansasCity   Facility = "ZKC"
	LosAngeles   Facility = "ZLA"
	Memphis      Facility = "ZME"
	Miami        Facility = "ZMA"
	Minneapolis  Facility = "ZMP"
	NewYork      Facility = "ZNY"
	Oakland      Facility = "ZOA"
	SaltLakeCity Facility = "ZLC"
	Seattle      Facility = "ZSE"
	Washington   Facility = "ZDC"
)

func IsRosterFacility(fac Facility) bool {
	return fac == Albuquerque ||
		fac == Anchorage ||
		fac == Atlanta ||
		fac == Boston ||
		fac == Chicago ||
		fac == Cleveland ||
		fac == Denver ||
		fac == FortWorth ||
		fac == Honolulu ||
		fac == Houston ||
		fac == Indianapolis ||
		fac == Jacksonville ||
		fac == KansasCity ||
		fac == LosAngeles ||
		fac == Memphis ||
		fac == Miami ||
		fac == Minneapolis ||
		fac == NewYork ||
		fac == Oakland ||
		fac == SaltLakeCity ||
		fac == Seattle ||
		fac == Washington
}

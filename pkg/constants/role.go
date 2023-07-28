package constants

type Role = string

// Facility Staff Roles
const (
	AirTrafficManager       Role = "ATM"
	DeputyAirTrafficManager Role = "DATM"
	TrainingAdministrator   Role = "TA"
	EventCoordinator        Role = "EC"
	FacilityEngineer        Role = "FC"
	WebMaster               Role = "WM"
	Instructor              Role = "INS"
	Mentor                  Role = "MTR"
)

// Division Team Roles
const (
	ACETeam  Role = "ACE"
	DICETeam Role = "DICE"
	DCCTeam  Role = "DCC"
	TechTeam Role = "TECH"
)

// Division Staff Roles
const (
	SystemAdministrator Role = "SYSADM"
	DivisionManagement  Role = "DIVISION_MANAGEMENT"
	DivisionStaff       Role = "DIVISION_STAFF"
)

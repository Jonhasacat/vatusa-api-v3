package role

type Role string

// Facility Staff Roles
const (
	RoleAirTrafficManager       Role = "ATM"
	RoleDeputyAirTrafficManager Role = "DATM"
	RoleTrainingAdministrator   Role = "TA"
	RoleEventCoordinator        Role = "EC"
	RoleFacilityEngineer        Role = "FC"
	RoleWebMaster               Role = "WM"
	RoleInstructor              Role = "INS"
	RoleMentor                  Role = "MTR"
)

// Division Team Roles
const (
	RoleACETeam  Role = "ACE"
	RoleDICETeam Role = "DICE"
	RoleDCCTeam  Role = "DCC"
	RoleTechTeam Role = "TECH"
)

// Division Staff Roles
const (
	RoleSystemAdministrator Role = "SYSADM"
	RoleDivisionManagement  Role = "DIVISION_MANAGEMENT"
	RoleDivisionStaff       Role = "DIVISION_STAFF"
)

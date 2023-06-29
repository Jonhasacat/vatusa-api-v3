package constants

type Rating int

const (
	RatingInactive      Rating = -1
	RatingSuspended     Rating = 0
	RatingObserver      Rating = 1
	RatingS1            Rating = 2
	RatingS2            Rating = 3
	RatingS3            Rating = 4
	RatingC1            Rating = 5
	RatingC2            Rating = 6
	RatingC3            Rating = 7
	RatingI1            Rating = 8
	RatingI2            Rating = 9
	RatingI3            Rating = 10
	RatingSupervisor    Rating = 11
	RatingAdministrator Rating = 12
)

var RatingShortMap = map[Rating]string{
	RatingInactive:      "AFK",
	RatingSuspended:     "SUS",
	RatingObserver:      "OBS",
	RatingS1:            "S1",
	RatingS2:            "S2",
	RatingS3:            "S3",
	RatingC1:            "C1",
	RatingC2:            "C2",
	RatingC3:            "C3",
	RatingI1:            "I1",
	RatingI2:            "I2",
	RatingI3:            "I3",
	RatingSupervisor:    "SUP",
	RatingAdministrator: "ADM",
}

var RatingLongMap = map[Rating]string{
	RatingInactive:      "Inactive",
	RatingSuspended:     "Suspended",
	RatingObserver:      "Observer",
	RatingS1:            "Student 1",
	RatingS2:            "Student 2",
	RatingS3:            "Student 3",
	RatingC1:            "Controller",
	RatingC2:            "Controller 2",
	RatingC3:            "Senior Controller",
	RatingI1:            "Instructor",
	RatingI2:            "Instructor 2",
	RatingI3:            "Senior Instructor",
	RatingSupervisor:    "Supervisor",
	RatingAdministrator: "Administrator",
}

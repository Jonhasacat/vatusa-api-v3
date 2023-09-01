package constants

type Rating = int

const (
	Inactive      Rating = -1
	Suspended     Rating = 0
	Observer      Rating = 1
	S1            Rating = 2
	S2            Rating = 3
	S3            Rating = 4
	C1            Rating = 5
	C2            Rating = 6
	C3            Rating = 7
	I1            Rating = 8
	I2            Rating = 9
	I3            Rating = 10
	Supervisor    Rating = 11
	Administrator Rating = 12
)

var RatingShortMap = map[Rating]string{
	Inactive:      "AFK",
	Suspended:     "SUS",
	Observer:      "OBS",
	S1:            "S1",
	S2:            "S2",
	S3:            "S3",
	C1:            "C1",
	C2:            "C2",
	C3:            "C3",
	I1:            "I1",
	I2:            "I2",
	I3:            "I3",
	Supervisor:    "SUP",
	Administrator: "ADM",
}

var RatingLongMap = map[Rating]string{
	Inactive:      "Inactive",
	Suspended:     "Suspended",
	Observer:      "Observer",
	S1:            "Student 1",
	S2:            "Student 2",
	S3:            "Student 3",
	C1:            "Controller",
	C2:            "Controller 2",
	C3:            "Senior Controller",
	I1:            "Instructor",
	I2:            "Instructor 2",
	I3:            "Senior Instructor",
	Supervisor:    "Supervisor",
	Administrator: "Administrator",
}

func RatingShort(rating Rating) string {
	return RatingShortMap[rating]
}

func RatingLong(rating Rating) string {
	return RatingLongMap[rating]
}

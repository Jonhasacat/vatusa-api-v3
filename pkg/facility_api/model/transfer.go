package model

import "time"

type Transfer struct {
	FromFacility string
	ToFacility   string
	Reason       string
	DateTime     time.Time
}

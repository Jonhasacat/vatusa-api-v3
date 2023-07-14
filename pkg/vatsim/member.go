package vatsim

import "time"

type Member struct {
	ID               uint64     `json:"id"`
	NameFirst        *string    `json:"name_first"`
	NameLast         *string    `json:"name_last"`
	Email            *string    `json:"email"`
	CountyState      string     `json:"countystate"`
	Country          string     `json:"country"`
	Rating           uint64     `json:"rating"`
	PilotRating      uint64     `json:"pilotrating"`
	MilitaryRating   uint64     `json:"militaryrating"`
	SuspendDate      *time.Time `json:"susp_date"`
	RegistrationDate *time.Time `json:"reg_date"`
	Region           string     `json:"region_id"`
	Division         string     `json:"division_id"`
	SubDivision      string     `json:"subdivision_id"`
	LastRatingChange *time.Time `json:"lastratingchange"`
}

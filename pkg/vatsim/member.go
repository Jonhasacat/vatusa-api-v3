package vatsim

import "time"

type Member struct {
	ID               uint64  `json:"id"`
	NameFirst        *string `json:"name_first"`
	NameLast         *string `json:"name_last"`
	Email            *string `json:"email"`
	CountyState      string  `json:"countystate"`
	Country          string  `json:"country"`
	Rating           int     `json:"rating"`
	PilotRating      int     `json:"pilotrating"`
	MilitaryRating   int     `json:"militaryrating"`
	SuspendDate      *string `json:"susp_date"`
	RegistrationDate *string `json:"reg_date"`
	Region           string  `json:"region_id"`
	Division         string  `json:"division_id"`
	SubDivision      string  `json:"subdivision_id"`
	LastRatingChange *string `json:"lastratingchange"`
}

func (m *Member) SuspendTime() *time.Time {
	if m.SuspendDate == nil {
		return nil
	}
	t, err := time.Parse("2006-01-02T15:04:05", *m.SuspendDate)
	if err != nil {
		return nil
	}
	return &t
}

func (m *Member) RegistrationTime() *time.Time {
	if m.RegistrationDate == nil {
		return nil
	}
	t, err := time.Parse("2006-01-02T15:04:05", *m.RegistrationDate)
	if err != nil {
		return nil
	}
	return &t
}

func (m *Member) LastRatingChangeTime() *time.Time {
	if m.LastRatingChange == nil {
		return nil
	}
	t, err := time.Parse("2006-01-02T15:04:05", *m.LastRatingChange)
	if err != nil {
		return nil
	}
	return &t
}

package model

type RosterRemoveRequest struct {
	CID          uint64  `json:"cid"`
	Reason       string  `json:"reason"`
	RequesterCID *uint64 `json:"requester_cid"`
}

type ProcessRosterRequestRequest struct {
	ID           uint `json:"id"`
	Accept       bool
	Reason       *string
	RequesterCID uint64
}

type RosterRequestType string

const (
	RosterRequestTransfer RosterRequestType = "transfer"
	RosterRequestVisit    RosterRequestType = "visit"
)

type RosterRequest struct {
	ID         uint
	Controller Controller
	Type       RosterRequestType
	Reason     string
}

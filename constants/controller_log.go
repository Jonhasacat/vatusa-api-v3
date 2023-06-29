package constants

type LogMessageType uint64

const (
	LogAction LogMessageType = iota
	LogComment
)

type LogVisibility uint64

const (
	VisibilityGeneral LogVisibility = iota
	VisibilitySeniorStaff
	VisibilityDivisionStaff
)

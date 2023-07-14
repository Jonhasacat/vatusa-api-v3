package action_log

type LogMessageType uint64

const (
	Action LogMessageType = iota
	Comment
)

type LogVisibility uint64

const (
	VisibilityGeneral LogVisibility = iota
	VisibilitySeniorStaff
	VisibilityDivisionStaff
)

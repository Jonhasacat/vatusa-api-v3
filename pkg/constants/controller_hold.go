package constants

type Hold = string

const (
	BasicAcademy    Hold = "ACADEMY"
	RecentTransfer  Hold = "RECENT_TRANSFER"
	RecentPromotion Hold = "RECENT_PROMOTION"
	PendingTransfer Hold = "PENDING_TRANSFER"
	RCEExam         Hold = "RCE_EXAM"
	Administrative  Hold = "ADMINISTRATIVE"
)

type Meta struct {
	ID               Hold
	DisplayName      string
	PreventTransfer  bool
	PreventVisit     bool
	PreventPromotion bool
}

var Map = map[Hold]Meta{
	BasicAcademy: {
		ID:               BasicAcademy,
		DisplayName:      "VATUSA Academy Required",
		PreventTransfer:  true,
		PreventVisit:     true,
		PreventPromotion: true,
	},
	RecentTransfer: {
		ID:               RecentTransfer,
		DisplayName:      "Recent Transfer",
		PreventTransfer:  true,
		PreventVisit:     true,
		PreventPromotion: false,
	},
	RecentPromotion: {
		ID:               RecentPromotion,
		DisplayName:      "Recent Promotion",
		PreventTransfer:  true,
		PreventVisit:     true,
		PreventPromotion: false,
	},
	PendingTransfer: {
		ID:               PendingTransfer,
		DisplayName:      "Pending Transfer",
		PreventTransfer:  true,
		PreventVisit:     false,
		PreventPromotion: false,
	},
	RCEExam: {
		ID:               RCEExam,
		DisplayName:      "RCE Exam Required",
		PreventTransfer:  true,
		PreventVisit:     true,
		PreventPromotion: true,
	},
	Administrative: {
		ID:               Administrative,
		DisplayName:      "Administrative HoldMeta",
		PreventTransfer:  true,
		PreventVisit:     true,
		PreventPromotion: true,
	},
}

func Get(id string) Meta {
	return Map[Hold(id)]
}

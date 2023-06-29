package constants

type Hold string

const (
	HoldAcademy         Hold = "ACADEMY"
	HoldRecentTransfer  Hold = "RECENT_TRANSFER"
	HoldRecentPromotion Hold = "RECENT_PROMOTION"
	HoldPendingTransfer Hold = "PENDING_TRANSFER"
	HoldRCEExam         Hold = "RCE_EXAM"
	HoldAdministrative  Hold = "ADMINISTRATIVE"
)

type HoldMeta struct {
	ID               Hold
	DisplayName      string
	PreventTransfer  bool
	PreventVisit     bool
	PreventPromotion bool
}

var HoldMap = map[Hold]HoldMeta{
	HoldAcademy: {
		ID:               HoldAcademy,
		DisplayName:      "VATUSA Academy Required",
		PreventTransfer:  true,
		PreventVisit:     true,
		PreventPromotion: true,
	},
	HoldRecentTransfer: {
		ID:               HoldRecentTransfer,
		DisplayName:      "Recent Transfer",
		PreventTransfer:  true,
		PreventVisit:     true,
		PreventPromotion: false,
	},
	HoldRecentPromotion: {
		ID:               HoldRecentPromotion,
		DisplayName:      "Recent Promotion",
		PreventTransfer:  true,
		PreventVisit:     true,
		PreventPromotion: false,
	},
	HoldPendingTransfer: {
		ID:               HoldPendingTransfer,
		DisplayName:      "Pending Transfer",
		PreventTransfer:  true,
		PreventVisit:     false,
		PreventPromotion: false,
	},
	HoldRCEExam: {
		ID:               HoldRCEExam,
		DisplayName:      "RCE Exam Required",
		PreventTransfer:  true,
		PreventVisit:     true,
		PreventPromotion: true,
	},
	HoldAdministrative: {
		ID:               HoldAdministrative,
		DisplayName:      "Administrative HoldMeta",
		PreventTransfer:  true,
		PreventVisit:     true,
		PreventPromotion: true,
	},
}

func GetHold(id Hold) HoldMeta {
	return HoldMap[id]
}

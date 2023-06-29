package legacydb

import "time"

type Controller struct {
	CID                  uint64     `gorm:"column:cid"`
	FName                string     `gorm:"column:fname"`
	LName                string     `gorm:"column:lname"`
	Email                string     `gorm:"column:email"`
	Facility             string     `gorm:"column:facility"`
	Rating               int        `gorm:"column:rating"`
	CreatedAt            time.Time  `gorm:"column:created_at"`
	UpdatedAd            time.Time  `gorm:"column:updated_at"`
	FlagNeedBasic        bool       `gorm:"column:flag_needbasic"`
	FlagTransferOverride bool       `gorm:"column:flag_xferOverride"`
	FacilityJoin         *time.Time `gorm:"column:facility_join"`
	FlagHomeController   bool       `gorm:"column:flag_homecontroller"`
	FlagBroadcastOptedIn bool       `gorm:"column:flag_broadcastOptedIn"`
	DiscordId            *string    `gorm:"column:discord_id"`
}

//  `gorm:"column:"`

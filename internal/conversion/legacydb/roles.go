package legacydb

import "time"

type Role struct {
	ID        uint64    `gorm:"column:id"`
	CID       uint64    `gorm:"column:cid"`
	Facility  string    `gorm:"column:facility"`
	Role      string    `gorm:"column:role"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

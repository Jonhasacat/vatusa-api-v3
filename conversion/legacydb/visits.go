package legacydb

import "time"

type Visit struct {
	ID        uint64    `gorm:"column:id"`
	CID       uint64    `gorm:"column:cid"`
	Facility  string    `gorm:"column:facility"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

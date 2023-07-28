package legacydb

import "time"

type Promotion struct {
	ID         uint64    `gorm:"column:id"`
	CID        uint64    `gorm:"column:cid"`
	Grantor    uint64    `gorm:"column:grantor"`
	ToRating   int       `gorm:"column:to"`
	FromRating int       `gorm:"column:from"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	Exam       string    `gorm:"column:exam"`
	Examiner   uint64    `gorm:"column:examiner"`
	Position   string    `gorm:"column:position"`
	EvalID     uint64    `gorm:"column:eval_id"`
}

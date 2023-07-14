package response

import (
	"time"
)

type SoloCertification struct {
	ID         uint64
	Controller *Controller
	Position   string
	IssuedAt   time.Time
	ExpiresAt  time.Time
}

package database

import "gorm.io/gorm"

type WebhookEventController struct {
	gorm.Model
	ControllerID      uint64
	Controller        *Controller
	WebhookConsumerID uint
	WebhookConsumer   *WebhookConsumer
	Lock              *string
	IsProcessed       bool
}

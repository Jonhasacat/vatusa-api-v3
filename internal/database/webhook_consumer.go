package database

import "gorm.io/gorm"

type WebhookConsumer struct {
	gorm.Model
	Facility                 string
	BaseURL                  string
	IsControllerEnabled      bool
	IsTransferRequestEnabled bool
	IsVisitRequestEnabled    bool
}

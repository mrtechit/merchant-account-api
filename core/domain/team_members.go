package domain

import (
	"time"
)

type TeamMember struct {
	ID                   uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name                 string    `gorm:"size:255;not null;" json:"merchant_name"`
	Email                string    `gorm:"size:100;not null;unique" json:"email"`
	CreatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	MerchantMerchantCode uint32
	Merchant             Merchant
}

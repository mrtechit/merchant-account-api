package domain

import (
	"time"
)

type Merchant struct {
	MerchantCode         uint32    `gorm:"primary_key;auto_increment" json:"merchant_code"`
	MerchantName         string    `gorm:"size:255;not null;" json:"merchant_name"`
	MerchantEmail        string    `gorm:"size:100;not null;" json:"email"`
	MerchantCity         string    `gorm:"size:100;not null;" json:"merchant_city"`
	MerchantCategoryCode string    `gorm:"size:100;not null;" json:"merchant_category_code"`
	CreatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

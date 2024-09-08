package models

import (
	"github.com/volatiletech/null/v8"
)

type ParcelMerchant struct {
	ID          uint        `json:"id" gorm:"Column:id"`
	MerchantId  uint        `json:"merchant_id" gorm:"Column:merchant_id"`
	LastName    null.String `json:"last_name" gorm:"Column:last_name"`
}

func (ParcelMerchant) TableName() string {
    return "parcel_merchants"
}

func (ParcelMerchant) Relationships() []string {
    return []string{}
}

func (ParcelMerchant) Connection() string {
    return ""
}

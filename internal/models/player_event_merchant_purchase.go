package models

import (
	"github.com/volatiletech/null/v8"
)

type PlayerEventMerchantPurchase struct {
	ID                      uint64      `json:"id" gorm:"Column:id"`
	NpcId                   null.Uint   `json:"npc_id" gorm:"Column:npc_id"`
	MerchantName            null.String `json:"merchant_name" gorm:"Column:merchant_name"`
	MerchantType            null.Uint   `json:"merchant_type" gorm:"Column:merchant_type"`
	ItemId                  null.Uint   `json:"item_id" gorm:"Column:item_id"`
	ItemName                null.String `json:"item_name" gorm:"Column:item_name"`
	Charges                 null.Int    `json:"charges" gorm:"Column:charges"`
	Cost                    null.Uint   `json:"cost" gorm:"Column:cost"`
	AlternateCurrencyId     null.Uint   `json:"alternate_currency_id" gorm:"Column:alternate_currency_id"`
	PlayerMoneyBalance      null.Uint64 `json:"player_money_balance" gorm:"Column:player_money_balance"`
	PlayerCurrencyBalance   null.Uint64 `json:"player_currency_balance" gorm:"Column:player_currency_balance"`
	CreatedAt               null.Time   `json:"created_at" gorm:"Column:created_at"`
}

func (PlayerEventMerchantPurchase) TableName() string {
    return "player_event_merchant_purchase"
}

func (PlayerEventMerchantPurchase) Relationships() []string {
    return []string{}
}

func (PlayerEventMerchantPurchase) Connection() string {
    return ""
}

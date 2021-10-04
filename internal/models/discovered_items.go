package models

type DiscoveredItem struct {
	ItemId          uint   `json:"item_id" gorm:"Column:item_id"`
	CharName        string `json:"char_name" gorm:"Column:char_name"`
	DiscoveredDate  uint   `json:"discovered_date" gorm:"Column:discovered_date"`
	AccountStatus   int    `json:"account_status" gorm:"Column:account_status"`
}

func (DiscoveredItem) TableName() string {
    return "discovered_items"
}

func (DiscoveredItem) Relationships() []string {
    return []string{}
}

func (DiscoveredItem) Connection() string {
    return ""
}

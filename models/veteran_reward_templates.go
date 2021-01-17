package models

type VeteranRewardTemplate struct {
	ClaimId     uint   `json:"claim_id" gorm:"Column:claim_id"`
	Name        string `json:"name" gorm:"Column:name"`
	ItemId      uint   `json:"item_id" gorm:"Column:item_id"`
	Charges     uint16 `json:"charges" gorm:"Column:charges"`
	RewardSlot  uint8  `json:"reward_slot" gorm:"Column:reward_slot"`
}

func (VeteranRewardTemplate) TableName() string {
    return "veteran_reward_templates"
}

func (VeteranRewardTemplate) Relationships() []string {
    return []string{}
}

func (VeteranRewardTemplate) Connection() string {
    return "eqemu_content"
}

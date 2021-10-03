package models

type AccountReward struct {
	AccountId  uint `json:"account_id" gorm:"Column:account_id"`
	RewardId   uint `json:"reward_id" gorm:"Column:reward_id"`
	Amount     uint `json:"amount" gorm:"Column:amount"`
}

func (AccountReward) TableName() string {
    return "account_rewards"
}

func (AccountReward) Relationships() []string {
    return []string{}
}

func (AccountReward) Connection() string {
    return ""
}

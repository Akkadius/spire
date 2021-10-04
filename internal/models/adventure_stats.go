package models

type AdventureStat struct {
	PlayerId   uint   `json:"player_id" gorm:"Column:player_id"`
	GukWins    uint32 `json:"guk_wins" gorm:"Column:guk_wins"`
	MirWins    uint32 `json:"mir_wins" gorm:"Column:mir_wins"`
	MmcWins    uint32 `json:"mmc_wins" gorm:"Column:mmc_wins"`
	RujWins    uint32 `json:"ruj_wins" gorm:"Column:ruj_wins"`
	TakWins    uint32 `json:"tak_wins" gorm:"Column:tak_wins"`
	GukLosses  uint32 `json:"guk_losses" gorm:"Column:guk_losses"`
	MirLosses  uint32 `json:"mir_losses" gorm:"Column:mir_losses"`
	MmcLosses  uint32 `json:"mmc_losses" gorm:"Column:mmc_losses"`
	RujLosses  uint32 `json:"ruj_losses" gorm:"Column:ruj_losses"`
	TakLosses  uint32 `json:"tak_losses" gorm:"Column:tak_losses"`
}

func (AdventureStat) TableName() string {
    return "adventure_stats"
}

func (AdventureStat) Relationships() []string {
    return []string{}
}

func (AdventureStat) Connection() string {
    return ""
}

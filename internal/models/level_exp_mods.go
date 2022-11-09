package models

import (
	"github.com/volatiletech/null/v8"
)

type LevelExpMod struct {
	Level      int          `json:"level" gorm:"Column:level"`
	ExpMod     float32 `json:"exp_mod" gorm:"Column:exp_mod"`
	AaExpMod   float32 `json:"aa_exp_mod" gorm:"Column:aa_exp_mod"`
}

func (LevelExpMod) TableName() string {
    return "level_exp_mods"
}

func (LevelExpMod) Relationships() []string {
    return []string{}
}

func (LevelExpMod) Connection() string {
    return ""
}

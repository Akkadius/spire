package models

import (
	"github.com/volatiletech/null/v8"
)

type TradeskillRecipe struct {
	ID                     int         `json:"id" gorm:"Column:id"`
	Name                   string      `json:"name" gorm:"Column:name"`
	Tradeskill             int16       `json:"tradeskill" gorm:"Column:tradeskill"`
	Skillneeded            int16       `json:"skillneeded" gorm:"Column:skillneeded"`
	Trivial                int16       `json:"trivial" gorm:"Column:trivial"`
	Nofail                 int8        `json:"nofail" gorm:"Column:nofail"`
	ReplaceContainer       int8        `json:"replace_container" gorm:"Column:replace_container"`
	Notes                  null.String `json:"notes" gorm:"Column:notes"`
	MustLearn              int8        `json:"must_learn" gorm:"Column:must_learn"`
	Quest                  int8        `json:"quest" gorm:"Column:quest"`
	Enabled                int8        `json:"enabled" gorm:"Column:enabled"`
	MinExpansion           uint8       `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8       `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
}

func (TradeskillRecipe) TableName() string {
    return "tradeskill_recipe"
}

func (TradeskillRecipe) Relationships() []string {
    return []string{}
}

func (TradeskillRecipe) Connection() string {
    return "eqemu_content"
}

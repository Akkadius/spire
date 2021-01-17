package models

type CharCreatePointAllocation struct {
	ID        uint `json:"id" gorm:"Column:id"`
	BaseStr   uint `json:"base_str" gorm:"Column:base_str"`
	BaseSta   uint `json:"base_sta" gorm:"Column:base_sta"`
	BaseDex   uint `json:"base_dex" gorm:"Column:base_dex"`
	BaseAgi   uint `json:"base_agi" gorm:"Column:base_agi"`
	BaseInt   uint `json:"base_int" gorm:"Column:base_int"`
	BaseWis   uint `json:"base_wis" gorm:"Column:base_wis"`
	BaseCha   uint `json:"base_cha" gorm:"Column:base_cha"`
	AllocStr  uint `json:"alloc_str" gorm:"Column:alloc_str"`
	AllocSta  uint `json:"alloc_sta" gorm:"Column:alloc_sta"`
	AllocDex  uint `json:"alloc_dex" gorm:"Column:alloc_dex"`
	AllocAgi  uint `json:"alloc_agi" gorm:"Column:alloc_agi"`
	AllocInt  uint `json:"alloc_int" gorm:"Column:alloc_int"`
	AllocWis  uint `json:"alloc_wis" gorm:"Column:alloc_wis"`
	AllocCha  uint `json:"alloc_cha" gorm:"Column:alloc_cha"`
}

func (CharCreatePointAllocation) TableName() string {
    return "char_create_point_allocations"
}

func (CharCreatePointAllocation) Relationships() []string {
    return []string{}
}

func (CharCreatePointAllocation) Connection() string {
    return "eqemu_content"
}

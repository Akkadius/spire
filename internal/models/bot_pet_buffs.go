package models

type BotPetBuff struct {
	PetBuffsIndex   uint `json:"pet_buffs_index" gorm:"Column:pet_buffs_index"`
	PetsIndex       uint `json:"pets_index" gorm:"Column:pets_index"`
	SpellId         uint `json:"spell_id" gorm:"Column:spell_id"`
	CasterLevel     uint `json:"caster_level" gorm:"Column:caster_level"`
	Duration        uint `json:"duration" gorm:"Column:duration"`
}

func (BotPetBuff) TableName() string {
    return "bot_pet_buffs"
}

func (BotPetBuff) Relationships() []string {
    return []string{}
}

func (BotPetBuff) Connection() string {
    return ""
}

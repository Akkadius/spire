package models

type BotSpellCastingChance struct {
	ID               int   `json:"id" gorm:"Column:id"`
	SpellTypeIndex   uint8 `json:"spell_type_index" gorm:"Column:spell_type_index"`
	ClassId          uint8 `json:"class_id" gorm:"Column:class_id"`
	StanceIndex      uint8 `json:"stance_index" gorm:"Column:stance_index"`
	NHSNDValue       uint8 `json:"n_hsnd_value" gorm:"Column:nHSND_value"`
	PHValue          uint8 `json:"p_h_value" gorm:"Column:pH_value"`
	PSValue          uint8 `json:"p_s_value" gorm:"Column:pS_value"`
	PHSValue         uint8 `json:"p_hs_value" gorm:"Column:pHS_value"`
	PNValue          uint8 `json:"p_n_value" gorm:"Column:pN_value"`
	PHNValue         uint8 `json:"p_hn_value" gorm:"Column:pHN_value"`
	PSNValue         uint8 `json:"p_sn_value" gorm:"Column:pSN_value"`
	PHSNValue        uint8 `json:"p_hsn_value" gorm:"Column:pHSN_value"`
	PDValue          uint8 `json:"p_d_value" gorm:"Column:pD_value"`
	PHDValue         uint8 `json:"p_hd_value" gorm:"Column:pHD_value"`
	PSDValue         uint8 `json:"p_sd_value" gorm:"Column:pSD_value"`
	PHSDValue        uint8 `json:"p_hsd_value" gorm:"Column:pHSD_value"`
	PNDValue         uint8 `json:"p_nd_value" gorm:"Column:pND_value"`
	PHNDValue        uint8 `json:"p_hnd_value" gorm:"Column:pHND_value"`
	PSNDValue        uint8 `json:"p_snd_value" gorm:"Column:pSND_value"`
	PHSNDValue       uint8 `json:"p_hsnd_value" gorm:"Column:pHSND_value"`
}

func (BotSpellCastingChance) TableName() string {
    return "bot_spell_casting_chances"
}

func (BotSpellCastingChance) Relationships() []string {
    return []string{}
}

func (BotSpellCastingChance) Connection() string {
    return ""
}

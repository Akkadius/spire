package models

type AdventureTemplateEntry struct {
	ID          uint `json:"id" gorm:"Column:id"`
	TemplateId  uint `json:"template_id" gorm:"Column:template_id"`
}

func (AdventureTemplateEntry) TableName() string {
    return "adventure_template_entry"
}

func (AdventureTemplateEntry) Relationships() []string {
    return []string{}
}

func (AdventureTemplateEntry) Connection() string {
    return "eqemu_content"
}

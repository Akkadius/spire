package models

type AdventureTemplateEntryFlavor struct {
	ID   uint   `json:"id" gorm:"Column:id"`
	Text string `json:"text" gorm:"Column:text"`
}

func (AdventureTemplateEntryFlavor) TableName() string {
    return "adventure_template_entry_flavor"
}

func (AdventureTemplateEntryFlavor) Relationships() []string {
    return []string{}
}

func (AdventureTemplateEntryFlavor) Connection() string {
    return "eqemu_content"
}

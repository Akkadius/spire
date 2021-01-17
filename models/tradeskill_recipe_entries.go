package models

type TradeskillRecipeEntry struct {
	ID             int  `json:"id" gorm:"Column:id"`
	RecipeId       int  `json:"recipe_id" gorm:"Column:recipe_id"`
	ItemId         int  `json:"item_id" gorm:"Column:item_id"`
	Successcount   int8 `json:"successcount" gorm:"Column:successcount"`
	Failcount      int8 `json:"failcount" gorm:"Column:failcount"`
	Componentcount int8 `json:"componentcount" gorm:"Column:componentcount"`
	Salvagecount   int8 `json:"salvagecount" gorm:"Column:salvagecount"`
	Iscontainer    int8 `json:"iscontainer" gorm:"Column:iscontainer"`
}

func (TradeskillRecipeEntry) TableName() string {
    return "tradeskill_recipe_entries"
}

func (TradeskillRecipeEntry) Relationships() []string {
    return []string{}
}

func (TradeskillRecipeEntry) Connection() string {
    return "eqemu_content"
}

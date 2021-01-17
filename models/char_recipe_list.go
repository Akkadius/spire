package models

type CharRecipeList struct {
	CharId    int `json:"char_id" gorm:"Column:char_id"`
	RecipeId  int `json:"recipe_id" gorm:"Column:recipe_id"`
	Madecount int `json:"madecount" gorm:"Column:madecount"`
}

func (CharRecipeList) TableName() string {
    return "char_recipe_list"
}

func (CharRecipeList) Relationships() []string {
    return []string{}
}

func (CharRecipeList) Connection() string {
    return ""
}

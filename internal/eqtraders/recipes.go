package eqtraders

// Item is a struct that represents an item
type Item struct {
	ItemId   int    `json:"item_id"`
	ItemName string `json:"item_name"`
	Count    int    `json:"count"`
}

// Recipe is a struct that represents a recipe
type Recipe struct {
	RecipeName         string `json:"recipe_name"`
	Skill              Skill  `json:"skill"`
	ExpansionId        int    `json:"expansion_id"`
	ExpansionName      string `json:"expansion_name"`
	Trivial            int    `json:"trivial"`
	RequiredSkillLevel int    `json:"required_skill_level"`
	ConsumeContainer   bool   `json:"consume_container"`
	NoFail             bool   `json:"no_fail"`
	RecipeItemId       int    `json:"recipe_item_id"`
	Components         []Item `json:"components"`
	In                 []Item `json:"in"`
	Yield              int    `json:"yield"`
	Returns            []Item `json:"returns"`
	FailureReturns     []Item `json:"failure_returns"`
	LearnedByItem      Item   `json:"learned_by_item"`
}

// recipes is a slice of Recipe
var recipes []Recipe

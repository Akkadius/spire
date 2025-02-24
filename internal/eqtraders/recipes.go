package eqtraders

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/models"
	"os"
)

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

func loadRecipes() error {
	file, err := os.ReadFile("./data/eqtraders/recipes.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &recipes)
	if err != nil {
		return err
	}

	return nil
}

func GetRecipeSignature(r Recipe) string {
	itemSummation := 0
	for _, entry := range r.Components {
		itemSummation += entry.ItemId
	}

	return fmt.Sprintf(
		"%v-%v-%v",
		r.RecipeName,
		r.Skill.SkillId,
		itemSummation,
	)
}

func GetDbRecipeSignature(r models.TradeskillRecipe) string {
	// create a unique key for each recipe
	itemSummation := 0
	for _, entry := range r.TradeskillRecipeEntries {
		if entry.Componentcount > 0 {
			itemSummation += entry.ItemId
		}
	}

	return fmt.Sprintf(
		"%v-%v-%v",
		r.Name,
		r.Tradeskill,
		itemSummation,
	)
}

// GetDbRecipeSignatureDeep is for diffing the entire model
func GetDbRecipeSignatureDeep(r models.TradeskillRecipe) string {
	r.ID = 0

	return fmt.Sprintf(
		"%v",
		r,
	)
}

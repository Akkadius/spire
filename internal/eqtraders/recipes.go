package eqtraders

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/models"
	"os"
	"sort"
	"strings"
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
	var items []string

	// sort r.Components by ItemId
	sort.Slice(r.Components, func(i, j int) bool {
		return r.Components[i].ItemId < r.Components[j].ItemId
	})

	for _, entry := range r.Components {
		if entry.Count == 0 {
			entry.Count = 1
		}
		items = append(items, fmt.Sprintf("%v-%v", entry.ItemId, entry.Count))
	}

	return fmt.Sprintf(
		"%v-%v-%v",
		r.RecipeName,
		r.Skill.SkillId,
		strings.Join(items, "-"),
	)
}

func GetDbRecipeSignature(r models.TradeskillRecipe) string {
	// create a unique key for each recipe
	var items []string

	sort.Slice(r.TradeskillRecipeEntries, func(i, j int) bool {
		return r.TradeskillRecipeEntries[i].ItemId < r.TradeskillRecipeEntries[j].ItemId
	})

	for _, entry := range r.TradeskillRecipeEntries {
		if entry.Componentcount > 0 {
			items = append(items, fmt.Sprintf("%v-%v", entry.ItemId, entry.Componentcount))
		}
	}

	return fmt.Sprintf(
		"%v-%v-%v",
		r.Name,
		r.Tradeskill,
		strings.Join(items, "-"),
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

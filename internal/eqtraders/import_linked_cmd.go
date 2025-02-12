package eqtraders

import (
	"encoding/json"
	"github.com/Akkadius/spire/internal/models"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"os"
)

type ImportLinkedCommand struct {
	db      *gorm.DB
	command *cobra.Command
}

func (c *ImportLinkedCommand) Command() *cobra.Command {
	return c.command
}

func NewImportLinkedCommand(
	db *gorm.DB,
) *ImportLinkedCommand {
	i := &ImportLinkedCommand{
		db: db,
		command: &cobra.Command{
			Use:   "eq-traders:import-linked",
			Short: "Imports data from eqtraders.com using data scraped via eq-traders:scrape",
		},
	}

	//i.command.Args = cobra.MinimumNArgs(1)
	i.command.Run = i.Handle
	i.command.Flags().StringVarP(&singleRecipe, "single-recipe", "r", "", "Scrape a single recipe by name")

	return i
}

func (c *ImportLinkedCommand) Handle(cmd *cobra.Command, args []string) {
	//itemLookupCache := make(map[string]models.Item)

	file, err := os.ReadFile("./data/eqtraders/recipes.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &recipes)
	if err != nil {
		return
	}

	// bulk query for all recipes
	var existingRecipes []models.TradeskillRecipe
	c.db.Preload("TradeskillRecipeEntries.TradeskillRecipe").Find(&existingRecipes)

	for _, recipe := range recipes {
		if len(singleRecipe) > 0 && recipe.RecipeName != singleRecipe {
			continue
		}

		for _, existingRecipe := range existingRecipes {
			for _, entry := range existingRecipe.TradeskillRecipeEntries {
				if entry.Componentcount > 0 && entry.ItemId == recipe.RecipeItemId {
					//fmt.Println("Found existing recipe: ", recipe.RecipeName)
					continue
				}
			}
		}
	}
}

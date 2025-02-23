package eqtraders

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/models"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"os"
	"strconv"
)

type ImportLinkedCommand struct {
	db      *gorm.DB
	command *cobra.Command
	logger  *logger.AppLogger
}

func (c *ImportLinkedCommand) Command() *cobra.Command {
	return c.command
}

func NewImportLinkedCommand(
	db *gorm.DB,
	logger *logger.AppLogger,
) *ImportLinkedCommand {
	i := &ImportLinkedCommand{
		db: db,
		command: &cobra.Command{
			Use:   "eq-traders:import-linked [expansion_number]",
			Short: "Imports data from eqtraders.com using data scraped via eq-traders:scrape",
		},
		logger: logger,
	}

	i.command.Args = cobra.MinimumNArgs(1)
	i.command.Run = i.Handle
	i.command.Flags().StringVarP(&singleRecipe, "single-recipe", "r", "", "Scrape a single recipe by name")

	return i
}

func (c *ImportLinkedCommand) Handle(cmd *cobra.Command, args []string) {
	//itemLookupCache := make(map[string]models.Item)

	expansion := os.Args[2]
	expansionNumber, err := strconv.Atoi(expansion)
	if err != nil {
		expansionNumber = -1
	}

	c.logger.Info().
		Any("expansion", expansionNumber).
		Msg("Importing linked recipes")

	file, err := os.ReadFile("./data/eqtraders/recipes.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &eqtradersRecipes)
	if err != nil {
		return
	}

	// bulk query for all recipes
	var existingRecipes []models.TradeskillRecipe
	c.db.Preload("TradeskillRecipeEntries.TradeskillRecipe").Find(&existingRecipes)

	// create a lookup map for existing recipes
	existingRecipeLookup := make(map[string]models.TradeskillRecipe)

	// loop through all recipes
	for _, recipe := range existingRecipes {
		key := GetDbRecipeSignature(recipe)
		if _, ok := existingRecipeLookup[key]; ok {
			continue
		}

		existingRecipeLookup[key] = recipe
	}

	// database
	for _, dbRecipe := range existingRecipes {
		if expansionNumber != -1 && dbRecipe.MinExpansion != int8(expansionNumber) {
			continue
		}

		//if len(singleRecipe) > 0 && recipe.RecipeName != singleRecipe {
		//	continue
		//}

		// TODO: make sure we handle if there is an existing recipe already linked

		//recipeToImport := Recipe{}
		existingRecipeModel := models.TradeskillRecipe{}
		for _, e := range dbRecipe.TradeskillRecipeEntries {
			// json (eqtraders)
			for _, recipe := range eqtradersRecipes {
				if e.Componentcount > 0 && e.ItemId == recipe.RecipeItemId {
					// from the recipe get the database object
					lookupKey := GetRecipeSignature(recipe)
					if existingRecipe, ok := existingRecipeLookup[lookupKey]; ok {
						c.logger.Info().
							Any("id", dbRecipe.ID).
							Any("parent", dbRecipe.Name).
							Any("recipe", recipe.RecipeName).
							Any("skill", recipe.Skill.SkillName).
							Any("skill id", recipe.Skill.SkillId).
							Any("expansion", recipe.ExpansionId).
							Msg("Found existing recipe")
						existingRecipeModel = existingRecipe
						if existingRecipeModel.ID > 0 {

						}
						break
					}

					fmt.Println("Found linked recipe:", recipe.RecipeName)
					//recipeToImport = recipe
				}
			}
		}

	}
}

package eqtraders

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/models"
	"github.com/spf13/cobra"
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
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
	itemLookupCache := make(map[string]models.Item)

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

	err = json.Unmarshal(file, &recipes)
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

		recipeToImport := Recipe{}
		existingRecipeModel := models.TradeskillRecipe{}
		for _, e := range dbRecipe.TradeskillRecipeEntries {
			// json (eqtraders)
			for _, recipe := range recipes {
				if e.Componentcount > 0 && e.ItemId == recipe.RecipeItemId {
					fmt.Println("Found linked recipe:", recipe.RecipeName)
					recipeToImport = recipe
					existingRecipeModel = dbRecipe
					break
				}
			}
		}

		if recipeToImport.RecipeName != "" {
			c.logger.Info().Any("recipe", recipeToImport.RecipeName).Msg("Importing linked recipe")

			var r models.TradeskillRecipe

			recipeKey := GetRecipeSignature(recipeToImport)

			hasExistingRecipe := false
			if existingRecipe, ok := existingRecipeLookup[recipeKey]; ok {
				fmt.Printf("Found existing recipe: %v\n", existingRecipe.Name)
				r = existingRecipe
				hasExistingRecipe = true
			}

			r.Name = recipeToImport.RecipeName
			r.Tradeskill = int16(recipeToImport.Skill.SkillId)
			r.Trivial = int16(dbRecipe.Trivial)
			r.Nofail = int8(0)
			//r.MinExpansion = int8(recipe.ExpansionId)
			r.MinExpansion = int8(expansionNumber)
			existingRecipeModel.MinExpansion = r.MinExpansion
			r.MaxExpansion = 99

			var nofail int8
			if recipeToImport.NoFail {
				nofail = 1
			}

			r.Nofail = nofail

			var additionalNotes []string
			if recipeToImport.LearnedByItem.ItemName != "" {
				additionalNotes = append(
					additionalNotes,
					fmt.Sprintf("Learned by item [%v]", recipeToImport.LearnedByItem.ItemName),
				)
			}

			r.Notes = null.StringFrom(
				fmt.Sprintf("%v - %v (eqtraders import) %v",
					recipeToImport.ExpansionName,
					recipeToImport.Skill.SkillName,
					strings.Join(additionalNotes, ", "),
				),
			)
			r.Enabled = int8(1)
			r.MustLearn = int8(0)
			r.Quest = int8(0)
			// replace container appear to be mostly quest based recipes
			// todo handle this with quest cases
			r.ReplaceContainer = int8(0)
			if recipeToImport.ConsumeContainer {
				r.ReplaceContainer = int8(1)
			}

			r.ContentFlags = null.StringFrom("")
			r.ContentFlagsDisabled = null.StringFrom("")
			r.Skillneeded = int16(recipeToImport.RequiredSkillLevel)

			if len(recipeToImport.LearnedByItem.ItemName) > 0 {
				if i, ok := itemLookupCache[recipeToImport.LearnedByItem.ItemName]; ok {
					r.LearnedByItemId = i.ID
				} else {
					var item models.Item
					c.db.Where("Name = ?", recipeToImport.LearnedByItem.ItemName).First(&item)

					if item.ID > 0 {
						itemLookupCache[recipeToImport.LearnedByItem.ItemName] = item
						r.LearnedByItemId = item.ID
					}
				}
				if r.LearnedByItemId > 0 {
					r.MustLearn = int8(1)
					fmt.Printf("Learned by item: %v (%v)\n", recipeToImport.LearnedByItem.ItemName, r.LearnedByItemId)
				}
			}

			fmt.Println("Recipe: ", GetDbRecipeSignatureDeep(r))
			fmt.Println("Recipe2: ", GetDbRecipeSignatureDeep(existingRecipeModel))
			fmt.Println("")

			// make sure the recipe doesn't already exist before we do a bunch of work
			if hasExistingRecipe && GetDbRecipeSignatureDeep(r) == GetDbRecipeSignatureDeep(existingRecipeModel) {
				c.logger.Info().Any("recipe", r.Name).Msg("Same exact recipe already exists in database, skipping update")
				continue
			}

			// insert recipe into database
			c.db.Save(&r)

			if r.ID == 0 {
				c.logger.Fatal().Msgf("Error inserting recipe id: %v name: %v into database", r.ID, r.Name)
			}

			existing := "re-import"
			if !hasExistingRecipe {
				existing = "new"
			}

			fmt.Println(strings.Repeat("-", 80))
			fmt.Printf("> Importing recipe: %v (%v)\n", recipeToImport.RecipeName, existing)
			fmt.Println(strings.Repeat("-", 80))

			fmt.Printf("> Skill %17v | Trivial (%v)\n",
				fmt.Sprintf("%v (%v)", recipeToImport.Skill.SkillName, recipeToImport.Skill.SkillId),
				dbRecipe.Trivial,
			)
			fmt.Println(strings.Repeat("-", 80))

			var components []models.TradeskillRecipeEntry

			// insert components into database
			for _, component := range recipeToImport.Components {
				var e models.TradeskillRecipeEntry
				e.RecipeId = r.ID
				e.ItemId = component.ItemId
				e.Componentcount = int8(component.Count)
				if e.Componentcount == 0 {
					e.Componentcount = 1
				}

				for _, returns := range recipeToImport.FailureReturns {
					if returns.ItemId == component.ItemId {
						e.Failcount = int8(returns.Count)
						break
					}
				}

				for _, returns := range recipeToImport.Returns {
					if returns.ItemId == component.ItemId {
						e.Successcount = int8(returns.Count)
						break
					}
				}

				if e.Failcount == 0 && !recipeToImport.NoFail {
					e.Salvagecount = int8(component.Count)
				}

				components = append(components, e)

				fmt.Printf(
					"|--- Component %10v | %-40v | Count (%v) Successcount (%v) Failcount (%v) Salvage (%v)\n",
					component.ItemId,
					component.ItemName,
					component.Count,
					e.Successcount,
					e.Failcount,
					e.Salvagecount,
				)
			}

			// insert in into database
			for _, in := range recipeToImport.In {
				alreadyExists := false
				for _, component := range components {
					if component.ItemId == in.ItemId {
						alreadyExists = true
						break
					}
				}

				if alreadyExists {
					continue
				}

				var e models.TradeskillRecipeEntry
				e.RecipeId = r.ID
				e.ItemId = in.ItemId
				e.Iscontainer = int8(1)
				components = append(components, e)
				fmt.Printf("|--- Container %10v | %-40v\n", in.ItemId, in.ItemName)
			}

			// insert item result itself
			if recipeToImport.RecipeItemId > 0 {
				var e models.TradeskillRecipeEntry
				e.RecipeId = r.ID
				e.ItemId = recipeToImport.RecipeItemId
				e.Successcount = int8(recipeToImport.Yield)
				components = append(components, e)
				fmt.Printf("|--- Item Result %8v | %-40v | Count (%v)\n", recipeToImport.RecipeItemId, recipeToImport.RecipeName, recipeToImport.Yield)
			}

			if hasExistingRecipe {
				c.db.Where("recipe_id = ?", r.ID).Delete(&models.TradeskillRecipeEntry{})
			}

			c.db.CreateInBatches(components, 100)

			var lookup models.TradeskillRecipe
			lookup.ID = r.ID
			c.db.Preload("TradeskillRecipeEntries.TradeskillRecipe").First(&lookup)
			existingRecipeLookup[GetDbRecipeSignature(lookup)] = lookup
			c.logger.Info().Any("recipe", r.Name).Msg("Adding recipe to lookup")
		}

		//for _, existingRecipe := range existingRecipes {
		//	for _, entry := range existingRecipe.TradeskillRecipeEntries {
		//		if entry.Componentcount > 0 && entry.ItemId == recipe.RecipeItemId {
		//fmt.Println("Found existing recipe: ", recipe.RecipeName)
		//continue
		//}
		//}
		//}
	}
}

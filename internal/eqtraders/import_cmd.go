package eqtraders

import (
	"fmt"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/models"
	"github.com/spf13/cobra"
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"strings"
)

type ImportCommand struct {
	db      *gorm.DB
	command *cobra.Command
	logger  *logger.AppLogger

	// storage
	recipeLookup map[string]models.TradeskillRecipe // used to look up existing eqtradersRecipes by signature
	itemLookup   map[string]models.Item             // used to look up learned by item etc.
}

func (c *ImportCommand) Command() *cobra.Command {
	return c.command
}

func NewImportCommand(
	db *gorm.DB,
	logger *logger.AppLogger,
) *ImportCommand {
	i := &ImportCommand{
		db: db,
		command: &cobra.Command{
			Use:   "eq-traders:import [expansion_number]",
			Short: "Imports data from eqtraders.com using data scraped via eq-traders:scrape",
		},
		logger:       logger,
		recipeLookup: make(map[string]models.TradeskillRecipe),
		itemLookup:   make(map[string]models.Item),
	}

	i.command.Args = cobra.MinimumNArgs(1)
	i.command.Run = i.Handle
	i.command.Flags().StringVarP(&singleRecipe, "single-recipe", "r", "", "Scrape a single recipe by name")

	return i
}

func (c *ImportCommand) Handle(cmd *cobra.Command, args []string) {

	// inputs
	expansion := os.Args[2]
	expansionNumber, err := strconv.Atoi(expansion)
	if err != nil {
		expansionNumber = -1
	}

	// load data
	err = LoadEqtradersRecipes()
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to load eqtraders recipes")
		return
	}

	c.loadDatabaseRecipes()

	for _, eqtradersRecipe := range eqtradersRecipes {
		if len(singleRecipe) > 0 && eqtradersRecipe.RecipeName != singleRecipe {
			continue
		}

		// ignore recipes that are not for the expansion we are importing
		if eqtradersRecipe.ExpansionId != expansionNumber && expansion != "all" {
			continue
		}

		var r models.TradeskillRecipe
		recipeKey := GetRecipeSignature(eqtradersRecipe)

		hasExistingRecipe := false
		existingRecipeModel := models.TradeskillRecipe{}

		if existingRecipe, ok := c.recipeLookup[recipeKey]; ok {
			fmt.Printf("Found existing recipe: %v\n", existingRecipe.Name)
			r = existingRecipe
			existingRecipeModel = existingRecipe
			hasExistingRecipe = true
		}

		r.Name = eqtradersRecipe.RecipeName
		r.Tradeskill = int16(eqtradersRecipe.Skill.SkillId)
		r.Trivial = int16(eqtradersRecipe.Trivial)
		r.Nofail = int8(0)
		r.MinExpansion = int8(eqtradersRecipe.ExpansionId)
		r.MaxExpansion = 99
		r.Quest = int8(0)
		if eqtradersRecipe.Skill.SkillId == 75 {
			r.Quest = int8(1)
		}

		var nofail int8
		if eqtradersRecipe.NoFail {
			nofail = 1
		}

		r.Nofail = nofail

		// notes
		var additionalNotes []string
		if eqtradersRecipe.LearnedByItem.ItemName != "" {
			additionalNotes = append(
				additionalNotes,
				fmt.Sprintf("Learned by item [%v]", eqtradersRecipe.LearnedByItem.ItemName),
			)
		}

		r.Notes = null.StringFrom(
			fmt.Sprintf("%v - %v (eqtraders import) %v",
				eqtradersRecipe.ExpansionName,
				eqtradersRecipe.Skill.SkillName,
				strings.Join(additionalNotes, ", "),
			),
		)

		// other fields
		r.Enabled = int8(1)
		r.MustLearn = int8(0)
		// replace container appear to be mostly quest based recipes
		// todo handle this with quest cases
		r.ReplaceContainer = int8(0)
		if eqtradersRecipe.ConsumeContainer {
			r.ReplaceContainer = int8(1)
		}

		r.ContentFlags = null.StringFrom("")
		r.ContentFlagsDisabled = null.StringFrom("")
		r.Skillneeded = int16(eqtradersRecipe.RequiredSkillLevel)

		// learned by item lookups
		if len(eqtradersRecipe.LearnedByItem.ItemName) > 0 {
			if i, ok := c.itemLookup[eqtradersRecipe.LearnedByItem.ItemName]; ok {
				r.LearnedByItemId = i.ID
			} else {
				var item models.Item
				c.db.Where("Name = ?", eqtradersRecipe.LearnedByItem.ItemName).First(&item)

				if item.ID > 0 {
					c.itemLookup[eqtradersRecipe.LearnedByItem.ItemName] = item
					r.LearnedByItemId = item.ID
				}
			}
			if r.LearnedByItemId > 0 {
				r.MustLearn = int8(1)
				fmt.Printf("Learned by item: %v (%v)\n", eqtradersRecipe.LearnedByItem.ItemName, r.LearnedByItemId)
			}
		}

		// make sure the recipe doesn't already exist before we do a bunch of work
		if hasExistingRecipe && GetDbRecipeSignatureDeep(r) == GetDbRecipeSignatureDeep(existingRecipeModel) {
			c.logger.Info().Any("recipe", r.Name).Msg("Same exact recipe already exists in database, skipping update")
			continue
		}

		// insert recipe into database
		c.db.Save(&r)

		if r.ID == 0 {
			log.Fatalf("Error inserting recipe id: %v name: %v into database", r.ID, r.Name)
		}

		existing := "re-import"
		if !hasExistingRecipe {
			existing = "new"
		}

		fmt.Println(strings.Repeat("-", 80))
		fmt.Printf("> Importing recipe: %v (%v)\n", eqtradersRecipe.RecipeName, existing)
		fmt.Println(strings.Repeat("-", 80))

		fmt.Printf("> Skill %17v | Trivial (%v)\n",
			fmt.Sprintf("%v (%v)", eqtradersRecipe.Skill.SkillName, eqtradersRecipe.Skill.SkillId),
			eqtradersRecipe.Trivial,
		)
		fmt.Println(strings.Repeat("-", 80))

		var components []models.TradeskillRecipeEntry

		// insert components into database
		for _, component := range eqtradersRecipe.Components {
			var e models.TradeskillRecipeEntry
			e.RecipeId = r.ID
			e.ItemId = component.ItemId
			e.Componentcount = int8(component.Count)
			if e.Componentcount == 0 {
				e.Componentcount = 1
			}

			for _, returns := range eqtradersRecipe.FailureReturns {
				if returns.ItemId == component.ItemId {
					e.Failcount = int8(returns.Count)
					break
				}
			}

			for _, returns := range eqtradersRecipe.Returns {
				if returns.ItemId == component.ItemId {
					e.Successcount = int8(returns.Count)
					break
				}
			}

			if e.Failcount == 0 && !eqtradersRecipe.NoFail {
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
		for _, in := range eqtradersRecipe.In {
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
		if eqtradersRecipe.RecipeItemId > 0 {
			var e models.TradeskillRecipeEntry
			e.RecipeId = r.ID
			e.ItemId = eqtradersRecipe.RecipeItemId
			e.Successcount = int8(eqtradersRecipe.Yield)
			components = append(components, e)
			fmt.Printf("|--- Item Result %8v | %-40v | Count (%v)\n", eqtradersRecipe.RecipeItemId, eqtradersRecipe.RecipeName, eqtradersRecipe.Yield)
		}

		if hasExistingRecipe {
			c.db.Where("recipe_id = ?", r.ID).Delete(&models.TradeskillRecipeEntry{})
		}

		c.db.CreateInBatches(components, 100)
	}
}

func (c *ImportCommand) loadDatabaseRecipes() {
	// bulk query for all recipes
	var existingRecipes []models.TradeskillRecipe
	c.db.Preload("TradeskillRecipeEntries.TradeskillRecipe").Find(&existingRecipes)

	// loop through all recipes
	for _, recipe := range existingRecipes {
		key := GetDbRecipeSignature(recipe)
		if _, ok := c.recipeLookup[key]; ok {
			continue
		}

		c.recipeLookup[key] = recipe
	}
}

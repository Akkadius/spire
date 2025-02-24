package eqtraders

import (
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

type ImportCommand struct {
	db      *gorm.DB
	command *cobra.Command
	logger  *logger.AppLogger

	// storage
	recipeLookup     map[string]models.TradeskillRecipe // used to look up existing recipes by signature
	itemLookup       map[string]models.Item             // used to look up learned by item etc.
	existingRecipes  []models.TradeskillRecipe          // used to look up existing recipes
	lastTradeskillId int                                // used to assign a new tradeskill id
}

func (c *ImportCommand) Command() *cobra.Command {
	return c.command
}

var skipLinked bool
var purgeExisting bool
var printRecipes bool

// NewImportCommand returns a new ImportCommand
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
	i.command.Flags().BoolVarP(&skipLinked, "skip-linked", "s", false, "Skip linked recipes")
	i.command.Flags().BoolVarP(&purgeExisting, "purge-existing", "d", false, "Purge existing")
	i.command.Flags().BoolVarP(&printRecipes, "print-recipes", "p", false, "Print recipes")

	return i
}

const insertBatchSize = 500
const linkedMaxPasses = 10

func (c *ImportCommand) Handle(cmd *cobra.Command, args []string) {

	// inputs
	expansion := os.Args[2]
	expansionNumber, err := strconv.Atoi(expansion)
	if err != nil {
		expansionNumber = -1
	}

	if purgeExisting {
		err = c.deleteTradeskillRecipesByExpansion(expansionNumber)
		if err != nil {
			c.logger.Error().Err(err).Msg("Failed to delete tradeskill recipes")
			return
		}
	}

	// load data
	err = loadRecipes()
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to load eqtraders recipes")
		return
	}

	// preload
	c.loadItemNameLookup()
	c.loadDatabaseRecipes()

	var recipesToCreate []models.TradeskillRecipe

	// top level recipe import
	recipesExistCount := make(map[string]bool)
	for _, recipe := range recipes {
		if len(singleRecipe) > 0 && recipe.RecipeName != singleRecipe {
			continue
		}

		// ignore recipes that are not for the expansion we are importing
		if recipe.ExpansionId != expansionNumber && expansion != "all" {
			continue
		}

		sig := GetRecipeSignature(recipe)

		existing := c.GetDbRecipeFromEqTradersSignature(recipe)
		if existing.ID > 0 {
			c.logger.DebugVvv().
				Any("recipe", recipe.RecipeName).
				Any("id", existing.ID).
				Msg("Recipe already exists in database")
			recipesExistCount[sig] = true
			continue
		}

		// build the recipe
		r := c.BuildDbRecipeFromEqtradersRecipe(recipe)
		c.AssignRecipeId(&r)

		recipesToCreate = append(recipesToCreate, r)

		// fake caching the recipe so we don't create it twice
		c.recipeLookup[sig] = r
		recipesExistCount[sig] = true

		// print the recipe
		if printRecipes {
			c.PrintEqtradersRecipeFromDbRecipe(r)
		}
	}

	if len(recipesToCreate) > 0 {
		err := c.db.CreateInBatches(recipesToCreate, insertBatchSize).Error
		if err != nil {
			c.logger.Fatal().Err(err).Msg("Failed to create recipes")
		}
	}

	// linked recipes (pre-reqs to other recipes in the same era)
	linkedRecipesExistCount := make(map[string]bool)
	linkedToBeCreatedCount := 0
	linkedPasses := 0
	if !skipLinked {
		c.logger.Info().
			Any("expansion", expansionNumber).
			Msg("Importing linked recipes")

		for pass := 0; pass < linkedMaxPasses; pass++ {
			toBeCreatedInPass := 0
			c.loadDatabaseRecipes()

			// reset
			recipesToCreate = []models.TradeskillRecipe{}

			// linked recipes
			// this is for creating recipes that may not be tagged within an era but are a recipe that
			// is pre-requisites for another recipe within the era
			for _, r := range c.existingRecipes {
				if expansionNumber != -1 && r.MinExpansion != int8(expansionNumber) {
					continue
				}

				for _, e := range r.TradeskillRecipeEntries {
					for _, recipe := range recipes {
						if e.Componentcount > 0 && e.ItemId == recipe.RecipeItemId {
							lookupKey := GetRecipeSignature(recipe)
							if existingRecipe, ok := c.recipeLookup[lookupKey]; ok {
								c.logger.DebugVvv().
									Any("parent_id", r.ID).
									Any("parent", r.Name).
									Any("id", existingRecipe.ID).
									Any("recipe", recipe.RecipeName).
									Any("expansion", recipe.ExpansionId).
									Msg("Component is a pre-req recipe for another recipe (Already created)")
								linkedRecipesExistCount[lookupKey] = true
								continue
							} else {
								c.logger.DebugVvv().
									Any("parent_id", r.ID).
									Any("parent", r.Name).
									Any("recipe", recipe.RecipeName).
									Msg("Component is a pre-req recipe for another recipe (To be created)")

								linkedToBeCreatedCount++
								toBeCreatedInPass++

								// fake caching the recipe so we don't create it twice
								c.recipeLookup[GetRecipeSignature(recipe)] = r

								// build the recipe
								nr := c.BuildDbRecipeFromEqtradersRecipe(recipe)
								c.AssignRecipeId(&r)
								nr.Notes = null.StringFrom(
									fmt.Sprintf("%v (Linked pre-req recipe to %v)",
										nr.Notes.String,
										r.Name,
									),
								)

								if nr.MinExpansion > int8(expansionNumber) {
									nr.MinExpansion = int8(expansionNumber)
								}

								if printRecipes {
									c.PrintEqtradersRecipeFromDbRecipe(nr)
								}

								recipesToCreate = append(recipesToCreate, nr)
							}
						}
					}
				}
			}

			if len(recipesToCreate) > 0 {
				err := c.db.CreateInBatches(recipesToCreate, insertBatchSize).Error
				if err != nil {
					c.logger.Fatal().Err(err).Msg("Failed to create linked recipes")
				}
				c.logger.Info().
					Any("recipes created", len(recipesToCreate)).
					Any("pass", pass).
					Msg("Linked recipes created")
			}

			if toBeCreatedInPass == 0 {
				linkedPasses = pass
				break
			}
		}
	}

	c.logger.Info().
		Any("recipes created", len(recipesToCreate)).
		Any("recipes already created", len(recipesExistCount)).
		Any("linked passes", linkedPasses).
		Any("linked recipes already created", len(linkedRecipesExistCount)).
		Any("linked recipes to be created", linkedToBeCreatedCount).
		Msg("Imported recipes summary")
}

// loadDatabaseRecipes loads all database recipes into a cache
func (c *ImportCommand) loadDatabaseRecipes() {
	c.existingRecipes = []models.TradeskillRecipe{}
	c.db.Preload("TradeskillRecipeEntries.TradeskillRecipe").Find(&c.existingRecipes)

	c.recipeLookup = make(map[string]models.TradeskillRecipe)

	// loop through all recipes
	for _, recipe := range c.existingRecipes {
		key := GetDbRecipeSignature(recipe)
		if _, ok := c.recipeLookup[key]; ok {
			continue
		}

		c.recipeLookup[key] = recipe
	}
}

// BuildDbRecipeFromEqtradersRecipe builds a database recipe model from an eqtraders recipe
// it also includes building the tradeskill_recipe_entries
// it does not have an ID assigned yet
// it does not interact with the database directly
func (c *ImportCommand) BuildDbRecipeFromEqtradersRecipe(recipe Recipe) models.TradeskillRecipe {
	var r models.TradeskillRecipe
	r.Name = recipe.RecipeName
	r.Tradeskill = int16(recipe.Skill.SkillId)
	r.Trivial = int16(recipe.Trivial)
	r.Nofail = int8(0)
	r.MinExpansion = int8(recipe.ExpansionId)
	r.MaxExpansion = 99
	r.Quest = int8(0)
	if recipe.Skill.SkillId == 75 {
		r.Quest = int8(1)
	}

	var nofail int8
	if recipe.NoFail {
		nofail = 1
	}

	r.Nofail = nofail

	// notes
	var additionalNotes []string
	if recipe.LearnedByItem.ItemName != "" {
		additionalNotes = append(
			additionalNotes,
			fmt.Sprintf("Learned by item [%v]", recipe.LearnedByItem.ItemName),
		)
	}

	r.Notes = null.StringFrom(
		fmt.Sprintf("%v - %v (eqtraders import) %v",
			recipe.ExpansionName,
			recipe.Skill.SkillName,
			strings.Join(additionalNotes, ", "),
		),
	)

	// other fields
	r.Enabled = int8(1)
	r.MustLearn = int8(0)
	// replace container appear to be mostly quest based recipes
	// todo handle this with quest cases
	r.ReplaceContainer = int8(0)
	if recipe.ConsumeContainer {
		r.ReplaceContainer = int8(1)
	}

	r.ContentFlags = null.StringFrom("")
	r.ContentFlagsDisabled = null.StringFrom("")
	r.Skillneeded = int16(recipe.RequiredSkillLevel)

	// learned by item lookups
	if len(recipe.LearnedByItem.ItemName) > 0 {
		if i, ok := c.itemLookup[recipe.LearnedByItem.ItemName]; ok {
			r.LearnedByItemId = i.ID
		} else {
			var item models.Item
			c.db.Where("Name = ?", recipe.LearnedByItem.ItemName).First(&item)

			if item.ID > 0 {
				c.itemLookup[recipe.LearnedByItem.ItemName] = item
				r.LearnedByItemId = item.ID
			}
		}
		if r.LearnedByItemId > 0 {
			r.MustLearn = int8(1)
		}
	}

	var components []models.TradeskillRecipeEntry

	// tradeskill_recipe_entries | components (componentcount > 0)
	for _, component := range recipe.Components {
		var e models.TradeskillRecipeEntry
		e.ItemId = component.ItemId
		e.Componentcount = int8(component.Count)
		if e.Componentcount == 0 {
			e.Componentcount = 1
		}

		for _, returns := range recipe.FailureReturns {
			if returns.ItemId == component.ItemId {
				e.Failcount = int8(returns.Count)
				break
			}
		}

		for _, returns := range recipe.Returns {
			if returns.ItemId == component.ItemId {
				e.Successcount = int8(returns.Count)
				break
			}
		}

		if e.Failcount == 0 && !recipe.NoFail {
			e.Salvagecount = int8(component.Count)
		}

		components = append(components, e)
	}

	// tradeskill_recipe_entries | container(s) combined in (iscontainer = 1)
	for _, in := range recipe.In {
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
		e.ItemId = in.ItemId
		e.Iscontainer = int8(1)
		components = append(components, e)
	}

	// tradeskill_recipe_entries | item result row (successcount > 0)
	if recipe.RecipeItemId > 0 {
		var e models.TradeskillRecipeEntry
		e.ItemId = recipe.RecipeItemId
		e.Successcount = int8(recipe.Yield)
		components = append(components, e)
	}

	r.TradeskillRecipeEntries = components

	return r
}

// PrintEqtradersRecipeFromDbRecipe prints a recipe to the console
func (c *ImportCommand) PrintEqtradersRecipeFromDbRecipe(r models.TradeskillRecipe) {
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("> Recipe | %v\n", r.Name)
	fmt.Println(strings.Repeat("-", 80))

	fmt.Printf("> Skill %17v | Trivial (%v)\n",
		fmt.Sprintf("%v (%v)", r.Tradeskill, r.Tradeskill),
		r.Trivial,
	)
	fmt.Println(strings.Repeat("-", 80))

	for _, component := range r.TradeskillRecipeEntries {
		if component.Componentcount > 0 {
			fmt.Printf(
				"|--- Component %10v | %-40v | Count (%v) Successcount (%v) Failcount (%v) Salvage (%v)\n",
				component.ItemId,
				c.LookupItemNameById(component.ItemId),
				component.Componentcount,
				component.Successcount,
				component.Failcount,
				component.Salvagecount,
			)
		} else if component.Iscontainer > 0 {
			fmt.Printf("|--- Container %10v | %-40v\n", component.ItemId, c.LookupItemNameById(component.ItemId))
		} else if component.Successcount > 0 {
			fmt.Printf("|--- Item Result %8v | %-40v | Count (%v)\n", component.ItemId, c.LookupItemNameById(component.ItemId), component.Successcount)
		} else {
			fmt.Printf("|--- Unknown %10v\n", component.ItemId)
		}
	}

	if r.LearnedByItemId > 0 {
		fmt.Printf("|--- Learned by item: %v (%v)\n", c.LookupItemNameById(r.LearnedByItemId), r.LearnedByItemId)
	}
}

var itemLookupById map[int]string

// loadItemNameLookup loads all items into a cache
func (c *ImportCommand) loadItemNameLookup() {
	if itemLookupById == nil {
		itemLookupById = make(map[int]string)
		var items []models.Item
		c.db.Select("id", "Name").Find(&items)
		for _, item := range items {
			itemLookupById[item.ID] = item.Name
		}

		c.logger.Info().
			Any("items", len(itemLookupById)).
			Msg("Loaded items into cache")
	}
}

// LookupItemNameById returns an item from the database by ID
func (c *ImportCommand) LookupItemNameById(itemId int) string {

	// use itemLookupById if it exists
	if itemLookupById != nil {
		if item, ok := itemLookupById[itemId]; ok {
			return item
		}
	}

	for _, objectType := range objectTypes {
		if objectType.Type == itemId {
			return objectType.Name + " (World container)"
		}
	}

	return fmt.Sprintf("Unknown Item %v", itemId)
}

// GetDbRecipeFromEqTradersSignature returns a database recipe from an eqtraders tradeskill signature
func (c *ImportCommand) GetDbRecipeFromEqTradersSignature(recipe Recipe) models.TradeskillRecipe {
	if recipe, ok := c.recipeLookup[GetRecipeSignature(recipe)]; ok {
		return recipe
	}

	return models.TradeskillRecipe{}
}

// GetNextTradeskillId returns the next tradeskill id
func (c *ImportCommand) GetNextTradeskillId() int {
	if c.lastTradeskillId == 0 {
		var tradeskill models.TradeskillRecipe
		c.db.Last(&tradeskill)
		c.lastTradeskillId = tradeskill.ID + 1
	}

	c.lastTradeskillId++

	return c.lastTradeskillId
}

// AssignRecipeId assigns a recipe id to a tradeskill recipe
func (c *ImportCommand) AssignRecipeId(m *models.TradeskillRecipe) {
	m.ID = c.GetNextTradeskillId()
	for i, _ := range m.TradeskillRecipeEntries {
		m.TradeskillRecipeEntries[i].RecipeId = m.ID
	}
}

// deleteTradeskillRecipesByExpansion deletes all tradeskill recipes by expansion
func (c *ImportCommand) deleteTradeskillRecipesByExpansion(expansionID int) error {
	// Delete tradeskill_recipe_entries where recipe_id exists in tradeskill_recipe with the given min_expansion
	err := c.db.Exec(`
		DELETE FROM tradeskill_recipe_entries 
		WHERE recipe_id IN (
			SELECT id FROM tradeskill_recipe WHERE min_expansion = ?
		)`, expansionID).Error
	if err != nil {
		return err
	}

	// Delete tradeskill_recipe where min_expansion matches the given expansionID
	err = c.db.Where("min_expansion = ?", expansionID).Delete(&models.TradeskillRecipe{}).Error
	if err != nil {
		return err
	}

	c.logger.Info().Any("expansion", expansionID).Msg("Deleted tradeskill recipes by expansion")

	return nil
}

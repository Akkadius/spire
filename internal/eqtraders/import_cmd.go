package eqtraders

import (
	"encoding/json"
	"fmt"
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
}

func (c *ImportCommand) Command() *cobra.Command {
	return c.command
}

func NewImportCommand(
	db *gorm.DB,
) *ImportCommand {
	i := &ImportCommand{
		db: db,
		command: &cobra.Command{
			Use:   "eq-traders:import [expansion_number]",
			Short: "Imports data from eqtraders.com using data scraped via eq-traders:scrape",
		},
	}

	i.command.Args = cobra.MinimumNArgs(1)
	i.command.Run = i.Handle
	i.command.Flags().StringVarP(&singleRecipe, "single-recipe", "r", "", "Scrape a single recipe by name")

	return i
}

func (c *ImportCommand) Handle(cmd *cobra.Command, args []string) {
	expansion := os.Args[2]

	itemLookupCache := make(map[string]models.Item)

	expansionNumber, err := strconv.Atoi(expansion)
	if err != nil {
		expansionNumber = -1
	}

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

	existingRecipeLookup := make(map[string]bool)

	// loop through all recipes
	for _, recipe := range existingRecipes {

		// create a unique key for each recipe
		itemSummation := 0
		for _, entry := range recipe.TradeskillRecipeEntries {
			if entry.Componentcount > 0 {
				itemSummation += entry.ItemId
			}
		}

		key := fmt.Sprintf(
			"%v-%v-%v-%v",
			recipe.Name,
			recipe.Tradeskill,
			recipe.MinExpansion,
			itemSummation,
		)

		if _, ok := existingRecipeLookup[key]; ok {
			continue
		}

		existingRecipeLookup[key] = true
	}

	for _, recipe := range recipes {
		if len(singleRecipe) > 0 && recipe.RecipeName != singleRecipe {
			continue
		}

		//pp.Println(recipe)

		// ignore recipes that are not for the expansion we are importing
		if recipe.ExpansionId != expansionNumber && expansion != "all" {
			continue
		}

		var r models.TradeskillRecipe

		itemSummation := 0
		for _, entry := range recipe.Components {
			itemSummation += entry.ItemId
		}

		key := fmt.Sprintf(
			"%v-%v-%v-%v",
			recipe.RecipeName,
			recipe.Skill.SkillId,
			recipe.ExpansionId,
			itemSummation,
		)

		hasExistingRecipe := false
		for _, existingRecipe := range existingRecipes {
			if _, ok := existingRecipeLookup[key]; ok {
				fmt.Printf("Found existing recipe: %v\n", existingRecipe.Name)
				r = existingRecipe
				hasExistingRecipe = true
				break
			}
		}
		//
		//if recipe.RecipeName == "Eminent Boot Symbol of the Warmonger" {
		//	pp.Println(recipe)
		//	pp.Println("Traders")
		//	pp.Println(key)
		//	pp.Println("Database")
		//	pp.Println(r)
		//}

		r.Name = recipe.RecipeName
		r.Tradeskill = int16(recipe.Skill.SkillId)
		r.Trivial = int16(recipe.Trivial)
		r.Nofail = int8(0)
		r.MinExpansion = int8(recipe.ExpansionId)
		r.MaxExpansion = 99

		var nofail int8
		if recipe.NoFail {
			nofail = 1
		}

		r.Nofail = nofail

		var additionalNotes []string
		if recipe.LearnedByItem.ItemName != "" {
			additionalNotes = append(
				additionalNotes,
				fmt.Sprintf(" - Learned by item [%v]", recipe.LearnedByItem.ItemName),
			)
		}

		r.Notes = null.StringFrom(
			fmt.Sprintf("%v-%v%v (eqtraders import)",
				recipe.ExpansionName,
				recipe.Skill.SkillName,
				strings.Join(additionalNotes, ", "),
			),
		)
		r.Enabled = int8(1)
		r.MustLearn = int8(0)
		r.Quest = int8(0)
		r.ReplaceContainer = int8(0) // todo - figure out what this is
		r.ContentFlags = null.StringFrom("")
		r.ContentFlagsDisabled = null.StringFrom("")
		r.Skillneeded = 0 // todo - figure out what this is

		if len(recipe.LearnedByItem.ItemName) > 0 {
			if i, ok := itemLookupCache[recipe.LearnedByItem.ItemName]; ok {
				r.LearnedByItemId = i.ID
			} else {
				var item models.Item
				c.db.Where("Name = ?", recipe.LearnedByItem.ItemName).First(&item)

				if item.ID > 0 {
					itemLookupCache[recipe.LearnedByItem.ItemName] = item
					r.LearnedByItemId = item.ID
				}
			}
			if r.LearnedByItemId > 0 {
				r.MustLearn = int8(1)
				fmt.Printf("Learned by item: %v (%v)\n", recipe.LearnedByItem.ItemName, r.LearnedByItemId)
			}
		}

		// insert recipe into database
		c.db.Save(&r)

		// insert into cache
		existingRecipes = append(existingRecipes, r)

		if r.ID == 0 {
			log.Fatalf("Error inserting recipe id: %v name: %v into database", r.ID, r.Name)
		}

		existing := "re-import"
		if !hasExistingRecipe {
			existing = "new"
		}

		fmt.Println(strings.Repeat("-", 80))
		fmt.Printf("> Importing recipe: %v (%v)\n", recipe.RecipeName, existing)
		fmt.Println(strings.Repeat("-", 80))

		fmt.Printf("> Skill %17v | Trivial (%v)\n",
			fmt.Sprintf("%v (%v)", recipe.Skill.SkillName, recipe.Skill.SkillId),
			recipe.Trivial,
		)
		fmt.Println(strings.Repeat("-", 80))

		var components []models.TradeskillRecipeEntry

		// insert components into database
		for _, component := range recipe.Components {
			var e models.TradeskillRecipeEntry
			e.RecipeId = r.ID
			e.ItemId = component.ItemId
			e.Componentcount = int8(component.Count)

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

			if e.Failcount == 0 {
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
			e.RecipeId = r.ID
			e.ItemId = in.ItemId
			e.Iscontainer = int8(1)
			components = append(components, e)
			fmt.Printf("|--- Container %10v | %-40v\n", in.ItemId, in.ItemName)
		}

		// insert item result itself
		if recipe.RecipeItemId > 0 {
			var e models.TradeskillRecipeEntry
			e.RecipeId = r.ID
			e.ItemId = recipe.RecipeItemId
			e.Successcount = int8(recipe.Yield)
			components = append(components, e)
			fmt.Printf("|--- Item Result %8v | %-40v | Count (%v)\n", recipe.RecipeItemId, recipe.RecipeName, recipe.Yield)
		}

		if hasExistingRecipe {
			c.db.Where("recipe_id = ?", r.ID).Delete(&models.TradeskillRecipeEntry{})
		}

		c.db.CreateInBatches(components, 100)

		//if expansion

	}
}

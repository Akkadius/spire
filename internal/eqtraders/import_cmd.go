package eqtraders

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
)

type ImportCommand struct {
	db      *gorm.DB
	logger  *logrus.Logger
	command *cobra.Command
}

func (c *ImportCommand) Command() *cobra.Command {
	return c.command
}

func NewImportCommand(
	db *gorm.DB,
	logger *logrus.Logger,
) *ImportCommand {
	i := &ImportCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use:   "eq-traders:import [expansion_number]",
			Short: "Imports data from eqtraders.com using data scraped via eq-traders:scrape",
		},
	}

	i.command.Args = cobra.MinimumNArgs(1)
	i.command.Run = i.Handle

	return i
}

func (c *ImportCommand) Handle(cmd *cobra.Command, args []string) {
	expansion := os.Args[2]
	if expansion == "all" {

	}

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
	c.db.Find(&existingRecipes)

	for _, recipe := range recipes {
		// ignore recipes that are not for the expansion we are importing
		if recipe.ExpansionId != expansionNumber && expansion != "all" {
			continue
		}

		var r models.TradeskillRecipe

		hasExistingRecipe := false
		for _, existingRecipe := range existingRecipes {
			if existingRecipe.Name == recipe.RecipeName {
				fmt.Printf("Found existing recipe: %v\n", existingRecipe.Name)
				r = existingRecipe
				hasExistingRecipe = true
				break
			}
		}

		r.Name = recipe.RecipeName
		r.Tradeskill = int16(recipe.Skill.SkillId)
		r.Trivial = int16(recipe.Trivial)
		r.Nofail = int8(0)
		r.MinExpansion = int8(recipe.ExpansionId)
		r.MaxExpansion = 99
		r.Notes = null.StringFrom(fmt.Sprintf("%v-%v (eqtraders import)", recipe.ExpansionName, recipe.Skill.SkillName))
		r.Enabled = int8(1)
		r.MustLearn = int8(0)
		r.Quest = int8(0)
		r.ReplaceContainer = int8(0) // todo - figure out what this is
		r.ContentFlags = null.StringFrom("")
		r.ContentFlagsDisabled = null.StringFrom("")
		r.Skillneeded = 0 // todo - figure out what this is

		// insert recipe into database
		c.db.Save(&r)

		if r.ID == 0 {
			c.logger.Fatalf("Error inserting recipe id: %v name: %v into database", r.ID, r.Name)
		}

		existing := "re-import"
		if !hasExistingRecipe {
			existing = "new"
		}

		fmt.Println(strings.Repeat("-", 40))
		fmt.Printf("> Importing recipe: %v (%v)\n", recipe.RecipeName, existing)
		fmt.Println(strings.Repeat("-", 40))

		var components []models.TradeskillRecipeEntry

		// insert components into database
		for _, component := range recipe.Components {
			var e models.TradeskillRecipeEntry
			e.RecipeId = r.ID
			e.ItemId = component.ItemId
			e.Componentcount = int8(component.Count)
			components = append(components, e)
			fmt.Printf("|--- Component: %v (%v) Count %v\n", component.ItemId, component.ItemName, component.Count)
		}

		// insert in into database
		for _, in := range recipe.In {
			var e models.TradeskillRecipeEntry
			e.RecipeId = r.ID
			e.ItemId = in.ItemId
			e.Iscontainer = int8(1)
			components = append(components, e)
			fmt.Printf("|--- In: %v (%v)\n", in.ItemId, in.ItemName)
		}

		// insert returns into database
		for _, returns := range recipe.Returns {
			var e models.TradeskillRecipeEntry
			e.RecipeId = r.ID
			e.ItemId = returns.ItemId
			e.Successcount = int8(returns.Count)
			components = append(components, e)
			fmt.Printf("|--- Returns: %v (%v) Count %v\n", returns.ItemId, returns.ItemName, returns.Count)
		}

		// insert failure returns into database
		for _, failureReturns := range recipe.FailureReturns {
			var e models.TradeskillRecipeEntry
			e.RecipeId = r.ID
			e.ItemId = failureReturns.ItemId
			e.Failcount = int8(failureReturns.Count)
			components = append(components, e)
			fmt.Printf("|--- Failure Returns: %v (%v) Count %v\n", failureReturns.ItemId, failureReturns.ItemName, failureReturns.Count)
		}

		if hasExistingRecipe {
			c.db.Where("recipe_id = ?", r.ID).Delete(&models.TradeskillRecipeEntry{})
		}

		c.db.CreateInBatches(components, 100)

		//if expansion

	}
}

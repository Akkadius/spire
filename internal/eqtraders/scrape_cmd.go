package eqtraders

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/models"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/gammazero/workerpool"
	"github.com/gosimple/slug"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type ScrapeCommand struct {
	db      *gorm.DB
	command *cobra.Command
	logger  *logger.AppLogger
}

func (c *ScrapeCommand) Command() *cobra.Command {
	return c.command
}

var skipLookups bool
var singleRecipe string

func NewScrapeCommand(
	db *gorm.DB,
	logger *logger.AppLogger,
) *ScrapeCommand {
	i := &ScrapeCommand{
		db: db,
		command: &cobra.Command{
			Use:   "eq-traders:scrape [expansion_number]",
			Short: "A command for scraping / downloading eq traders recipes. Use eq-traders:import to import the data into the database.",
		},
		logger: logger,
	}

	i.command.Args = cobra.MinimumNArgs(1)
	i.command.Run = i.Handle

	i.command.Flags().BoolVarP(&skipLookups, "skip-lookups", "s", false, "Skip lookups for items that are not found in the database")
	i.command.Flags().StringVarP(&singleRecipe, "single-recipe", "r", "", "Scrape a single recipe by name")

	return i
}

// Handle implementation of the Command interface
func (c *ScrapeCommand) Handle(cmd *cobra.Command, args []string) {
	err := os.MkdirAll("data/eqtraders", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll("data/eqtraders/site-cache/", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	list := []ExpansionRecipe{
		{ExpId: -1, ExpName: "All", PageTitle: "Smithing Recipes", Url: "https://eqtraders.com/recipes/recipe_page.php?article=338&rsa=Smithing"},
		{ExpId: -1, ExpName: "All", PageTitle: "Tailoring Recipes", Url: "https://eqtraders.com/recipes/recipe_page.php?article=338&rsa=Tailoring"},
		{ExpId: -1, ExpName: "All", PageTitle: "Tinkering Recipes", Url: "https://eqtraders.com/recipes/recipe_page.php?article=338&rsa=Tinkering"},
		{ExpId: -1, ExpName: "All", PageTitle: "Pottery Recipes", Url: "https://eqtraders.com/recipes/recipe_page.php?article=338&rsa=Pottery"},
		{ExpId: -1, ExpName: "All", PageTitle: "Bakery Recipes", Url: "https://eqtraders.com/recipes/recipe_page.php?article=338&rsa=Baking"},
		{ExpId: -1, ExpName: "All", PageTitle: "Brewing Recipes", Url: "https://eqtraders.com/recipes/recipe_page.php?article=338&rsa=Brewing"},
		{ExpId: -1, ExpName: "All", PageTitle: "Alchemy Recipes", Url: "https://eqtraders.com/recipes/recipe_page.php?article=338&rsa=Alchemy"},
		{ExpId: -1, ExpName: "All", PageTitle: "Fletching Recipes", Url: "https://eqtraders.com/recipes/recipe_page.php?article=338&rsa=Fletching"},
		{ExpId: -1, ExpName: "All", PageTitle: "Jewelcraft Recipes", Url: "https://eqtraders.com/recipes/recipe_page.php?article=338&rsa=Jewelcraft"},
		{ExpId: -1, ExpName: "All", PageTitle: "Make Poison Recipes", Url: "https://eqtraders.com/recipes/recipe_page.php?article=338&rsa=make%20poison"},
		{ExpId: -1, ExpName: "All", PageTitle: "Spell Research Recipes", Url: "https://eqtraders.com/recipes/recipe_page.php?article=338&rsa=Spell%20Research"},
		{ExpId: -1, ExpName: "All", PageTitle: "No Skill Recipes", Url: "https://eqtraders.com/recipes/recipe_page.php?article=338&rsa=No%20Skill"},
		{ExpId: 2, ExpName: "Scars of Velious", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=148&rsa=Baking&sub=velluc&sb=item&menustr=080020040000"},
		{ExpId: 2, ExpName: "Scars of Velious", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=135&rsa=Brewing&sb=item&sub=SoV&menustr=080110120000"},
		{ExpId: 2, ExpName: "Scars of Velious", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1463&rsa=Jewelcraft&sub=SoV&sb=item&menustr=080070000000"},
		{ExpId: 2, ExpName: "Scars of Velious", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1457&rsa=Smithing&sb=item&sub=SoV&menustr=080100000000"},
		{ExpId: 2, ExpName: "Scars of Velious", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=134&rsa=Tailoring&sb=item&sub=SoV&menustr=080110100000"},
		{ExpId: 2, ExpName: "Scars of Velious", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=136&rsa=Tinkering&sb=item&sub=SoV&menustr=080110140000"},
		{ExpId: 3, ExpName: "Shadows of Luclin", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1467&rsa=Baking&sub=velluc&sb=item&menustr=080020040000"},
		{ExpId: 3, ExpName: "Shadows of Luclin", PageTitle: "Smithing Chain Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=128&rsa=Smithing&sb=item&sub=lchain&menustr=080100120000"},
		{ExpId: 3, ExpName: "Shadows of Luclin", PageTitle: "Smithing Plate Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=129&rsa=Smithing&sb=item&sub=lplate&menustr=080100140000"},
		{ExpId: 3, ExpName: "Shadows of Luclin", PageTitle: "Tailoring Leather Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=170&rsa=Tailoring&sb=item&sub=llthr&menustr=080110146000"},
		{ExpId: 3, ExpName: "Shadows of Luclin", PageTitle: "Tailoring Silk Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=171&rsa=Tailoring&sb=item&sub=lslk&menustr=080110164000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=149&rsa=Baking&sub=pop&sb=item&menustr=080020060000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=149&rsa=Baking&sub=pop&sb=item&menustr=080020060000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=166&rsa=Jewelcraft&sb=item&sub=pop&menustr=080070050000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=166&rsa=Jewelcraft&sb=item&sub=pop&menustr=080070050000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Planes of Power Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=240&rsa=Smithing&sb=item&sub=pop&menustr=080100160000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Planes of Power Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=240&rsa=Smithing&sb=item&sub=pop&menustr=080100160000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=182&rsa=Pottery&sub=pop&sb=item&menustr=080090050000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=182&rsa=Pottery&sub=pop&sb=item&menustr=080090050000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Smithing Chain Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=130&rsa=Smithing&sb=item&sub=pchain&menustr=080100160000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Smithing Chain Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=130&rsa=Smithing&sb=item&sub=pchain&menustr=080100160000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Smithing Plate Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=131&rsa=Smithing&sb=item&sub=pplate&menustr=080100180000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Smithing Plate Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=131&rsa=Smithing&sb=item&sub=pplate&menustr=080100180000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Smithing Quest Chain Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=168&rsa=Smithing&sb=item&sub=pqchain&menustr=080100200000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Smithing Quest Chain Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=168&rsa=Smithing&sb=item&sub=pqchain&menustr=080100200000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Smithing Quest Plate Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=169&rsa=Smithing&sb=item&sub=pqplate&menustr=080100220000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Smithing Quest Plate Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=169&rsa=Smithing&sb=item&sub=pqplate&menustr=080100220000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Tailoring Leather Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=137&rsa=Tailoring&sb=item&sub=plthr&menustr=080110160000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Tailoring Leather Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=137&rsa=Tailoring&sb=item&sub=plthr&menustr=080110160000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Tailoring Quest Leather Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=172&rsa=Tailoring&sb=item&sub=pqlthr&menustr=080110200000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Tailoring Quest Leather Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=172&rsa=Tailoring&sb=item&sub=pqlthr&menustr=080110200000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Tailoring Quest Silk Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=173&rsa=Tailoring&sb=item&sub=pqslk&menustr=080110220000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Tailoring Quest Silk Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=173&rsa=Tailoring&sb=item&sub=pqslk&menustr=080110220000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Tailoring Silk Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=138&rsa=Tailoring&sb=item&sub=pslk&menustr=080110180000"},
		{ExpId: 4, ExpName: "Planes of Power", PageTitle: "Tailoring Silk Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=138&rsa=Tailoring&sb=item&sub=pslk&menustr=080110180000"},
		{ExpId: 5, ExpName: "Legacy of Ykesha", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=150&rsa=Baking&sub=loyek&sb=item&menustr=080020080000"},
		{ExpId: 5, ExpName: "Legacy of Ykesha", PageTitle: "Brewing Ribbon Dye Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=180&rsa=Brewing&sb=item&sub=dye&menustr=080030035000"},
		{ExpId: 5, ExpName: "Legacy of Ykesha", PageTitle: "Spell Research Priest Spell Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=124&rsa=Spell%20Research&sub=deity&sb=item&menustr=080105100000"},
		{ExpId: 5, ExpName: "Legacy of Ykesha", PageTitle: "Spell Research Ribbon Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=125&rsa=Spell%20Research&sub=loyek&sb=item&menustr=080105120000"},
		{ExpId: 5, ExpName: "Legacy of Ykesha", PageTitle: "Tailoring Ribbon Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=133&rsa=Tailoring&sb=item&sub=ribbon&menustr=080110040000"},
		{ExpId: 6, ExpName: "Lost Dungeons of Norrath", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=222&rsa=Baking&sub=LDON&sb=item&menustr=080020120000"},
		{ExpId: 6, ExpName: "Lost Dungeons of Norrath", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=223&rsa=Brewing&sub=LDON&sb=item&menustr=080030038000"},
		{ExpId: 6, ExpName: "Lost Dungeons of Norrath", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=224&rsa=Fletching&sub=LDON&sb=item&menustr=080060023000"},
		{ExpId: 6, ExpName: "Lost Dungeons of Norrath", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=225&rsa=Jewelcraft&sub=LDON&sb=item&menustr=080070060000"},
		{ExpId: 6, ExpName: "Lost Dungeons of Norrath", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=226&rsa=Smithing&sub=LDON&sb=item&menustr=080100280000"},
		{ExpId: 6, ExpName: "Lost Dungeons of Norrath", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=430&rsa=Tailoring&sub=LDON&sb=item&menustr=080110700000"},
		{ExpId: 7, ExpName: "Gates of Discord", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=228&rsa=Baking&sub=GoD&sb=item&menustr=080020160000"},
		{ExpId: 7, ExpName: "Gates of Discord", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=229&rsa=Brewing&sb=item&sub=GoD&menustr=080030039000"},
		{ExpId: 7, ExpName: "Gates of Discord", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=231&rsa=Fletching&sb=item&sub=GoD&menustr=080060024000"},
		{ExpId: 7, ExpName: "Gates of Discord", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=232&rsa=Jewelcraft&sb=item&sub=GoD&menustr=080070065000"},
		{ExpId: 7, ExpName: "Gates of Discord", PageTitle: "Make Poison Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=233&rsa=Make%20Poison&sb=item&sub=GoD&menustr=080080040000"},
		{ExpId: 7, ExpName: "Gates of Discord", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=234&rsa=Pottery&sub=GoD&sb=item&menustr=080090055000"},
		{ExpId: 7, ExpName: "Gates of Discord", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=235&rsa=Smithing&sb=item&sub=GoD&menustr=080100320000"},
		{ExpId: 7, ExpName: "Gates of Discord", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=236&rsa=Tailoring&sb=item&sub=GoD&menustr=080110230000"},
		{ExpId: 7, ExpName: "Gates of Discord", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=237&rsa=Tinkering&rc=GNM&sub=GoD&sb=item&menustr=080120080000"},
		{ExpId: 8, ExpName: "Omens of War", PageTitle: "Alchemy Augment Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=258&rsa=Alchemy&rc=SHM&sub=OoW&sb=item&menustr=080010500000"},
		{ExpId: 8, ExpName: "Omens of War", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=282&rsa=Baking&sub=OoW&sb=item&menustr=080020140000"},
		{ExpId: 8, ExpName: "Omens of War", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=265&rsa=Brewing&sub=OoW&sb=item&menustr=080030060000"},
		{ExpId: 8, ExpName: "Omens of War", PageTitle: "Fletching", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=272&rsa=Fletching&sub=OoW&sb=item&menustr=080060042000"},
		{ExpId: 8, ExpName: "Omens of War", PageTitle: "Jewelcraft Stone Cutting Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=264&rsa=Jewelcraft&sub=OoW&sb=item&menustr=080070485000"},
		{ExpId: 8, ExpName: "Omens of War", PageTitle: "Poison Making Recipes, for Augments and Actual Poisons", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=262&rsa=Make%20Poison&rc=ROG&sub=OoW&sb=item&menustr=080080045000"},
		{ExpId: 8, ExpName: "Omens of War", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=283&rsa=Pottery&sub=OoW&sb=item&menustr=080090058000"},
		{ExpId: 8, ExpName: "Omens of War", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=263&rsa=Smithing&sub=OoW&sb=item&menustr=080100610000"},
		{ExpId: 8, ExpName: "Omens of War", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=271&rsa=Tailoring&sb=item&sub=OoW&menustr=080110150000"},
		{ExpId: 8, ExpName: "Omens of War", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=266&rsa=Tinkering&rc=GNM&sub=OoW&sb=item&menustr=080120095000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=268&rsa=Baking&sb=item&sub=dron&menustr=080020000000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=295&rsa=Brewing&sub=dron&sb=item&menustr=080030000000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=296&rsa=Fletching&sb=item&sub=dron&menustr=080060000000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=298&rsa=Jewelcraft&sb=item&sub=dron&menustr=080070000000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=299&rsa=Pottery&sub=Dron"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=297&rsa=Tinkering&sub=dron&sb=item&menustr=080120000000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Barbarian Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=301&rsa=Smithing&rc=BAR&sb=item&sub=Dron&menustr=080040010000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Barbarian Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=302&rsa=Tailoring&rc=BAR&sb=item&sub=Dron&menustr=080040010000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Dark Elf Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=324&rsa=Smithing&rc=DEF&sb=item&sub=dron&menustr=080040020030"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Dark Elf Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=325&rsa=Tailoring&rc=DEF&sb=item&sub=dron&menustr=080040020060"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Drakkin Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=278&rsa=Smithing&rc=DRK&sub=DRoN&sb=item&menustr=080040025040"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Drakkin Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=227&rsa=Tailoring&rc=DRK&sub=DRoN&sb=item&menustr=080040025000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Dwarven Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=326&rsa=Smithing&rc=DWF&sb=item&sub=dron&menustr=080040030080"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Erudite Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=330&rsa=Smithing&rc=ERU&sb=item&sub=dron&menustr=080040040030"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Erudite Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=329&rsa=Tailoring&rc=ERU&sb=item&sub=dron&menustr=080040040120"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Froglok Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=313&rsa=Smithing&rc=FRG&sub=dron&sb=item&menustr=080040043020"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Froglok Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=314&rsa=Tailoring&rc=FRG&sb=item&menustr=080040043040"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Gnome Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=305&rsa=Tailoring&rc=GNM&sb=item&sub=dron&menustr=080040045060"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Gnome Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=304&rsa=Smithing&rc=GNM&sb=item&sub=dron&menustr=080040045020"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Half Elf Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=310&rsa=Smithing&rc=HEF&sb=item&sub=dron&menustr=080040048020"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Half Elf Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=311&rsa=Tailoring&rc=HEF&sb=item&sub=dron&menustr=080040048040"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Halfling Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=331&rsa=Smithing&rc=HFL&sb=item&sub=dron&menustr=080040050030"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Halfling Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=332&rsa=Tailoring&rc=HFL&sb=item&sub=dron&menustr=080040050060"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "High Elf Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=333&rsa=Smithing&rc=HIE&sb=item&sub=dron&menustr=080040060030"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "High Elf Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=334&rsa=Tailoring&rc=HIE&sb=item&sub=dron&menustr=080040060060"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Human Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=307&rsa=Smithing&rc=HUM&sb=item&sub=dron&menustr=080040070000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Human Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=308&rsa=Tailoring&rc=HUM&sb=item&sub=dron&menustr=080040070000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Iksar Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=269&rsa=Smithing&rc=IKS&sb=item&sub=dron&menustr=080040090030"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Iksar Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=335&rsa=Tailoring&rc=IKS&sb=item&sub=dron&menustr=080040090060"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Ogre Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=316&rsa=Smithing&rc=OGR&sb=item&sub=dron&menustr=080040100030"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Ogre Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=317&rsa=Tailoring&rc=OGR&sb=item&sub=dron&menustr=080040100040"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Troll Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=318&rsa=Smithing&rc=TRL&sb=item&sub=dron&menustr=080040110030"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Troll Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=320&rsa=Tailoring&rc=TRL&sb=item&sub=dron&menustr=080040110040"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Vah Shir Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=322&rsa=Smithing&rc=VAH&sub=DroN&sb=item&menustr=080040140020"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Vah Shir Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=323&rsa=Tailoring&rc=VAH&sub=DroN&sb=item&menustr=080040140040"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Wood Elf Cultural Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=336&rsa=Smithing&rc=ELF&sb=item&sub=dron&menustr=080040120050"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Wood Elf Cultural Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=337&rsa=Tailoring&rc=ELF&sb=item&sub=dron&menustr=080040120080"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Agnostic Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=agnostic&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Agnostic Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=agnostic&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Agnostic Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=agnostic&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Bertoxxulous Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=bertox&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Bertoxxulous Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=bertox&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Bertoxxulous Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=bertox&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Brell Serilis Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=brell&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Brell Serilis Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=brell&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Brell Serilis Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=brell&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Bristlebane Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=bristlebane&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Bristlebane Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=bristlebane&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Bristlebane Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=bristlebane&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Cazic Thule Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=cazic&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Cazic Thule Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=cazic&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Cazic Thule Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=cazic&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Erollisi Marr Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=erollisi&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Erollisi Marr Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=erollisi&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Erollisi Marr Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=erollisi&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Innoruuk Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=innoruuk&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Innoruuk Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=innoruuk&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Innoruuk Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=innoruuk&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Karana Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=karana&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Karana Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=karana&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Karana Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=karana&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Mithaniel Marr Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=mithaniel&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Mithaniel Marr Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=mithaniel&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Mithaniel Marr Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=mithaniel&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Prexus Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=prexus&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Prexus Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=prexus&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Prexus Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=prexus&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Quellious Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=quellious&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Quellious Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=quellious&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Quellious Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=quellious&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Rallos Zek Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=rallos&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Rallos Zek Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=rallos&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Rallos Zek Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=rallos&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Rodcet Nife Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=rodcet&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Rodcet Nife Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=rodcet&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Rodcet Nife Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=rodcet&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Solusek Ro Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=solusek&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Solusek Ro Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=solusek&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Solusek Ro Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=solusek&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "The Tribunal Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=tribunal&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "The Tribunal Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=tribunal&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "The Tribunal Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=tribunal&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Tunare Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=tunare&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Tunare Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=tunare&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Tunare Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=tunare&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Veeshan Cultural Augment Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tailoring&sb=item&sub=veeshan&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Veeshan Cultural Augment Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Smithing&sb=item&sub=veeshan&menustr=080110235000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Veeshan Cultural Augment Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=339&rsa=Tinkering&sb=item&sub=veeshan&menustr=080110235000"},
		{ExpId: 10, ExpName: "Depths of Darkhollow", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=346&rsa=Baking&sb=item&sub=dodh&menustr=080020220000"},
		{ExpId: 10, ExpName: "Depths of Darkhollow", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=347&rsa=Brewing&sb=item&sub=dodh&menustr=080030044000"},
		{ExpId: 10, ExpName: "Depths of Darkhollow", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=348&rsa=Fletching&sb=item&sub=dodh&menustr=080060030000"},
		{ExpId: 10, ExpName: "Depths of Darkhollow", PageTitle: "No Skill Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=356&rsa=No%20Skill&sb=item&sub=dodh&menustr=080085140000"},
		{ExpId: 10, ExpName: "Depths of Darkhollow", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=354&rsa=Pottery&sb=item&sub=dodh&menustr=080090000000"},
		{ExpId: 10, ExpName: "Depths of Darkhollow", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=350&rsa=Smithing&sb=item&sub=dodh&menustr=080100770000"},
		{ExpId: 10, ExpName: "Depths of Darkhollow", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=351&rsa=Tailoring&sb=item&sub=dodh&menustr=080110237000"},
		{ExpId: 10, ExpName: "Depths of Darkhollow", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=352&rsa=Tinkering&sub=dodh&sb=item&menustr=080120000000"},
		{ExpId: 11, ExpName: "Prophecy of Ro", PageTitle: "Alchemy Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=29&rsa=Alchemy&sb=item&sub=&menustr=080010000000"},
		{ExpId: 11, ExpName: "Prophecy of Ro", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=373&rsa=Baking&sub=por&sb=item&menustr=080020240000"},
		{ExpId: 11, ExpName: "Prophecy of Ro", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=372&rsa=Brewing&sub=por&sb=item&menustr=080030046000"},
		{ExpId: 11, ExpName: "Prophecy of Ro", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=371&rsa=Fletching&sub=por&sb=item&menustr=080060032000"},
		{ExpId: 11, ExpName: "Prophecy of Ro", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=359&rsa=Pottery&sub=por&sb=item&menustr=080090061000"},
		{ExpId: 11, ExpName: "Prophecy of Ro", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=370&rsa=Tinkering&sub=por&sb=item&menustr=080120098000"},
		{ExpId: 12, ExpName: "The Serpent's Spine", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=398&rsa=Baking&sb=item&sub=tss&menustr=080020260000"},
		{ExpId: 12, ExpName: "The Serpent's Spine", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=399&rsa=Brewing&sb=item&sub=tss&menustr=080030048000"},
		{ExpId: 12, ExpName: "The Serpent's Spine", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=400&rsa=Jewelcraft&sub=tss&sb=item&menustr=080070090000"},
		{ExpId: 12, ExpName: "The Serpent's Spine", PageTitle: "No Skill", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=511&rsa=No%20Skill&sb=item&sub=TSS&menustr=080085000000"},
		{ExpId: 12, ExpName: "The Serpent's Spine", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=397&rsa=Pottery&sub=tss&sb=item&menustr=080090062000"},
		{ExpId: 12, ExpName: "The Serpent's Spine", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=394&rsa=Tinkering&sb=item&sub=tss&menustr=080120098100"},
		{ExpId: 13, ExpName: "The Buried Sea", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=405&rsa=Baking&sb=item&sub=tbs&menustr=080020280000"},
		{ExpId: 13, ExpName: "The Buried Sea", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=409&rsa=Brewing&sb=item&sub=tbs&menustr=080030050000"},
		{ExpId: 13, ExpName: "The Buried Sea", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=406&rsa=Jewelcraft&sb=item&sub=tbs&menustr=080070100000"},
		{ExpId: 13, ExpName: "The Buried Sea", PageTitle: "No Skill", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=510&rsa=No%20Skill&sb=item&sub=TBS&menustr=080085000000"},
		{ExpId: 13, ExpName: "The Buried Sea", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=407&rsa=Pottery&sub=tbs&sb=item&menustr=080090063000"},
		{ExpId: 13, ExpName: "The Buried Sea", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=408&rsa=Smithing&sb=item&sub=tbs&menustr=080100780000"},
		{ExpId: 13, ExpName: "The Buried Sea", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=408&rsa=Tailoring&sb=item&sub=tbs&menustr=080110238000"},
		{ExpId: 13, ExpName: "The Buried Sea", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=406&rsa=Tinkering&sb=item&sub=tbs&menustr=080120098200"},
		{ExpId: 14, ExpName: "Secrets of Faydwer", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=422&rsa=Baking&sb=item&sub=sof&menustr=080020000000"},
		{ExpId: 14, ExpName: "Secrets of Faydwer", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=423&rsa=Brewing&sb=item&sub=sof&menustr=080030000000"},
		{ExpId: 14, ExpName: "Secrets of Faydwer", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=420&rsa=Jewelcraft&sb=item&sub=sof&menustr=080070000000"},
		{ExpId: 14, ExpName: "Secrets of Faydwer", PageTitle: "No Skill", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=508&rsa=No%20Skill&sb=item&sub=SOF&menustr=080085000000"},
		{ExpId: 14, ExpName: "Secrets of Faydwer", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=408&rsa=Smithing&sb=item&sub=sof&menustr=080100000000"},
		{ExpId: 14, ExpName: "Secrets of Faydwer", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=408&rsa=Tailoring&sb=item&sub=sof&menustr=080110000000"},
		{ExpId: 15, ExpName: "Seeds of Destruction", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=433&rsa=Baking&sb=item&sub=Sod&menustr=080020000000"},
		{ExpId: 15, ExpName: "Seeds of Destruction", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=434&rsa=Brewing&sb=item&sub=Sod&menustr=080030000000"},
		{ExpId: 15, ExpName: "Seeds of Destruction", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=519&rsa=Fletching&sb=item&sub=Sod&menustr=080060000000"},
		{ExpId: 15, ExpName: "Seeds of Destruction", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=421&rsa=Jewelcraft&sb=item&sub=Sod&menustr=080070000000"},
		{ExpId: 15, ExpName: "Seeds of Destruction", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=432&rsa=Pottery&sb=item&sub=SoD&menustr=080090000000"},
		{ExpId: 15, ExpName: "Seeds of Destruction", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=435&rsa=Smithing&sb=item&sub=Sod&menustr=080100000000"},
		{ExpId: 15, ExpName: "Seeds of Destruction", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=436&rsa=Tailoring&sb=item&sub=Sod&menustr=080110000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=437&rsa=Baking&sb=item&sub=UF&menustr=080020000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Blessed Reaching Augments (Rec 20)", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1524&rsa=Fletching&sb=item&sub=UFraidseals&ins=Blessed&menustr=080060000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=438&rsa=Brewing&sb=item&sub=UF&menustr=080030000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Eminent Reaching Augments (66 - 70)", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1552&rsa=Fletching&sb=item&sub=UFraidseals&ins=Eminent&menustr=080060000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Exalted Reaching Augments (71 - 75)", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1531&rsa=Fletching&sb=item&sub=UFraidseals&ins=Exalted&menustr=080060000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=520&rsa=Fletching&sb=item&sub=UF&menustr=080060000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=431&rsa=Jewelcraft&sb=item&sub=UF&menustr=080070000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "No Skill", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=507&rsa=No%20Skill&sb=item&sub=UF&menustr=080085000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=439&rsa=Pottery&sb=item&sub=UF&menustr=080090000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Revered Reaching Augments (31-40)", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1547&rsa=Fletching&sb=item&sub=UFraidseals&ins=Revered&menustr=080060000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Sacred Reaching Augments (51 - 60)", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1551&rsa=Fletching&sb=item&sub=UFraidseals&ins=Sacred&menustr=080060000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=440&rsa=Smithing&sb=item&sub=UF&menustr=080100000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Sublime Reaching Augments (76 - 80)", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1542&rsa=Fletching&sb=item&sub=UFraidseals&ins=Sublime&menustr=080060000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=441&rsa=Tailoring&sb=item&sub=UF&menustr=080110000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=442&rsa=Tinkering&sb=item&sub=UF&menustr=080120000000"},
		{ExpId: 16, ExpName: "Underfoot", PageTitle: "Venerable Reaching Augments (81 - 85)", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1545&rsa=Fletching&sb=item&sub=UFraidseals&ins=Venerable&menustr=080060000000"},
		{ExpId: 17, ExpName: "House of Thule", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=517&rsa=Baking&sb=item&sub=HOT&menustr=080020000000"},
		{ExpId: 17, ExpName: "House of Thule", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=518&rsa=Brewing&sb=item&sub=HOT&menustr=080030000000"},
		{ExpId: 17, ExpName: "House of Thule", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=521&rsa=Fletching&sb=item&sub=HOT&menustr=080060000000"},
		{ExpId: 17, ExpName: "House of Thule", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=428&rsa=Jewelcraft&sb=item&sub=HOT&menustr=080070000000"},
		{ExpId: 17, ExpName: "House of Thule", PageTitle: "Make Poison", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=585&rsa=Make%20Poison&sub=item&sub=HoT&menustr=080080000000"},
		{ExpId: 17, ExpName: "House of Thule", PageTitle: "No Skill", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=509&rsa=No%20Skill&sb=item&sub=HOT&menustr=080085000000"},
		{ExpId: 17, ExpName: "House of Thule", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=522&rsa=Pottery&sb=item&sub=HoT&menustr=080090000000"},
		{ExpId: 17, ExpName: "House of Thule", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=429&rsa=Smithing&sb=item&sub=HOT&menustr=080100000000"},
		{ExpId: 17, ExpName: "House of Thule", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=523&rsa=Tailoring&sb=item&sub=HOT&menustr=080110000000"},
		{ExpId: 17, ExpName: "House of Thule", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=524&rsa=Tinkering&sb=item&sub=HOT&menustr=080120000000"},
		{ExpId: 18, ExpName: "Veil of Alaris", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=545&rsa=Baking&sb=item&sub=VOA&menustr=080020000000"},
		{ExpId: 18, ExpName: "Veil of Alaris", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=546&rsa=Brewing&sb=item&sub=VoA&menustr=080030000000"},
		{ExpId: 18, ExpName: "Veil of Alaris", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=547&rsa=Fletching&sb=item&sub=VOA&menustr=080060000000"},
		{ExpId: 18, ExpName: "Veil of Alaris", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=548&rsa=Jewelcraft&sb=item&sub=VOA&menustr=080070000000"},
		{ExpId: 18, ExpName: "Veil of Alaris", PageTitle: "Make Poison", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=584&rsa=Make%20Poison&sub=item&sub=VOA&menustr=080080000000"},
		{ExpId: 18, ExpName: "Veil of Alaris", PageTitle: "No Skill", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=549&rsa=No%20Skill&sb=item&sub=VoA&menustr=080085000000"},
		{ExpId: 18, ExpName: "Veil of Alaris", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=550&rsa=Pottery&sb=item&sub=VoA&menustr=080090000000"},
		{ExpId: 18, ExpName: "Veil of Alaris", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=551&rsa=Smithing&sb=item&sub=VoA&menustr=080100000000"},
		{ExpId: 18, ExpName: "Veil of Alaris", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=552&rsa=Tailoring&sb=item&sub=VoA&menustr=080110000000"},
		{ExpId: 18, ExpName: "Veil of Alaris", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=553&rsa=Tinkering&sb=item&sub=VoA&menustr=080120000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=560&rsa=Baking&sb=item&sub=ROF&menustr=080020000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=560&rsa=Baking&sb=item&sub=ROF&menustr=080020000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=561&rsa=Brewing&sb=item&sub=ROF&menustr=080030000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=561&rsa=Brewing&sb=item&sub=ROF&menustr=080030000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=562&rsa=Fletching&sb=item&sub=ROF&menustr=080060000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=562&rsa=Fletching&sb=item&sub=ROF&menustr=080060000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=563&rsa=Jewelcraft&sb=item&sub=ROF&menustr=080070000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=563&rsa=Jewelcraft&sb=item&sub=ROF&menustr=080070000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Make Poison", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=583&rsa=Make%20Poison&sub=item&sub=ROF&menustr=080080000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Make Poison", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=583&rsa=Make%20Poison&sub=item&sub=ROF&menustr=080080000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "No Skill", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=564&rsa=No%20Skill&sb=item&sub=ROF&menustr=080085000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "No Skill", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=564&rsa=No%20Skill&sb=item&sub=ROF&menustr=080085000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=564&rsa=Pottery&sb=item&sub=ROF&menustr=080090000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=564&rsa=Pottery&sb=item&sub=ROF&menustr=080090000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=568&rsa=Smithing&sb=item&sub=ROF&menustr=080100000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=568&rsa=Smithing&sb=item&sub=ROF&menustr=080100000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=565&rsa=Tailoring&sb=item&sub=ROF&menustr=080110000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=565&rsa=Tailoring&sb=item&sub=ROF&menustr=080110000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=566&rsa=Tinkering&sb=item&sub=ROF&menustr=080120000000"},
		{ExpId: 19, ExpName: "Rain of Fear", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=566&rsa=Tinkering&sb=item&sub=ROF&menustr=080120000000"},
		{ExpId: 20, ExpName: "Call of the Forsaken", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=572&rsa=Baking&sb=item&sub=CoF&menustr=080020000000"},
		{ExpId: 20, ExpName: "Call of the Forsaken", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=573&rsa=Brewing&sb=item&sub=CoF&menustr=080030000000"},
		{ExpId: 20, ExpName: "Call of the Forsaken", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=574&rsa=Fletching&sb=item&sub=CoF&menustr=080060000000"},
		{ExpId: 20, ExpName: "Call of the Forsaken", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=575&rsa=Jewelcraft&sb=item&sub=CoF&menustr=080070000000"},
		{ExpId: 20, ExpName: "Call of the Forsaken", PageTitle: "Make Poison", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=582&rsa=Make%20poison&sub=item&sub=CoF&menustr=080080000000"},
		{ExpId: 20, ExpName: "Call of the Forsaken", PageTitle: "No Skill", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=576&rsa=No%20Skill&sb=item&sub=CoF&menustr=080085000000"},
		{ExpId: 20, ExpName: "Call of the Forsaken", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=577&rsa=Pottery&sb=item&sub=CoF&menustr=080090000000"},
		{ExpId: 20, ExpName: "Call of the Forsaken", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=578&rsa=Smithing&sb=item&sub=CoF&menustr=080100000000"},
		{ExpId: 20, ExpName: "Call of the Forsaken", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=579&rsa=Tailoring&sb=item&sub=CoF&menustr=080110000000"},
		{ExpId: 20, ExpName: "Call of the Forsaken", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=580&rsa=Tinkering&sb=item&sub=CoF&menustr=080120000000"},
		{ExpId: 21, ExpName: "The Darkened Sea", PageTitle: "Alchemy Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=652&rsa=Alchemy&sb=item&sub=TDS&menustr=080010000000"},
		{ExpId: 21, ExpName: "The Darkened Sea", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=607&rsa=Baking&sb=item&sub=TDS&menustr=080020000000"},
		{ExpId: 21, ExpName: "The Darkened Sea", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=608&rsa=Brewing&sb=item&sub=TDS&menustr=080030000000"},
		{ExpId: 21, ExpName: "The Darkened Sea", PageTitle: "Cultural Raid Seal Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=656&rsa=Pottery&sb=item&sub=TDSseals&menustr=080085000000"},
		{ExpId: 21, ExpName: "The Darkened Sea", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=609&rsa=Fletching&sb=item&sub=TDS&menustr=080060000000"},
		{ExpId: 21, ExpName: "The Darkened Sea", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=610&rsa=Jewelcraft&sb=item&sub=TDS&menustr=080070000000"},
		{ExpId: 21, ExpName: "The Darkened Sea", PageTitle: "Make Poison", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=611&rsa=Make%20poison&sub=item&sub=TDS&menustr=080080000000"},
		{ExpId: 21, ExpName: "The Darkened Sea", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=612&rsa=Pottery&sb=item&sub=TDS&menustr=080090000000"},
		{ExpId: 21, ExpName: "The Darkened Sea", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=613&rsa=Smithing&sb=item&sub=TDS&menustr=080100000000"},
		{ExpId: 21, ExpName: "The Darkened Sea", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=614&rsa=Tailoring&sb=item&sub=TDS&menustr=080110000000"},
		{ExpId: 21, ExpName: "The Darkened Sea", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=615&rsa=Tinkering&sb=item&sub=TDS&menustr=080120000000"},
		{ExpId: 22, ExpName: "The Broken Mirror", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1611&rsa=Smithing&sb=item&sub=TBM&menustr=080100000000"},
		{ExpId: 22, ExpName: "The Broken Mirror", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1609&rsa=Tailoring&sb=item&sub=TBM&menustr=080110000000"},
		{ExpId: 23, ExpName: "Empires of Kunark", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=738&rsa=Baking&sb=item&sub=Eok&menustr=080020000000"},
		{ExpId: 23, ExpName: "Empires of Kunark", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=742&rsa=Brewing&sb=item&sub=Eok&menustr=080030000000"},
		{ExpId: 23, ExpName: "Empires of Kunark", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=734&rsa=Jewelcraft&sb=item&sub=EoK&menustr=080070000000"},
		{ExpId: 23, ExpName: "Empires of Kunark", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=730&rsa=Pottery&sb=item&sub=EoK&menustr=080090000000"},
		{ExpId: 23, ExpName: "Empires of Kunark", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=722&rsa=Smithing&sb=item&sub=EoK&menustr=080100000000"},
		{ExpId: 23, ExpName: "Empires of Kunark", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=726&rsa=Tailoring&sb=item&sub=EoK&menustr=080110000000"},
		{ExpId: 24, ExpName: "Ring of Scale", PageTitle: "Alchemy Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1071&rsa=Alchemy&sb=item&sub=RoS"},
		{ExpId: 24, ExpName: "Ring of Scale", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1075&rsa=Baking&sb=item&sub=RoS"},
		{ExpId: 24, ExpName: "Ring of Scale", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1081&rsa=Brewing&sb=item&sub=RoS"},
		{ExpId: 24, ExpName: "Ring of Scale", PageTitle: "Fishing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1085&rsa=Fishing&sb=item&sub=RoS"},
		{ExpId: 24, ExpName: "Ring of Scale", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1089&rsa=Fletching&sb=item&sub=RoS"},
		{ExpId: 24, ExpName: "Ring of Scale", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1093&rsa=Jewelcraft&sb=item&sub=RoS"},
		{ExpId: 24, ExpName: "Ring of Scale", PageTitle: "Make Poison Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1096&rsa=Make%20Poison&sb=item&sub=RoS"},
		{ExpId: 24, ExpName: "Ring of Scale", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=1100&rsa=Pottery&sb=item&sub=RoS"},
		{ExpId: 24, ExpName: "Ring of Scale", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1101&rsa=Smithing&sb=item&sub=RoS"},
		{ExpId: 24, ExpName: "Ring of Scale", PageTitle: "Spell Research Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1104&rsa=Spell%20Research&sb=item&sub=RoS"},
		{ExpId: 24, ExpName: "Ring of Scale", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1107&rsa=Tailoring&sb=item&sub=RoS"},
		{ExpId: 24, ExpName: "Ring of Scale", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1111&rsa=Tinkering&sb=item&sub=RoS"},
		{ExpId: 25, ExpName: "The Burning Lands", PageTitle: "Armor Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1137&rsa=No%20Skill&sb=item&sub=TBLArmor"},
		{ExpId: 25, ExpName: "The Burning Lands", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1129&rsa=Baking&sb=item&sub=TBL"},
		{ExpId: 25, ExpName: "The Burning Lands", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1135&rsa=Brewing&sb=item&sub=TBL"},
		{ExpId: 25, ExpName: "The Burning Lands", PageTitle: "Fishing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1141&rsa=Fishing&sb=item&sub=TBL"},
		{ExpId: 25, ExpName: "The Burning Lands", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1139&rsa=Fletching&sb=item&sub=TBL"},
		{ExpId: 25, ExpName: "The Burning Lands", PageTitle: "Gear Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1388&rsa=No%20Skill&sb=item&sub=TBLGear"},
		{ExpId: 25, ExpName: "The Burning Lands", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1146&rsa=Jewelcraft&sb=item&sub=TBL"},
		{ExpId: 25, ExpName: "The Burning Lands", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=1152&rsa=Pottery&sb=item&sub=TBL"},
		{ExpId: 25, ExpName: "The Burning Lands", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1156&rsa=Smithing&sb=item&sub=TBL"},
		{ExpId: 25, ExpName: "The Burning Lands", PageTitle: "Spell Research Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1162&rsa=Spell%20Research&sb=item&sub=TBL"},
		{ExpId: 25, ExpName: "The Burning Lands", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1157&rsa=Tailoring&sb=item&sub=TBL"},
		{ExpId: 25, ExpName: "The Burning Lands", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1166&rsa=Tinkering&sb=item&sub=TBL"},
		{ExpId: 25, ExpName: "The Burning Lands", PageTitle: "Weapon Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1392&rsa=No%20Skill&sb=item&sub=TBLWeapons"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Alchemy Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1182&rsa=Alchemy&sb=item&sub=ToV"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Armor Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1417&rsa=No%20Skill&sb=item&sub=ToVArmor"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1187&rsa=Baking&sb=item&sub=Tov"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1193&rsa=Brewing&sb=item&sub=Tov"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Fishing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1200&rsa=Fishing&sb=item&sub=Tov"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1202&rsa=Fletching&sb=item&sub=Tov"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Gear Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1424&rsa=No%20Skill&sb=item&sub=ToVGear"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1204&rsa=Jewelcraft&sb=item&sub=Tov"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Make Poison", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1207&rsa=Make%20poison&sub=item&sub=TOv"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=1210&rsa=Pottery&sb=item&sub=ToV"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Quest Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1619&rsa=Jewelcraft&sb=item&sub=ToVQuest"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1212&rsa=Smithing&sb=item&sub=Tov"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Spell Research Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1214&rsa=Spell%20Research&sb=item&sub=Tov"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1216&rsa=Tailoring&sb=item&sub=Tov"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1219&rsa=Tinkering&sb=item&sub=Tov"},
		{ExpId: 26, ExpName: "Torment of Velious", PageTitle: "Weapon Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1411&rsa=No%20Skill&sb=item&sub=ToVWeapon"},
		{ExpId: 27, ExpName: "Claws of Veeshan", PageTitle: "Armor Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1282&rsa=No%20Skill&sb=item&sub=CovArmor"},
		{ExpId: 27, ExpName: "Claws of Veeshan", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1245&rsa=Baking&sb=item&sub=CoV"},
		{ExpId: 27, ExpName: "Claws of Veeshan", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1250&rsa=Brewing&sb=item&sub=CoV"},
		{ExpId: 27, ExpName: "Claws of Veeshan", PageTitle: "Fishing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1253&rsa=Fishing&sb=item&sub=CoV"},
		{ExpId: 27, ExpName: "Claws of Veeshan", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1257&rsa=Fletching&sb=item&sub=CoV"},
		{ExpId: 27, ExpName: "Claws of Veeshan", PageTitle: "Gear Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1286&rsa=No%20Skill&sb=item&sub=CovGear"},
		{ExpId: 27, ExpName: "Claws of Veeshan", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1261&rsa=Jewelcraft&sb=item&sub=Cov"},
		{ExpId: 27, ExpName: "Claws of Veeshan", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=1262&rsa=Pottery&sb=item&sub=CoV"},
		{ExpId: 27, ExpName: "Claws of Veeshan", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1265&rsa=Smithing&sb=item&sub=Cov"},
		{ExpId: 27, ExpName: "Claws of Veeshan", PageTitle: "Spell Research Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1276&rsa=Spell%20Research&sb=item&sub=Cov"},
		{ExpId: 27, ExpName: "Claws of Veeshan", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1272&rsa=Tailoring&sb=item&sub=Cov"},
		{ExpId: 27, ExpName: "Claws of Veeshan", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1278&rsa=Tinkering&sb=item&sub=Cov"},
		{ExpId: 27, ExpName: "Claws of Veeshan", PageTitle: "Weapon Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1294&rsa=No%20Skill&sb=item&sub=CovWeapon"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Armor Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1379&rsa=No%20Skill&sb=item&sub=ToLArmor"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1353&rsa=Baking&sb=item&sub=ToL"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1359&rsa=Brewing&sb=item&sub=ToL"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Fishing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1360&rsa=Fishing&sb=item&sub=ToL"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1363&rsa=Fletching&sb=item&sub=ToL"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Gear Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1383&rsa=No%20Skill&sb=item&sub=ToLGear"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1364&rsa=Jewelcraft&sb=item&sub=ToL"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Make Poison Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1405&rsa=Make%20poison&sb=item&sub=ToL"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=1365&rsa=Pottery&sb=item&sub=ToL"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1369&rsa=Smithing&sb=item&sub=ToL"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Spell Research Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1373&rsa=Spell%20Research&sb=item&sub=ToL"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1376&rsa=Tailoring&sb=item&sub=ToL"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1377&rsa=Tinkering&sb=item&sub=ToL"},
		{ExpId: 28, ExpName: "Terror of Luclin", PageTitle: "Weapon Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1385&rsa=No%20Skill&sb=item&sub=ToLWeapon"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Alchemy Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1398&rsa=Alchemy&sb=item&sub=NoS"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Armor Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1505&rsa=No%20Skill&sb=item&sub=NoSArmor"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Baking Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1480&rsa=Baking&sb=item&sub=NoS"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Brewing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1482&rsa=Brewing&sb=item&sub=NoS"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Fishing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1483&rsa=Fishing&sb=item&sub=NoS"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Fletching Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1485&rsa=Fletching&sb=item&sub=NoS"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Gear Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1507&rsa=No%20Skill&sb=item&sub=NoSGear"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Jewelcraft Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1474&rsa=Jewelcraft&sb=item&sub=NoS"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=1489&rsa=Pottery&sb=item&sub=NoS"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Smithing Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1496&rsa=Smithing&sb=item&sub=NoS"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Spell Research Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1492&rsa=Spell%20Research&sb=item&sub=NoS"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Tailoring Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1499&rsa=Tailoring&sb=item&sub=NoS"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1501&rsa=Tinkering&sb=item&sub=NoS"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Type 18/19 Augment Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1616&rsa=Jewelcraft&sb=item&sub=NoSQuest"},
		{ExpId: 29, ExpName: "Night of Shadows", PageTitle: "Weapon Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=1510&rsa=No%20Skill&sb=item&sub=NoSWeapon"},
	}

	expansion := os.Args[2]

	c.LoadItemCache()

	if expansion == "all" {
		for _, r := range list {
			c.parseRecipePage(r)
			c.SaveItemCache()
		}
	} else {
		expansionId, err := strconv.Atoi(expansion)
		if err != nil {
			fmt.Println(err)
		}
		for _, r := range list {
			if r.ExpId == expansionId {
				c.parseRecipePage(r)
			}
		}
	}

	c.SaveItemCache()

	// dump recipes to json file recipes.json
	data, err := json.MarshalIndent(recipes, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("data/eqtraders/recipes.json", data, 0644)
	if err != nil {
		fmt.Println(err)
	}

	//pp.Println(string(data))
}

func removeDigits(r rune) rune {
	if unicode.IsDigit(r) {
		r = -1
	}
	return r
}

func (c *ScrapeCommand) FindExpansionFromText(text string) Expansion {
	sort.Slice(expansions, func(i, j int) bool {
		return expansions[i].ExpansionName > expansions[j].ExpansionName
	})

	for _, e := range expansions {
		if strings.Contains(text, e.ExpansionName) {
			return e
		}
	}

	return Expansion{}
}

func (c *ScrapeCommand) getStringInBetween(str, before, after string) string {
	if !strings.Contains(str, before) {
		return ""
	}

	a := strings.SplitAfterN(str, before, 2)
	b := strings.SplitAfterN(a[len(a)-1], after, 2)
	if 1 == len(b) {
		return b[0]
	}
	return b[0][0 : len(b[0])-len(after)]
}

var recipeWriteMutex = &sync.Mutex{}

func (c *ScrapeCommand) parseRecipePage(r ExpansionRecipe) {
	fmt.Printf("> Parsing recipe page: %v&printer=normal\n", r.Url)

	// get page slug
	hash := md5.Sum([]byte(r.Url))
	pageSlug := slug.Make(fmt.Sprintf("%v-%v-%v", r.ExpName, r.PageTitle, hex.EncodeToString(hash[:])))
	file := fmt.Sprintf("data/eqtraders/site-cache/%v.html", pageSlug)
	contents := ""

	// check if cache file exists
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		// read cache file
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
		}

		contents = string(data)
		contents = strings.ReplaceAll(contents, "<td>&nbsp;&nbsp;</td>", "")
	}

	// if cache doesn't exist, fetch page and write cache
	if len(contents) == 0 {
		resp, err := http.Get(r.Url + "&printer=normal")
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		// write site cache
		err = os.WriteFile(file, data, 0644)
		if err != nil {
			fmt.Println(err)
		}

		// some pages are formatted strange and have an extra <td> beginning of the row
		// https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=1365&rsa=Pottery&sb=item&sub=ToL&printer=normal
		contents = string(data)
		contents = strings.ReplaceAll(contents, "<td>&nbsp;&nbsp;</td>", "")
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(contents)))
	if err != nil {
		fmt.Println(err)
	}

	wp := workerpool.New(50)

	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}

		wp.Submit(func() {

			// this is a td that contains <td>&nbsp;&nbsp;</td> and messed up the parsing
			messedUpRow := false
			messedUpRowStr, _ := s.Find("td").First().Html()
			if len(messedUpRowStr) == 4 {
				messedUpRow = true
			}

			recipeName := strings.TrimSpace(s.Find("td a").First().Text())
			recipeNameHtml, err := s.Find("td").First().Html()
			if err != nil {
				fmt.Println(err)
			}
			recipe, err := s.Find("td").Next().Html()
			if err != nil {
				fmt.Println(err)
			}
			recipeText := s.Find("td").Next().Text()

			// some pages are formatted strange and have an extra <td> beginning of the row
			if messedUpRow {
				recipeName = s.Find("td").First().Next().Text()
				recipeNameHtml = strings.TrimSpace(s.Find("td").First().Next().Text())
				recipe, err = s.Find("td").First().Next().Next().Html()
				if err != nil {
					c.logger.Error().Err(err).Msg("error parsing recipe html")
				}
				recipeText = s.Find("td").First().Next().Next().Text()

				//fmt.Printf("messed up recipeName [%v] recipeNameHtml [%v] recipeText [%v] recipe [%v]\n", recipeName, recipeNameHtml, recipeText, recipe)
			}

			consumeContainer := false

			if len(singleRecipe) > 0 && recipeName != singleRecipe {
				return
			}

			// second attempt to get recipe name
			if recipeName == "" {
				recipeName = strings.Split(recipeName, "\n")[0]
			}

			if recipeName == "" {
				return
			}

			// trivial
			trivialInt := 0
			trivial := s.Find("td").Next().Next().Text()
			//pp.Println(trivial)

			if messedUpRow {
				trivial = s.Find("td").First().Next().Next().Next().Text()
			}

			trivialInt, _ = strconv.Atoi(trivial)

			noFail := false
			if strings.Contains(trivial, "(No-Fail)") {
				trivial = strings.TrimSpace(strings.ReplaceAll(trivial, "(No-Fail)", ""))
				noFail = true
				trivialInt, err = strconv.Atoi(trivial)
				if err != nil {
					fmt.Println("error parsing trivial [%v] err [%v]", trivial, err.Error())
				}
			} else if strings.Contains(trivial, "no fail") {
				noFail = true
				trivialInt = 0
			}

			components := c.getStringInBetween(recipe, "Components:", "In:")
			in := c.getStringInBetween(recipe, "In:", "Yield:")

			var requiredSkillLevel int
			if strings.Contains(recipe, "Required Skill Level:") {
				reqStr := c.getStringInBetween(recipe, "Required Skill Level:", "<")
				reqStr = strings.ReplaceAll(reqStr, "&gt;", "")
				if strings.Contains(reqStr, "&") {
					reqStr = strings.Split(reqStr, "&")[0]
				}

				reqStr = strings.TrimSpace(reqStr)

				// parse out "300" in Required Skill Level: 300
				requiredSkillLevel, err = strconv.Atoi(reqStr)
				if err != nil {
					c.logger.Info().Msgf("error parsing required skill level [%v] err [%v]", c.getStringInBetween(recipe, "Required Skill Level:", "<"), err.Error())
				}
			}

			var componentsList []Item
			for _, s := range strings.Split(components, ",") {
				s = c.stripTradersComments(s)

				quantity := 1
				if strings.Contains(s, "(") {
					qty := strings.TrimSpace(c.getStringInBetween(s, "(", ")"))

					qty = strings.ReplaceAll(qty, "x", "")

					if strings.Contains(qty, "right") {
						continue
					}

					quantity, err = strconv.Atoi(qty)
					if err != nil {
						fmt.Println("error parsing component quantity [%v] err [%v]", qty, err.Error())
					}
				}

				name := c.getStringInBetween(s, ">", "<")
				name = html.UnescapeString(name)

				componentsList = append(componentsList, Item{
					ItemId:   c.getItemIdFromHtml(s),
					ItemName: name,
					Count:    quantity,
				})
			}
			var inList []Item
			for _, s := range strings.Split(in, ",") {
				if strings.TrimSpace(s) == "" {
					continue
				}

				i := c.getItemIdFromHtml(s)
				if i == 0 {

					// attempt to resolve to object type
					name := c.getStringInBetween(s, ">", "<")
					name = strings.ReplaceAll(name, " (Stationary)", "")
					name = strings.ReplaceAll(name, " (Formerly Ak&#39;Anon Forge)", "")
					name = strings.TrimSpace(name)

					// html decode name
					name = html.UnescapeString(name)

					objectType := c.getObjectTypeFromName(name)
					if objectType.Type > 0 {
						//fmt.Println("world container [%v] resolved to object type [%v] for recipe [%v]", name, objectType.Type, recipeName)
						inList = append(inList, Item{
							ItemId:   objectType.Type,
							ItemName: name,
							Count:    1,
						})
						continue
					}
					fmt.Println("world container [%v] not found, attempting to resolve to object type for recipe [%v]", name, recipeName)
					continue
				}

				inList = append(inList, Item{
					ItemId:   i,
					ItemName: c.getStringInBetween(s, ">", "<"),
					Count:    1,
				})
			}

			// yield
			yieldInt := 1
			yield := strings.TrimSpace(strings.ReplaceAll(c.getStringInBetween(recipe, "Yield:", "<b"), "</b>", ""))
			split := strings.Split(yield, "<")
			if len(split) > 1 {
				yield = split[0]
			}
			if len(yield) > 0 {
				yieldInt, err = strconv.Atoi(yield)
				if err != nil {
					fmt.Println("error parsing yield [%v] for recipe [%v] err [%v]", yield, recipeName, err.Error())
				}
			}

			// return items
			returns := c.getStringInBetween(recipe, "Also Returns:", "On Failure Returns")
			var returnItems []Item
			for _, s := range strings.Split(returns, ",") {
				i := c.getItemIdFromHtml(s)
				if i == 0 {
					continue
				}

				s = c.stripTradersComments(s)

				quantity := 1
				if strings.Contains(s, "(") {
					qty := c.getStringInBetween(s, "(", ")")
					quantity, err = strconv.Atoi(qty)
					if err != nil {
						fmt.Println("error parsing returns quantity [%v] err [%v]", qty, err.Error())
					}
				}

				name := c.getStringInBetween(s, ">", "<")
				name = html.UnescapeString(name)

				returnItems = append(returnItems, Item{
					ItemId:   i,
					ItemName: name,
					Count:    quantity,
				})
			}

			onFailureReturns := c.getStringInBetween(recipe, "On Failure Returns:", ",")
			var failureReturnItems []Item
			for _, s := range strings.Split(onFailureReturns, ",") {
				i := c.getItemIdFromHtml(s)
				if i == 0 {
					continue
				}

				s = c.stripTradersComments(s)

				// remove everything after Notes: (if it exists)
				s = strings.Split(s, "Notes:")[0]

				quantity := 1
				if strings.Contains(s, "(") {
					qty := c.getStringInBetween(s, "(", ")")
					quantity, err = strconv.Atoi(qty)
					if err != nil {
						fmt.Println("error parsing failure quantity [%v] err [%v]", qty, err.Error())
					}
				}

				failureReturnItems = append(failureReturnItems, Item{
					ItemId:   i,
					ItemName: c.getStringInBetween(s, ">", "<"),
					Count:    quantity,
				})
			}

			notesSplit := strings.Split(recipeText, "Notes:")
			var notes string
			learnedItem := ""

			if len(notesSplit) > 1 {
				notes = strings.TrimSpace(notesSplit[1])
				notes = strings.ReplaceAll(notes, "&#39;", "")
				notes = strings.ReplaceAll(notes, "&#39;", "")

				if strings.Contains(notes, "is consumed") {
					consumeContainer = true
					//fmt.Println("Notes ", notes, " consumed in recipe ", recipeName)
				}

				if strings.Contains(notes, "You may need to purchase and scribe") {
					learnedItem = strings.TrimSpace(c.getStringInBetween(notes, "You may need to purchase and scribe", "in order to perform this combine"))
				} else if strings.Contains(notes, "You must purchase and scribe") && strings.Contains(notes, "from") {
					learnedItem = strings.TrimSpace(c.getStringInBetween(notes, "You must purchase and scribe", "from"))
				} else if strings.Contains(notes, "You must purchase and scribe") {
					learnedItem = strings.TrimSpace(c.getStringInBetween(notes, "You must purchase and scribe", "in order to perform this combine"))
				} else if strings.Contains(notes, "You can purchase and scribe") {
					learnedItem = strings.TrimSpace(c.getStringInBetween(notes, "You can purchase and scribe", "in order to learn this recipe"))
				}

				learnedItem = strings.ReplaceAll(learnedItem, " to learn this recipe without experimenting.", "")
				learnedItem = strings.ReplaceAll(learnedItem, " in order to learn this recipe without experiementing.", "")

				// strip ' on outside of learned item
				learnedItem = strings.TrimPrefix(learnedItem, "'")
				learnedItem = strings.TrimSuffix(learnedItem, "'")

				// extract Simple Charms from Can scribe to Learn:
				if strings.Contains(notes, "Can scribe to Learn:") {
					learnedItem = strings.TrimSpace(strings.Split(notes, "Can scribe to Learn:")[1])
				}

				learnedItem = strings.ReplaceAll(learnedItem, "Compendium on", "Compendium of")
				learnedItem = strings.ReplaceAll(learnedItem, "Regal froglok Tailoring", "Regal Froglok Cultural Tailoring")
				learnedItem = strings.ReplaceAll(learnedItem, "Regal froglok Smithing", "Regal Froglok Cultural Smithing")
				learnedItem = strings.ReplaceAll(learnedItem, "Numinous Reaching Augments for Fletchers", "Numinous Reaching Weapon Augments for Fletchers")
				learnedItem = strings.ReplaceAll(learnedItem, "Numinous Reaching Augments for Jewelers", "Numinous Reaching Weapon Augments for Jewelers")
				learnedItem = strings.ReplaceAll(learnedItem, "Numinous Weapon Augmenting for Jewelers", "Numinous Reaching Weapon Augments for Jewelers")
				learnedItem = strings.ReplaceAll(learnedItem, "Numnious", "Numinous")
				learnedItem = strings.ReplaceAll(learnedItem, "Glorious Erudian", "Glorious Erudite")

				learnedItem = strings.TrimPrefix(learnedItem, "the ")
				learnedItem = strings.TrimSuffix(learnedItem, " scroll")
				learnedItem = strings.ReplaceAll(learnedItem, " in order to complete this combine", "")
				learnedItem = strings.ReplaceAll(learnedItem, " in order to learn this combine without experimentation", "")
				learnedItem = strings.ReplaceAll(learnedItem, " before you can complete this combine", "")
				learnedItem = strings.ReplaceAll(learnedItem, " to learn these recipes without experimenting", "")
				learnedItem = strings.ReplaceAll(learnedItem, " in order to learn the recipe without experimentation", "")
				learnedItem = strings.ReplaceAll(learnedItem, " in order to make this combine", "")
				learnedItem = strings.ReplaceAll(learnedItem, " in order to make this item", "")
				learnedItem = strings.ReplaceAll(learnedItem, " in order to make this item", "")
				learnedItem = strings.ReplaceAll(learnedItem, " without experimentation scribing may be required", "")
				learnedItem = strings.ReplaceAll(learnedItem, " to learn these recipes", "")
				learnedItem = strings.ReplaceAll(learnedItem, " in order", "")
				learnedItem = strings.ReplaceAll(learnedItem, " to make this recipe", "")
				learnedItem = strings.ReplaceAll(learnedItem, "  to learn this recipes", "")
				learnedItem = strings.ReplaceAll(learnedItem, "  to learn this recipe", "")
				learnedItem = strings.ReplaceAll(learnedItem, " to learn this recipes", "")
				learnedItem = strings.ReplaceAll(learnedItem, " to learn this recipe", "")
				learnedItem = strings.ReplaceAll(learnedItem, "Cultural Drakkin", "Drakkin Cultural")
				learnedItem = strings.ReplaceAll(learnedItem, " Erudian ", " Erudite ")
				learnedItem = strings.ReplaceAll(learnedItem, " No-Fail", "")
				learnedItem = strings.ReplaceAll(learnedItem, " (No-Fail)", "")
				learnedItem = strings.ReplaceAll(learnedItem, " without experimentation (scribing may be required).", "")
				learnedItem = strings.ReplaceAll(learnedItem, "A Guide.", "A Guide")

				if strings.Contains(learnedItem, "Glorious") && !strings.Contains(learnedItem, "Cultural") {
					learnedItem = strings.ReplaceAll(learnedItem, " Smithing", " Cultural Smithing")
				}

				if strings.Contains(learnedItem, "Glorious") && !strings.Contains(learnedItem, "Cultural") {
					learnedItem = strings.ReplaceAll(learnedItem, " Tailoring", " Cultural Tailoring")
				}

				if strings.Contains(learnedItem, "Regal") && !strings.Contains(learnedItem, "Cultural") {
					learnedItem = strings.ReplaceAll(learnedItem, " Smithing", " Cultural Smithing")
				}

				if strings.Contains(learnedItem, "Regal") && !strings.Contains(learnedItem, "Cultural") {
					learnedItem = strings.ReplaceAll(learnedItem, " Tailoring", " Cultural Tailoring")
				}

				learnedItem = strings.ReplaceAll(learnedItem, "Regal Weapon Cultural Tailoring", "Regal Weapon Tailoring")
				learnedItem = strings.ReplaceAll(learnedItem, "Regal Weapon Cultural Smithing", "Regal Weapon Smithing")

				// strip numbers end of learned item
				//learnedItem = strings.Map(removeDigits, learnedItem)
				re := regexp.MustCompile(`[^a-zA-Z \-:]+`)
				//fmt.Println(Input)
				learnedItem = re.ReplaceAllString(learnedItem, "")

				if len(learnedItem) > 0 {
					fmt.Printf("-- Learned by item [%v] for recipe [%v]\n", learnedItem, recipeName)
				}
			}

			// remove everything in recipeName in parentheses
			// Resilient Supplicant's Earring of Rallos Zek (Evolving Level 1/2)
			recipeName = strings.Split(recipeName, "(Evolving Level")[0]
			recipeName = strings.TrimSpace(recipeName)

			fmt.Printf(
				"- Recipe [%v] Expansion [%v] (%v) Skill [%v] Trivial [%v] Required Skill [%v] Components [%v] Containers [%v] Returns [%v] \n",
				recipeName,
				r.ExpName,
				r.ExpId,
				c.getSkillFromName(r.PageTitle).SkillName,
				trivialInt,
				requiredSkillLevel,
				len(componentsList),
				len(inList),
				len(returnItems),
			)

			skill := c.getSkillFromName(r.PageTitle)
			if skill.SkillId == 0 {
				skill = c.getSkillFromName(r.Url)
			}
			if skill.SkillId == 0 {
				c.logger.Info().Msgf("skill not found for recipe [%v] [%v]", recipeName, r.Url)
			}

			if len(componentsList) == 0 {
				c.logger.Info().Msgf("components not found for recipe [%v] [%v]", recipeName, r.Url)
				return
			}

			for _, component := range componentsList {
				if component.ItemId == 0 {
					c.logger.Info().Msgf("component item not found for recipe [%v] [%v]", recipeName, r.Url)
					return
				}
			}

			recipeItemId := 0
			if messedUpRow {
				var item models.Item
				c.db.Where("Name = ?", strings.TrimSpace(recipeName)).First(&item)

				if item.ID > 0 {
					recipeItemId = item.ID
				}
			} else {
				recipeItemId = c.getItemIdFromHtml(recipeNameHtml)
			}

			r := Recipe{
				RecipeName:         recipeName,
				Skill:              skill,
				ExpansionId:        r.ExpId,
				ExpansionName:      r.ExpName,
				Trivial:            trivialInt,
				RequiredSkillLevel: requiredSkillLevel,
				ConsumeContainer:   consumeContainer,
				NoFail:             noFail,
				RecipeItemId:       recipeItemId,
				Components:         componentsList,
				In:                 inList,
				Yield:              yieldInt,
				Returns:            returnItems,
				FailureReturns:     failureReturnItems,
				LearnedByItem: Item{
					ItemName: learnedItem,
				},
			}

			recipeWriteMutex.Lock()
			recipes = append(recipes, r)
			recipeWriteMutex.Unlock()
		})
	})

	wp.StopWait()
}

var itemLookupCache = make(map[string]int, 0)
var itemLookupReadOnlyCache = make(map[string]int, 0) // conncurent

var lookupWriteMutex = &sync.Mutex{}

func (c *ScrapeCommand) getItemIdFromHtml(html string) int {
	url := c.getStringInBetween(html, "href=\"", "\">")
	tradersItemId := c.getStringInBetween(url, "item=", "&amp;")

	val, ok := itemLookupReadOnlyCache[tradersItemId]
	if ok {
		return val
	}

	url = "https://www.eqtraders.com/" + url

	if skipLookups {
		return 0
	}

	resp, err := http.Get(url + "&printer=normal")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	itemId := c.getStringInBetween(string(data), "EQ item ID: ", "<")

	i, err := strconv.Atoi(itemId)
	if err != nil {
		return 0
	}

	lookupWriteMutex.Lock()
	itemLookupCache[tradersItemId] = i
	lookupWriteMutex.Unlock()

	return i
}

func (c *ScrapeCommand) SaveItemCache() {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)

	// Encoding the map
	err := e.Encode(itemLookupCache)
	if err != nil {
		fmt.Println(err)
	}

	// write to file
	err = os.WriteFile(filepath.Join("./data/eqtraders/item-lookup-cache.gob"), b.Bytes(), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *ScrapeCommand) LoadItemCache() {
	file := filepath.Join("./data/eqtraders/item-lookup-cache.gob")
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return
	}

	b, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	d := gob.NewDecoder(bytes.NewReader(b))

	// Decoding the serialized data
	err = d.Decode(&itemLookupCache)
	if err != nil {
		fmt.Println(err)
	}

	// copy to read only cache
	for k, v := range itemLookupCache {
		itemLookupReadOnlyCache[k] = v
	}
}

type Skill struct {
	SkillId   int    `json:"skill_id"`
	SkillName string `json:"skill_name"`
}

var skills = []Skill{
	{SkillId: 56, SkillName: "Make Poison"},
	{SkillId: 56, SkillName: "Make%20Poison"},
	{SkillId: 57, SkillName: "Tinkering"},
	{SkillId: 58, SkillName: "Research"},
	{SkillId: 59, SkillName: "Alchemy"},
	{SkillId: 60, SkillName: "Baking"},
	{SkillId: 61, SkillName: "Tailoring"},
	{SkillId: 63, SkillName: "Smithing"},
	{SkillId: 64, SkillName: "Fletching"},
	{SkillId: 65, SkillName: "Brewing"},
	{SkillId: 68, SkillName: "Jewelcraft"},
	{SkillId: 69, SkillName: "Pottery"},
	// add fishing
}

func (c *ScrapeCommand) getSkillFromName(name string) Skill {
	for _, s := range skills {
		if strings.Contains(name, s.SkillName) {
			return s
		}
	}

	return Skill{}
}

func (c *ScrapeCommand) getObjectTypeFromName(name string) ObjectType {
	for _, o := range objectTypes {
		if strings.Contains(name, o.Name) {
			return o
		}
	}

	for _, o := range objectTypes {
		for _, n := range o.TradersNames {
			if strings.Contains(name, n) {
				return o
			}
		}
	}

	return ObjectType{}
}

func (c *ScrapeCommand) stripTradersComments(s string) string {
	s = strings.ReplaceAll(s, "(temporary)", "")
	s = strings.ReplaceAll(s, "(Cannot Scribe)", "")
	s = strings.ReplaceAll(s, "(Legacy)", "")
	s = strings.ReplaceAll(s, "(removed)", "")
	s = strings.ReplaceAll(s, "(foraged)", "")
	s = strings.ReplaceAll(s, "(looted)", "")
	s = strings.ReplaceAll(s, "(Pattern)", "")
	return s
}

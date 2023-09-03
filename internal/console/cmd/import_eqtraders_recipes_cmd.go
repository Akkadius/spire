package cmd

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type ImportEqTradersCommand struct {
	db      *gorm.DB
	logger  *logrus.Logger
	command *cobra.Command
}

func (c *ImportEqTradersCommand) Command() *cobra.Command {
	return c.command
}

func NewImportEqTradersCommand(
	db *gorm.DB,
	logger *logrus.Logger,
) *ImportEqTradersCommand {
	i := &ImportEqTradersCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use:   "import:eq-traders-recipes [expansion_number]",
			Short: "A command for importing eq traders recipes",
		},
	}

	i.command.Args = cobra.MinimumNArgs(1)
	i.command.Run = i.Handle

	return i
}

type Expansion struct {
	ExpansionNumber int    `json:"expansion_number"`
	ExpansionName   string `json:"expansion_name"`
	ShortName       string `json:"short_name"`
	MaxLevel        int    `json:"max_level"`
}

// fill in expansions from json data above
var expansions = []Expansion{
	{
		ExpansionNumber: 0,
		ExpansionName:   "Original",
		ShortName:       "",
		MaxLevel:        50,
	},
	{
		ExpansionNumber: 1,
		ExpansionName:   "Ruins of Kunark",
		ShortName:       "RoK",
		MaxLevel:        60,
	},
	{
		ExpansionNumber: 2,
		ExpansionName:   "Scars of Velious",
		ShortName:       "SoV",
		MaxLevel:        60,
	},
	{
		ExpansionNumber: 3,
		ExpansionName:   "Shadows of Luclin",
		ShortName:       "SoL",
		MaxLevel:        60,
	},
	{
		ExpansionNumber: 4,
		ExpansionName:   "Planes of Power",
		ShortName:       "PoP",
		MaxLevel:        65,
	},
	{
		ExpansionNumber: 5,
		ExpansionName:   "Legacy of Ykesha",
		ShortName:       "LoY",
		MaxLevel:        65,
	},
	{
		ExpansionNumber: 6,
		ExpansionName:   "Lost Dungeons of Norrath",
		ShortName:       "LDoN",
		MaxLevel:        65,
	},
	{
		ExpansionNumber: 7,
		ExpansionName:   "Gates of Discord",
		ShortName:       "GoD",
		MaxLevel:        65,
	},
	{
		ExpansionNumber: 8,
		ExpansionName:   "Omens of War",
		ShortName:       "OoW",
		MaxLevel:        70,
	},
	{
		ExpansionNumber: 9,
		ExpansionName:   "Dragons of Norrath",
		ShortName:       "DoN",
		MaxLevel:        70,
	},
	{
		ExpansionNumber: 10,
		ExpansionName:   "Depths of Darkhollow",
		ShortName:       "DoDH",
		MaxLevel:        70,
	},
	{
		ExpansionNumber: 11,
		ExpansionName:   "Prophecy of Ro",
		ShortName:       "PoR",
		MaxLevel:        70,
	},
	{
		ExpansionNumber: 12,
		ExpansionName:   "The Serpent's Spine",
		ShortName:       "TSS",
		MaxLevel:        75,
	},
	{
		ExpansionNumber: 13,
		ExpansionName:   "The Buried Sea",
		ShortName:       "TBS",
		MaxLevel:        75,
	},
	{
		ExpansionNumber: 14,
		ExpansionName:   "Secrets of Faydwer",
		ShortName:       "SoF",
		MaxLevel:        80,
	},
	{
		ExpansionNumber: 15,
		ExpansionName:   "Seeds of Destruction",
		ShortName:       "SoD",
		MaxLevel:        85,
	},
	{
		ExpansionNumber: 16,
		ExpansionName:   "Underfoot",
		ShortName:       "UF",
		MaxLevel:        85,
	},
	{
		ExpansionNumber: 17,
		ExpansionName:   "House of Thule",
		ShortName:       "HoT",
		MaxLevel:        90,
	},
	{
		ExpansionNumber: 18,
		ExpansionName:   "Veil of Alaris",
		ShortName:       "VoA",
		MaxLevel:        95,
	},
	{
		ExpansionNumber: 19,
		ExpansionName:   "Rain of Fear",
		ShortName:       "RoF",
		MaxLevel:        100,
	},
	{
		ExpansionNumber: 20,
		ExpansionName:   "Call of the Forsaken",
		ShortName:       "CoTF",
		MaxLevel:        100,
	},
	{
		ExpansionNumber: 21,
		ExpansionName:   "The Darkened Sea",
		ShortName:       "TDS",
		MaxLevel:        105,
	},
	{
		ExpansionNumber: 22,
		ExpansionName:   "The Broken Mirror",
		ShortName:       "TBM",
		MaxLevel:        105,
	},
	{
		ExpansionNumber: 23,
		ExpansionName:   "Empires of Kunark",
		ShortName:       "EoK",
		MaxLevel:        105,
	},
	{
		ExpansionNumber: 24,
		ExpansionName:   "Ring of Scale",
		ShortName:       "RoS",
		MaxLevel:        110,
	},
	{
		ExpansionNumber: 25,
		ExpansionName:   "The Burning Lands",
		ShortName:       "TBL",
		MaxLevel:        110,
	},
	{
		ExpansionNumber: 26,
		ExpansionName:   "Torment of Velious",
		ShortName:       "ToV",
		MaxLevel:        115,
	},
	{
		ExpansionNumber: 27,
		ExpansionName:   "Claws of Veeshan",
		ShortName:       "CoV",
		MaxLevel:        115,
	},
	{
		ExpansionNumber: 28,
		ExpansionName:   "Terror of Luclin",
		ShortName:       "ToL",
		MaxLevel:        120,
	},
	{
		ExpansionNumber: 29,
		ExpansionName:   "Night of Shadows",
		ShortName:       "NoS",
		MaxLevel:        120,
	},
}

type ExpansionRecipe struct {
	ExpId     int
	ExpName   string
	PageTitle string
	Url       string
}

// Handle implementation of the Command interface
func (c *ImportEqTradersCommand) Handle(cmd *cobra.Command, args []string) {
	list := []ExpansionRecipe{
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
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Pottery Recipes", Url: "https://www.eqtraders.com/recipes/pottery_recipe_page.php?article=299&rsa=Pottery&sb=item&sub=dron&menustr=080010000000"},
		{ExpId: 9, ExpName: "Dragons of Norrath", PageTitle: "Tinkering Recipes", Url: "https://www.eqtraders.com/recipes/recipe_page.php?article=297&rsa=Tinkering&sub=dron&sb=item&menustr=080120000000"},
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
			c.logger.Error(err)
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
		c.logger.Error(err)
	}
	err = os.WriteFile("recipes.json", data, 0644)
	if err != nil {
		c.logger.Error(err)
	}

	//pp.Println(string(data))
}

func (c *ImportEqTradersCommand) FindExpansionFromText(text string) Expansion {

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

func (c *ImportEqTradersCommand) getStringInBetween(str, before, after string) string {
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

type Item struct {
	ItemId   int    `json:"item_id"`
	ItemName string `json:"item_name"`
	Count    int    `json:"count"`
}

type Recipe struct {
	RecipeName     string `json:"recipe_name"`
	Skill          Skill  `json:"skill"`
	ExpansionId    int    `json:"expansion_id"`
	ExpansionName  string `json:"expansion_name"`
	Trivial        int    `json:"trivial"`
	NoFail         bool   `json:"no_fail"`
	RecipeItemId   int    `json:"recipe_item_id"`
	Components     []Item `json:"components"`
	In             []Item `json:"in"`
	Yield          int    `json:"yield"`
	Returns        []Item `json:"returns"`
	FailureReturns []Item `json:"failure_returns"`
}

var recipes []Recipe

func (c *ImportEqTradersCommand) parseRecipePage(r ExpansionRecipe) {
	c.logger.Info("Parsing recipe page: ", r.Url+"&printer=normal")
	resp, err := http.Get(r.Url + "&printer=normal")
	if err != nil {
		c.logger.Error(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error(err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		c.logger.Error(err)
	}

	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}

		recipeName := strings.TrimSpace(s.Find("td").First().Text())
		recipeNameHtml, err := s.Find("td").First().Html()
		if err != nil {
			c.logger.Error(err)
		}
		recipe, err := s.Find("td").Next().Html()
		if err != nil {
			c.logger.Error(err)
		}

		if recipeName == "" {
			return
		}

		// trivial
		trivialInt := 0
		trivial := s.Find("td").Next().Next().Text()
		noFail := false
		if strings.Contains(trivial, "(No-Fail)") {
			trivial = strings.TrimSpace(strings.ReplaceAll(trivial, "(No-Fail)", ""))
			noFail = true
			trivialInt, err = strconv.Atoi(trivial)
			if err != nil {
				c.logger.Errorf("error parsing trivial [%v] err [%v]", trivial, err.Error())
			}
		} else if strings.Contains(trivial, "no fail") {
			noFail = true
			trivialInt = 0
		}

		components := c.getStringInBetween(recipe, "Components:", "In:")
		in := c.getStringInBetween(recipe, "In:", "Yield:")

		var componentsList []Item
		for _, s := range strings.Split(components, ",") {
			s = strings.ReplaceAll(s, "(temporary)", "")
			s = strings.ReplaceAll(s, "(Cannot Scribe)", "")

			quantity := 1
			if strings.Contains(s, "(") {
				qty := strings.TrimSpace(c.getStringInBetween(s, "(", ")"))

				quantity, err = strconv.Atoi(qty)
				if err != nil {
					c.logger.Errorf("error parsing quantity [%v] err [%v]", qty, err.Error())
				}
			}

			componentsList = append(componentsList, Item{
				ItemId:   c.getItemIdFromHtml(s),
				ItemName: c.getStringInBetween(s, ">", "<"),
				Count:    quantity,
			})
		}
		var inList []Item
		for _, s := range strings.Split(in, ",") {
			i := c.getItemIdFromHtml(s)
			if i == 0 {
				continue
			}

			inList = append(inList, Item{
				ItemId:   i,
				ItemName: c.getStringInBetween(s, ">", "<"),
				Count:    1,
			})
		}

		// yield
		yield := strings.TrimSpace(strings.ReplaceAll(c.getStringInBetween(recipe, "Yield:", "<b"), "</b>", ""))
		split := strings.Split(yield, "<")
		if len(split) > 1 {
			yield = split[0]
		}
		yieldInt, err := strconv.Atoi(yield)
		if err != nil {
			c.logger.Errorf("error parsing yield [%v] for recipe [%v] err [%v]", yield, recipeName, err.Error())
		}

		// return items
		returns := c.getStringInBetween(recipe, "Also Returns:", "On Failure Returns")
		var returnItems []Item
		for _, s := range strings.Split(returns, ",") {
			i := c.getItemIdFromHtml(s)
			if i == 0 {
				continue
			}

			s = strings.ReplaceAll(s, "(temporary)", "")
			s = strings.ReplaceAll(s, "(Cannot Scribe)", "")

			quantity := 1
			if strings.Contains(s, "(") {
				qty := c.getStringInBetween(s, "(", ")")
				quantity, err = strconv.Atoi(qty)
				if err != nil {
					c.logger.Errorf("error parsing quantity [%v] err [%v]", qty, err.Error())
				}
			}

			returnItems = append(returnItems, Item{
				ItemId:   i,
				ItemName: c.getStringInBetween(s, ">", "<"),
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

			s = strings.ReplaceAll(s, "(temporary)", "")
			s = strings.ReplaceAll(s, "(Cannot Scribe)", "")

			quantity := 1
			if strings.Contains(s, "(") {
				qty := c.getStringInBetween(s, "(", ")")
				quantity, err = strconv.Atoi(qty)
				if err != nil {
					c.logger.Errorf("error parsing quantity [%v] err [%v]", qty, err.Error())
				}
			}

			failureReturnItems = append(failureReturnItems, Item{
				ItemId:   i,
				ItemName: c.getStringInBetween(s, ">", "<"),
				Count:    quantity,
			})
		}

		r := Recipe{
			RecipeName:     strings.Split(recipeName, "\n")[0],
			Skill:          c.getSkillFromName(r.PageTitle),
			ExpansionId:    r.ExpId,
			ExpansionName:  r.ExpName,
			Trivial:        trivialInt,
			NoFail:         noFail,
			RecipeItemId:   c.getItemIdFromHtml(recipeNameHtml),
			Components:     componentsList,
			In:             inList,
			Yield:          yieldInt,
			Returns:        returnItems,
			FailureReturns: failureReturnItems,
		}

		pp.Println(r)
		fmt.Println("")

		recipes = append(recipes, r)
	})
}

var itemLookupCache = make(map[string]int, 0)

func (c *ImportEqTradersCommand) getItemIdFromHtml(html string) int {
	url := c.getStringInBetween(html, "href=\"", "\">")
	tradersItemId := c.getStringInBetween(url, "item=", "&amp;")
	val, ok := itemLookupCache[tradersItemId]
	if ok {
		//c.logger.Info("Found item in cache: ", val)
		return val
	}

	url = "https://www.eqtraders.com/" + url

	resp, err := http.Get(url + "&printer=normal")
	if err != nil {
		c.logger.Error(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error(err)
	}

	itemId := c.getStringInBetween(string(data), "EQ item ID: ", "<")

	i, err := strconv.Atoi(itemId)
	if err != nil {
		return 0
	}

	itemLookupCache[tradersItemId] = i

	return i
}

func (c *ImportEqTradersCommand) SaveItemCache() {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)

	// Encoding the map
	err := e.Encode(itemLookupCache)
	if err != nil {
		c.logger.Error(err)
	}

	// write to file
	err = os.WriteFile(filepath.Join("itemLookupCache.gob"), b.Bytes(), 0644)
	if err != nil {
		c.logger.Error(err)
	}
}

func (c *ImportEqTradersCommand) LoadItemCache() {
	file := filepath.Join("itemLookupCache.gob")

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return
	}

	b, err := os.ReadFile(file)
	if err != nil {
		c.logger.Error(err)
	}

	d := gob.NewDecoder(bytes.NewReader(b))

	// Decoding the serialized data
	err = d.Decode(&itemLookupCache)
	if err != nil {
		c.logger.Error(err)
	}
}

type Skill struct {
	SkillId   int    `json:"skill_id"`
	SkillName string `json:"skill_name"`
}

var skills = []Skill{
	{SkillId: 56, SkillName: "Make Poison"},
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
}

func (c *ImportEqTradersCommand) getSkillFromName(name string) Skill {
	for _, s := range skills {
		if strings.Contains(name, s.SkillName) {
			return s
		}
	}

	return Skill{}
}

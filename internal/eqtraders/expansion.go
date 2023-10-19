package eqtraders

type Expansion struct {
	ExpansionNumber int    `json:"expansion_number"`
	ExpansionName   string `json:"expansion_name"`
	ShortName       string `json:"short_name"`
	MaxLevel        int    `json:"max_level"`
}

type ExpansionRecipe struct {
	ExpId     int
	ExpName   string
	PageTitle string
	Url       string
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
